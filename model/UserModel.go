// UserModel
package Model

import (
	"aresgo/text"
	"time"
)

type (
	User struct {
		ID         int32
		Name       string
		Mobile     string
		Age        int
		Pwd        string
		money      int
		CreateTime time.Time
	}
)

//根据用户id获取用户详情
func (user *User) GetUserDetail(uid int32) *User {
	if uid == 1 {
		return &User{
			ID:         1,
			Name:       "test1",
			Mobile:     "13912346789",
			Age:        20,
			Pwd:        Text.Md5("123456"),
			money:      100,
			CreateTime: time.Now(),
		}
	} else {
		return &User{
			ID:         2,
			Name:       "test2",
			Mobile:     "13887654321",
			Age:        30,
			Pwd:        Text.Md5("111111"),
			money:      56,
			CreateTime: time.Now(),
		}
	}

}
