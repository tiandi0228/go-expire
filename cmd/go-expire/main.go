/**********************************************
** @Des:
** @Author: hong cha
** @Date:   2022/2/19
***********************************************/
package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hongcha/go-expire/internal/base/conf"
	"hongcha/go-expire/internal/base/db"
	"hongcha/go-expire/internal/base/logger"
	"hongcha/go-expire/internal/base/router"
	"os"
)

const (
	banner            = "expire v%v build %v\n"
	defaultConfigPath = "/etc/go-expire.yaml"
)

var (
	version   = "dev"
	buildTime = "unknown"

	printVersion = flag.Bool("v", false, "Print version and exit")
	configPath   = flag.String("c", defaultConfigPath, "config file path")
)

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(banner, version, buildTime))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *printVersion {
		flag.Usage()
		os.Exit(0)
	}

	confInit()
	logInit()
	dbInit()
	webInit()
}

func logInit() {
	logConf := conf.GlobalConfig.Logger
	logger.InitLogger("hongcha", logConf.Level, logConf.Path, logConf.MaxAge, logConf.RotationTime, logConf.StdOut)
}

func confInit() {
	conf.InitConfig(*configPath)
}

func dbInit() {
	repoConf := conf.GlobalConfig.Repo
	db.InitDB(repoConf.Connection)
	redisConf := conf.GlobalConfig.Redis
	db.InitRedis(redisConf.Connection, redisConf.Username, redisConf.Password)
}

func webInit() {
	baseConf := conf.GlobalConfig.Base
	router.CornInit()
	router.InitWebRouter(baseConf.WebPort)
}
