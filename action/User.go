/*
	此实例主要是通过自动注册struct而实现的自动路由。
	只要在服务器监听前将需要注册的struct通过方法router.Register注册即可
	注意：方法名必须是首字母大写其他小写，不能用驼峰方法
*/
package Action

import (
	. "Demo/model"
	"fmt"

	"github.com/aresgo"
)

var (
	userMod *User
)

type (
	UserAction struct{}
)

func (u *UserAction) Reg(ctx *aresgo.Context) {
	str := fmt.Sprintf("这里是注册页!当前时间：%s", aresgo.StartTime)
	fmt.Fprint(ctx, str)
}

func (u *UserAction) Login(ctx *aresgo.Context) {
	poststr := string(ctx.PostBody())
	fmt.Fprint(ctx, poststr+"这里是登录页!\r\n")
}

func (u *UserAction) Usercenter(ctx *aresgo.Context) {
	fmt.Fprint(ctx, "---------------这里是用户中心-------------\r\n")
	currentUser := userMod.GetUserDetail(1)
	fmt.Fprint(ctx, currentUser)
}
