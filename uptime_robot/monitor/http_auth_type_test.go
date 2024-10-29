package monitor

import "testing"

func TestHttpAuthType_String(t *testing.T) {
	tests := map[string]struct {
		in   HttpAuthType
		want string
	}{
		"none": {
			in:   HttpAuthTypeNone,
			want: "None",
		},
		"basic": {
			in:   HttpAuthTypeHttpBasic,
			want: "HTTP Basic Auth",
		},
		"digest": {
			in:   HttpAuthTypeDigest,
			want: "Digest",
		},
		"unknown": {
			in:   HttpAuthType(999),
			want: "Unknown",
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
