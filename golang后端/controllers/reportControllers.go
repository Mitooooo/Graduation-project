package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"os"
)

func ReportDownload(c *gin.Context) {
	//filpath := "templates/index.tmpl"
	//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "1.html"))
	//c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//c.File(filpath)
	GenerateReport("http://localhost:81/cms/")
}

type AllInfo struct {
	AssetsName    string
	AssetsAddress string
	Ipdomain      string
	Port          string
	Urladdress    string
	Dir           string
	Buginfo       string
}

func GenerateReport(address string) {
	infocol := GetRepostWithBaseInfo(address)
	buginfo := GetReportBugInfo(address)
	AllInfo := AllInfo{
		AssetsName:    infocol[0].AssetsName,
		AssetsAddress: infocol[0].AssetsAddress,
		Port:          infocol[0].Port,
		Urladdress:    infocol[0].Urladdress,
		Dir:           infocol[0].Dir,
		Buginfo:       buginfo,
	}
	//const text = `
	//<!DOCTYPE html>
	//<html>
	//<head>
	//	<title></title>
	//</head>
	//<body align="center">
	//<div> 检测报告</div>
	//<div >
	//	<div>
	//		<p>信息概况</p>
	//AssetsName:
	//{{range .AssetsName}}
	//	{{.}}
	//{{end}}
	//AssetsAddress: {{.AssetsAddress}}
	//Port: {{.Port}}
	//Urladdress: {{.Urladdress}}
	//Dir: {{.Dir}}
	//	</div>
	//	<div>
	//		<p>漏洞概况</p>
	//	</div>
	//</div>
	//
	//</body>
	//</html>
	//`
	const text = `
<!DOCTYPE html>
<html>
<head>
	<title>信息展示平台</title>
	<script src="https://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
	
</head>
<body>
<div>
	<div>
		<strong>站点ip</strong>
		<br>
		<ip></ip>
	</div>
	<div>
		<strong>站点链接</strong>
		<br>
		<urls></urls>
	</div>
	<div>
		<strong>开放端口</strong>
		<br>
		<ports></ports>
	</div>
	<div>
		<strong>目录扫描</strong>
		<br>
		<dirs></dirs>
	</div>
</div>
    <table border="1px" id="tab">
      <tr>
        <td>ipdomain</td>
        <td>urladdress</td>
        <td>bugname</td>
		<td>bugpoc</td>
      </tr>
    </table>
</body>
</html>
<!-- <script src="info.js"></script> -->
<script type="text/javascript">
	var ip = "{{.AssetsAddress}}";
	var urls = "{{.Urladdress}}"
	var ports = "{{.Port}}";
	var dirs = "{{.Dir}}";
	var buginfo = "{{.Buginfo}}"` + `
	$(function(){
	   dirs = dirs.replace("[","");
	   dirs = dirs.replace("]","");
	   var reg = new RegExp(" ","g");
	   console.log(dirs.replace(reg,"<br>"));
	   $("dirs").append(dirs.replace(reg,"<br>"));
	   ports = ports.replace("[","");
	   ports = ports.replace("]","");
	   $("ports").append(ports);
	   urls = urls.replace("[","");
	   urls = urls.replace("]","");
	   var reg = new RegExp(" ","g");
	   console.log(urls.replace(reg,"<br>"));
	   $("urls").append(urls.replace(reg,"<br>"));
	   $("ip").append(ip);
        var s = "";
        for(var i = 0; i < buginfo.length; i++) {
          s = "<tr><td>" + buginfo[i].ipdomain + "</td><td>" + buginfo[i].urladdress + "</td><td>" +
            buginfo[i].bugname + "</td><td>" + buginfo[i].bugpoc + "</td></tr>";
          $("#tab").append(s);
        }
      });
</script>
`
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.OpenFile("templates/index.html", os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// 渲染并写入文件
	if err := tmpl.Execute(f, AllInfo); err != nil {
		log.Fatalln(err)
	}
}
func GetRepostWithBaseInfo(address string) []models.PrimaryInfoCol {
	var rows *sql.Rows

	rows = models.GetBaseInfoAll(address)
	fmt.Println(rows)
	PrimaryInfoColData := make([]models.PrimaryInfoCol, 0)
	for rows.Next() {
		var infocol models.PrimaryInfoCol
		rows.Scan(&infocol.AssetsName, &infocol.AssetsAddress, &infocol.Ipdomain, &infocol.Port, &infocol.Urladdress, &infocol.Dir)
		PrimaryInfoColData = append(PrimaryInfoColData, infocol)
	}
	//fmt.Println(PrimaryInfoColData)
	return PrimaryInfoColData

}
func GetReportBugInfo(address string) string {

	var rows *sql.Rows
	rows = models.GetBugInfo(address)
	BugInfoData := make([]models.BugInfo, 0)
	fmt.Println(BugInfoData)
	for rows.Next() {
		var BugInfo models.BugInfo
		rows.Scan(&BugInfo.Ipdomain, &BugInfo.Urladdress, &BugInfo.Bugname, &BugInfo.Bugpoc)
		BugInfoData = append(BugInfoData, BugInfo)
	}
	//fmt.Println(BugInfoData)
	str, _ := json.Marshal(BugInfoData)
	fmt.Println(string(str))
	return string(str)
}
