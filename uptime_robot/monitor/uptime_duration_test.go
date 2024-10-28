package monitor

import (
	"github.com/go-test/deep"
	"testing"
	"time"
)

func TestUptimeDuration_UnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		in       string
		want     UptimeDuration
		hasError bool
	}{
		"invalid: too few items": {
			in:       "\"123-234\"",
			want:     UptimeDuration{},
			hasError: true,
		},
		"invalid: too many items": {
			in:       "\"123-234-345-456\"",
			want:     UptimeDuration{},
			hasError: true,
		},
		"invalid: contains strings": {
			in:       "\"123-asdf-234\"",
			want:     UptimeDuration{},
			hasError: true,
		},
		"invalid: contains floats": {
			in:       "\"123-asdf-234\"",
			want:     UptimeDuration{},
			hasError: true,
		},
		"valid": {
			in: "\"123-234-345\"",
			want: UptimeDuration{
				Up:     123 * time.Second,
				Down:   234 * time.Second,
				Paused: 345 * time.Second,
			},
			hasError: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := &UptimeDuration{}
			//err := json.Unmarshal([]byte(tt.in), got)
			err := got.UnmarshalJSON([]byte(tt.in))
			if err != nil && !tt.hasError {
				t.Errorf("unexpected error: %s", err)
				return
			}

			if diff := deep.Equal(tt.want, *got); diff != nil {
				t.Error(diff)
			}
		})
	}
}
