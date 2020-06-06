package infoCollectionModel

type InfocolAll struct {
	Id         int    `json:"id" form:"id"`
	Ipdomain   string `json:"ipdomain" form:"ipdomain"`
	Port       string `json:"port" form:"port"`
	Urladdress string `json:"urladdress" form:"urladdress"`
	Dir        string `json:"dir" form:"dir"`
}