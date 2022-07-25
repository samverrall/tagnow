package semvar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    SemvarTag
		wantErr bool
	}{
		{
			name:  "Valid new semvar tag",
			input: "v1.0.1",
			want: SemvarTag{
				Major: 1,
				Minor: 0,
				Patch: 1,
			},
			wantErr: false,
		},
		{
			name:    "In-valid new semvar tag",
			input:   "vA.0.1",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			semvar, err := NewFromString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, semvar.ToString(), tt.want.ToString(), "The two words should be the same.")
		})
	}
}
