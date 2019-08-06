package main

import (
    "fmt"
    "net/smtp"
    "strings"
        "encoding/json"
    "io/ioutil"
   // "fmt"
  //  "os"
    "sync"
)



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


func main() {
    auth := smtp.PlainAuth("", "zukeqiang@163.com", "z1k6q3", "smtp.163.com")

    to := []string{"Keqiang.Zu@luxshare-ict.com"}
    nickname := "test"
    user := "zukeqiang@163.com"
    //user := "software"
    subject := "test mail"
    content_type := "Content-Type: text/plain; charset=UTF-8"
    body := "This is the email body."
    msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
        "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    err := smtp.SendMail("smtp.163.com:25", auth, user, to, msg)
    if err != nil {
        fmt.Printf("send mail error: %v", err)
    }
}
