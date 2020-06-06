package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"gin/controllers/infoCollection"
	"gin/models"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

// 任务的执行顺序 先进行信息收集 然后再进行漏洞检测
type UrlStruct []struct {
	URL string `json:"url"`
}

func ExecTask(c *gin.Context) {
	cCp := c.Copy()
	assetsAddress := cCp.PostForm("assetsAddress")
	fmt.Println(assetsAddress)
	if models.UpdateExecStatus(assetsAddress, -1) {
		fmt.Println("更新状态成功")
		fmt.Println("正在执行。。。")
	} else {
		fmt.Println("更新状态失败")
	}
	realIp := judgeIpAddressOrurlAddress(assetsAddress)
	if realIp {
		fmt.Println("ip地址为：", assetsAddress)
		//进行端口扫描
		//port_scan := infoCollection.PortScan(assetsAddress)
		//if models.UpdateUrlAddressWithPrimaryInfoCol(assetsAddress, port_scan) {
		//	fmt.Println("端口扫描完成")
		//}

	} else {
		go func() {
			fmt.Println(fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05")))
			if models.InsertStartTimeInfoCol(fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05")), assetsAddress) {
				fmt.Println("开始时间更新成功")
			}
			var links []string
			cmd := exec.Command("controllers/crawlergo.exe", "-c", "controllers/chrome-win/chrome.exe", "-t", "20", "-o", "json", assetsAddress)
			fmt.Println(2)
			output, _ := cmd.CombinedOutput()
			str1 := strings.Split(string(output), "--[Mission Complete]--")[1]
			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(str1), &dat); err == nil {
				for _, v := range dat["req_list"].([]interface{}) {
					links = append(links, fmt.Sprintf("%v", v.(map[string]interface{})["url"]))
				}
			}
			if models.UpdateUrlAddressWithPrimaryInfoCol(fmt.Sprintf("%v", links), assetsAddress) {
				fmt.Println("更新数据成功：：", assetsAddress)
			}
			a1 := strings.Split(assetsAddress, "//")[1]
			a2 := strings.Split(a1, "/")[0]
			a3 := strings.Split(a2, ":")[0]
			ipAddr, err := net.ResolveIPAddr("ip", a3)
			if err != nil {
				fmt.Println(err)
			}
			ip := fmt.Sprintf("%v", ipAddr.IP)
			if models.UpdateIpdomainWithPrimaryInfoCol(ip, assetsAddress) {
				fmt.Println("ip开放情况收集完成")
			}
			//ip = "127.0.0.1"
			var ports []string
			portScanResult := infoCollection.PortScan(ip, ports)
			if models.UpdatePortWithPrimaryInfoCol(portScanResult, assetsAddress) {
				fmt.Println("端口开放情况收集完成")
			}
			time.Sleep(time.Duration(121) * time.Second)
			var dirs []string
			dirsReturn := infoCollection.DirScan(dirs, assetsAddress)
			if models.UpdateDirWithPrimaryInfoCol(fmt.Sprintf("%v", dirsReturn), assetsAddress) {
				fmt.Println("目录收集完成")
			}
			if models.InsertEndTimeInfoCol(fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05")), assetsAddress) {
				fmt.Println("结束时间更新成功")
			}
			fmt.Println(fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05")))

			for _, v := range links {
				ExecTarget(assetsAddress, v)
			}
			ExecTarget(assetsAddress, ip)
		}()
	}
}

type res struct {
	Ipdomain   string `json:"ipdomain"`
	Urladdress string `json:"urladdress"`
	Bugname    string `json:"bugname"`
	Bugpoc     string `json:"bugpoc"`
}

// 脚本名taskName 单个的资产地址target 单个的资产地址内部的目标地址 targetAddressAll
func goTask(taskName, targetAddressAll, target string) {
	var ResGo res
	go func() {
		var outInfo bytes.Buffer
		cmd := exec.Command("go", "run", taskName, targetAddressAll)
		cmd.Stdout = &outInfo
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(outInfo.String())
		jsonStr := outInfo.String()
		json.Unmarshal([]byte(jsonStr), &ResGo)
		if (ResGo.Urladdress == "" || ResGo.Bugname == "" || ResGo.Bugpoc == "") == false {
			reciveResult(target, ResGo.Urladdress, ResGo.Bugname, ResGo.Bugpoc)
		}
	}()

}

func pyTask(taskName, targetAddressAll, target, bugname string) {
	var Res res
	go func() {
		var outInfo bytes.Buffer
		cmd := exec.Command("python3", taskName, targetAddressAll)
		cmd.Stdout = &outInfo
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(taskName, targetAddressAll, target, bugname)
		//jsonStr := `{"ipdomain": "127.0.0.1", "urladdress": "127.0.0.1:3306", "bugname": "\u68c0\u6d4bmysql\u5f31\u53e3\u4ee4","bugpoc": "123456"}`
		jsonStr := outInfo.String()
		json.Unmarshal([]byte(jsonStr), &Res)
		fmt.Println(Res)
		if (Res.Urladdress == "" || bugname == "" || Res.Bugpoc == "") == false {
			reciveResult(target, Res.Urladdress, bugname, Res.Bugpoc)
		}
		fmt.Println("---------------------------", target, Res.Urladdress, bugname, Res.Bugpoc)
	}()

}

func javaTask(taskName, targetAddressAll, target string) {
	go func() {
		var outInfo bytes.Buffer
		cmd := exec.Command("java", "-jar", taskName, targetAddressAll)
		cmd.Stdout = &outInfo
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(outInfo.String())
		//fmt.Println(reflect.TypeOf(PocFilePathData[i])) 打印类型
	}()
}

func reciveResult(ipdomain, urladdress, bugname, bugpoc string) bool {

	fmt.Println("buginfo:::", ipdomain, urladdress, bugname, bugpoc)
	//将结果插入到数据库中
	flag := models.InsertBugInfo(ipdomain, urladdress, bugname, bugpoc)
	if flag {
		fmt.Println("bug信息插入成功")
	} else {
		fmt.Println("bug信息插入失败")
	}
	return true
}
func scriptClassDistribute(taskName, target, bugname string, address string, suffix []string, scriptClass []string) {
	fmt.Println("bugname start :: ", bugname)
	var scriptName string
	scriptName = bugname
	//脚本类别分发
	if judgeIpAddressOrurlAddress(address) { //执行host插件
		if scriptClass[1] == "host" {
			if suffix[1] == "go" {
				fmt.Println("goPoc host开始执行")
				goTask(taskName, address, target)
			}
			if suffix[1] == "py" {
				fmt.Println("pythonPoc host开始执行")
				fmt.Println("py bug name::", scriptName)
				pyTask(taskName, address, target, scriptName)
			}
			return
		}
	} else { //加载web的执行脚本
		if scriptClass[1] == "applicationOfWeb" {
			// 执行web插件
			if suffix[1] == "go" {
				fmt.Println("goPoc web插件 开始执行")
				goTask(taskName, address, target)
			}
			if suffix[1] == "py" {
				fmt.Println("pythonPoc web插件 开始执行")
				pyTask(taskName, address, target, bugname)
			}
			return
		}
	}
}
func judgeIpAddressOrurlAddress(target string) bool {
	ipv4 := target
	// ParseIP 这个方法 可以用来检查 ip 地址是否正确，如果不正确，该方法返回 nil
	address := net.ParseIP(ipv4)
	if address == nil {
		return false
	} else {
		fmt.Println("正确的ip地址", address.String())
		return true
	}
}
func ExecTarget(target, address string) {
	fmt.Println("target", address)
	var rows *sql.Rows
	rows = models.GetPocFilePath()
	PocFilePathData := make([]models.PocInfo, 0)
	fmt.Println(PocFilePathData)
	for rows.Next() {
		var pocInfo models.PocInfo
		rows.Scan(&pocInfo.PocName, &pocInfo.PocFilePath)
		//fmt.Println(pocInfo.PocFilePath)
		PocFilePathData = append(PocFilePathData, pocInfo)
	}
	fmt.Println(PocFilePathData)
	PocFilePathDataLen := len(PocFilePathData)
	for i := 0; i < PocFilePathDataLen; i++ {
		fmt.Println(PocFilePathData[i].PocFilePath)
		taskName := PocFilePathData[i].PocFilePath
		suffix := strings.Split(taskName, ".")
		scriptClass := strings.Split(suffix[0], "/") //脚本类别 go java 。。。
		fmt.Println(scriptClass)
		fmt.Println(scriptClass[2])
		fmt.Println("PocFilePathData[i].PocName:::", PocFilePathData[i].PocName)
		scriptClassDistribute(taskName, target, PocFilePathData[i].PocName, address, suffix, scriptClass)
		//此处的PocFilePathData[i].PocName没问题
	}
}

// 上面写的都是垃圾
func DoWebSomething(target string) {
	fmt.Println("目标为链接：地址为：", target)
}

// =================================================
//func urlscol(target string) []string{
//
//	return url
//}

func WriteWithIo(name, content string) {
	fileObj, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0644)
	//os.O_WRONLY | os.O_CREATE
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	if _, err := io.WriteString(fileObj, content); err == nil {
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.", content)
	}
}
