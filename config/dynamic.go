package config

// ValueChange 监听的动态变量改变时候的回调函数
type ValueChange func(newV interface{})

// DynamicConfig 动态配置接口
type DynamicConfig interface {
	// Registry 对指定路径的变量注册监听,若变量发生改变，则将最新的值赋值
	// 其中val应该是一个指针,handler是变量那个发生变化之后的回调函数
	Registry(path string, val interface{}, handler ValueChange) error
}
