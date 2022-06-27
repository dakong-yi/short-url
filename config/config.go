package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	// MongoCfg    *.Config `json:"mongo,omitempty" yaml:"mongo,omitempty" mapstructure:"mongo,omitempty"`
	ServerCfg *ServerConfig `json:"server,omitempty" yaml:"server,omitempty" mapstructure:"server,omitempty"`
	MysqlCfg  *Mysql        `json:"mysql,omitempty" yaml:"mysql,omitempty" mapstructure:"mysql,omitempty"`
	RedisCfg  *Redis        `json:"redis,omitempty" yaml:"redis,omitempty" mapstructure:"redis,omitempty"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Mysql struct {
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	User        string        `yaml:"user"`
	Password    string        `yaml:"password"`
	Database    string        `yaml:"database"`
	Charset     string        `yaml:"charset"`     // 要支持完整的UTF-8编码,需设置成: utf8mb4
	AutoMigrate bool          `yaml:"autoMigrate"` // 初始化时调用数据迁移
	ParseTime   bool          `yaml:"parseTime"`   // 解析time.Time类型
	TimeZone    string        `yaml:"timeZone"`    // 时区,若设置 Asia/Shanghai,需写成: Asia%2fShanghai
	SlowSql     time.Duration `yaml:"slowSql"`     // 慢SQL
	LogLevel    string        `yaml:"logLevel"`    // 日志记录级别
}

type ServerConfig struct {
	Host string `json:"host,omitempty" yaml:"host,omitempty" mapstructure:"host,omitempty"`
	Port int    `json:"port,omitempty" yaml:"port,omitempty" mapstructure:"port,omitempty"`
}

func New(defaultPath string) *viper.Viper {
	var file string
	pflag.StringVarP(&file, "config", "c", defaultPath, "")
	vp := viper.New()
	vp.SetConfigFile(file)
	return vp
}

var Cfg *Config

func Setup() {
	vp := New("./config/config.yaml")
	if err := vp.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config. %v\n", err)
		os.Exit(8)
	}
	err := vp.Unmarshal(&Cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse configuration. %v\n", err)
		os.Exit(1)
	}
}
