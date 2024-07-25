package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
)

func GetChartStatistic(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")

	dayNumMap := make(map[string]string)
	result := models.CountVisitorsEveryDay(kefuName.(string))
	for _, item := range result {
		dayNumMap[item.Day] = tools.Int2Str(item.Num)
	}

	nowTime := time.Now()
	list := make([]map[string]string, 0)
	for i := 0; i > -46; i-- {
		getTime := nowTime.AddDate(0, 0, i)   //年，月，日   獲得一天前的時间
		resTime := getTime.Format("06-01-02") //獲得的時间的格式
		tmp := make(map[string]string)
		tmp["day"] = resTime
		tmp["num"] = dayNumMap[resTime]
		list = append(list, tmp)
	}

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": list,
	})

}
