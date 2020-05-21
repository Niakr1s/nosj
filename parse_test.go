package nosj

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	obj := &Object{items: map[string]Node{"name": &String{"Pavel"}}}
	obj2 := &Object{items: map[string]Node{"name": &String{"Pavel"}, "name1": &String{"Pavel1"}}}
	arr := &Array{items: []Node{obj, obj}}
	arr2 := &Array{items: []Node{obj2, obj2}}

	tests := []struct {
		name                string
		inputStr            string
		want                Node
		expectedRemainedStr string
	}{
		{"object with 1 key, value - string",
			`{"name" : "Pavel"}`,
			obj, ""},

		{"object with multiple keys, value - string",
			`{"name" : "Pavel", "name1" : "Pavel1"}`,
			obj2, ""},

		{"array with 2 objects with 1 key: value - string",
			`[{"name" : "Pavel"},{"name" : "Pavel"}]`,
			arr, ""},

		{"array with 2 objects with 2 key: value - string",
			`[{"name" : "Pavel", "name1" : "Pavel1"},{"name" : "Pavel", "name1" : "Pavel1"}]`,
			arr2, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ScannerFromString(tt.inputStr)

			if got := Parse(s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}

			if rem := ConsumeScanner(s); rem != tt.expectedRemainedStr {
				t.Errorf("Parse(): remained %s, want %s", rem, tt.expectedRemainedStr)
			}
		})
	}
}
