package infoCollection

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func checkPort(portScanResult *[]string, ip net.IP, port int, wg *sync.WaitGroup) {
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	ch := make(chan bool)
	timeout := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		timeout <- true
	}()
	go func() {
		conn, err := net.DialTCP("tcp", nil, &tcpAddr)
		ch <- true
		if err == nil {
			fmt.Printf("ip: %v port: %v \n", ip, port)
			//infoCollectionModel.PortScanModel(fmt.Sprintf("%v",ip),fmt.Sprintf("%v",port))
			*portScanResult = append(*portScanResult, fmt.Sprintf("%v,", port))
			defer func() {
				if conn != nil {
					e := conn.Close()
					if e != nil {
						fmt.Println(e)
					}
				}
			}()
		}
	}()
	select {
	case <-timeout:
		wg.Done()
	case <-ch:
		wg.Done()
	}
}
func checkIp(ip string) bool {
	if net.ParseIP(ip) == nil {
		fmt.Println("非法ip地址")
		return false
	} else {
		return true
	}
}
func PortScan(ip string, portScanResult []string) string {

	startTime := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(65534)
	//ip := os.Args[1]
	if checkIp(ip) {
		for port := 1; port <= 65534; port++ {
			go checkPort(&portScanResult, net.ParseIP(ip), port, &wg)
		}
	}
	wg.Wait()
	endTime := time.Now()
	fmt.Printf("执行时间%v", endTime.Sub(startTime))
	//fmt.Println(portScanResult)
	//infoCollectionModel.PortScanModel(fmt.Sprintf("%v",ip),fmt.Sprintf("%v",portScanResult))
	fmt.Println(portScanResult)
	return fmt.Sprintf("%v", portScanResult)

}
