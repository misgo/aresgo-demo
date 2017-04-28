/*
	Aresgo框架的http路由及Web服务实例
	author: hyperion
	since:2017-2-8

*/
package main

import (
	action "Demo/action"

	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aresgo"
)

func server(ip string) {
	timeNow := time.Now()
	aresgo.StartTime = timeNow.Format("2006-01-02 15:04:05")

	//初始化路由
	router := aresgo.Routing()

	//路由地址
	router.Get("/", Index)
	router.Get("/404.html", NotFound)  //404错误页
	router.Get("/hello/:name", Hello)  //欢迎页
	router.Get("/detail/:id", Detail)  //详情页
	router.Get("/multi/*.html", Multi) //

	router.Register("/passport/", &action.UserAction{}, aresgo.ActionGet) //注册对象，注册后对象所有公共方法可以被调用

	router.NotFound = NotFound                 //404页面方法
	router.MethodNotAllowed = DisAllowedMethod //POST or GET or ...请求被拒绝时执行的方法，取决于路由方法的设置

	router.Listen(ip)

}

//页面不存在
func NotFound(ctx *aresgo.Context) {
	fmt.Fprint(ctx, "页面不存在!\n")
}

//请求方法被拒绝
func DisAllowedMethod(ctx *aresgo.Context) {
	htmlStr := fmt.Sprintf("当前访问不允许!<br />此页面不允许【%s】请求！", ctx.Method())
	fmt.Fprint(ctx, htmlStr)

}

// 首页
func Index(ctx *aresgo.Context) {
	fmt.Fprint(ctx, "欢迎，这里是首页!\n")
}

// 欢迎页
func Hello(ctx *aresgo.Context) {
	fmt.Fprintf(ctx, "hello, 欢迎【%s】光临!\n", ctx.UserValue("name"))
}

// 欢迎页
func Multi(ctx *aresgo.Context) {
	para := ctx.UserValue(".html")
	fmt.Printf("%v\r\n", para)
	if para != nil {
		fmt.Fprintf(ctx, "hello, 欢迎【%s】光临!这是页面2\n", para.(string))
	} else {
		NotFound(ctx)
	}

}

func Detail(ctx *aresgo.Context) {
	Id := ctx.UserValue("id")
	if Id != nil {
		IdStr := Id.(string)
		IdStr = strings.Replace(IdStr, ".html", "", -1)
		IdStr = strings.Replace(IdStr, ".htm", "", -1)
		IdStr = strings.Replace(IdStr, ".php", "", -1)
		IdStr = strings.Replace(IdStr, ".jsp", "", -1)
		IdStr = strings.Replace(IdStr, ".aspx", "", -1)
		objId, err := strconv.ParseInt(IdStr, 10, 64)
		if err == nil {
			goods := action.GetGoodsInfo(objId)
			html := fmt.Sprintf("商品id：%d;\r\n商品名称：%s;\r\n商品价格：%f;\r\n商品数量：%d;\r\n所属店铺名称：%s;\r\n所属店铺地址：%s;\r\n",
				goods.GoodsId, goods.GoodsName, goods.Price, goods.Amount, goods.Shop.ShopName, goods.Shop.ShopAddress)
			fmt.Fprintf(ctx, html)
		} else {
			NotFound(ctx)
		}

	} else {
		NotFound(ctx)
	}

}
