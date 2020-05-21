package nosj

import (
	"strings"
	"testing"
)

func TestObject_getKey(t *testing.T) {
	tests := []struct {
		name                 string
		o                    *Object
		inputStr             string
		expectedRemainingStr string
		expectedKey          string
	}{
		{"parses valid string", NewObject(), `"name" : "pavel"}`, `"pavel"}`, "name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := prepareScanner((strings.NewReader(tt.inputStr)))

			key := tt.o.getKey(s)

			if key != tt.expectedKey {
				t.Errorf("expected key %s, got %s", tt.expectedKey, key)
			}

			if rem := ConsumeScanner(s); rem != tt.expectedRemainingStr {
				t.Errorf("expected remaining <<%s>>, got <<%s>>", tt.expectedRemainingStr, rem)
			}
		})
	}
}
