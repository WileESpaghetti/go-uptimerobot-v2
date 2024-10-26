package number

import (
	"encoding/json"
	"testing"
)

func TestOptionalNumber_UnmarshalJSON(t *testing.T) {
	/*
	 * other cases I was too lazy to test
	 * floating point (unquoted, quoted, all zeros after the decimal, locales that don't use `.` as the separator)
	 * other number formats supported by JavaScript: octal, hex, bin, exponential (are these even valid JSON?)
	 */
	// TODO test floating point number
	// TODO test floating point number (all zeros
	tests := map[string]struct {
		in           string
		want         Optional
		wantErr      bool
		hasNumberErr bool // json.Number encoding error
		hasIntErr    bool
		hasStringErr bool
	}{
		"empty string (double quotes)": {
			in:           "\"\"",
			want:         Optional(0),
			wantErr:      false,
			hasNumberErr: true,
			hasIntErr:    true,
			hasStringErr: false,
		},
		//"empty string (single quotes)": {
		// single quotes are not valid JSON
		//	in:           "''",
		//	want:         Optional(0),
		//	wantErr:      true,
		//	hasNumberErr: true,
		//	hasIntErr:    true,
		//	hasStringErr: true,
		//},
		"no input": {
			in:           "", // empty string is considered invalid JSON by the encoder
			want:         Optional(0),
			wantErr:      true,
			hasNumberErr: true,
			hasIntErr:    true,
			hasStringErr: true,
		},
		"null": {
			in:           "null",
			want:         Optional(0),
			wantErr:      false,
			hasNumberErr: false,
			hasIntErr:    false,
			hasStringErr: false,
		},
		"0": {
			in:           "0",
			want:         Optional(0),
			wantErr:      false,
			hasNumberErr: false,
			hasIntErr:    false,
			hasStringErr: true,
		},
		"0 (quoted)": {
			in:           "\"0\"",
			want:         Optional(0),
			wantErr:      false,
			hasNumberErr: false,
			hasIntErr:    true,
			hasStringErr: false,
		},
		"1337": {
			in:           "1337",
			want:         Optional(1337),
			wantErr:      false,
			hasNumberErr: false,
			hasIntErr:    false,
			hasStringErr: true,
		},
		"1337 (quoted)": {
			in:           "\"1337\"",
			want:         Optional(1337),
			wantErr:      false,
			hasNumberErr: false,
			hasIntErr:    true,
			hasStringErr: false,
		},
		"-42": {
			in:           "-42",
			want:         Optional(-42),
			wantErr:      false,
			hasNumberErr: false,
			hasIntErr:    false,
			hasStringErr: true,
		},
		"-42 (quoted)": {
			in:           "\"-42\"",
			want:         Optional(-42),
			wantErr:      false,
			hasNumberErr: false,
			hasIntErr:    true,
			hasStringErr: false,
		},
		"string": {
			in:           "\"1h3ll0 w0rld\"",
			want:         Optional(0),
			wantErr:      true,
			hasNumberErr: true,
			hasIntErr:    true,
			hasStringErr: false,
		},
		"83.1": {
			in:           "83.1",
			want:         Optional(0),
			wantErr:      true,
			hasNumberErr: false,
			hasIntErr:    true,
			hasStringErr: true,
		},
		"83.1 (quoted)": {
			in:           "\"83.1\"",
			want:         Optional(0),
			wantErr:      true,
			hasNumberErr: false,
			hasIntErr:    true,
			hasStringErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got Optional
			err := json.Unmarshal([]byte(test.in), &got)
			gotErr := err != nil
			if test.wantErr != gotErr {
				t.Errorf("encoding Optional: wanted %t, got %t", test.wantErr, gotErr)
			}

			var gotNumber json.Number
			numErr := json.Unmarshal([]byte(test.in), &gotNumber)
			gotNumErr := numErr != nil
			if test.hasNumberErr != gotNumErr {
				t.Errorf("encoding json.Number: wanted %t, got %t", test.hasNumberErr, gotNumErr)
			}

			var gotInt int64
			intErr := json.Unmarshal([]byte(test.in), &gotInt)
			gotIntErr := intErr != nil
			if test.hasIntErr != gotIntErr {
				t.Errorf("encoding int64: wanted %t, got %t", test.hasIntErr, gotIntErr)
			}

			var gotStr string
			strErr := json.Unmarshal([]byte(test.in), &gotStr)
			gotStrErr := strErr != nil
			if test.hasStringErr != gotStrErr {
				t.Errorf("encoding string: wanted %t, got %t", test.hasStringErr, gotStrErr)
			}

			if got != test.want {
				t.Errorf("value: wanted %d, got %d", test.want, got)
			}
		})
	}
}
