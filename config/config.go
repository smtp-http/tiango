package config 

import (
    "encoding/json"
    "io/ioutil"
   // "fmt"
  //  "os"
    "sync"
)

///////////////////////////////////  configuration ////////////////////////////////////

/*config*/
type Configuration struct {
    HttpPort    string  `json:"http_port"`
    Version     string  `json:"version"`
    Database    string  `json:"database"`
    DataSourceName      string  `json:"data_source_name"`
}


var config *Configuration
var once_cfg sync.Once
 
func GetConfig() *Configuration {
    once_cfg.Do(func() {
        config = &Configuration{}
    })
    return config
}



//////////////////////////////////////  config loader ///////////////////////////////////

type ConfigLoader struct {

}


var loader *ConfigLoader
var once_loader sync.Once
 
func GetLoader() *ConfigLoader {
    once_loader.Do(func() {
        loader = &ConfigLoader{}
    })
    return loader
}



func (jst *ConfigLoader) Load(filename string, v interface{}) { 
//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回 
	data, err := ioutil.ReadFile(filename) 
	if err != nil { 
		return 
	} //读取的数据为json格式，需要进行解码 

	err = json.Unmarshal(data, v) 
	if err != nil { 
		return 
	} 
}