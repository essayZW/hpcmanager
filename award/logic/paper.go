package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/award/db"
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

func NewPaper(paperAwardDB *db.PaperAwardDB) *Paper {
	return &Paper{
		paperAwardDB: paperAwardDB,
	}
}
