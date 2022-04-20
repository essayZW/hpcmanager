package logic

import "github.com/essayZW/hpcmanager/award/db"

type Paper struct {
	paperAwardDB *db.PaperAwardDB
}

func NewPaper(paperAwardDB *db.PaperAwardDB) *Paper {
	return &Paper{
		paperAwardDB: paperAwardDB,
	}
}
