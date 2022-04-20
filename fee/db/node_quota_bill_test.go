package db

import (
	"context"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager/db"
)

func TestNodeQuotaBillDB_Insert(t *testing.T) {
	type fields struct {
		conn *db.DB
	}
	type args struct {
		ctx     context.Context
		newInfo *NodeQuotaBill
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
			name: "test 1",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &NodeQuotaBill{
					UserID:      1,
					UserName:    "testing",
					Username:    "testing",
					UserGroupID: 2,
					OperType:    1,
					OldSize:     1,
					NewSize:     2,
					OldEndTime:  time.Now(),
					NewEndTime:  time.Now(),
					Fee:         12.3,
					CreateTime:  time.Now(),
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "test 2",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &NodeQuotaBill{
					UserID:      2,
					UserName:    "testing",
					Username:    "testing",
					UserGroupID: 3,
					OperType:    2,
					OldSize:     2,
					NewSize:     3,
					OldEndTime:  time.Now(),
					NewEndTime:  time.Now(),
					Fee:         13.3,
					CreateTime:  time.Now(),
				},
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NodeQuotaBillDB{
				conn: tt.fields.conn,
			}
			got, err := this.Insert(tt.args.ctx, tt.args.newInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeQuotaBillDB.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeQuotaBillDB.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
