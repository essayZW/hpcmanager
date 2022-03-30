package db

import (
	"context"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager/db"
)

func TestNodeUsageTimeDB_Insert(t *testing.T) {
	type fields struct {
		conn *db.DB
	}
	type args struct {
		ctx  context.Context
		info *HpcUsageTime
	}
	conn := getDBConn()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "success1",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				info: &HpcUsageTime{
					UserID:     1,
					QueueName:  "testing_queue",
					WallTime:   123,
					GWallTime:  234,
					StartTime:  time.Now(),
					EndTime:    time.Now(),
					CreateTime: time.Now(),
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "success2",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				info: &HpcUsageTime{
					UserID:     2,
					QueueName:  "testing_queue2",
					WallTime:   130,
					GWallTime:  239,
					StartTime:  time.Now(),
					EndTime:    time.Now(),
					CreateTime: time.Now(),
				},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "success2",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				info: &HpcUsageTime{
					UserID:    2,
					QueueName: "testing_queue2",
					WallTime:  130,
					GWallTime: 239,
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NodeUsageTimeDB{
				conn: tt.fields.conn,
			}
			got, err := n.Insert(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeUsageTimeDB.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got != tt.want {
				t.Errorf("NodeUsageTimeDB.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
