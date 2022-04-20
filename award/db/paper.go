package db

import "github.com/essayZW/hpcmanager/db"

type PaperAwardDB struct {
	conn *db.DB
}

func NewPaperAward(conn *db.DB) *PaperAwardDB {
	return &PaperAwardDB{
		conn: conn,
	}
}
