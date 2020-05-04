package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		lvl  Level
		want string
	}{
		{
			name: "must returns debug string",
			lvl:  LevelDebug,
			want: "Debug",
		},
		{
			name: "must returns info string",
			lvl:  LevelInfo,
			want: "Info",
		},
		{
			name: "must returns warning string",
			lvl:  LevelWarning,
			want: "Warning",
		},
		{
			name: "must returns error string",
			lvl:  LevelError,
			want: "Error",
		},
		{
			name: "must returns fatal string",
			lvl:  LevelFatal,
			want: "Fatal",
		},
		{
			name: "must returns unknown string",
			lvl:  10,
			want: "Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.lvl.String(), tt.want)
		})
	}
}
