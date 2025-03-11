package must

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	type args struct {
		fn func() (int, error)
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantPanic bool
	}{
		{
			name: "success: returns value",
			args: args{
				fn: func() (int, error) {
					return 1, nil
				},
			},
			want: 1,
		},
		{
			name: "failure: panics",
			args: args{
				fn: func() (int, error) {
					panic("error")
				},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					Do(tt.args.fn)
				})
			} else {
				got := Do(tt.args.fn)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDo2(t *testing.T) {
	type args struct {
		fn func() (int, int, error)
	}
	tests := []struct {
		name      string
		args      args
		want      int
		want1     int
		wantPanic bool
	}{
		{
			name: "success: returns values",
			args: args{
				fn: func() (int, int, error) {
					return 1, 2, nil
				},
			},
			want:  1,
			want1: 2,
		},
		{
			name: "failure: panics",
			args: args{
				fn: func() (int, int, error) {
					panic("error")
				},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					Do2(tt.args.fn)
				})
			} else {
				got, got1 := Do2(tt.args.fn)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.want1, got1)
			}
		})
	}
}

func TestDo3(t *testing.T) {
	type args struct {
		fn func() (int, int, int, error)
	}
	tests := []struct {
		name      string
		args      args
		want      int
		want1     int
		want2     int
		wantPanic bool
	}{
		{
			name: "success: returns values",
			args: args{
				fn: func() (int, int, int, error) {
					return 1, 2, 3, nil
				},
			},
			want:  1,
			want1: 2,
			want2: 3,
		},
		{
			name: "failure: panics",
			args: args{
				fn: func() (int, int, int, error) {
					panic("error")
				},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					Do3(tt.args.fn)
				})
			} else {
				got, got1, got2 := Do3(tt.args.fn)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.want1, got1)
				assert.Equal(t, tt.want2, got2)
			}
		})
	}
}
