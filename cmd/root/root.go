package root

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"io"

	"github.com/hjson/hjson-go"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	"github.com/Eun/minigo/pkg/minigo"
)

var (
	rootCmd = &cobra.Command{
		Use:   "minigo [flags] <.go files to run>",
		Short: "A mini golang interpreter",
		Long: `A mini golang interpreter based on yaegi-template and yaegi.
https://github.com/Eun/minigo`,
		RunE: run,
		Args: cobra.MinimumNArgs(1),
	}

	templatingFlag bool
	writeToFlag    string
	contextFlag    string
)

// Execute executes the root cmd.
func Execute() {
	rootCmd.Flags().BoolVarP(&templatingFlag, "template", "t", false, "enable templating")
	rootCmd.Flags().StringVarP(&writeToFlag, "out", "o", "", "write output to file")
	rootCmd.Flags().StringVarP(&contextFlag, "context", "c", "", "set context to the specified json object")
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// Version sets the rootCmd version to the specified value.
func Version(version string) {
	rootCmd.Version = version
}

func run(cmd *cobra.Command, args []string) error {
	var w io.Writer = os.Stdout
	if writeToFlag != "" {
		f, err := os.Create(writeToFlag)
		if err != nil {
			return xerrors.Errorf("unable to write file `%s': %w\n", writeToFlag, err)
		}
		defer f.Close()
		w = f
	}

	context, err := makeContext()
	if err != nil {
		return err
	}

	amountOfGoFiles := len(args)
	if amountOfGoFiles <= 0 {
		return nil
	}

	g, err := minigo.New(minigo.Config{
		TemplateMode: templatingFlag,
	})
	if err != nil {
		return xerrors.Errorf("unable to create minigo instance: %w", err)
	}

	for _, s := range args {
		f, err := os.Open(s)
		if err != nil {
			return xerrors.Errorf("unable to open file `%s': %v", s, err)
		}
		err = g.Run(f, context, w)
		_ = f.Close()
		if err != nil {
			return xerrors.Errorf("unable to exec file `%s': %v", s, err)
		}
	}
	return nil
}

func makeContext() (interface{}, error) {
	if contextFlag == "" {
		return nil, nil
	}

	contextFlag = strings.TrimFunc(contextFlag, func(r rune) bool {
		return unicode.IsSpace(r) || r == '\'' || r == '"'
	})
	if contextFlag == "" {
		return nil, nil
	}
	var dat map[string]interface{}
	if err := hjson.Unmarshal([]byte(contextFlag), &dat); err != nil {
		return nil, xerrors.Errorf("unable to decode context `%s': %w", contextFlag, err)
	}
	context, err := minigo.ConvertMapToStruct(dat)
	if err != nil {
		return nil, xerrors.Errorf("unable to convert context map to struct `%s': %w", contextFlag, err)
	}
	return context, nil
}
