package logic

import (
	"context"
	"testing"
	"time"

	hpcdb "github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/node/db"
)

func getNodeUsageTimeDB() *db.NodeUsageTimeDB {
	dbConn, err := hpcdb.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	return db.NewNodeUsageTime(dbConn)
}

func TestNodeUsageTime_AddRecord(t *testing.T) {
	type fields struct {
		nodeUsageTimeDB *db.NodeUsageTimeDB
	}
	type args struct {
		ctx  context.Context
		info *db.HpcUsageTime
	}
	nodeUsageTimeDB := getNodeUsageTimeDB()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test error1",
			fields: fields{
				nodeUsageTimeDB: nodeUsageTimeDB,
			},
			args: args{
				ctx:  context.Background(),
				info: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test error2",
			fields: fields{
				nodeUsageTimeDB: nodeUsageTimeDB,
			},
			args: args{
				ctx:  context.Background(),
				info: &db.HpcUsageTime{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test error3",
			fields: fields{
				nodeUsageTimeDB: nodeUsageTimeDB,
			},
			args: args{
				ctx: context.Background(),
				info: &db.HpcUsageTime{
					QueueName: "testing",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test error3",
			fields: fields{
				nodeUsageTimeDB: nodeUsageTimeDB,
			},
			args: args{
				ctx: context.Background(),
				info: &db.HpcUsageTime{
					QueueName: "testing",
					UserID:    1,
					WallTime:  -1,
					GWallTime: -2,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test error4",
			fields: fields{
				nodeUsageTimeDB: nodeUsageTimeDB,
			},
			args: args{
				ctx: context.Background(),
				info: &db.HpcUsageTime{
					QueueName: "testing",
					UserID:    1,
					WallTime:  1,
					GWallTime: 2,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test error4",
			fields: fields{
				nodeUsageTimeDB: nodeUsageTimeDB,
			},
			args: args{
				ctx: context.Background(),
				info: &db.HpcUsageTime{
					QueueName: "testing",
					UserID:    1,
					WallTime:  1,
					GWallTime: 2,
					StartTime: time.Now().Add(time.Duration(10) * time.Second),
					EndTime:   time.Now(),
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test error5",
			fields: fields{
				nodeUsageTimeDB: nodeUsageTimeDB,
			},
			args: args{
				ctx: context.Background(),
				info: &db.HpcUsageTime{
					QueueName: "testing",
					UserID:    1,
					WallTime:  1,
					GWallTime: 2,
					StartTime: time.Now(),
					EndTime:   time.Now(),
				},
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NodeUsageTime{
				nodeUsageTimeDB: tt.fields.nodeUsageTimeDB,
			}
			got, err := n.AddRecord(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeUsageTime.AddRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeUsageTime.AddRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
