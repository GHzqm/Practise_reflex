package main

import (
	"Practise_reflex/model"
	"Practise_reflex/tool"
	"fmt"
	"io/ioutil"
)

// 解析文件
func parseFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	var conf model.Config
	err = tool.UnMarshal(data, &conf)
	if err != nil {
		return
	}
	fmt.Printf("反序列化成功  conf: %#v\n  port: %#v\n", conf, conf.ServerConf.Port)

}

func parseFile2(filename string)  {
	// 有一些假数据
	var conf model.Config
	conf.ServerConf.Ip="127.0.0.1"
	conf.ServerConf.Port=8000
	conf.MysqlConf.Port=9000
	err := tool.MarshalFile(filename,conf)
	if err != nil{
		return
	}
}

func main() {
	parseFile("./config.ini")
	//parseFile2("/Users/zqm/GolandProjects/src/Practise_reflex/my2.ini")
}
