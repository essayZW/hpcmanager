package db

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	hpcdb "github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
}

func getDBConn() *db.DB {
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	return dbConn

}

func TestNodeDistributeDB_Insert(t *testing.T) {
	type fields struct {
		conn *hpcdb.DB
	}
	type args struct {
		ctx  context.Context
		info *NodeDistribute
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
			name: "insert success 1",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				info: &NodeDistribute{
					ApplyID:         1,
					CreateTime:      time.Now(),
					ExtraAttributes: nil,
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "insert success 2",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				info: &NodeDistribute{
					ApplyID:         2,
					CreateTime:      time.Now(),
					ExtraAttributes: nil,
				},
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ndb := &NodeDistributeDB{
				conn: tt.fields.conn,
			}
			got, err := ndb.Insert(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeDistributeDB.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeDistributeDB.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeDistributeDB_QueryByApplyID(t *testing.T) {
	type fields struct {
		conn *hpcdb.DB
	}
	type args struct {
		ctx     context.Context
		applyID int
	}
	conn := getDBConn()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NodeDistribute
		wantErr bool
	}{
		{
			name: "select success",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:     context.Background(),
				applyID: 1,
			},
			want: &NodeDistribute{
				ID:               1,
				ApplyID:          1,
				HandlerFlag:      0,
				DistributeBillID: 0,
				CreateTime: func() time.Time {
					location, _ := time.LoadLocation("Asia/Shanghai")
					t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-03-23 20:04:07", location)
					return t
				}(),
			},
		},
		{
			name: "select success",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:     context.Background(),
				applyID: 2,
			},
			want: &NodeDistribute{
				ID:               2,
				ApplyID:          2,
				HandlerFlag:      0,
				DistributeBillID: 0,
				CreateTime: func() time.Time {
					location, _ := time.LoadLocation("Asia/Shanghai")
					t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-03-23 20:04:07", location)
					return t
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ndb := &NodeDistributeDB{
				conn: tt.fields.conn,
			}
			got, err := ndb.QueryByApplyID(tt.args.ctx, tt.args.applyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeDistributeDB.QueryByApplyID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeDistributeDB.QueryByApplyID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeDistributeDB_QueryCountByApply(t *testing.T) {
	type fields struct {
		conn *hpcdb.DB
	}
	type args struct {
		ctx     context.Context
		applyID int
	}
	conn := getDBConn()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{

			name: "select 1",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:     context.Background(),
				applyID: 1,
			},
			want: 1,
		},
		{
			name: "select 2",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:     context.Background(),
				applyID: 2,
			},
			want: 1,
		},
		{
			name: "select 0",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:     context.Background(),
				applyID: 8,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ndb := &NodeDistributeDB{
				conn: tt.fields.conn,
			}
			got, err := ndb.QueryCountByApply(tt.args.ctx, tt.args.applyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeDistributeDB.QueryCountByApply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeDistributeDB.QueryCountByApply() = %v, want %v", got, tt.want)
			}
		})
	}
}
