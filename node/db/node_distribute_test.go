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
	"gopkg.in/guregu/null.v4"
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

func TestNodeDistributeDB_QueryLimit(t *testing.T) {
	type fields struct {
		conn *hpcdb.DB
	}
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	conn := getDBConn()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantLen int
		wantErr bool
	}{
		{
			name: "limit 0, 2",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:    context.Background(),
				limit:  0,
				offset: 2,
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "limit 0, 3",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:    context.Background(),
				limit:  0,
				offset: 3,
			},
			wantLen: 2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ndb := &NodeDistributeDB{
				conn: tt.fields.conn,
			}
			got, err := ndb.QueryLimit(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeDistributeDB.QueryLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantLen != len(got) {
				t.Errorf("NodeDistributeDB.QueryLimit() len = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}

func TestNodeDistributeDB_QueryCount(t *testing.T) {
	type fields struct {
		conn *hpcdb.DB
	}
	type args struct {
		ctx context.Context
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
			name: "test success",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
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
			got, err := ndb.QueryCount(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeDistributeDB.QueryCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeDistributeDB.QueryCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeDistributeDB_UpdateHandlerFlag(t *testing.T) {
	type fields struct {
		conn *hpcdb.DB
	}
	type args struct {
		ctx     context.Context
		newInfo *NodeDistribute
	}
	conn := getDBConn()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test success",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &NodeDistribute{
					ID:              1,
					HandlerUserID:   null.IntFrom(1),
					HandlerUserName: null.StringFrom("testing"),
					HandlerUsername: null.StringFrom("testing"),
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test success",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &NodeDistribute{
					ID:              0,
					HandlerUserID:   null.IntFrom(1),
					HandlerUserName: null.StringFrom("testing"),
					HandlerUsername: null.StringFrom("testing"),
				},
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ndb := &NodeDistributeDB{
				conn: tt.fields.conn,
			}
			got, err := ndb.UpdateHandlerFlag(tt.args.ctx, tt.args.newInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeDistributeDB.UpdateHandlerFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeDistributeDB.UpdateHandlerFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
