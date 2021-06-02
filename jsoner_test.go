package jsoner

import (
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	type args struct {
		input map[string]interface{}
	}
	type MapType map[string]interface{}
	type ArrayType []interface{}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			"t1",
			args{input: MapType{
				"a": MapType{
					"x": MapType{
						"p": ArrayType{"l", "o", "l"},
					},
				},
			},
			},
			`{"a":{"x":{"p":["l","o","l"]}}}`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := ToJSON(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.TrimSpace(gotOutput) != strings.TrimSpace(tt.wantOutput) {
				t.Errorf("ToJSON() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
