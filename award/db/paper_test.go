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

func TestPaperAwardDB_Insert(t *testing.T) {
	type fields struct {
		conn *db.DB
	}
	type args struct {
		ctx     context.Context
		newInfo *PaperApply
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
				newInfo: &PaperApply{
					CreaterID:                1,
					CreaterUsername:          "testing",
					CreaterName:              "testing",
					CreateTime:               time.Now(),
					UserGroupID:              2,
					TutorID:                  3,
					TutorUsername:            "testing_tutor",
					TutorName:                "testing_tutor",
					PaperTitle:               "测试论文题目",
					PaperCategory:            "测试的论文分类",
					PaperPartition:           "测试的论文分区",
					PaperFirstPageImageName:  "xss.jpg",
					PaperThanksPageImageName: "thanks.jpg",
					RemarkMessage:            "测试的备注",
				},
			},
			wantErr: false,
			want:    1,
		},
		{
			name: "test 2",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				newInfo: &PaperApply{
					CreaterID:                2,
					CreaterUsername:          "testing",
					CreaterName:              "testing",
					CreateTime:               time.Now(),
					UserGroupID:              3,
					TutorID:                  4,
					TutorUsername:            "testing_tutor",
					TutorName:                "testing_tutor",
					PaperTitle:               "测试论文题目",
					PaperCategory:            "测试的论文分类",
					PaperPartition:           "测试的论文分区",
					PaperFirstPageImageName:  "xss.jpg",
					PaperThanksPageImageName: "thanks.jpg",
				},
			},
			wantErr: false,
			want:    2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &PaperAwardDB{
				conn: tt.fields.conn,
			}
			got, err := this.Insert(tt.args.ctx, tt.args.newInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaperAwardDB.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PaperAwardDB.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
