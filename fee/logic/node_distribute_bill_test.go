package logic

import (
	"testing"
	"time"
)

func parseTime(t *testing.T, timeStr string) int64 {
	timeLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Error(err)
		return 0
	}

	time, err := time.ParseInLocation("2006-01-02", timeStr, timeLocation)
	if err != nil {
		t.Error(err)
		return 0
	}
	return time.Unix()
}

func TestNodeDistributeBill_calTimeDurationYear(t *testing.T) {
	type args struct {
		startTimeUnix int64
		endTimeUnix   int64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 1 year",
			args: args{
				startTimeUnix: parseTime(t, "2021-04-22"),
				endTimeUnix:   parseTime(t, "2022-04-22"),
			},
			want: 1.0,
		},
		{
			name: "test 0.5 year",
			args: args{
				startTimeUnix: parseTime(t, "2021-01-01"),
				endTimeUnix:   parseTime(t, "2021-07-01"),
			},
			want: 0.5,
		},
		{
			name: "test 1.25 year",
			args: args{
				startTimeUnix: parseTime(t, "2021-01-01"),
				endTimeUnix:   parseTime(t, "2022-04-01"),
			},
			want: 1.25,
		},
		{
			name: "test 1.25 year",
			args: args{
				startTimeUnix: parseTime(t, "2021-01-01"),
				endTimeUnix:   parseTime(t, "2022-04-30"),
			},
			want: 1.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ndbl := &NodeDistributeBill{}
			if got := ndbl.calTimeDurationYear(tt.args.startTimeUnix, tt.args.endTimeUnix); got != tt.want {
				t.Errorf("NodeDistributeBill.calTimeDurationYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
