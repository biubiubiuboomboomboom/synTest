package main

import (
	"synTest/defs"
	"synTest/t/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	//"github.com/golang/glog"
)


/*
index ：   返回页面json
		页面效果根据前端来显示
 */

func nowindex(writer http.ResponseWriter, request *http.Request) {

	if time.Now().Unix() - defs.TestUser_record.UserLastTime >2 && defs.TestUser_record.UserCount>0 {
		defs.TestUser_record.UserCount -= 1
		defs.Mylooger.Printf("用户%v断开连接",defs.TestUser_record.Userinfo.Username)
	}

	if bs,err := json.Marshal(defs.TestUser_record);err == nil {
		req := bytes.NewBuffer(bs)
		fmt.Fprint(writer, req)
	}
}



/*

对用户接口 ， 给用户传数据  以及  接受用户信息

 */



func ServerTest(w http.ResponseWriter, r *http.Request)  {
	//连接数据库
	sdb := helper.NewSqlCon()
	s_num := helper.Sql_SgetMax(sdb)

	//读
	body,_ := ioutil.ReadAll(r.Body)
	var info defs.Info

	if err:= json.Unmarshal(body,&info);err == nil{
		var count = int(helper.Get_UserCount())-1
		defs.Mylooger.Printf("向用户%v发送一条数据,当前目标已接收 %d条数据,目前已接收到%d条数据",info.Username,info.Recive,info.Already)
		helper.ChangeUserRecode(info,time.Now().Unix(),count)
		//写
		if info.Already !=  s_num{
			println("需要同步")
			fmt.Fprint(w,helper.Sql_getOne(sdb,info.Already))
		}
	}else {
		fmt.Fprint(w,"Request请求不正确")
	}
}


func main()  {
	defs.Mylooger = helper.LogInfo()
	http.HandleFunc("/", nowindex)
	http.HandleFunc("/test",ServerTest)
	if err:=http.ListenAndServe("127.0.0.1:8080",nil) ; err!=nil {
		log.Fatalf("ListenAndServe : ",err)
	}
}


