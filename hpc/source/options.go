package source

// Options source配置的options
type Options struct {
	// CmdLocation 脚本文件的根目录
	CmdBaseDir string
}

// Option 选项
type Option func(*Options)

// WithCmdBaseDir 配置可选的脚本文件路径
func WithCmdBaseDir(dir string) Option {
	return func(o *Options) {
		o.CmdBaseDir = dir
	}
}
