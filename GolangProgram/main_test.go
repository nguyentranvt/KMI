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

func TestPrintSex(t *testing.T) {
	type args struct {
		Sex string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Happy case 6",
			args: args{
				Sex: "Nam",
			},
			want: "Gioi tinh: Nam",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintSex(tt.args.Sex); got != tt.want {
				t.Errorf("PrintSex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintReview(t *testing.T) {
	type args struct {
		Review string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Happy case 7",
			args: args{
				Review: "Mê gái",
			},
			want: "Review của sếp: Mê gái",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintReview(tt.args.Review); got != tt.want {
				t.Errorf("PrintReview() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestPrintUpdateTime(t *testing.T) {
// 	type args struct {
// 		Time int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.

// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := PrintUpdateTime(tt.args.Time); got != tt.want {
// 				t.Errorf("PrintUpdateTime() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestPrintMyName(t *testing.T) {
	type args struct {
		ten string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case 1",
			args: args{
				ten: "huynh bao toan",
			},
			want: "toi ten la: huynh bao toan",
		}, {
			name: "Happy case 2",
			args: args{
				ten: "pham minh duc",
			},
			want: "toi ten la: pham minh duc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintMyName(tt.args.ten); got != tt.want {
				t.Errorf("PrintMyName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintGirlFriendCurrentNumber(t *testing.T) {
	type args struct {
		GirlFrendNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case 8",
			args: args{
				GirlFrendNumber: 3,
			},
			want: "Số lượng bạn gái hiện tại: 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintGirlFriendCurrentNumber(tt.args.GirlFrendNumber); got != tt.want {
				t.Errorf("PrintGirlFriendCurrentNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintIntroduce(t *testing.T) {
	type args struct {
		Introduce string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case 9",
			args: args{
				Introduce: "Yêu màu tím, thích đạp xe dưới mưa",
			},
			want: "Giới thiệu bản thân: Yêu màu tím, thích đạp xe dưới mưa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintIntroduce(tt.args.Introduce); got != tt.want {
				t.Errorf("PrintIntroduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintManager(t *testing.T) {
	type args struct {
		Manager string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case 10",
			args: args{
				Manager: "Duy Minh",
			},
			want: "Sếp hiện tại: Duy Minh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintManager(tt.args.Manager); got != tt.want {
				t.Errorf("PrintManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintManagerEmail(t *testing.T) {
	type args struct {
		ManagerEmail string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case 11",
			args: args{
				ManagerEmail: "Duyminh@gmail.com",
			},
			want: "Email sếp: Duyminh@gmail.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintManagerEmail(tt.args.ManagerEmail); got != tt.want {
				t.Errorf("PrintManagerEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
