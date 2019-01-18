package config

import (
	"io"
)

//Config web config
type Config struct {
	HTTPServer HTTPServerInfo
	Log        LogInfo
	Profile    ProfileInfo
	DataBase   DataBaseInfo
	Redis      RedisInfo
}

//HTTPServerInfo config for [httpserver]
type HTTPServerInfo struct {
	IP   string
	Port int
}

// LogInfo config for [log]
type LogInfo struct {
	Name    string    `default:"../log/going"` //日志文件路径前缀，文件名为 going.2018-05-25.log
	Level   string    `default:"debug"`        //默认debug 配置文件应写 level="debug" or "info" "error" ...
	MaxSize int64     `default:"1073741824"`   //默认1G 单个日志文件最大size，单位 B K M G 配置文件应写 maxSize="2G"
	MaxNum  int       `default:"10"`           //默认最多保存10个日志文件
	Mode    int       `default:"1"`            //默认1  0：going特色请求块日志模式 1：传统日志级别模式
	Writer  io.Writer //日志写接口,不允许配置，通过GetConfig自动生成
}

//ProfileInfo config for [profile]
type ProfileInfo struct {
	CPU      string //cpu性能分析输出profile文件路径 分析工具：go tool pprof exefile cpuprofile 或者：go-torch --binaryinput=cpuprofile --binaryname=exefile
	Mem      string //内存性能分析输出profile文件路径 go tool pprof -alloc_space exefile memprofile 或者：go-torch -inuse_space exefile memprofile
	Stack    string //pmg模型导致goroutine会在不同线程上浮动，无法用传统pstack/strace分析问题，可以通过配置这里，打印所有goroutine调用栈
	Duration int    `default:"10"` //性能分析时间，默认10分钟
}

//DataBaseInfo config for [database]
type DataBaseInfo struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Charset      string
	Host         string
	Port         uint16
	SQLLog       bool
	MaxIdelConns int
	MaxOpenConns int
	URL          string
}

//RedisInfo config for [redis]
type RedisInfo struct {
	IP   string
	Port uint16
}
