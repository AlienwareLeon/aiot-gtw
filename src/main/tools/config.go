package tools

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database           *Database `json:"database"`
}
var GlobalConfigSetting = &Config{}
func SetUp() {
	filePtr, err := os.Open("src/main/sqlConfig.json") //config的文件目录
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]\n", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(GlobalConfigSetting)
	DatabaseSetting = GlobalConfigSetting.Database
	fmt.Println(DatabaseSetting)
}
type Database struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}
var DatabaseSetting = &Database{}