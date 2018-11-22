package main

import (
	"flag"
	"runtime"
	"log"
	"github.com/larspensjo/config"
	"fmt"
	"path"
)

var(
	//https://studygolang.com/articles/686
	//支持命令行输入格式为-configfile=name, 默认为config.ini
	//配置文件一般获取到都是类型
	configFile = "../conf/config.ini" //flag.String("../conf/config.ini","../conf/config.ini","General configuration file")
	TOPIC = make(map[string]string)
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//_, filename, _, _ := runtime.Caller(1)
	//filepath := path.Join(path.Dir(filename), configFile)

	filepath := `./config.ini` //`D:\GOPATH\src\github.com\gorepomaker\conf\config.ini`
	cfg,err := config.ReadDefault(path.Join(filepath))   //读取配置文件，并返回其Config

	if err != nil {
		log.Fatalf("Fail to find %v,%v",filepath,err)
	}

	if	cfg.HasSection("mqtt") {   //判断配置文件中是否有section（一级标签）
		options,err := cfg.SectionOptions("mqtt")    //获取一级标签的所有子标签options（只有标签没有值）
		if err == nil {
			for _,v := range options{
				optionValue,err := cfg.String("mqtt",v)  //根据一级标签section和option获取对应的值
				if err == nil {
					TOPIC[v] =optionValue
				}
			}
		}
	}
	fmt.Println(TOPIC)
}
