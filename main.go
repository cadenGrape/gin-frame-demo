package main

import (
	"errors"
	"fmt"
	"gin-frame/db"
	"gin-frame/graceful"
	"gin-frame/logger"
	"gin-frame/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"local.com/gin-frame/config"
	"log"
	"net/http"
	"time"
)

var (
	conf = pflag.StringP("config", "c", "", "config filepath")
)

func main() {
	pflag.Parse()
	//初始化配置
	if err := config.Run(*conf); err!=nil {
		panic(err)
	}
	//初始化连接池
	btn := db.GetInstance().InitPool()
	if !btn {
		log.Println("init database pool failure...")
		panic(errors.New("init database pool failure"))
	}
	//初始化redis
	db.InitRedis()

	gin.SetMode(viper.GetString("mode"))
	g := gin.New()
	g = router.Load(g)
	//_ = g.Run(viper.GetString("addr"))

	go func() {
		if err := pingServer(); err != nil {
			fmt.Println("fail:健康检查失败", err)
		}
		fmt.Println("success:健康检查成功")
	}()
	logger.Info("启动http服务端口%s\n", viper.GetString("addr"))

	if err := graceful.ListenAndServe(viper.GetString("addr"), g); err != nil && err != http.ErrServerClosed {
		logger.Error("fail:http服务启动失败: %s\n", err)
	}
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		url := fmt.Sprintf("%s%s%s", "http://127.0.0.1", viper.GetString("addr"), viper.GetString("healthCheck"))
		fmt.Println(url)
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("健康检测404")
}

