package db

import (
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

var userPermissionDB *UserPermissionDB

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	userPermissionDB = NewUserPermission(dbConn)
}

func TestQueryUserPermissionLevel(t *testing.T) {
	tests := []struct {
		Name string

		ID     int
		Except []int
	}{
		{
			Name:   "test1",
			ID:     1,
			Except: []int{0, 1},
		},
		{
			Name:   "test2",
			ID:     2,
			Except: []int{0},
		},
		{
			Name:   "test3",
			ID:     3,
			Except: []int{},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := userPermissionDB.QueryUserPermissionLevel(test.ID)
			if err != nil {
				t.Error(err)
				return
			}
			logger.Debug(res)
			for index, p := range res {
				if int(p.Level) != test.Except[index] {
					t.Errorf("Except %v Get %v", test.Except, p.Level)
					break
				}
			}
		})
	}
}
