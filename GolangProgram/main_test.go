package main

import (
	"testing"
)

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

func TestPrintTeamNumber(t *testing.T) {
	type args struct {
		Number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Happy case 3",
			args: args{
				Number: 2,
			},
			want: "Thanh vien nhom 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintTeamNumber(tt.args.Number); got != tt.want {
				t.Errorf("PrintTeamNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintAddress(t *testing.T) {
	type args struct {
		Number  int
		Address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case 4",
			args: args{
				Number:  10,
				Address: "Ha Noi",
			},
			want: "Dia chi: Ha Noi So nha: 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintAddress(tt.args.Number, tt.args.Address); got != tt.want {
				t.Errorf("PrintAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintTellNumber(t *testing.T) {
	type args struct {
		Number string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case 5",
			args: args{
				Number: "0123456789",
			},
			want: "So dien thoai: 0123456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintTellNumber(tt.args.Number); got != tt.want {
				t.Errorf("PrintTellNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}