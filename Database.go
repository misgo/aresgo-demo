/*
	Aresgo框架的数据操作实例
	author: hyperion
	since:2017-2-7

	1.数据库表结构：
		CREATE TABLE `t_user` (
		   `Uid` int(10) unsigned NOT NULL AUTO_INCREMENT,
		   `Username` varchar(30) NOT NULL DEFAULT '',
		   `Email` varchar(50) NOT NULL DEFAULT '',
		   `Mobile` varchar(20) NOT NULL DEFAULT '',
		   `Password` varchar(32) NOT NULL,
		   `Nickname` varchar(30) NOT NULL DEFAULT '',
		   `Gender` tinyint(4) NOT NULL DEFAULT '1',
		   `Birth` date DEFAULT NULL,
		   `Createtime` int(11) NOT NULL DEFAULT '0',
		   PRIMARY KEY (`Uid`)
		 ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8


	2.struct对象的Tag标签属性（struct小米号``中的部分）说明：
		table:对应数据库表的名称，默认取第一个设置table标签的属性
		key:用来标识一些字段特性。值---pk:代表是主键，notfield:表示不是数据库字段在进行curd操作是不对其进行处理
		field:当前struct属性对应的数据库字段的名称，不设置则取struct属性名
		type:数据库字段类型，date:数据表date类型，datetime:数据表datetime类型，int:数据表tinyint,int,bigint
		auto:是否为数据库自动生成数据字段，如数据自增字段。1：是；0：否。如果是1，则在添加数据或修改数据是则不进行处理
*/
package main

import (
	"fmt"
	"time"

	"github.com/aresgo"
	"github.com/aresgo/text"
)

type (
	//用户对象
	UserInfo struct {
		Uid        int64  `table:"t_user" key:"pk" auto:"1" `
		UserName   string `field:"Username"`
		Email      string
		Mobile     string
		Pwd        string `field:"Password"`
		NickName   string `field:"Nickname"`
		Gender     int8
		Birth      time.Time `type:"date"`
		CreateTime time.Time `field:"Createtime" type:"int"`
		Group      GroupInfo `key:"notfield"`
	}
	//用户组对象
	GroupInfo struct {
		Gid  int32
		Name string
	}
)

func database(confpath string) {
	aresgo.DbConfigPath = confpath //设置数据库配置
	curd_select()                  //CURD操作-查询
	//	curd_update()                  //CURD操作-更新
	//	curd_insert()                  //CURD操作-添加
	//	curd_delete()                  //CURD操作-删除
	//	curd_model_find()              //CURD操作-根据struct获取数据
	//	curd_model_findList()          //CURD操作-根据struct获取struct的数据列表
	//	curd_model_add()               //CURD操作-通过struct赋值添加对象数据
	//	curd_model_update()            //CURD操作-通过struct赋值修改对象数据
	//	getSqlList()                   //获取列表
	//	getSqlRow()                    //根据SQL获取一行数据
	//	sqlExcute()                    //数据库修改操作

}

func curd_select() {
	db := aresgo.D("dev")
	res, _ := db.Field("Uid,UserName").Table("t_user").Select()
	//	res, _ := db.Field("Uid,UserName").Table("t_user").Where("Uid <? ", 3).Select()
	//	res, _ := db.Field("Uid,UserName").Table("t_user").Limit(0, 2).Select()
	//	res, _ := db.Field("Uid,UserName,Count(Uid) as num").Table("t_user").GroupBy("Gender").Having("Gender>0").Select()
	//	res, _ := db.Field("Uid,UserName,Nickname,Gender").Table("t_user").OrderBy("Birth desc", "Uid asc").Select()

	for _, row := range *res {
		fmt.Printf("%v\r\n", row)
	}

}

func curd_update() {
	fields := make(map[string]interface{})
	fields["Username"] = "administrator"
	fields["Password"] = "21232f297a57a5a743894a0e4a801fc3"
	fields["Createtime"] = 1486463479
	fields["Gender"] = 2
	res, _ := aresgo.D("dev").Table("t_user").Where("Uid = ? ", 1).Update(fields)
	fmt.Printf("更新行数:%d\r\n", res)
}

func curd_insert() {
	fields := make(map[string]interface{})
	fields["UserName"] = "go1"
	fields["Password"] = "21232f297a57a5a743894a0e4a801fc3"
	fields["Mobile"] = "13988888888"
	fields["Createtime"] = time.Now().Unix()
	res, _ := aresgo.D("dev").Table("t_user").Insert(fields)
	fmt.Printf("用户Id:%d", res)
}

func curd_delete() {
	res, _ := aresgo.D("dev").Table("t_user").Where("Uid =?", 9).Delete() //通过条件删除
	//	res, _ := aresgo.D("dev").Table("t_user").SetPK("Uid").Delete(9)  //通过主键删除

	fmt.Printf("删除成功的行数:%d\r\n", res)
}

func curd_model_find() {
	var user UserInfo
	err := aresgo.D("dev").Where("Uid = ?", 3).Find(&user)
	//	err := aresgo.D("dev").FindByPK(&user, 1)

	if err == nil {
		fmt.Printf("user:%v\r\n", user)
	} else {
		fmt.Printf(err.Error())
	}
}

func curd_model_findList() {
	var users []UserInfo
	err := aresgo.D("dev").OrderBy("Uid").FindList(&users)
	if err == nil {
		for _, user := range users {
			fmt.Printf("%v\r\n", user)

		}
	} else {
		fmt.Printf("查询错误：%v\r\n", err.Error())
	}

}

func curd_model_add() {
	var birthday time.Time
	birthday, _ = time.Parse("2006-01-02", "1986-06-01")
	user := &UserInfo{
		UserName:   "go10",
		Mobile:     "15000001234",
		Pwd:        Text.Md5("111111"),
		NickName:   "go测试账号",
		Birth:      birthday,
		CreateTime: time.Now(),
	}

	userId, _ := aresgo.D("dev").Add(user)
	fmt.Printf("User_old:%v\r\n", user)
	fmt.Printf("Uid:%d\r\n", userId)
	user.Uid = userId
	fmt.Printf("User_new:%v\r\n", user)
}

func curd_model_update() {
	/*此处需要注意
	更新数据应将struct所有字段赋值，例如：struct有10个属性，但是只更新1个对应的数据库字段属性，如果只传一个值，其他值将按照初始值添加到数据库。
	这里应该注意数据覆盖问题
	*/
	user := &UserInfo{
		Uid:        18,
		UserName:   "go2",
		Mobile:     "13999999999",
		CreateTime: time.Now(),
		Gender:     1,
	}
	res, err := aresgo.D("dev").Save(user)
	if err == nil {
		fmt.Printf("更新影响的行数：%d\r\n", res)
	} else {
		fmt.Printf("%v\r\n", err.Error())
	}
}

func getSqlList() {
	sqlstr := "SELECT Uid,Username,Email,Gender FROM t_user WHERE Uid<10"
	res, _ := aresgo.D("dev").Query(sqlstr)
	for _, row := range *res {
		fmt.Printf("%v\r\n", row)
	}
}
func getSqlRow() {
	sqlstr := "SELECT Uid,Username,Email,Gender FROM t_user WHERE Uid<10"
	res, _ := aresgo.D("dev").GetRow(sqlstr)
	fmt.Printf("%v\r\n", res)
}

func sqlExcute() {
	//添加
	insertSql := fmt.Sprintf("insert t_user set Username='test1',Password='%s',Createtime=%d", Text.Md5("123456"), time.Now().Unix())
	res, _ := aresgo.D("dev").Execute(aresgo.DbInsert, insertSql)
	fmt.Printf("Uid:%d\r\n", res)

	//更新
	//	updateSql := fmt.Sprintf("update t_user set Username = 'test2',Password='%s' where Uid=7", Text.Md5("111111"))
	//	res, _ := aresgo.D("dev").Execute(aresgo.DbUpdate, updateSql)
	//	fmt.Printf("更新行数:%d\r\n", res)
}
