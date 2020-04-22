package minigo

import (
	"os"
	"sort"

	"io"
	"strings"

	"bytes"

	"reflect"

	yaegi_template "github.com/Eun/yaegi-template"
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"golang.org/x/xerrors"
)

// Config holds configuration values for Minigo
type Config struct {
	TemplateMode bool
}

// Minigo can be used to run .go files
type Minigo struct {
	t      *yaegi_template.Template
	config Config
}

// New creates a new Minigo instance
func New(config Config) (*Minigo, error) {
	t, err := yaegi_template.New(interp.Options{
		GoPath: os.Getenv("GOPATH"),
	}, stdlib.Symbols)
	if err != nil {
		return nil, err
	}
	if !config.TemplateMode {
		t.StartTokens = []rune{}
		t.EndTokens = []rune{}
	}
	return &Minigo{
		t:      t,
		config: config,
	}, nil
}

// Run runs the go interpreter on src and writes the output to output
func (minigo *Minigo) Run(src io.ReadSeeker, context interface{}, output io.Writer) error {
	t := yaegi_template.MustNew(interp.Options{
		GoPath: os.Getenv("GOPATH"),
	}, stdlib.Symbols)
	if !minigo.config.TemplateMode {
		t.StartTokens = []rune{}
		t.EndTokens = []rune{}
	}

	if err := skipShebang(src); err != nil {
		return xerrors.Errorf("unable to skip shebang: %w", err)
	}

	if err := t.Parse(src); err != nil {
		return xerrors.Errorf("unable to parse source: %w", err)
	}

	if _, err := t.Exec(output, context); err != nil {
		return xerrors.Errorf("unable to exec source: %w", err)
	}

	return nil
}

func skipShebang(rws io.ReadSeeker) error {
	var buf [128]byte
	var firstLine bytes.Buffer
	for {
		n, err := rws.Read(buf[:])
		if n <= 0 {
			if err == nil {
				return io.ErrUnexpectedEOF
			}
			if err == io.EOF {
				break
			}
			return err
		}
		firstLine.Write(buf[:n])
		if bytes.ContainsRune(buf[:n], '\n') {
			break
		}
	}
	if strings.HasPrefix(firstLine.String(), "#!") {
		// seek to the first \n
		newLine := bytes.IndexRune(firstLine.Bytes(), '\n')
		_, err := rws.Seek(int64(newLine), io.SeekStart)
		return err
	}
	_, err := rws.Seek(0, io.SeekStart)
	return err
}

// ConvertMapToStruct converts a map into a struct representation
func ConvertMapToStruct(v map[string]interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return nil, xerrors.Errorf("%#v is invalid", v)
	}

	instance := dynamicstruct.NewStruct()

	type KeyValue struct {
		Name  string
		Value reflect.Value
	}

	keys := rv.MapKeys()
	collectedKeys := make([]KeyValue, len(keys))
	for i, key := range keys {
		value := rv.MapIndex(key)
		if !value.CanInterface() {
			return nil, xerrors.Errorf("cannot make interface for value of %s", key.String())
		}

		if subMap, ok := value.Interface().(map[string]interface{}); ok {
			subStruct, err := ConvertMapToStruct(subMap)
			if err != nil {
				return nil, xerrors.Errorf("cannot convert child map (key=%s) to struct: %w", key.String(), err)
			}
			value = reflect.ValueOf(subStruct)
			if !value.CanInterface() {
				return nil, xerrors.Errorf("cannot make interface for value of %s", key.String())
			}
		}

		collectedKeys[i] = KeyValue{
			Name:  strings.Title(key.String()),
			Value: value.Elem(),
		}
	}

	// sort keys
	sort.Slice(collectedKeys, func(i, j int) bool {
		return strings.Compare(collectedKeys[i].Name, collectedKeys[j].Name) < 0
	})

	for _, key := range collectedKeys {
		instance = instance.AddField(key.Name, key.Value.Interface(), "")
	}

	st := instance.Build().New()
	rst := reflect.ValueOf(st).Elem()

	for _, key := range collectedKeys {
		f := rst.FieldByName(key.Name)
		if !f.IsValid() {
			return nil, xerrors.Errorf("unable to find %s in new struct", key.Name)
		}
		if !f.CanSet() {
			return nil, xerrors.Errorf("unable to set value for %s in new struct", key.Name)
		}
		f.Set(key.Value)
	}

	return st, nil
}
