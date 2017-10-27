package public

import (
	"flag"
	"fmt"
)

// GetEnvFileName 根据命令行传来的生产环境：prod或者测试环境：dev的值，来判断读相应的配置
func GetEnvFileName() string {
	// flag
	var envFileName string
	envName := GetEnvName()
	if envName == "dev" {
		envFileName = "./cdnboss-middle-conf/proxy.dev.json"
	} else {
		envFileName = "./cdnboss-middle-conf/proxy.prod.json"
	}
	fmt.Println("conf file: ", envFileName)
	return envFileName
}

// GetEnvName 获得环境变量： env or prod
func GetEnvName() string {
	envName := flag.String("envName", "", " Please input your environment name: 'prod' or 'dev'")
	flag.Parse()
	if *envName == "" {
		*envName = "env"
	}
	fmt.Println("environment: ", *envName)
	return *envName
}
