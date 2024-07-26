package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/freshtoastman/imaptool/common"
	"github.com/freshtoastman/imaptool/models"
	"github.com/freshtoastman/imaptool/tools"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "安装导入數據",
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}

func install() {
	if ok, _ := tools.IsFileNotExist("./install.lock"); !ok {
		log.Println("請先刪除./install.lock")
		os.Exit(1)
	}
	sqlFile := "import.sql"
	isExit, _ := tools.IsFileExist(common.MysqlConf)
	dataExit, _ := tools.IsFileExist(sqlFile)
	if !isExit || !dataExit {
		log.Println("config/mysql.json 數據庫配置文件或者數據庫文件go-fly.sql不存在")
		os.Exit(1)
	}
	sqls, _ := ioutil.ReadFile(sqlFile)
	sqlArr := strings.Split(string(sqls), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err := models.Execute(sql)
		if err == nil {
			log.Println(sql, "\t success!")
		} else {
			log.Println(sql, err, "\t failed!", "數據庫导入失敗")
			os.Exit(1)
		}
	}
	installFile, _ := os.OpenFile("./install.lock", os.O_RDWR|os.O_CREATE, os.ModePerm)
	installFile.WriteString("gofly live chat")
}
