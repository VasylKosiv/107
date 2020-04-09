package main

import (
	"testing"
)

func Test_findk(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{

		{
			name: "CorrectInput",
			args: 567,
			want: 4,
		},
		{
			name: "ZeroInput",
			args: 0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findk(tt.args); got != tt.want {
				t.Errorf("findk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputArg(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "CorrectInput",
			args: "567",
			want: 567,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inputArg(tt.args); got != tt.want {
				t.Errorf("inputArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputConsole(t *testing.T) {
	tests := []struct {
		name string

		want string
	}{
		{
			name: "CorrectInput",

			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inputConsole(); got != tt.want {
				t.Errorf("inputConsole() = %v, want %v", got, tt.want)
			}
		})
	}
}
