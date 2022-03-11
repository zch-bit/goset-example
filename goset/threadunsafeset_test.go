package goset

import "testing"

func Test_threadUnsafeSet_Add(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		s    threadUnsafeSet
		args args
		want bool
	}{
		{
			name: "test add a value",
			s:    threadUnsafeSet{},
			args: args{
				"a",
			},
			want: true,
		},
		{
			name: "test add an existing value",
			s:    threadUnsafeSet{"a": struct{}{}},
			args: args{
				"a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Add(tt.args.i); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
