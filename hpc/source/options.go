package source

// Options source配置的options
type Options struct {
	// CmdLocation 脚本文件的根目录
	CmdBaseDir string

	// DevMode 是否是开发模式下
	DevMode bool
}

// Option 选项
type Option func(*Options)

// WithCmdBaseDir 配置可选的脚本文件路径
func WithCmdBaseDir(dir string) Option {
	return func(o *Options) {
		o.CmdBaseDir = dir
	}
}

// WithDevSource 配置是否使用dev模式下的source
func WithDevSource(dev bool) Option {
	return func(o *Options) {
		o.DevMode = dev
	}
}
