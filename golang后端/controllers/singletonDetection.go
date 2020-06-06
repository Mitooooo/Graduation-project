package controllers

import (
	"bytes"
	"database/sql"
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strings"
)

func SingletonCheck(c *gin.Context) {
	pluginname := c.PostForm("pluginname")
	urladdress := c.PostForm("urladdress")
	pluginpath := c.PostForm("pluginpath")

	scriptclass := strings.Split(pluginpath, ".")[1]
	fmt.Println(scriptclass)
	var lanuage string
	//var arg string
	if scriptclass == "go" {
		lanuage = "go"
		arg := "run"
		var outInfo bytes.Buffer
		fmt.Println(lanuage)
		cmd := exec.Command(lanuage, arg, pluginpath, urladdress)
		cmd.Stdout = &outInfo
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		if models.InsertSingleBugInfo(pluginname, pluginpath, urladdress, outInfo.String()) {
			fmt.Println("单条执行结果插入成功")
		} else {
			fmt.Println("单条执行结果插入失败")
		}
	}
	if scriptclass == "py" {
		lanuage = "python3"
		var outInfo bytes.Buffer
		fmt.Println(lanuage)
		cmd := exec.Command(lanuage, pluginpath, urladdress)
		cmd.Stdout = &outInfo
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		if models.InsertSingleBugInfo(pluginname, pluginpath, urladdress, outInfo.String()) {
			fmt.Println("单条执行结果插入成功")
		} else {
			fmt.Println("单条执行结果插入失败")
		}
	}

}

func SelectSingkeAllInfo(c *gin.Context) {
	var rows *sql.Rows
	rows = models.SelectSingkeAllInfo()
	SelectSingkeAll := make([]models.SelectSingkeAll, 0)
	var SelectAllInfo models.SelectSingkeAll

	for rows.Next() {
		rows.Scan(&SelectAllInfo.Pluginname, &SelectAllInfo.Pluginpath, &SelectAllInfo.Urladdress, &SelectAllInfo.Returninfo)
		SelectSingkeAll = append(SelectSingkeAll, SelectAllInfo)
	}
	fmt.Println(SelectSingkeAll)
	c.JSON(200, gin.H{
		"status": 200,
		"data":   SelectSingkeAll,
	})
}
func SingletonCheckResult(c *gin.Context) {
	pluginname := c.PostForm("pluginname")
	urladdress := c.PostForm("urladdress")
	pluginpath := c.PostForm("pluginpath")
	var rows *sql.Rows
	rows = models.SelectSingkeInfo(pluginname, pluginpath, urladdress)
	SelectSingke := make([]models.SelectSingkeAll, 0)
	var SelectInfo models.SelectSingkeAll
	for rows.Next() {
		rows.Scan(&SelectInfo.Pluginname, &SelectInfo.Pluginpath, &SelectInfo.Urladdress, &SelectInfo.Returninfo)
		SelectSingke = append(SelectSingke, SelectInfo)
	}
	fmt.Println(SelectSingke)
	c.JSON(200, gin.H{
		"status": 200,
		"data":   SelectSingke,
	})
}
