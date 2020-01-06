package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	//config 文件
	Cfg *ini.File
	RunMode string
	JwtSecret string
	PageSize int
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
)

func init() {
	//错误变量
	var err error

	//加载config文件
	Cfg, err = ini.Load("conf/app.ini")

	//判断是否加载失败
	if err != nil {
		log.Fatalf("app.ini 文件加载失败 : %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

//加载基础配置
func LoadBase() {
	//链式调用
	//取出默认分区下的 RUN_MODE值, 如果该值不存在 默认返回debug
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//加载server配置
func LoadServer() {
	//返回server分区
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}

//加载app配置文件
func LoadApp() {
	//返回server分区
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
