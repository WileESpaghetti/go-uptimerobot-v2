package monitor

import "testing"

func TestKeywordType_String(t *testing.T) {
	tests := map[string]struct {
		in   KeywordType
		want string
	}{
		"exists": {
			in:   KeywordTypeExists,
			want: "Exists",
		},
		"not exists": {
			in:   KeywordTypeNotExists,
			want: "Not Exists",
		},
		"zero value": {
			in:   KeywordType(0),
			want: "",
		},
		"unknown": {
			in:   KeywordType(999),
			want: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			//t.Parallel() // marks each test case as capable of running in parallel with each other
			if got := tt.in.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeywordCaseType_String(t *testing.T) {
	tests := map[string]struct {
		in   KeywordCaseType
		want string
	}{
		"case sensitive": {
			in:   KeywordCaseSensitive,
			want: "Case Sensitive",
		},
		"case insensitive": {
			in:   KeywordCaseInsensitive,
			want: "Case Insensitive",
		},
		"other": {
			in:   KeywordCaseType(999),
			want: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			//t.Parallel() // marks each test case as capable of running in parallel with each other
			if got := tt.in.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
