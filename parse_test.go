package nosj

import (
	"bufio"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name                string
		inputStr            string
		want                Node
		expectedRemainedStr string
	}{
		{"simple object: value - string", `{"name" : "Pavel"}`,
			&Object{items: map[string]Node{"name": &String{"Pavel"}}},
			""},
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
