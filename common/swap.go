package common

import "encoding/json"

// SwapTo 通过JsonTag进行结构体赋值，直接将data(request or category结构体)中的内容按照tag映射到mapping(request or category结构体)中
func SwapTo(data interface{}, mapping interface{}) (err error) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, mapping)
}
