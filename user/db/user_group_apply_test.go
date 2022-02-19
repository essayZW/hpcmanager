package db

import (
	"context"
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

var userGroupApplyDB *UserGroupApplyDB

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	userGroupApplyDB = NewUserGroupApply(dbConn)
}

func TestAdminLimitQueryApplyCount(t *testing.T) {
	count, err := userGroupApplyDB.AdminLimitQueryApplyCount(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	if count != 1 {
		t.Errorf("Get: %v, Except: 1", count)
	}
}

func TestTutorLimitQueryApplyCount(t *testing.T) {
	count, err := userGroupApplyDB.TutorLimitQueryApplyCount(context.Background(), 2)
	if err != nil {
		t.Error(err)
		return
	}
	if count != 2 {
		t.Errorf("Get: %v, Except: 2", count)
	}
}

func TestCommonLimitQueryApplyCount(t *testing.T) {
	count, err := userGroupApplyDB.CommonLimitQueryApplyCount(context.Background(), 20)
	if err != nil {
		t.Error(err)
		return
	}
	if count != 1 {
		t.Errorf("Get: %v, Except: 1", count)
	}
}
