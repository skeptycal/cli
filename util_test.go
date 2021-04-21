package cli

import (
	"os"
	"testing"
)

func TestBasicEncode(t *testing.T) {
	type args struct {
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"byte", args{byte(8)}, "\033[8m"},
		{"int", args{int(8)}, "\033[8m"},
		{"big int", args{int(0xFFFF)}, "\033[255m"},
		{"uint", args{uint(8)}, "\033[8m"},
		{"big uint", args{uint(0xFFFF)}, "\033[255m"},
		{"float64", args{float64(8)}, "\033[8m"},
		{"pi", args{3.14159}, "\033[3m"},
		{"string", args{"8"}, "\033[8m"},
		{"invalid string", args{"a"}, ""},  // no output
		{"nil", args{nil}, ""},             // no output
		{"os.Stdout", args{os.Stdout}, ""}, // no output

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BasicEncode(tt.args.b); got != tt.want {
				t.Errorf("BasicEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
