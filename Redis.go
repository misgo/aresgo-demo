// Redis
package main

import (
	"fmt"

	"github.com/misgo/aresgo"
)

func redis(confPath string) {
	redisPath := fmt.Sprintf("%s/redis.json", confPath)
	aresgo.RedisConfigPath = redisPath
	r := aresgo.R("dev")
	r.Select(5)
	//	fmt.Printf("错误：%v\r\n", r.Ping())

	//############string类型操作##################
	//-----Get-单条------
	g1 := r.GetString("c")
	fmt.Printf("%s\r\n", g1)
	g2 := r.GetInt("a")
	fmt.Printf("%d\r\n", g2)
	g3 := r.GetBool("a")
	fmt.Printf("%t\r\n", g3)
	//-----Get-多条------
	//	g4 := r.GetStrList("a", "b", "c", "d", "e")
	//	fmt.Printf("%s\r\n", g4)
	//-----Set-单条------
	//	s1 := r.Set("a", 2258874)
	//	s1 := r.Set("abcd", 123456, 60)  //加入过期时间（s）
	//	fmt.Printf("执行结果：%v\r\n", s1)
	//-----Set-多条------
	//	var setVals map[string]interface{} = make(map[string]interface{}, 0)
	//	setVals["v1"] = 1234567
	//	setVals["v2"] = "abcdefg"
	//	setVals["v3"] = false //布尔型保存后：true:1;false:0
	//	s2 := r.SetValues(setVals)
	//	fmt.Printf("共执行条数：%d\r\n", s2)
	//-----Delete-单条------
	//	d1 := r.Del("a")
	//	fmt.Printf("执行结果：%t\r\n", d1)

	//###############hash类型操作#################
	//-----Get 单条----
	//		hg1 := r.GetString("h1", "v1") //其他操作间string操作
	//		fmt.Printf("hash val:%s\r\n", hg1)
	//-----Get hash list(多个key多个hash key,获取Hash map)-----
	//	keys := make(map[string][]string)
	//	keys["h1"] = []string{}
	//	keys["h2"] = []string{"v1", "v2", "v3", "v4"}
	//	keys["h3"] = []string{"v2"}
	//	hg, _ := r.GetHashList(keys)
	//	fmt.Printf("%s\r\n", hg)
	//	fmt.Printf("%s\r\n", hg["h2"]["v3"])
	//-----Set-单条----
	//	hs1 := r.HSet("h", "v1", "aaaaa")
	//	fmt.Printf("执行结果：%t\r\n", hs1)
	//-----Set-多条----
	//	var setVals map[string]interface{} = make(map[string]interface{}, 0)
	//	setVals["v1"] = 100001
	//	setVals["v2"] = "123456"
	//	setVals["v3"] = "2017-01-01 13:45:12"
	//	setVals["v4"] = false
	//	hs2 := r.SetValues(setVals, "h")
	//	fmt.Printf("共执行条数：%d\r\n", hs2)
	//-----Delete(多个hashKey)-----
	//	hd1 := r.Del("h", "v1", "v2")
	//	fmt.Printf("执行结果：%t\r\n", hd1)

	//#####################其他操作#########################
	//设置键的失效时间
	t1 := r.SetTimeout("a", 60)
	fmt.Printf("执行结果：%v\r\n", t1)

	//获取键的失效时间
	//	t2 := r.GetTimeout("a")
	//	fmt.Printf("剩余失效时间（s）：%v\r\n", t2)

	//执行redis命令
	//	resSet, _ := r.Do("set", "a", 123456)
	//	rs, _ := r.String(resSet, nil)
	//	fmt.Printf("执行结果：%s\r\n", rs)
	//	vals, _ := r.Do("mget", "a", "b", "c", "d")
	//	resGet := vals.([]interface{})
	//	vg, _ := r.Int(resGet[0], nil)
	//	fmt.Printf("结果：%#s\r\n", resGet)
}
