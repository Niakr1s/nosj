package nosj

import (
	"reflect"
	"testing"
)

func TestString_Parse(t *testing.T) {
	tests := []struct {
		name                string
		str                 *String
		inputStr            string
		want                Node
		expectedRemainedStr string
	}{
		{"valid string", NewString(), `"Pavel",`, &String{s: "Pavel"}, `,`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ScannerFromString(tt.inputStr)

			if got := tt.str.Parse(s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String.Parse() = %v, want %v", got, tt.want)
			}

			if rem := ConsumeScanner(s); rem != tt.expectedRemainedStr {
				t.Errorf("String.Parse(): remained %s, want %s", rem, tt.expectedRemainedStr)
			}
		})
	}
}
