package main

import "testing"

func Test_ftoc(t *testing.T) {
	type args struct {
		temperature float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{temperature: 212.0}, want: 100},
		{name: "test2", args: args{temperature: 214.0}, want: 101.11111111111111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ftoc(tt.args.temperature); got != tt.want {
				t.Errorf("ftoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
