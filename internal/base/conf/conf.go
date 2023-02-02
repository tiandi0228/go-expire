package conf

import (
	"github.com/spf13/viper"
	"time"
)

// GlobalConfig 全局配置索引
var GlobalConfig = &Conf{
	Base: Base{WebPort: "80", CallbackURL: "http://127.0.0.1:80/test-callback"},
}

// Conf 配置文件映射
type Conf struct {
	Base    Base
	Logger  Logger
	Repo    Repo
	Redis   Redis
	Message Message
}

// Base 基本配置，包括相关端口号地址等
type Base struct {
	WebPort     string // web 服务端口
	CallbackURL string
	UserURL     string
}

// Logger 日志配置
type Logger struct {
	Level        string        // 日志打印级别
	Path         string        // 日志存放路径
	MaxAge       time.Duration // 最大存放时间
	RotationTime time.Duration // 日志分割时间
	StdOut       bool          // 控制台输出
}

// Repo 数据仓库配置
type Repo struct {
	Connection      string // Data Source Name
	ConnMaxLifeTime int    // 连接池中每个连接的最大生存时间，单位秒。
	MaxOpenConn     int    // 连接池中允许同时打开的最大连接数
	MaxIdleConn     int    // 连接池中允许存在的最大空闲连接数
}

// Redis配置
type Redis struct {
	Connection string // Data Source Name
	Username   string
	Password   string
}

// Message 消息配置
type Message struct {
	Connection string // Data Source Name
	Icon       string
	Title      string
}

// InitConfig 读取配置文件，读取配置文件异常直接panic提示 path: 文件路径
func InitConfig(path string) {
	configVip := viper.New()
	configVip.SetConfigFile(path)

	// 读取配置
	if err := configVip.ReadInConfig(); err != nil {
		panic(err)
	}

	// 配置映射到结构体
	GlobalConfig = &Conf{}
	if err := configVip.Unmarshal(GlobalConfig); err != nil {
		panic(err)
	}

	// 设置日志时间 配置文件中填的是以天和小时为单位的，所以要进行换算
	GlobalConfig.Logger.MaxAge *= 24 * time.Hour
	GlobalConfig.Logger.RotationTime *= time.Hour
}
