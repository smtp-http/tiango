package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/smtp-http/tiango/config"
    "github.com/smtp-http/tiango/datastorage"
    "net/http"
    "github.com/smtp-http/tiango/dataanalysis"
)

type GinServer struct {
    Proxy       *datastorage.StorageProxy
    Analysiser  *dataanalysis.DataAnalysiser
}

func (s *GinServer)StartHttpServer() {

    // start database

    s.Proxy = datastorage.CreateStorageProxy(config.GetConfig().Database, config.GetConfig().DataSourceName)
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

    if err := s.Proxy.SyncTable(new(datastorage.DomSizeTable)); err != nil {
        fmt.Printf("Fail to sync database DomSizeTable: %v\n", err)
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


    s.Analysiser = new(dataanalysis.DataAnalysiser)
    s.Analysiser.SetProxy(s.Proxy)


    gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
    router := gin.Default()    //获得路由实例

    //添加中间件
    // router.Use(Middleware)
    //注册接口
    
    router.POST("/api/" + config.GetConfig().Version +"/productinformation", s.Productinformation)

    // /api/v1/dpsize
    router.POST("/api/" + config.GetConfig().Version +"/dpsize_right", s.DpsizeRight)
    router.POST("/api/" + config.GetConfig().Version +"/dpsize_left", s.DpsizeLeft)

    router.POST("/api/" + config.GetConfig().Version +"/domsize_right", s.DomsizeRight)
    router.POST("/api/" + config.GetConfig().Version +"/domsize_left", s.DomsizeLeft)

    //   /api/v1/param-material-input-guidance

    router.POST("/api/" + config.GetConfig().Version +"/param-material-input-guidance", s.ParamMaterialInputGuidance)

    // /api/v1/param-send-material

     router.POST("/api/" + config.GetConfig().Version +"/param-send-material", s.ParamSendMaterial)


    // event
    router.POST("/api/" + config.GetConfig().Version +"/add_event", s.AddEvent)
    router.POST("/api/" + config.GetConfig().Version +"/del_event", s.DelEvent)
    router.POST("/api/" + config.GetConfig().Version +"/get_event_list", s.AddEvent)
   
    //监听端口
    http.ListenAndServe(":" + config.GetConfig().HttpPort, router)
}

func Middleware(c *gin.Context) {
    fmt.Println("this is a middleware!")
}

type JsonRes struct {
    ReqId   int32       `json:"req_id"`
    ResCode int32       `json:"rescode"`
    Result  string      `json:"result"`
}


/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type ProductInformationReq struct {
    ReqId   int32                          `json:"req_id"`
    Data    datastorage.ProductInformation  `json:"data"`
}



func (s *GinServer)Productinformation(c *gin.Context) {

    c.Header("Access-Control-Allow-Origin", "*")
    var proInfo ProductInformationReq
    err := c.BindJSON(&proInfo)
    if err != nil {
        fmt.Printf("==== %v\n",err)
        res := JsonRes{ReqId: proInfo.ReqId, ResCode: 1,Result:"bind json error"}
        c.JSON(200,res)
        return
    } else {
        var proInfoTable datastorage.ProductInformationTable
        proInfoTable.DomSupplier            =   proInfo.Data.DomSupplier
        proInfoTable.DpSupplier             =   proInfo.Data.DpSupplier 
        proInfoTable.ProductCn              =   proInfo.Data.ProductCn
        proInfoTable.LRstationDifference    =   proInfo.Data.LRstationDifference

        proInfoTable.A_B                    =   proInfo.Data.A_B
        proInfoTable.B_D                    =   proInfo.Data.B_D
        proInfoTable.E_F                    =   proInfo.Data.E_F
        proInfoTable.G_H                    =   proInfo.Data.G_H


        proInfoTable.Result                 =   proInfo.Data.Result
        proInfoTable.Angle                  =   proInfo.Data.Angle
        proInfoTable.SizeA                  =   proInfo.Data.SizeA
        proInfoTable.SizeB                  =   proInfo.Data.SizeB
        proInfoTable.SizeC                  =   proInfo.Data.SizeC
        proInfoTable.SizeD                  =   proInfo.Data.SizeD
        proInfoTable.SizeE                  =   proInfo.Data.SizeE
        proInfoTable.SizeF                  =   proInfo.Data.SizeF
        proInfoTable.SizeG                  =   proInfo.Data.SizeG
        proInfoTable.SizeH                  =   proInfo.Data.SizeH


        var res JsonRes
        

        e := s.Proxy.Insert(proInfoTable)
        if e != nil {
            fmt.Printf("ProInfo data insert error!\n")
            res = JsonRes{ReqId: proInfo.ReqId, ResCode: 2,Result:"Proinfo data insert err!"}
            return
        }

        res = JsonRes{ReqId: proInfo.ReqId, ResCode: 0,Result:""}
    //若返回json数据，可以直接使用gin封装好的JSON方法
        c.JSON(http.StatusOK, res)
        return
    }

    
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
type DpSizeReq struct {
    ReqId   int32               `json:"req_id"`
    Result  string              `json:"Result"`
    Data    datastorage.DpSize  `json:"data"`
}

func (s *GinServer)DpsizeLeft(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    var dpsize DpSizeReq
    err := c.BindJSON(&dpsize)
    if err != nil {
        fmt.Printf("==== %v\n",err)
        res := JsonRes{ReqId: dpsize.ReqId, ResCode: 1,Result:"dpsize bind json error"}
        c.JSON(200,res)
        return
    } else {

        var dpsizeTable datastorage.DpSizeTable

        dpsizeTable.Position = "left"


        dpsizeTable.Result                 =   dpsize.Result
        dpsizeTable.DomSupplier            =   dpsize.Data.DomSupplier
        dpsizeTable.DpSupplier             =   dpsize.Data.DpSupplier 
        dpsizeTable.ProductCn              =   dpsize.Data.ProductCn
        dpsizeTable.LRstationDifference    =   dpsize.Data.LRstationDifference

        dpsizeTable.Length                 =   dpsize.Data.Length
        dpsizeTable.Width                  =   dpsize.Data.Width
        dpsizeTable.LongSideAngle          =   dpsize.Data.LongSideAngle
        dpsizeTable.ShortSideAngle         =   dpsize.Data.ShortSideAngle
        dpsizeTable.Angle1                 =   dpsize.Data.Angle1
        dpsizeTable.Angle2                 =   dpsize.Data.Angle2
        dpsizeTable.Angle3                 =   dpsize.Data.Angle3
        dpsizeTable.Angle4                 =   dpsize.Data.Angle4
        dpsizeTable.DX                     =   dpsize.Data.DX
        dpsizeTable.DY                     =   dpsize.Data.DY
        dpsizeTable.DR                     =   dpsize.Data.DR
       

        var res JsonRes
        

        e := s.Proxy.Insert(dpsizeTable)
        if e != nil {
            fmt.Printf("DpSize data insert error!\n")
            res = JsonRes{ReqId: dpsize.ReqId, ResCode: 2,Result:"DpSize data insert err!"}
            return
        }

        res = JsonRes{ReqId: dpsize.ReqId, ResCode: 0,Result:""}
    //若返回json数据，可以直接使用gin封装好的JSON方法
        c.JSON(http.StatusOK, res)
        return
    }

}

func (s *GinServer)DpsizeRight(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    var dpsize DpSizeReq
    err := c.BindJSON(&dpsize)
    if err != nil {
        fmt.Printf("==== %v\n",err)
        res := JsonRes{ReqId: dpsize.ReqId, ResCode: 1,Result:"dpsize bind json error"}
        c.JSON(200,res)
        return
    } else {

        var dpsizeTable datastorage.DpSizeTable

        dpsizeTable.Position = "right"

        dpsizeTable.Result                 =   dpsize.Result
        dpsizeTable.DomSupplier            =   dpsize.Data.DomSupplier
        dpsizeTable.DpSupplier             =   dpsize.Data.DpSupplier 
        dpsizeTable.ProductCn              =   dpsize.Data.ProductCn
        dpsizeTable.LRstationDifference    =   dpsize.Data.LRstationDifference

        dpsizeTable.Length                 =   dpsize.Data.Length
        dpsizeTable.Width                  =   dpsize.Data.Width
        dpsizeTable.LongSideAngle          =   dpsize.Data.LongSideAngle
        dpsizeTable.ShortSideAngle         =   dpsize.Data.ShortSideAngle
        dpsizeTable.Angle1                 =   dpsize.Data.Angle1
        dpsizeTable.Angle2                 =   dpsize.Data.Angle2
        dpsizeTable.Angle3                 =   dpsize.Data.Angle3
        dpsizeTable.Angle4                 =   dpsize.Data.Angle4
        dpsizeTable.DX                     =   dpsize.Data.DX
        dpsizeTable.DY                     =   dpsize.Data.DY
        dpsizeTable.DR                     =   dpsize.Data.DR
       

        var res JsonRes
        

        e := s.Proxy.Insert(dpsizeTable)
        if e != nil {
            fmt.Printf("DpSize data insert error!\n")
            res = JsonRes{ReqId: dpsize.ReqId, ResCode: 2,Result:"DpSize data insert err!"}
            return
        }

        res = JsonRes{ReqId: dpsize.ReqId, ResCode: 0,Result:""}
    //若返回json数据，可以直接使用gin封装好的JSON方法
        c.JSON(http.StatusOK, res)
        return
    }

}


//////////////////////////////////////////////////////////////////////////////////////////////////////////////
type DomSizeReq struct {
    ReqId   int32               `json:"req_id"`
    Result  string              `json:"Result"`
    Data    datastorage.DomSize  `json:"data"`
}

func (s *GinServer)DomsizeLeft(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    var domsize DomSizeReq
    err := c.BindJSON(&domsize)
    if err != nil {
        fmt.Printf("==== %v\n",err)
        res := JsonRes{ReqId: domsize.ReqId, ResCode: 1,Result:"domsize bind json error"}
        c.JSON(200,res)
        return
    } else {

        var domsizeTable datastorage.DomSizeTable

        domsizeTable.Position = "left"

        domsizeTable.Result                 =   domsize.Result
        domsizeTable.DomSupplier            =   domsize.Data.DomSupplier
        domsizeTable.DpSupplier             =   domsize.Data.DpSupplier 
        domsizeTable.ProductCn              =   domsize.Data.ProductCn
        domsizeTable.LRstationDifference    =   domsize.Data.LRstationDifference

        domsizeTable.Length                 =   domsize.Data.Length
        domsizeTable.Width                  =   domsize.Data.Width
        domsizeTable.LongSideAngle          =   domsize.Data.LongSideAngle
        domsizeTable.ShortSideAngle         =   domsize.Data.ShortSideAngle
        domsizeTable.Angle1                 =   domsize.Data.Angle1
        domsizeTable.Angle2                 =   domsize.Data.Angle2
        domsizeTable.Angle3                 =   domsize.Data.Angle3
        domsizeTable.Angle4                 =   domsize.Data.Angle4
        domsizeTable.DX                     =   domsize.Data.DX
        domsizeTable.DY                     =   domsize.Data.DY
        domsizeTable.DR                     =   domsize.Data.DR
       

        var res JsonRes
        

        e := s.Proxy.Insert(domsizeTable)
        if e != nil {
            fmt.Printf("domsize data insert error!\n")
            res = JsonRes{ReqId: domsize.ReqId, ResCode: 2,Result:"domsize data insert err!"}
            return
        }

        res = JsonRes{ReqId: domsize.ReqId, ResCode: 0,Result:""}
    //若返回json数据，可以直接使用gin封装好的JSON方法
        c.JSON(http.StatusOK, res)
        return
    }
}

func (s *GinServer)DomsizeRight(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    var domsize DomSizeReq
    err := c.BindJSON(&domsize)
    if err != nil {
        fmt.Printf("==== %v\n",err)
        res := JsonRes{ReqId: domsize.ReqId, ResCode: 1,Result:"domsize bind json error"}
        c.JSON(200,res)
        return
    } else {

        var domsizeTable datastorage.DomSizeTable

        domsizeTable.Position = "right"

        domsizeTable.Result                 =   domsize.Result
        domsizeTable.DomSupplier            =   domsize.Data.DomSupplier
        domsizeTable.DpSupplier             =   domsize.Data.DpSupplier 
        domsizeTable.ProductCn              =   domsize.Data.ProductCn
        domsizeTable.LRstationDifference    =   domsize.Data.LRstationDifference

        domsizeTable.Length                 =   domsize.Data.Length
        domsizeTable.Width                  =   domsize.Data.Width
        domsizeTable.LongSideAngle          =   domsize.Data.LongSideAngle
        domsizeTable.ShortSideAngle         =   domsize.Data.ShortSideAngle
        domsizeTable.Angle1                 =   domsize.Data.Angle1
        domsizeTable.Angle2                 =   domsize.Data.Angle2
        domsizeTable.Angle3                 =   domsize.Data.Angle3
        domsizeTable.Angle4                 =   domsize.Data.Angle4
        domsizeTable.DX                     =   domsize.Data.DX
        domsizeTable.DY                     =   domsize.Data.DY
        domsizeTable.DR                     =   domsize.Data.DR
       

        var res JsonRes
        

        e := s.Proxy.Insert(domsizeTable)
        if e != nil {
            fmt.Printf("domsize data insert error!\n")
            res = JsonRes{ReqId: domsize.ReqId, ResCode: 2,Result:"domsize data insert err!"}
            return
        }

        res = JsonRes{ReqId: domsize.ReqId, ResCode: 0,Result:""}
    //若返回json数据，可以直接使用gin封装好的JSON方法
        c.JSON(http.StatusOK, res)
        return
    }
}


//=======================================================================================================================

type ParamMaterialInputReq struct {
    ReqId   int32               `json:"req_id"`
    Data    datastorage.ParamMaterialInputGuidance  `json:"data"`
}

func (s *GinServer)ParamMaterialInputGuidance(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    var paramMaterialInput ParamMaterialInputReq
    err := c.BindJSON(&paramMaterialInput)
    if err != nil {
        fmt.Printf("==== %v\n",err)
        res := JsonRes{ReqId: paramMaterialInput.ReqId, ResCode: 1,Result:"paramMaterialInput bind json error"}
        c.JSON(200,res)
        return
    } else {

        var paramMaterialInputGuidanceTable datastorage.ParamMaterialInputGuidanceTable
        paramMaterialInputGuidanceTable.PhotoDelay             =   paramMaterialInput.Data.PhotoDelay
        paramMaterialInputGuidanceTable.CompensationX1         =   paramMaterialInput.Data.CompensationX1 
        paramMaterialInputGuidanceTable.CompensationY1         =   paramMaterialInput.Data.CompensationY1
        paramMaterialInputGuidanceTable.CompensationR1         =   paramMaterialInput.Data.CompensationR1

        paramMaterialInputGuidanceTable.CompensationX2         =   paramMaterialInput.Data.CompensationX2
        paramMaterialInputGuidanceTable.CompensationY2         =   paramMaterialInput.Data.CompensationY2
        paramMaterialInputGuidanceTable.CompensationR2         =   paramMaterialInput.Data.CompensationR2

        paramMaterialInputGuidanceTable.CompensationX3         =   paramMaterialInput.Data.CompensationX3 
        paramMaterialInputGuidanceTable.CompensationY3         =   paramMaterialInput.Data.CompensationY3
        paramMaterialInputGuidanceTable.CompensationR3         =   paramMaterialInput.Data.CompensationR3

        paramMaterialInputGuidanceTable.CompensationX4         =   paramMaterialInput.Data.CompensationX4 
        paramMaterialInputGuidanceTable.CompensationY4         =   paramMaterialInput.Data.CompensationY4
        paramMaterialInputGuidanceTable.CompensationR4         =   paramMaterialInput.Data.CompensationR4

        paramMaterialInputGuidanceTable.MaterialInputReferenceX  =   paramMaterialInput.Data.MaterialInputReferenceX
        paramMaterialInputGuidanceTable.MaterialInputReferenceY  =   paramMaterialInput.Data.MaterialInputReferenceY
        paramMaterialInputGuidanceTable.MaterialInputReferenceZ  =   paramMaterialInput.Data.MaterialInputReferenceZ
        paramMaterialInputGuidanceTable.MaterialInputReferenceR  =   paramMaterialInput.Data.MaterialInputReferenceR

        paramMaterialInputGuidanceTable.FallingInitialSpeed      =   paramMaterialInput.Data.FallingInitialSpeed
        paramMaterialInputGuidanceTable.FallingAcceleration      =   paramMaterialInput.Data.FallingAcceleration
        paramMaterialInputGuidanceTable.FallingDeceleration      =   paramMaterialInput.Data.FallingDeceleration
        paramMaterialInputGuidanceTable.FallingSpeed             =   paramMaterialInput.Data.FallingSpeed

        paramMaterialInputGuidanceTable.PutOnTableInitialSpeed   =   paramMaterialInput.Data.PutOnTableInitialSpeed
        paramMaterialInputGuidanceTable.PutOnTableAcceleration   =   paramMaterialInput.Data.FallingAcceleration
        paramMaterialInputGuidanceTable.PutOnTableDeceleration   =   paramMaterialInput.Data.PutOnTableDeceleration
        paramMaterialInputGuidanceTable.PutOnTableSpeed          =   paramMaterialInput.Data.PutOnTableSpeed
       
        paramMaterialInputGuidanceTable.MaterialInputDelay       =   paramMaterialInput.Data.MaterialInputDelay

        var res JsonRes
        

        e := s.Proxy.Insert(paramMaterialInputGuidanceTable)
        if e != nil {
            fmt.Printf("paramMaterialInput data insert error!\n")
            res = JsonRes{ReqId: paramMaterialInput.ReqId, ResCode: 2,Result:"paramMaterialInput data insert err!"}
            return
        }

        res = JsonRes{ReqId: paramMaterialInput.ReqId, ResCode: 0,Result:""}
    //若返回json数据，可以直接使用gin封装好的JSON方法
        c.JSON(http.StatusOK, res)
        return
    }
}




////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ParamSendMaterialReq struct {
    ReqId   int32               `json:"req_id"`
    Data    datastorage.ParamSendMaterial  `json:"data"`
}


func (s *GinServer)ParamSendMaterial(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    var paramSendMaterial ParamSendMaterialReq
    err := c.BindJSON(&paramSendMaterial)
    if err != nil {
        fmt.Printf("==== %v\n",err)
        res := JsonRes{ReqId: paramSendMaterial.ReqId, ResCode: 1,Result:"paramSendMaterial bind json error"}
        c.JSON(200,res)
        return
    } else {

        var paramSendMaterialTable datastorage.ParamSendMaterialTable
        paramSendMaterialTable.SendMaterialSpeed             =   paramSendMaterial.Data.SendMaterialSpeed
        paramSendMaterialTable.StopDelay         =   paramSendMaterial.Data.StopDelay 
        paramSendMaterialTable.FitBenchmarkX         =   paramSendMaterial.Data.FitBenchmarkX
        paramSendMaterialTable.FitBenchmarkY         =   paramSendMaterial.Data.FitBenchmarkY
        paramSendMaterialTable.FitBenchmarkZ         =   paramSendMaterial.Data.FitBenchmarkZ
        paramSendMaterialTable.FitBenchmarkR         =   paramSendMaterial.Data.FitBenchmarkR

        paramSendMaterialTable.CameraHeight         =   paramSendMaterial.Data.CameraHeight
        paramSendMaterialTable.FitCompensationX         =   paramSendMaterial.Data.FitCompensationX
        paramSendMaterialTable.FitCompensationY         =   paramSendMaterial.Data.FitCompensationY
        paramSendMaterialTable.FitCompensationZ         =   paramSendMaterial.Data.FitCompensationZ 
        paramSendMaterialTable.FitCompensationR         =   paramSendMaterial.Data.FitCompensationR

        paramSendMaterialTable.RemoveFitInitialSpeed  =   paramSendMaterial.Data.RemoveFitInitialSpeed
        paramSendMaterialTable.RemoveFitAcceleration  =   paramSendMaterial.Data.RemoveFitAcceleration
        paramSendMaterialTable.RemoveFitDeceleration  =   paramSendMaterial.Data.RemoveFitDeceleration
        paramSendMaterialTable.RemoveFitSpeed  =   paramSendMaterial.Data.RemoveFitSpeed

        paramSendMaterialTable.ReturnInitialSpeed      =   paramSendMaterial.Data.ReturnInitialSpeed
        paramSendMaterialTable.ReturnAcceleration      =   paramSendMaterial.Data.ReturnAcceleration
        paramSendMaterialTable.ReturnDeceleration      =   paramSendMaterial.Data.ReturnDeceleration
        paramSendMaterialTable.ReturnSpeed             =   paramSendMaterial.Data.ReturnSpeed

        paramSendMaterialTable.GotoPhotoInitialSpeed   =   paramSendMaterial.Data.GotoPhotoInitialSpeed
        paramSendMaterialTable.GotoPhotoAcceleration   =   paramSendMaterial.Data.GotoPhotoAcceleration
        paramSendMaterialTable.GotoPhotoDeceleration   =   paramSendMaterial.Data.GotoPhotoDeceleration
        paramSendMaterialTable.GotoPhotoSpeed          =   paramSendMaterial.Data.GotoPhotoSpeed

        paramSendMaterialTable.FitInitialSpeed   =   paramSendMaterial.Data.FitInitialSpeed
        paramSendMaterialTable.FitAcceleration   =   paramSendMaterial.Data.FitAcceleration
        paramSendMaterialTable.FitDeceleration   =   paramSendMaterial.Data.FitDeceleration
        paramSendMaterialTable.FitSpeed          =   paramSendMaterial.Data.FitSpeed

        paramSendMaterialTable.PutOnTableInitialSpeed   =   paramSendMaterial.Data.PutOnTableInitialSpeed
        paramSendMaterialTable.PutOnTableAcceleration   =   paramSendMaterial.Data.PutOnTableAcceleration
        paramSendMaterialTable.PutOnTableDeceleration   =   paramSendMaterial.Data.PutOnTableDeceleration
        paramSendMaterialTable.PutOnTableSpeed          =   paramSendMaterial.Data.PutOnTableSpeed
       
        paramSendMaterialTable.BlowingHeight       =   paramSendMaterial.Data.BlowingHeight
        paramSendMaterialTable.AdjustmentTimes       =   paramSendMaterial.Data.AdjustmentTimes

        var res JsonRes
        

        e := s.Proxy.Insert(paramSendMaterialTable)
        if e != nil {
            fmt.Printf("paramSendMaterial data insert error!\n")
            res = JsonRes{ReqId: paramSendMaterial.ReqId, ResCode: 2,Result:"paramSendMaterial data insert err!"}
            return
        }

        res = JsonRes{ReqId: paramSendMaterial.ReqId, ResCode: 0,Result:""}
    //若返回json数据，可以直接使用gin封装好的JSON方法
        c.JSON(http.StatusOK, res)
        return
    }
}




/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//                                                           EVENT                                                             //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//======================================================== ADD EVENT ============================================================

func (s *GinServer)AddEvent(c *gin.Context) {

}

//======================================================== DEL EVENT ============================================================

func (s *GinServer)DelEvent(c *gin.Context) {
    
}


//====================================================== GET EVENT LIST =========================================================
func (s *GinServer)GetEventList(c *gin.Context) {
    
}