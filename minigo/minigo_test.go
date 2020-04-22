package minigo_test

import (
	"testing"

	"github.com/Eun/minigo/minigo"
	"github.com/google/go-cmp/cmp"
)

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
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ConvertMapToStruct() %s", cmp.Diff(got, tt.want))
			}
		})
	}
}
