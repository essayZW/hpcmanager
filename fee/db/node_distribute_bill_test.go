package db

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
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

func TestNodeDistributeBillDB_Insert(t *testing.T) {
	type fields struct {
		conn *db.DB
	}
	type args struct {
		ctx     context.Context
		newInfo *NodeDistributeBill
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
				newInfo: &NodeDistributeBill{
					ApplyID:          1,
					NodeDistributeID: 2,
					Fee:              1234.43,
					UserID:           1,
					Username:         "testing",
					UserName:         "testingName",
					CreateTime:       time.Now(),
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ndb := &NodeDistributeBillDB{
				conn: tt.fields.conn,
			}
			got, err := ndb.Insert(tt.args.ctx, tt.args.newInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeDistributeBillDB.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeDistributeBillDB.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
