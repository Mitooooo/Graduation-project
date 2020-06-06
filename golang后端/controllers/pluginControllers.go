package controllers

import (
	"database/sql"
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetPocInfo(c *gin.Context) {
	var rows *sql.Rows
	rows = models.GetPocInfo()
	pocinfoall := make([]models.Plugin, 0)
	for rows.Next() {
		var pocinfo models.Plugin
		rows.Scan(&pocinfo.Pocname, &pocinfo.Pocclass, &pocinfo.PocfilePath)
		pocinfoall = append(pocinfoall, pocinfo)
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data":   pocinfoall,
	})
}
func DownloadInfo(c *gin.Context) {
	scriptname := c.PostForm("scriptname")
	fmt.Println(scriptname)
	fn := func(c rune) bool {
		return strings.ContainsRune("/", c)
	}
	sc := strings.FieldsFunc(scriptname, fn)
	fmt.Println(sc)
	if sc[0] != "verificationScript" {
		c.JSON(404, gin.H{
			"status": 404,
		})
	} else {
		scriptname = strings.Replace(scriptname, "..", "", -1)
		fmt.Println(scriptname)
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", sc[2])) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		filepath := "verificationScript/" + sc[1] + "/" + sc[2]
		c.File(filepath)
		c.JSON(200, gin.H{
			"status": 200,
		})
	}

}
func RemovePlugin(c *gin.Context) {
	pocname := c.PostForm("pocname")
	flag := models.DelPocInfo(pocname)
	if flag {
		fmt.Println("移除成功：", pocname)
		c.JSON(200, gin.H{
			"status": 200,
			"data":   "移除成功",
		})
	} else {
		fmt.Println("移除失败：", pocname)
		c.JSON(401, gin.H{
			"status": 401,
			"data":   "移除失败",
		})
	}
}

//GetPocNameAll
func GetPocNameAll(c *gin.Context) {
	var rows *sql.Rows
	rows = models.GetPocNameAll()
	pocnameall := make([]models.PluginNameAll, 0)
	for rows.Next() {
		var pocnameinfo models.PluginNameAll
		rows.Scan(&pocnameinfo.Pocname, &pocnameinfo.PocfilePath)
		pocnameall = append(pocnameall, pocnameinfo)
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data":   pocnameall,
	})
}
