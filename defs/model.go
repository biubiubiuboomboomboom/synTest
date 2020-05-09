package defs

import "log"

//  Request
type Info struct {
	Username     string `json:"username"` // 用户名
	Recive       int    `json:"recive"`	//接收到的数据
	Already      int    `json:"already"` // 已有数据量
}



var Mylooger *log.Logger

type Record struct {
	Userinfo Info
	UserCount int
	UserLastTime int64
}

var TestUser_record = &Record{
	Userinfo: Info{},
	UserCount: 0,
}






