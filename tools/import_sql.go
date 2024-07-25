package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ImportSqlTool struct {
	SqlPath                                    string
	Username, Password, Server, Port, Database string
}

func (this *ImportSqlTool) ImportSql() error {
	_, err := os.Stat(this.SqlPath)
	if os.IsNotExist(err) {
		log.Println("數據庫SQL文件不存在:", err)
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", this.Username, this.Password, this.Server, this.Port, this.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("數據庫連接失敗:", err)
		//panic("數據庫連接失敗!")
		return err
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(59 * time.Second)

	sqls, _ := ioutil.ReadFile(this.SqlPath)
	sqlArr := strings.Split(string(sqls), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err := db.Exec(sql).Error
		if err != nil {
			log.Println("數據庫导入失敗:" + err.Error())
			return err
		} else {
			log.Println(sql, "\t success!")
		}
	}
	return nil
}
