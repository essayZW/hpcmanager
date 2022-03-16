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

var projectDB *ProjectDB

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	projectDB = NewProject(dbConn)
}
func TestInsert(t *testing.T) {
	tests := []struct {
		Name string

		PName       string
		From        string
		Numbering   string
		Expense     string
		Description string

		UserID   int
		UserName string
		Username string

		Time time.Time

		ExceptID int64
		Error    bool
	}{
		{
			Name:      "insert 1",
			PName:     "单元测试1",
			From:      "来源testing",
			Numbering: "一个编号",
			Expense:   "好多经费",
			UserID:    1,
			UserName:  "测试用户1",
			Username:  "testing1",
			Time:      time.Now(),

			ExceptID: 1,
			Error:    false,
		},
		{
			Name:      "insert 2",
			PName:     "单元测试2",
			From:      "来源testing",
			Numbering: "一个编号",
			Expense:   "好多经费",
			UserID:    2,
			UserName:  "测试用户2",
			Username:  "testing2",
			Time:      time.Now(),

			ExceptID: 2,
			Error:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			id, err := projectDB.Insert(context.Background(), &Project{
				Name:            test.PName,
				From:            test.From,
				Numbering:       test.Numbering,
				Expenses:        test.Expense,
				Description:     test.Description,
				CreaterUserID:   test.UserID,
				CreaterUsername: test.Username,
				CreaterUserName: test.UserName,
				CreateTime:      test.Time,
				ModifyUserID:    test.UserID,
				ModifyUsername:  test.Username,
				ModifyUserName:  test.UserName,
				ModifyTime:      test.Time,
			})
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
			if test.ExceptID != id {
				t.Errorf("Get: %v, Except: %v", id, test.ExceptID)
			}
		})
	}
}

func TestQueryByID(t *testing.T) {
	tests := []struct {
		Name string

		ID int

		ExceptName string
		Error      bool
	}{
		{
			Name:       "query exists",
			ID:         1,
			ExceptName: "单元测试1",
			Error:      false,
		},
		{
			Name:       "query exists2",
			ID:         2,
			ExceptName: "单元测试2",
			Error:      false,
		},
		{
			Name:       "query not exists",
			ID:         10086,
			ExceptName: "",
			Error:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			info, err := projectDB.QueryByID(context.Background(), test.ID)
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
			if test.ExceptName != info.Name {
				t.Errorf("Get: %v, Except: %v", info, test.ExceptName)
			}
		})
	}
}

func TestLimitQueryByCreaterUserID(t *testing.T) {
	tests := []struct {
		Name string

		Limit  int
		Offset int
		UserID int

		Error       bool
		ExceptCount int
	}{
		{
			Name:        "limit 0, 2",
			Limit:       0,
			Offset:      2,
			Error:       false,
			UserID:      1,
			ExceptCount: 2,
		},
		{
			Name:        "limit 0, 5",
			Limit:       0,
			Offset:      5,
			UserID:      2,
			Error:       false,
			ExceptCount: 2,
		},
		{
			Name:        "limit 5, 1",
			Limit:       5,
			Offset:      1,
			Error:       false,
			ExceptCount: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			infos, err := projectDB.LimitQueryByCreaterUserID(context.Background(), test.Limit, test.Offset, test.UserID)
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
			if len(infos) != test.ExceptCount {
				t.Errorf("Get: %v, Except: %v", infos, test.ExceptCount)
			}
		})
	}
}

func TestQueryCount(t *testing.T) {
	count, err := projectDB.QueryCount(context.Background())
	if err != nil {
		t.Error(err)
	}
	if count != 4 {
		t.Error(count)
	}
}

func TestQueryCountByCreaterUserID(t *testing.T) {
	count, err := projectDB.QueryCountByCreaterUserID(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Error(count)
	}
}
func TestLimitQuery(t *testing.T) {
	tests := []struct {
		Name string

		Limit  int
		Offset int

		Error       bool
		ExceptCount int
	}{
		{
			Name:        "limit 0, 2",
			Limit:       0,
			Offset:      2,
			Error:       false,
			ExceptCount: 2,
		},
		{
			Name:        "limit 0, 5",
			Limit:       0,
			Offset:      5,
			Error:       false,
			ExceptCount: 4,
		},
		{
			Name:        "limit 5, 1",
			Limit:       5,
			Offset:      1,
			Error:       false,
			ExceptCount: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			infos, err := projectDB.LimitQuery(context.Background(), test.Limit, test.Offset)
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
			if len(infos) != test.ExceptCount {
				t.Errorf("Get: %v, Except: %v", infos, test.ExceptCount)
			}
		})
	}
}
