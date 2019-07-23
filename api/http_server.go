package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/smtp-http/tiango/config"
    "github.com/smtp-http/tiango/datastorage"
    "net/http"
)

func StartHttpServer() {

    gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
    router := gin.Default()    //获得路由实例

    //添加中间件
    router.Use(Middleware)
    //注册接口
    
    router.POST("/api/" + config.GetConfig().Version +"/productinformation", Productinformation)

    // /api/v1/dpsize
    router.POST("/api/" + config.GetConfig().Version +"/dpsize", Dpsize)

    router.POST("/api/" + config.GetConfig().Version +"/domsize", Domsize)

    //   /api/v1/param-material-input-guidance

    router.POST("/api/" + config.GetConfig().Version +"/param-material-input-guidance", ParamMaterialInputGuidance)

    // /api/v1/param-send-material

     router.POST("/api/" + config.GetConfig().Version +"/param-send-material", ParamSendMaterial)
   
    //监听端口
    http.ListenAndServe(":" + config.GetConfig().HttpPort, router)
}

func Middleware(c *gin.Context) {
    fmt.Println("this is a middleware!")
}

func Productinformation(c *gin.Context) {
    var proInfo datastorage.ProductInformation
    err := c.BindJSON(&proInfo)
    if err != nil {
        
        c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
        return
    } else {
        fmt.Println(proInfo)
        type JsonHolder struct {
            Id   int    `json:"id"`
            Name string `json:"name"`
        }
        holder := JsonHolder{Id: 1, Name: "my name"}
    //若返回json数据，可以直接使用gin封装好的JSON方法
        c.JSON(http.StatusOK, holder)
        return
    }

    
}

func Dpsize(c *gin.Context) {


}

func Domsize(c *gin.Context) {

}

func ParamMaterialInputGuidance(c *gin.Context) {
    type JsonHolder struct {
        Id   int    `json:"id"`
        Name string `json:"name"`
    }
    holder := JsonHolder{Id: 1, Name: "my name"}
    //若返回json数据，可以直接使用gin封装好的JSON方法
    c.JSON(http.StatusOK, holder)
    return
}

func ParamSendMaterial(c *gin.Context) {

}
