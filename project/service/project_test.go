package service

import (
	"context"
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	gateway "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/logger"
	projectdb "github.com/essayZW/hpcmanager/project/db"
	"github.com/essayZW/hpcmanager/project/logic"
	projectpb "github.com/essayZW/hpcmanager/project/proto"
	"github.com/essayZW/hpcmanager/verify"
)

var projectService *ProjectService
var baseRequest *gateway.BaseRequest

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
	projectLogic := logic.NewProject(projectDB)
	baseRequest = &gateway.BaseRequest{
		RequestInfo: &gateway.RequestInfo{
			Id: "testing",
		},
		UserInfo: &gateway.UserInfo{
			UserId: 0,
			Levels: []int32{
				int32(verify.SuperAdmin),
				int32(verify.CommonAdmin),
				int32(verify.Tutor),
				int32(verify.Common),
				int32(verify.Guest),
			},
		},
	}
	projectService = &ProjectService{
		projectLogic: projectLogic,
	}
}

func TestGetByID(t *testing.T) {
	tests := []struct {
		Name string

		ID int

		Error      bool
		ExceptName string
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
			var req projectpb.GetProjectInfoByIDRequest
			req.BaseRequest = baseRequest
			req.Id = int32(test.ID)
			var resp projectpb.GetProjectInfoByIDResponse
			err := projectService.GetProjectInfoByID(context.Background(), &req, &resp)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v Except: %v", err, test.Error)
				}
				return
			}
			if test.Error {
				t.Errorf("Get: %v Except: %v", err, test.Error)
				return
			}
			if resp.Data.Name != test.ExceptName {
				t.Errorf("Get: %v Except: %v", resp.Data, test.ExceptName)
			}
		})
	}
}

func TestPaginationGetProjectInfos(t *testing.T) {
	tests := []struct {
		Name string

		PageIndex int
		PageSize  int
		UserID    int
		Levels    []int32

		ExceptCount int
		ExceptLen   int
		Error       bool
	}{
		{
			Name:      "test admin",
			PageIndex: 1,
			PageSize:  3,
			UserID:    1,
			Levels: []int32{
				int32(verify.CommonAdmin),
			},
			ExceptCount: 4,
			ExceptLen:   3,
			Error:       false,
		},
		{
			Name:      "test common",
			PageIndex: 1,
			PageSize:  3,
			UserID:    1,
			Levels: []int32{
				int32(verify.Common),
			},
			ExceptCount: 2,
			ExceptLen:   2,
			Error:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var req projectpb.PaginationGetProjectInfosRequest
			req.BaseRequest = baseRequest
			req.BaseRequest.UserInfo.UserId = int32(test.UserID)
			req.BaseRequest.UserInfo.Levels = test.Levels
			req.PageIndex = int32(test.PageIndex)
			req.PageSize = int32(test.PageSize)
			var resp projectpb.PaginationGetProjectInfosResponse
			err := projectService.PaginationGetProjectInfos(context.Background(), &req, &resp)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v Except: %v", err, test.Error)
				}
				return
			}
			if test.Error {
				t.Errorf("Get: %v Except: %v", err, test.Error)
				return
			}
			if test.ExceptCount != int(resp.Count) {
				t.Errorf("Get: %v, Except: %v", resp.Count, test.ExceptCount)
				return
			}
			if test.ExceptLen != len(resp.Infos) {
				t.Errorf("Get: %v, Except: %v", resp.Infos, test.ExceptLen)
			}
		})
	}
}
