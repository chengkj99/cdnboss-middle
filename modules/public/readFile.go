package public

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ReadFile 读取并解析proxy.xx.json文件内容
func ReadFile(filename string, data interface{}) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}
	if err := json.Unmarshal(bytes, data); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return err
	}
	return nil
}
