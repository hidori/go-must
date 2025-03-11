package must

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		fn func() (int, error)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDo2(t *testing.T) {
	type args struct {
		fn func() (int, int, error)
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Do2(tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Do2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDo3(t *testing.T) {
	type args struct {
		fn func() (int, int, int, error)
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Do3(tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do3() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Do3() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Do3() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
