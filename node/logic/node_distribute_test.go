package logic

import (
	"context"
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	hpcdb "github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/node/db"
)

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
}

func getNodeDistributeDB() *db.NodeDistributeDB {
	dbConn, err := hpcdb.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	return db.NewNodeDistribute(dbConn)
}

func TestNodeDistribute_PaginationGet(t *testing.T) {
	type fields struct {
		nodeDistributeDB *db.NodeDistributeDB
	}
	type args struct {
		ctx       context.Context
		pageIndex int
		pageSize  int
	}

	nodeDistributeDB := getNodeDistributeDB()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PaginationQueryNodeDistribute
		wantLen int
		wantErr bool
	}{
		{
			name: "1, 2",
			fields: fields{
				nodeDistributeDB: nodeDistributeDB,
			},
			args: args{
				ctx:       context.Background(),
				pageIndex: 1,
				pageSize:  2,
			},
			want: &PaginationQueryNodeDistribute{
				Count: 2,
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "0, 2",
			fields: fields{
				nodeDistributeDB: nodeDistributeDB,
			},
			args: args{
				ctx:       context.Background(),
				pageIndex: 0,
				pageSize:  2,
			},
			want: &PaginationQueryNodeDistribute{
				Count: 2,
			},
			wantErr: true,
		},
		{
			name: "1, 0",
			fields: fields{
				nodeDistributeDB: nodeDistributeDB,
			},
			args: args{
				ctx:       context.Background(),
				pageIndex: 1,
				pageSize:  0,
			},
			want: &PaginationQueryNodeDistribute{
				Count: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodeDistribute := &NodeDistribute{
				nodeDistributeDB: tt.fields.nodeDistributeDB,
			}
			got, err := nodeDistribute.PaginationGet(tt.args.ctx, tt.args.pageIndex, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeDistribute.PaginationGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if tt.want.Count != got.Count {
				t.Errorf("NodeDistribute.PaginationGet() Count = %v, want %v", got.Count, tt.want.Count)
			}
			if tt.wantLen != len(got.Data) {
				t.Errorf("NodeDistribute.PaginationGet() Len = %v, want %v", len(got.Data), tt.wantLen)
			}
		})
	}
}
