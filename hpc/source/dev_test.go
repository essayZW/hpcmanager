package source

import (
	"testing"
)

func Test_hpcDev_parseKBToStr(t *testing.T) {
	type args struct {
		kb int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				kb: 1024,
			},
			want: "1M",
		},
		{
			name: "test 2",
			args: args{
				kb: 2048,
			},
			want: "2M",
		},
		{
			name: "test 3",
			args: args{
				kb: 512,
			},
			want: "512K",
		},
		{
			name: "test 4",
			args: args{
				kb: 524288,
			},
			want: "512M",
		},
		{
			name: "test 5",
			args: args{
				kb: 536870912,
			},
			want: "512G",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dev := &hpcDev{}
			if got := dev.parseKBToStr(tt.args.kb); got != tt.want {
				t.Errorf("hpcDev.parseKBToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
