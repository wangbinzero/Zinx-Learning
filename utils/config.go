package utils

import (
	"Zinx-Learning/ziface"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	TcpServer      ziface.IServer //当前全局server对象
	Host           string         //服务IP
	TcpPort        int            //监听端口
	Name           string         //服务名称
	Version        string         //服务版本
	MaxConn        int            //允许最大连接数
	MaxPackageSize uint32         //数据包最大值
}

//定义全局对象
var BaseConfig *Config

//初始化方法
func init() {

	//如果配置文件没有加载，则作为默认值
	BaseConfig = &Config{
		Host:           "0.0.0.0",
		TcpPort:        8999,
		Name:           "zinx",
		Version:        "v0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	//尝试从config.json 加载
	BaseConfig.Reload()
}

//尝试从配置文件加载参数
func (c *Config) Reload() {
	data, err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &BaseConfig)
	if err != nil {
		panic(err)
	}
}
