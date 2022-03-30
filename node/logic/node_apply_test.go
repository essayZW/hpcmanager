package logic

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	nodedb "github.com/essayZW/hpcmanager/node/db"
)

var nodeApplyLogic *NodeApply

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}

	nodeApplyDB := nodedb.NewNodeApply(dbConn)
	nodeApplyLogic = NewNodeApply(nodeApplyDB)
}

func TestCreateApplyNode(t *testing.T) {
	tests := []struct {
		Name string

		User      *ApplyItemUserInfo
		Tutor     *ApplyItemUserInfo
		NodeInfo  *ApplyNodeInfo
		ProjectID int

		ExceptID int
		Error    bool
	}{
		{
			Name: "test success",
			User: &ApplyItemUserInfo{
				ID:       4,
				Name:     "单元测试机器人",
				Username: "Robot",
			},
			Tutor: &ApplyItemUserInfo{
				ID:       8,
				Name:     "单元测试导师",
				Username: "TutorRobot",
			},
			NodeInfo: &ApplyNodeInfo{
				NodeType:  "type1",
				NodeNum:   12,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Duration(3600) * time.Minute),
			},
			ProjectID: 21,
			ExceptID:  7,
			Error:     false,
		},
		{
			Name: "test success2",
			User: &ApplyItemUserInfo{
				ID:       5,
				Name:     "单元测试机器人",
				Username: "Robot",
			},
			Tutor: &ApplyItemUserInfo{
				ID:       9,
				Name:     "单元测试导师",
				Username: "TutorRobot",
			},
			NodeInfo: &ApplyNodeInfo{
				NodeType:  "type1",
				NodeNum:   14,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Duration(3600) * time.Minute),
			},
			ProjectID: 23,
			ExceptID:  8,
			Error:     false,
		},
		{
			Name: "test error user",
			User: &ApplyItemUserInfo{
				ID:       0,
				Name:     "单元测试机器人",
				Username: "Robot",
			},
			Tutor: &ApplyItemUserInfo{
				ID:       8,
				Name:     "单元测试导师",
				Username: "TutorRobot",
			},
			NodeInfo: &ApplyNodeInfo{
				NodeType:  "type1",
				NodeNum:   12,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Duration(3600) * time.Minute),
			},
			ProjectID: 21,
			ExceptID:  0,
			Error:     true,
		},
		{
			Name: "test error tutor",
			User: &ApplyItemUserInfo{
				ID:       3,
				Name:     "单元测试机器人",
				Username: "Robot",
			},
			Tutor: &ApplyItemUserInfo{
				ID:       0,
				Name:     "单元测试导师",
				Username: "TutorRobot",
			},
			NodeInfo: &ApplyNodeInfo{
				NodeType:  "type1",
				NodeNum:   12,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Duration(3600) * time.Minute),
			},
			ProjectID: 21,
			ExceptID:  0,
			Error:     true,
		},
		{
			Name: "test error time",
			User: &ApplyItemUserInfo{
				ID:       3,
				Name:     "单元测试机器人",
				Username: "Robot",
			},
			Tutor: &ApplyItemUserInfo{
				ID:       6,
				Name:     "单元测试导师",
				Username: "TutorRobot",
			},
			NodeInfo: &ApplyNodeInfo{
				NodeType:  "type1",
				NodeNum:   12,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Duration(-3600) * time.Minute),
			},
			ProjectID: 21,
			ExceptID:  0,
			Error:     true,
		},
		{
			Name: "test error node info",
			User: &ApplyItemUserInfo{
				ID:       3,
				Name:     "单元测试机器人",
				Username: "Robot",
			},
			Tutor: &ApplyItemUserInfo{
				ID:       6,
				Name:     "单元测试导师",
				Username: "TutorRobot",
			},
			NodeInfo: &ApplyNodeInfo{
				NodeType:  "type1",
				NodeNum:   -1,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Duration(-3600) * time.Minute),
			},
			ProjectID: 21,
			ExceptID:  0,
			Error:     true,
		},
		{
			Name: "test error node name",
			User: &ApplyItemUserInfo{
				ID:       3,
				Name:     "单元测试机器人",
				Username: "Robot",
			},
			Tutor: &ApplyItemUserInfo{
				ID:       6,
				Name:     "单元测试导师",
				Username: "TutorRobot",
			},
			NodeInfo: &ApplyNodeInfo{
				NodeNum:   3,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Duration(-3600) * time.Minute),
			},
			ProjectID: 21,
			ExceptID:  0,
			Error:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			id, err := nodeApplyLogic.CreateNodeApply(
				context.Background(),
				test.User,
				test.Tutor,
				test.NodeInfo,
				test.ProjectID,
			)
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
