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
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
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
				ModifyTime:      null.NewTime(time.Now(), true),
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
				ModifyTime:      null.NewTime(time.Now(), true),
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
