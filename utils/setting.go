package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

//全局配置
var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

// init 初始化配置
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("文件读取出错，请检查文件路径与文件格式")
	}
	fmt.Println(file)
	LoadServer(file)
	LoadDatabase(file)
}

// LoadServer 加载服务器配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
}

// LoadDatabase 加载数据库配置
func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3360")
	DbUser = file.Section("database").Key("DbUser").MustString("govueblog")
	DbPassword = file.Section("database").Key("DbPassword").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("go_vue_blog")
}
