package Utills

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() { //只能是init函数，init是初始化的一个接口
	file, err := ini.Load("Config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，检查文件路径", err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug") //如果Appmodle无值就取debug
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("eqwr3425")
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("wwy040609")
	DbName = file.Section("database").Key("DbName").MustString("NewGinBlog")
}
