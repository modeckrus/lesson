package fib

import "testing"

func TestFib(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1 1 2 (3)",
			args{
				a: 3,
			},
			3,
		},
		{
			"1 1 2 3 (5)",
			args{
				a: 3,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.a); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
