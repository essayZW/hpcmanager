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
