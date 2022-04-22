package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/award/db"
	"gopkg.in/guregu/null.v4"
)

type Paper struct {
	paperAwardDB *db.PaperAwardDB
}

func (this *Paper) CreateNew(
	ctx context.Context,
	creater *UserInfoParam,
	tutor *UserInfoParam,
	groupID int,
	paper *PaperInfoParam,
) (int64, error) {
	if creater == nil {
		return 0, errors.New("creater can't be empty")
	}
	if tutor == nil {
		return 0, errors.New("tutor can't be empty")
	}
	if groupID <= 0 {
		return 0, errors.New("invalid groupID")
	}
	if paper == nil {
		return 0, errors.New("paper can't be empty")
	}
	return this.paperAwardDB.Insert(ctx, &db.PaperApply{
		CreaterID:                creater.ID,
		CreaterUsername:          creater.Username,
		CreaterName:              creater.Name,
		CreateTime:               time.Now(),
		UserGroupID:              groupID,
		TutorID:                  tutor.ID,
		TutorUsername:            tutor.Username,
		TutorName:                tutor.Name,
		PaperTitle:               paper.Title,
		PaperCategory:            paper.Category,
		PaperPartition:           paper.Partition,
		PaperFirstPageImageName:  paper.FirstPageImageName,
		PaperThanksPageImageName: paper.ThanksPageImageName,
		RemarkMessage:            paper.RemarkMessage,
	})
}

// PaginationGetPaperApplyResult 分页查询论文奖励申请的结果
type PaginationGetPaperApplyResult struct {
	Count int
	Data  []*db.PaperApply
}

// PaginationGetAll 分页查询所有范围内的记录
func (this *Paper) PaginationGetAll(ctx context.Context, pageIndex, pageSize int) (*PaginationGetPaperApplyResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.paperAwardDB.QueryAllCount(ctx)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.paperAwardDB.QueryAllWithLimit(ctx, limit, pageSize)
	return &PaginationGetPaperApplyResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetByCreaterID 分页查询某个用户创建的论文申请记录
func (this *Paper) PaginationGetByCreaterID(
	ctx context.Context,
	createrID int,
	pageIndex, pageSize int,
) (*PaginationGetPaperApplyResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.paperAwardDB.QueryCountByCreaterID(ctx, createrID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.paperAwardDB.QueryWithLimitByCreaterID(ctx, createrID, limit, pageSize)
	return &PaginationGetPaperApplyResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetByGroupID 分页查询某个用户组创建的论文奖励申请
func (this *Paper) PaginationGetByGroupID(
	ctx context.Context,
	groupID int,
	pageIndex, pageSize int,
) (*PaginationGetPaperApplyResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.paperAwardDB.QueryCountByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.paperAwardDB.QueryWithLimitByGroupID(ctx, groupID, limit, pageSize)
	return &PaginationGetPaperApplyResult{
		Count: count,
		Data:  data,
	}, nil
}

// CheckPaperApply 审核论文奖励申请
func (this *Paper) CheckPaperApply(
	ctx context.Context,
	id int,
	checkerInfo *UserInfoParam,
	money float64,
	message string,
) (bool, error) {
	if checkerInfo == nil {
		return false, errors.New("checker info can't be nil")
	}
	return this.paperAwardDB.UpdateCheckStatus(ctx, &db.PaperApply{
		ID:              id,
		CheckerID:       null.IntFrom(int64(checkerInfo.ID)),
		CheckerUsername: null.StringFrom(checkerInfo.Username),
		CheckerName:     null.StringFrom(checkerInfo.Name),
		CheckMoney:      money,
		CheckMessage:    null.StringFrom(message),
		CheckTime:       null.TimeFrom(time.Now()),
	})
}

// GetInfoByID 通过ID查询信息
func (this *Paper) GetInfoByID(ctx context.Context, id int) (*db.PaperApply, error) {
	return this.paperAwardDB.QueryByID(ctx, id)
}

func NewPaper(paperAwardDB *db.PaperAwardDB) *Paper {
	return &Paper{
		paperAwardDB: paperAwardDB,
	}
}
