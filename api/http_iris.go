//包主显示如何在您的Web应用程序中使用orm
//它只是插入一列并选择第一列。
package api

import (
    "time"

    "github.com/go-xorm/xorm"
    "github.com/kataras/iris"
    "github.com/smtp-http/tiango/config"
    "github.com/smtp-http/tiango/datastorage"
    _ "github.com/mattn/go-sqlite3"
)

/*
   go get -u github.com/mattn/go-sqlite3
   go get -u github.com/go-xorm/xorm
   如果您使用的是win64并且无法安装go-sqlite3：
       1.下载：https：//sourceforge.net/projects/mingw-w64/files/latest/download
       2.选择“x86_x64”和“posix”
       3.添加C:\Program Files\mingw-w64\x86_64-7.1.0-posix-seh-rt_v5-rev1\mingw64\bin
       到你的PATH env变量。
   手册: http://xorm.io/docs/
*/
//User是我们的用户表结构。
type User struct {
    ID        int64  // xorm默认自动递增
    Version   string `xorm:"varchar(200)"`
    Salt      string
    A用户名      string
    Password  string    `xorm:"varchar(200)"`
    Languages string    `xorm:"varchar(200)"`
    CreatedAt time.Time `xorm:"created"`
    UpdatedAt time.Time `xorm:"updated"`
}

type IrisServer struct {
    App     *iris.Application
    Proxy   *datastorage.StorageProxy
}


func (s *IrisServer) StartHttpServer() {
    s.App   = iris.New()
    s.Proxy = datastorage.CreateStorageProxy("sqlite3", config.GetConfig().DbName)
    if s.Proxy == nil {
        s.App.Logger().Fatalf("orm failed to initialized\n")
        return
    }

    // iris.RegisterOnInterrupt(func() {
    //     orm.Close()
    // })

    if err := s.Proxy.SyncTable(new(datastorage.ProductInformationTable)); err != nil {
        s.App.Logger().Fatalf("Fail to sync database ProductInformationTable: %v\n", err)
        return
    }

    if err := s.Proxy.SyncTable(new(datastorage.DpSizeTable)); err != nil {
        s.App.Logger().Fatalf("Fail to sync database DpSizeTable: %v\n", err)
        return
    }

    
    if err := s.Proxy.SyncTable(new(datastorage.ParamMaterialInputGuidanceTable)); err != nil {
        s.App.Logger().Fatalf("Fail to sync database ParamMaterialInputGuidanceTable: %v\n", err)
        return
    }

    if err := s.Proxy.SyncTable(new(datastorage.ParamSendMaterialTable)); err != nil {
        s.App.Logger().Fatalf("Fail to sync database ParamSendMaterialTable: %v\n", err)
        return
    }

    //s.App.Get()
}


func Test() {
    app := iris.New()
    orm, err := xorm.NewEngine("sqlite3", config.GetConfig().DbName)
    if err != nil {
        app.Logger().Fatalf("orm failed to initialized: %v", err)
    }
    iris.RegisterOnInterrupt(func() {
        orm.Close()
    })
    err = orm.Sync2(new(User))
    if err != nil {
        app.Logger().Fatalf("orm failed to initialized User table: %v", err)
    }
    app.Get("/insert", func(ctx iris.Context) {
        user := &User{A用户名: "大大", Salt: "hash---", Password: "hashed", CreatedAt: time.Now(), UpdatedAt: time.Now()}
        orm.Insert(user)
        ctx.Writef("user inserted: %#v", user)
    })
    app.Get("/get/{id:int}", func(ctx iris.Context) {
        id, _ := ctx.Params().GetInt("id")
        //int到int64
        id64 := int64(id)
        ctx.Writef("id is %#v", id64)
        user := User{ID: id64}
        if ok, _ := orm.Get(&user); ok {
            ctx.Writef("user found: %#v", user)
        }
    })
    app.Get("/delete", func(ctx iris.Context) {
        user := User{ID: 1}
        orm.Delete(user)
        ctx.Writef("user delete: %#v", user)
    })
    app.Get("/update", func(ctx iris.Context) {
        user := User{ID: 2, A用户名: "小小"}
        orm.Update(user)
        ctx.Writef("user update: %#v", user)
    })
    // http://localhost:8080/insert
    // http://localhost:8080/get/数字
    app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}