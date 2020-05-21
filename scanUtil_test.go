package nosj

import (
	"bufio"
	"strings"
	"testing"
)

func TestSkip(t *testing.T) {
	type args struct {
		s     *bufio.Scanner
		runes []string
	}
	tests := []struct {
		name             string
		args             args
		expectedRemained string
	}{
		{"don't skip", args{bufio.NewScanner(strings.NewReader(`"lalala"`)), []string{" "}}, `"lalala"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Skip(tt.args.s, tt.args.runes...)
			gotRemained := ConsumeScanner(tt.args.s)
			if gotRemained != tt.expectedRemained {
				t.Errorf("Expected remained: %s, got %s", tt.expectedRemained, gotRemained)
			}
		})
	}
}
