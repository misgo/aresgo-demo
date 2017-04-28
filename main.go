// 此项目是有关aresgo的实例展示
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	ConfPath   string = ""
	DbConfPath string = ""
	AppPath    string
)

func main() {
	initapp()

	//	fmt.Println("------配置文件测试--------")
	//	configs(ConfPath)
	//	fmt.Println("------数据库测试--------")
	//	database(DbConfPath)
	//	fmt.Println("------Redis测试--------")
	//	redis(ConfPath)
	//	fmt.Println("------服务器路由测试--------")
	server("127.0.0.1:8010")

}

func initapp() {
	procPath, _ := filepath.Abs(os.Args[0])
	procPath = strings.Replace(procPath, "\\", "/", -1)
	filePath, _ := os.Getwd()
	filePath = strings.Replace(filePath, "\\", "/", -1)
	//	parentPath := Text.CutStr(filePath, 0, strings.LastIndex(filePath, "/")) //上一级文件路径
	ConfPath = fmt.Sprintf("%s/config", filePath)
	DbConfPath = fmt.Sprintf("%s/db.json", ConfPath)
	//	fmt.Printf("dir:%v\r\n", ConfPath)
	AppPath = procPath
	fmt.Printf("------------运行时信息---start-------------\r\n")
	fmt.Printf("CPU个数:%d\r\n", runtime.NumCPU())
	fmt.Printf("操作系统:%s\r\n", runtime.GOOS)
	fmt.Printf("当前程序所在目录:%s\r\n", filePath)
	fmt.Printf("当前程序:%s\r\n", procPath)
	fmt.Printf("运行时信息：%v\r\n", os.Environ())
	fmt.Printf("------------运行时信息----end-------------\r\n")

}
