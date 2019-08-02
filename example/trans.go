package main


import (
  "errors"
  "os"
  //"utils"
  "log"
  "fmt"
  "github.com/go-xorm/xorm"
  //_ "github.com/mattn/go-sqlite3"
  _ "github.com/go-sql-driver/mysql"
)


type Account struct {
  Id int64 // 主键 int64
  Name string `xorm:"unique"` // 唯一索引
  Blance float64
  Version int `xorm:"version"` // 乐观锁
}
var engine *xorm.Engine
func init () {
  // 根据名称注册驱动并创建 ORM 引擎
  var err error
  engine, err = xorm.NewEngine("mysql", "root:luxshare123@/go?charset=utf8")
  
  //engine, err = xorm.NewEngine("sqlite3", "f:\\db\\trans.db")
  if err != nil {
    log.Fatalf("fail to create engine：%v", err)
  }
  if err := engine.Sync(new(Account)); err != nil {
    log.Fatalf("fail to sync database：%v", err)
  }
  // 记录SQL语句log
  f, err := os.Create("sql.log")
  if err != nil {
    log.Fatalf("fail to create log file %v", err)
  }
  // defer f.Close() 打开这行就记录不到，需要注视掉
  engine.SetLogger(xorm.NewSimpleLogger(f))
  engine.ShowSQL(true)
}
func transfer(id1, id2 int , blance float64) error {
  account_1 := &Account{}
  has1, err := engine.Id(id1).Get(account_1)
  if err != nil {
    return err
  } else if !has1 {
    return errors.New("account_1 not found")
  }
  account_2 := &Account{}
  has2, err := engine.Id(id2).Get(account_2)
  if err != nil {
    return err
  } else if !has2 {
    return errors.New("account_2 not found")
  }
  if account_1.Blance < blance {
    return errors.New("blance not enough")
  }
  account_1.Blance -= blance
  account_2.Blance += blance
  // 事务处理
  sess:= engine.NewSession()
  defer sess.Close()
  if err := sess.Begin() ; err != nil {
    return errors.New("fail to session begin")
  }
  if _, err := sess.Id(id1).Cols("blance").Update(account_1) ; err != nil {
    sess.Rollback()
    return errors.New("fail to update 1")
  }
  if _, err := sess.Id(id2).Cols("blance").Update(account_2); err != nil {
    sess.Rollback()
    return errors.New("fail to update 2")
  }
  return sess.Commit()
}
func main() {
  // 填充数据
  
    count, err := engine.Count(new(Account))
    if err != nil {
    log.Fatal(err.Error())
    }
    for i := count ; i < 10 ; i++ {
      demo := Account{
      Blance :float64(i * 100),
      Name : fmt.Sprintf("zero_%d", i),
    }
    if _, err := engine.InsertOne(&demo) ;err != nil {
      log.Fatal(err.Error())
    }
  }
  
  // 转账
  error := transfer(5, 8, 100)
  if error != nil {
    log.Fatalf("fail to transfer %v", error)
  }
  fmt.Println("transfer OK")
}

