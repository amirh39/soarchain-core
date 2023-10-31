package utility

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCalculateDPREndTime checks the determinism and correctness of CalculateDPREndTime function
func TestCalculateDPREndTime(t *testing.T) {
	// Define test cases
	tests := []struct {
		name      string
		blockTime time.Time
		epochs    int
		want      int64
		wantErr   bool
		repeat    int // Number of times to repeat the function call for determinism check
	}{
		{
			name:      "Positive Epochs",
			blockTime: time.Date(2023, 10, 31, 12, 0, 0, 0, time.UTC),
			epochs:    10,
			want:      1698758100, // Expected Unix time
			wantErr:   false,
			repeat:    5,
		},
		{
			name:      "Zero Epochs",
			blockTime: time.Date(2023, 10, 31, 12, 0, 0, 0, time.UTC),
			epochs:    0,
			want:      1698753600, // Expected Unix time
			wantErr:   false,
			repeat:    5,
		},
		{
			name:      "Negative Epochs",
			blockTime: time.Date(2023, 10, 31, 12, 0, 0, 0, time.UTC),
			epochs:    -1,
			want:      0,
			wantErr:   true,
			repeat:    5,
		},
		{
			name:      "Large Number of Epochs",
			blockTime: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			epochs:    100000,
			want:      1717531200, // Expected Unix time
			wantErr:   false,
			repeat:    5,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.repeat; i++ {
				got, err := CalculateDPREndTime(tt.blockTime, tt.epochs)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.want, got, "Run %d failed for test %s", i+1, tt.name)
				}
			}
		})
	}
}
