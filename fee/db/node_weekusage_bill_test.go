package db

import (
	"context"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager/db"
)

func TestNodeWeekUsageBillDB_Insert(t *testing.T) {
	type fields struct {
		conn *db.DB
	}
	type args struct {
		ctx     context.Context
		newInfo *NodeWeekUsageBill
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
			name: "test insert 1",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &NodeWeekUsageBill{
					UserID:      1,
					Username:    "testing",
					UserName:    "testingName",
					WallTime:    123,
					GWallTime:   345,
					Fee:         123.12,
					StartTime:   time.Now(),
					EndTime:     time.Now(),
					UserGroupID: 1,
					CreateTime:  time.Now(),
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NodeWeekUsageBillDB{
				conn: tt.fields.conn,
			}
			got, err := this.Insert(tt.args.ctx, tt.args.newInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeWeekUsageBillDB.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeWeekUsageBillDB.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
