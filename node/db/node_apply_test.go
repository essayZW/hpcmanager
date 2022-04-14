package db

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	"gopkg.in/guregu/null.v4"
)

var nodeApplyDB *NodeApplyDB

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
	dbConn := getDBConn()
	nodeApplyDB = NewNodeApply(dbConn)
}

func TestInsert(t *testing.T) {
	tests := []struct {
		Name string

		Info *NodeApply

		ExceptID int
		Error    bool
	}{
		{
			Name: "test success1",
			Info: &NodeApply{
				CreateTime:      time.Now(),
				CreaterID:       1,
				CreaterUsername: "username",
				CreaterName:     "xs",
				ProjectID:       1,
				TutorID:         2,
				TutorName:       "tutor",
				TutorUsername:   "tutorUsername",
				ModifyTime:      null.TimeFrom(time.Now()),
				ModifyUserID:    1,
				ModifyName:      "modify",
				ModifyUsername:  "modifyUsername",
				NodeType:        "node type",
				NodeNum:         10,
				StartTime:       time.Now(),
				EndTime:         time.Now(),
			},
			ExceptID: 3,
			Error:    false,
		},
		{
			Name: "test success2",
			Info: &NodeApply{
				CreateTime:      time.Now(),
				CreaterID:       2,
				CreaterUsername: "username",
				CreaterName:     "xs",
				ProjectID:       2,
				TutorID:         3,
				TutorName:       "tutor",
				TutorUsername:   "tutorUsername",
				ModifyTime:      null.TimeFrom(time.Now()),
				ModifyUserID:    2,
				ModifyName:      "modify",
				ModifyUsername:  "modifyUsername",
				NodeType:        "node type",
				NodeNum:         17,
				StartTime:       time.Now(),
				EndTime:         time.Now(),
			},
			ExceptID: 4,
			Error:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			id, err := nodeApplyDB.Insert(context.Background(), test.Info)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v, Except: %v", err, test.Error)
				}
				return
			}
			if test.Error {
				t.Errorf("Get: %v, Except: %v", err, test.Error)
				return
			}
			if test.ExceptID != int(id) {
				t.Errorf("Get: %v, Except: %v", id, test.ExceptID)
			}
		})
	}
}

func TestNodeApplyDB_QueryByID(t *testing.T) {
	type fields struct {
		conn *db.DB
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
		wantID  int
		wantErr bool
	}{
		{
			name: "test success",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:     context.Background(),
				applyID: 1,
			},
			wantID:  1,
			wantErr: false,
		},
		{
			name: "test success",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:     context.Background(),
				applyID: 2,
			},
			wantID:  2,
			wantErr: false,
		},
		{
			name: "test success",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:     context.Background(),
				applyID: 100086,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &NodeApplyDB{
				conn: tt.fields.conn,
			}
			got, err := node.QueryByID(tt.args.ctx, tt.args.applyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeApplyDB.QueryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if tt.wantID != got.ID {
				t.Errorf("NodeApplyDB.QueryByID() ID = %v, want %v", got, tt.wantID)
			}
		})
	}
}

func TestNodeApplyDB_Update(t *testing.T) {
	type fields struct {
		conn *db.DB
	}
	type args struct {
		ctx     context.Context
		newInfo *NodeApply
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
			name: "test success1",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &NodeApply{
					ID:        1,
					CreaterID: 1,
					NodeType:  "update by testing",
					NodeNum:   11,
					StartTime: time.Now(),
					EndTime:   time.Now(),
				},
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "test success2",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &NodeApply{
					ID:        2,
					CreaterID: 2,
					NodeType:  "update by testing",
					NodeNum:   14,
					StartTime: time.Now(),
					EndTime:   time.Now(),
				},
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "test fail",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &NodeApply{
					ID:        0,
					NodeType:  "update by testing",
					NodeNum:   14,
					StartTime: time.Now(),
					EndTime:   time.Now(),
				},
			},
			wantErr: false,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &NodeApplyDB{
				conn: tt.fields.conn,
			}
			got, err := node.UpdateByCreaterID(tt.args.ctx, tt.args.newInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeApplyDB.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NodeApplyDB.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
