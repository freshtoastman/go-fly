package tools

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var logrusObj *logrus.Logger

func Logger() *logrus.Logger {
	if logrusObj != nil {
		src, _ := setOutputFile()
		//設置输出
		logrusObj.Out = src
		return logrusObj
	}

	//实例化
	logger := logrus.New()
	src, _ := setOutputFile()
	//設置输出
	logger.Out = src
	//設置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//設置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrusObj = logger
	return logger
}
func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//寫入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}

//func LoggerToFile() gin.HandlerFunc {
//	logger := Logger()
//	return func(c *gin.Context) {
//		// 开始時间
//		startTime := time.Now()
//
//		// 處理請求
//		c.Next()
//
//		// 结束時间
//		endTime := time.Now()
//
//		// 执行時间
//		latencyTime := endTime.Sub(startTime)
//
//		// 請求方式
//		reqMethod := c.Request.Method
//
//		// 請求路由
//		reqUri := c.Request.RequestURI
//
//		// 狀態碼
//		statusCode := c.Writer.Status()
//
//		// 請求IP
//		clientIP := c.ClientIP()
//
//		//日志格式
//		logger.Infof("| %3d | %13v | %15s | %s | %s |",
//			statusCode,
//			latencyTime,
//			clientIP,
//			reqMethod,
//			reqUri,
//		)
//	}
//}
