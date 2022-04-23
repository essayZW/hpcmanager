package db

import (
	"context"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager/db"
)

func TestTechnologyAwardApplyDB_Insert(t *testing.T) {
	type fields struct {
		conn *db.DB
	}
	type args struct {
		ctx     context.Context
		newInfo *TechnologyApply
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
				newInfo: &TechnologyApply{
					CreaterID:       1,
					CreaterUsername: "testing username",
					CreaterName:     "testing name",
					UserGroupID:     2,
					TutorID:         1,
					TutorUsername:   "testing username",
					TutorName:       "testing name",
					ProjectID:       1,
					ProjectName:     "testing project",
					PrizeLevel:      "level 1",
					PrizeImageName:  "test image",
					CreateTime:      time.Now(),
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &TechnologyAwardApplyDB{
				conn: tt.fields.conn,
			}
			got, err := this.Insert(tt.args.ctx, tt.args.newInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TechnologyAwardApplyDB.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TechnologyAwardApplyDB.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
