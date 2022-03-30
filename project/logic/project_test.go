package logic

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	projectdb "github.com/essayZW/hpcmanager/project/db"
)

var projectLogic *Project

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	projectDB := projectdb.NewProject(dbConn)
	projectLogic = NewProject(projectDB)
}

func TestCreate(t *testing.T) {
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

			ExceptID: 3,
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

			ExceptID: 4,
			Error:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			id, err := projectLogic.Create(
				context.Background(),
				test.UserID,
				test.UserName,
				test.Username,
				&projectdb.Project{
					Name:        test.PName,
					From:        test.From,
					Numbering:   test.Numbering,
					Expenses:    test.Expense,
					Description: test.Description,
				},
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
			if test.ExceptID != id {
				t.Errorf("Get: %v, Except: %v", id, test.ExceptID)
			}
		})
	}
}

func TestPaginationGet(t *testing.T) {
	tests := []struct {
		Name string

		PageIndex int
		PageSize  int

		Error       bool
		ExceptCount int
		ExceptLen   int
	}{
		{
			Name:        "test1",
			PageIndex:   1,
			PageSize:    2,
			Error:       false,
			ExceptCount: 4,
			ExceptLen:   2,
		},
		{
			Name:        "test2",
			PageIndex:   2,
			PageSize:    4,
			Error:       false,
			ExceptCount: 4,
			ExceptLen:   0,
		},
		{
			Name:        "test error",
			PageIndex:   0,
			PageSize:    2,
			Error:       true,
			ExceptCount: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := projectLogic.PaginationGet(
				context.Background(),
				test.PageIndex,
				test.PageSize,
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
			if len(res.Data) != test.ExceptLen {
				t.Errorf("Get: %v Except: %v", res, test.ExceptLen)
			}
			if res.Count != test.ExceptCount {
				t.Errorf("Get: %v Except: %v", res, test.ExceptCount)
			}
		})
	}
}
