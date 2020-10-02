package minigo_test

import (
	"testing"

	"reflect"

	"github.com/google/go-cmp/cmp"

	"github.com/Eun/minigo/pkg/minigo"
)

func structEqual(a, b interface{}) bool {
	sta := reflect.ValueOf(a).Elem()
	stb := reflect.ValueOf(b).Elem()

	if sta.Type().NumField() != stb.Type().NumField() {
		return false
	}

	for i := 0; i < sta.Type().NumField(); i++ {
		fa := sta.Type().Field(i)
		fb := stb.FieldByName(fa.Name)
		if !fb.IsValid() {
			return false
		}

		if !cmp.Equal(sta.Field(i).Interface(), fb.Interface()) {
			return false
		}
	}
	return true
}

func TestStructEqual(t *testing.T) {
	tests := []struct {
		a    interface{}
		b    interface{}
		want bool
	}{
		{
			&struct {
				Name string
			}{
				Name: "Joe",
			},
			&struct {
				Name string
			}{
				Name: "Joe",
			},
			true,
		},
		{
			&struct {
				Name string
			}{
				Name: "Alice",
			},
			&struct {
				Name string
			}{
				Name: "Joe",
			},
			false,
		},
		{
			&struct {
				ID int
			}{
				ID: 10,
			},
			&struct {
				Name string
			}{
				Name: "Joe",
			},
			false,
		},

		{
			&struct {
				ID   int
				Name string
			}{
				ID:   10,
				Name: "Joe",
			},
			&struct {
				Name string
			}{
				Name: "Joe",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := structEqual(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("structEqual() got %v want %v", got, tt.want)
			}
		})
	}
}

func TestConvertMapToStruct(t *testing.T) {
	tests := []struct {
		name    string
		v       map[string]interface{}
		want    interface{}
		wantErr bool
	}{
		{
			"simple",
			map[string]interface{}{"Name": "Joe", "ID": 10},
			&struct {
				ID   int
				Name string
			}{
				ID:   10,
				Name: "Joe",
			},
			false,
		},
		{
			"slice",
			map[string]interface{}{"Name": "Joe", "ID": 10, "Roles": []string{"Admin", "Developer"}},
			&struct {
				ID    int
				Name  string
				Roles []string
			}{
				ID:    10,
				Name:  "Joe",
				Roles: []string{"Admin", "Developer"},
			},
			false,
		},
		{
			"nested",
			map[string]interface{}{"Name": "Joe", "ID": 10, "Details": map[string]interface{}{"Country": "US"}},
			&struct {
				ID      int
				Name    string
				Details struct {
					Country string
				}
			}{
				ID:   10,
				Name: "Joe",
				Details: struct {
					Country string
				}{
					Country: "US",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := minigo.ConvertMapToStruct(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertMapToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !structEqual(got, tt.want) {
				t.Errorf("ConvertMapToStruct() %s", cmp.Diff(got, tt.want))
			}
		})
	}
}
