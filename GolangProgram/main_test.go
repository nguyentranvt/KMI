package main

import "testing"

func TestPrintName(t *testing.T) {
	type args struct {
		Name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case 1",
			args: args{
				Name: "Nguyen Van A",
			},
			want: "Ho va ten Nguyen Van A",
		},
		{
			name: "Happy case 2",
			args: args{
				Name: "Pham Minh Duc",
			},
			want: "Ho va ten: Pham Minh Duc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintName(tt.args.Name); got != tt.want {
				t.Errorf("PrintName() = %v, want %v", got, tt.want)
			}
		})
	}
}
