package db

import (
	"time"

	"github.com/essayZW/hpcmanager/db"
	"gopkg.in/guregu/null.v4"
)

// PaperApply 论文奖励申请
type PaperApply struct {
	ID                       int         `db:"id"`
	CreaterID                int         `db:"creater_id"`
	CreaterUsername          string      `db:"creater_username"`
	CreaterName              string      `db:"creater_name"`
	CreateTime               time.Time   `db:"create_time"`
	UserGroupID              int         `db:"user_group_id"`
	TutorID                  int         `db:"tutor_id"`
	TutorUsername            string      `db:"tutor_username"`
	TutorName                string      `db:"tutor_name"`
	PaperTitle               string      `db:"paper_title"`
	PaperCategory            string      `db:"paper_category"`
	PaperPartition           string      `db:"paper_partition"`
	PaperFirstPageImageName  string      `db:"paper_firstpage_img"`
	PaperThanksPageImageName string      `db:"paper_thankspage_img"`
	RemarkMessage            string      `db:"remark_message"`
	CheckStatus              int8        `db:"check_status"`
	CheckerID                null.Int    `db:"checker_id"`
	CheckerName              null.String `db:"checker_name"`
	CheckerUsername          null.String `db:"checker_username"`
	CheckMoney               float64     `db:"check_money"`
	CheckMessage             null.String `db:"check_message"`
	CheckTime                null.Time   `db:"check_time"`
	ExtraAttributes          *db.JSON    `db:"extraAttributes"`
}

// TechnologyApply 科技奖励申请记录消息映射
type TechnologyApply struct {
	ID                 int         `db:"id"`
	CreaterID          int         `db:"creater_id"`
	CreaterUsername    string      `db:"creater_username"`
	CreaterName        string      `db:"creater_name"`
	CreateTime         time.Time   `db:"create_time"`
	UserGroupID        int         `db:"user_group_id"`
	TutorID            int         `db:"tutor_id"`
	TutorUsername      string      `db:"tutor_username"`
	TutorName          string      `db:"tutor_name"`
	ProjectID          int         `db:"project_id"`
	ProjectName        string      `db:"project_name"`
	ProjectDescription null.String `db:"project_description"`
	PrizeLevel         string      `db:"prize_level"`
	PrizeImageName     string      `db:"prize_img"`
	RemarkMessage      string      `db:"remark_message"`
	CheckStatus        int8        `db:"check_status"`
	CheckerID          null.Int    `db:"checker_id"`
	CheckerName        null.String `db:"checker_name"`
	CheckerUsername    null.String `db:"checker_username"`
	CheckMessage       null.String `db:"check_message"`
	CheckTime          null.Time   `db:"check_time"`
	CheckMoney         float64     `db:"check_money"`
	ExtraAttributes    *db.JSON    `db:"extraAttributes"`
}
