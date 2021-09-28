package cli

import (
	"testing"
)

func TestWrap(t *testing.T) {

	// var a500 = strings.Repeat("A", 500)

	type args struct {
		s     string
		width int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		// {"repeat 500 A's", args{a500, 20}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.s, tt.args.width); got != tt.want {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
