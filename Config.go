/*
	Aresgo框架的配置文件操作
	目前支持json文件，ini文件操作。
	注意：文件格式有误读取失败
	author: hyperion
	since:2017-2-10

*/
package main

import (
	"fmt"

	"github.com/misgo/aresgo"
	"github.com/misgo/aresgo/config"
)

var configer config.Configer = nil

func configs(confPath string) {
	jsonConf := fmt.Sprintf("%s/db.json", confPath)
	iniConf := fmt.Sprintf("%s/db.ini", confPath)
	jsonConfiger, _ := aresgo.LoadConfig("json", jsonConf)
	iniConfiger, _ := aresgo.LoadConfig("ini", iniConf)
	//json操作
	fmt.Println("\r\n---------json文件操作--------\r\n")
	ip := jsonConfiger.String("dev.master.ip")                         //取字符串值
	ip2 := jsonConfiger.DefaultString("dev.master.ip2", "192.168.0.1") //取字符串值,字符串不存在返回默认值
	obj, _ := jsonConfiger.GetVal("dev.master")                        //取到一段数据，注意json格式用GetVal,ini用GetSection
	fmt.Printf("%s\r\n%s\r\n%v\r\n", ip, ip2, obj)
	//ini操作
	fmt.Println("\r\n---------ini文件操作--------\r\n")
	db := iniConfiger.String("dev.slave->db")                        //取字符串值
	db2 := iniConfiger.DefaultString("dev.slave->db2", "not choose") //取字符串值,字符串不存在返回默认值
	obj2, _ := iniConfiger.GetSection("dev.slave")                   //取到一段数据，注意json格式用GetVal,ini用GetSection

	fmt.Printf("%s\r\n%s\r\n%v\r\n", db, db2, obj2)
	//	fmt.Printf("%v\r\n%s\r\n", jsonConfiger, iniConfiger)
}
