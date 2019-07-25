package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/smtp-http/tiango/config"
    "github.com/smtp-http/tiango/datastorage"
    "net/http"
)

type GinServer struct {
    Proxy   *datastorage.StorageProxy
}

func (s *GinServer)StartHttpServer() {

    // start database

    s.Proxy = datastorage.CreateStorageProxy("sqlite3", config.GetConfig().DbName)
    if s.Proxy == nil {
        fmt.Printf("orm failed to initialized\n")
        return
    }

    if err := s.Proxy.SyncTable(new(datastorage.ProductInformationTable)); err != nil {
        fmt.Printf("Fail to sync database ProductInformationTable: %v\n", err)
        return
    }

    if err := s.Proxy.SyncTable(new(datastorage.DpSizeTable)); err != nil {
        fmt.Printf("Fail to sync database DpSizeTable: %v\n", err)
        return
    }

    
    if err := s.Proxy.SyncTable(new(datastorage.ParamMaterialInputGuidanceTable)); err != nil {
        fmt.Printf("Fail to sync database ParamMaterialInputGuidanceTable: %v\n", err)
        return
    }

    if err := s.Proxy.SyncTable(new(datastorage.ParamSendMaterialTable)); err != nil {
        fmt.Printf("Fail to sync database ParamSendMaterialTable: %v\n", err)
        return
    }



    gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
    router := gin.Default()    //获得路由实例

    //添加中间件
    router.Use(Middleware)
    //注册接口
    
    router.POST("/api/" + config.GetConfig().Version +"/productinformation", s.Productinformation)

    // /api/v1/dpsize
    router.POST("/api/" + config.GetConfig().Version +"/dpsize", s.Dpsize)

    router.POST("/api/" + config.GetConfig().Version +"/domsize", s.Domsize)

    //   /api/v1/param-material-input-guidance

    router.POST("/api/" + config.GetConfig().Version +"/param-material-input-guidance", s.ParamMaterialInputGuidance)

    // /api/v1/param-send-material

     router.POST("/api/" + config.GetConfig().Version +"/param-send-material", s.ParamSendMaterial)
   
    //监听端口
    http.ListenAndServe(":" + config.GetConfig().HttpPort, router)
}

func Middleware(c *gin.Context) {
    fmt.Println("this is a middleware!")
}

func (s *GinServer)Productinformation(c *gin.Context) {
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

func (s *GinServer)Dpsize(c *gin.Context) {


}

func (s *GinServer)Domsize(c *gin.Context) {

}

func (s *GinServer)ParamMaterialInputGuidance(c *gin.Context) {
    type JsonHolder struct {
        Id   int    `json:"id"`
        Name string `json:"name"`
    }
    holder := JsonHolder{Id: 1, Name: "my name"}
    //若返回json数据，可以直接使用gin封装好的JSON方法
    c.JSON(http.StatusOK, holder)
    return
}

func (s *GinServer)ParamSendMaterial(c *gin.Context) {

}
