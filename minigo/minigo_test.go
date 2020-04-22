package minigo_test

import (
	"testing"

	"reflect"

	"github.com/Eun/minigo/minigo"
	"github.com/google/go-cmp/cmp"
)

func structEqual(a, b interface{}) bool {
	sta := reflect.ValueOf(a).Elem()
	stb := reflect.ValueOf(b).Elem()

	if sta.Type().NumField() != stb.Type().NumField() {
		return false
	}

	for i := 0; i < sta.Type().NumField(); i++ {
		fa := sta.Type().Field(i)
		fb := sta.FieldByName(fa.Name)

		if !cmp.Equal(sta.Field(i).Interface(), fb.Interface()) {
			return false
		}
	}
	return true
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
			map[string]interface{}{"Name": "Joe", "Id": 10},
			&struct {
				Id   int
				Name string
			}{
				Id:   10,
				Name: "Joe",
			},
			false,
		},
		{
			"slice",
			map[string]interface{}{"Roles": []string{"Admin", "Developer"}},
			&struct {
				Roles []string
			}{
				Roles: []string{"Admin", "Developer"},
			},
			false,
		},
		{
			"nested",
			map[string]interface{}{"Details": map[string]interface{}{"Country": "US"}},
			&struct {
				Details struct {
					Country string
				}
			}{
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
