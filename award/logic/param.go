package logic

// UserInfoParam 用户信息参数
type UserInfoParam struct {
	ID       int
	Username string
	Name     string
}

// PaperInfoParam 论文信息参数
type PaperInfoParam struct {
	Title               string
	Category            string
	Partition           string
	FirstPageImageName  string
	ThanksPageImageName string
	RemarkMessage       string
}

// TechnologyParam 科技奖励信息参数
type TechnologyParam struct {
	Level         string
	ImageName     string
	RemarkMessage string
}

// ProjectInfoParam 项目信息参数
type ProjectInfoParam struct {
	ID          int
	Name        string
	Description string
}
