package cli

import (
	"os"
	"testing"
)

// check will compare 'got' and 'want' values and report
// true if they are the same or false if they are different.
//
// 'name' is a descriptive name that will display in error messages.
// 'got' and 'want' should be of the same type.
// wantErr states whether an error is expected or not.
//
// Use 'nil' for *testing.T to skip reporting (not recommended)
func check(name string, got, want Any, wantErr bool, t *testing.T) bool {
	if want != got {
		if !wantErr {
			t.Errorf("%s = %v, want %v", name, got, want)
			return true
		}
	}
	return false
}

func TestBasicEncode(t *testing.T) {

	tests := []struct {
		name    string
		input   interface{}
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"byte", byte(8), "\033[8m", false},
		{"int", int(8), "\033[8m", false},
		{"big int", int(0xFFFF), "\033[255m", false},
		{"uint", uint(8), "\033[8m", false},
		{"big uint", uint(0xFFFF), "\033[255m", false},
		{"float64", float64(8.0042), "\033[8m", false},
		{"pi", 3.14159, "\033[3m", false},
		{"string", "8", "\033[8m", false},
		{"invalid string", "a", "", true},  // no output
		{"nil", nil, "", true},             // no output
		{"os.Stdout", os.Stdout, "", true}, // no output
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c, err := byteEncode(tt.input)
			if err != nil {
				if !tt.wantErr {
					t.Fatal(err)
				}
			} else {

				_ = check("NewColor()", NewColor(c).String(), tt.want, tt.wantErr, t)

				// check("BasicEncode()", BasicEncode(tt.input), tt.want, false, t)

			}
		})
	}
}

func TestColumns(t *testing.T) {

	t.Run("Columns()", func(t *testing.T) {

		ws, err := GetWinSize()

		if err != nil {
			t.Fatal(err)
		}

		got := Columns()
		if got < 1 || got > 1000 {
			t.Errorf("Columns() - expected int between 1 and 1000, got: %v (ws = %v)", got, ws)
		}
	})
}
