package main

import (
	"synTest/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"synTest/defs"
)


//功能点： 1、 接收数据落到本地
//		2、 将落地信息返回服务端
//


func main() {
	cdb := helper.CNewSqlCon()
	var info defs.Info
	info.Username = "Li"
	info.Recive = 0
	info.Already = helper.Sql_CgetMax(cdb)


	for {
		if bs,err := json.Marshal(info);err == nil{
			req := bytes.NewBuffer([]byte(bs))
			body_type := "application/json;charset=utf-8"
			resp, _ := http.Post("http://127.0.0.1:8080/test", body_type, req)

			//读
			body, _ := ioutil.ReadAll(resp.Body)
			var intbody int
			intbody, _ = strconv.Atoi(string(body))
			fmt.Println(intbody)
			helper.Sql_insert(cdb,intbody)  // 数据落库
			info.Already = helper.Sql_CgetMax(cdb)
			info.Recive += 1
		}else {
			fmt.Println(err)
		}
		time.Sleep(time.Second)

	}

}
