package cmd

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/freshtoastman/imaptool/common"
	"github.com/freshtoastman/imaptool/middleware"
	"github.com/freshtoastman/imaptool/router"
	"github.com/freshtoastman/imaptool/static"
	"github.com/freshtoastman/imaptool/tools"
	"github.com/freshtoastman/imaptool/ws"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/zh-five/xdaemon"
)

var (
	port   string
	daemon bool
)
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "启動http服務",
	Example: "go-fly server -c config/",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	serverCmd.PersistentFlags().StringVarP(&port, "port", "p", "8081", "监听端口號")
	serverCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "是否為守护進程模式")
}
func run() {
	if daemon == true {
		logFilePath := ""
		if dir, err := os.Getwd(); err == nil {
			logFilePath = dir + "/logs/"
		}
		_, err := os.Stat(logFilePath)
		if os.IsNotExist(err) {
			if err := os.MkdirAll(logFilePath, 0777); err != nil {
				log.Println(err.Error())
			}
		}
		d := xdaemon.NewDaemon(logFilePath + "go-fly.log")
		d.MaxCount = 10
		d.Run()
	}

	baseServer := "0.0.0.0:" + port
	log.Println("start server...\r\ngo：http://" + baseServer)
	tools.Logger().Println("start server...\r\ngo：http://" + baseServer)

	engine := gin.Default()
	if common.IsCompireTemplate {
		templ := template.Must(template.New("").ParseFS(static.TemplatesEmbed, "templates/*.html"))
		engine.SetHTMLTemplate(templ)
		engine.StaticFS("/assets", http.FS(static.JsEmbed))
	} else {
		engine.LoadHTMLGlob("static/templates/*")
		engine.Static("/assets", "./static")
	}

	engine.Static("/static", "./static")
	engine.Use(tools.Session("gofly"))
	engine.Use(middleware.CrossSite)
	//性能监控
	pprof.Register(engine)

	//紀錄日志
	engine.Use(middleware.NewMidLogger())
	router.InitViewRouter(engine)
	router.InitApiRouter(engine)
	//紀錄pid
	ioutil.WriteFile("gofly.sock", []byte(fmt.Sprintf("%d,%d", os.Getppid(), os.Getpid())), 0666)
	//限流类
	tools.NewLimitQueue()
	//清理
	ws.CleanVisitorExpire()
	//后端websocket
	go ws.WsServerBackend()

	engine.Run(baseServer)
}
