package main

import (
    "errors"
    "log"
    "fmt"
    "github.com/go-xorm/xorm"
    //_ "github.com/mattn/go-sqlite3"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

// 银行账户
type Account struct {
    Id      int64
    Name    string `xorm:"unique"`
    Balance float64
    Version int `xorm:"version"` // 乐观锁
}

// ORM 引擎
var x *xorm.Engine

func init() {
    // 创建 ORM 引擎与数据库
    var err error
    x, err = xorm.NewEngine("mysql", "root:luxshare123@/sys?charset=utf8")
    if err != nil {
        log.Fatalf("Fail to create engine: %v\n", err)
    }

    // 同步结构体与数据表
    if err = x.Sync(new(Account)); err != nil {
        log.Fatalf("Fail to sync database: %v\n", err)
    }
}

// 创建新的账户
func newAccount(name string, balance float64) error {
    // 对未存在记录进行插入
    _, err := x.Insert(&Account{Name: name, Balance: balance})
    return err
}

// 获取账户信息
func getAccount(id int64) (*Account, error) {
    a := &Account{}
    // 直接操作 ID 的简便方法
    has, err := x.Id(id).Get(a)
    // 判断操作是否发生错误或对象是否存在
    if err != nil {
        return nil, err
    } else if !has {
        return nil, errors.New("Account does not exist")
    }
    return a, nil
}

// 用户转账
func makeTransfer(id1, id2 int64, balance float64) error {
    // 创建 Session 对象
    sess := x.NewSession()
    defer sess.Close()
    // 启动事务
    if err := sess.Begin(); err != nil {
        return err
    }

    a1, err := getAccount(id1)
    if err != nil {
        return err
    }

    a2, err := getAccount(id2)
    if err != nil {
        return err
    }

    if a1.Balance < balance {
        return errors.New("Not enough balance")
    }

    a1.Balance -= balance

    a2.Balance += balance

    if _, err = sess.Update(a1); err != nil {
        // 发生错误时进行回滚
        sess.Rollback()
        return err
    }
    if _, err = sess.Update(a2); err != nil {
        sess.Rollback()
        return err
    }
    // 完成事务
    return sess.Commit()

    return nil
}

// 用户存款
func makeDeposit(id int64, deposit float64) (*Account, error) {
    a, err := getAccount(id)
    if err != nil {
        return nil, err
    }
    sess := x.NewSession()
    defer sess.Close()
    if err = sess.Begin(); err != nil {
        return nil, err
    }
    a.Balance += deposit
    // 对已有记录进行更新
    if _, err = sess.Update(a); err != nil {
        sess.Rollback()
        return nil, err
    }

    return a, sess.Commit()
}

// 用户取款
func makeWithdraw(id int64, withdraw float64) (*Account, error) {
    a, err := getAccount(id)
    if err != nil {
        return nil, err
    }
    if a.Balance < withdraw {
        return nil, errors.New("Not enough balance")
    }
    sess := x.NewSession()
    defer sess.Close()
    if err = sess.Begin(); err != nil {
        return nil, err
    }
    a.Balance -= withdraw
    if _, err = sess.Update(a); err != nil {
        return nil, err
    }
    return a, sess.Commit()
}

// 按照 ID 正序排序返回所有账户
func getAccountsAscId() (as []Account, err error) {
    // 使用 Find 方法批量获取记录
    err = x.Find(&as)
    return as, err
}

// 按照存款倒序排序返回所有账户
func getAccountsDescBalance() (as []Account, err error) {
    // 使用 Desc 方法使结果呈倒序排序
    err = x.Desc("balance").Find(&as)
    return as, err
}

// 删除账户
func deleteAccount(id int64) error {
    // 通过 Delete 方法删除记录
    _, err := x.Delete(&Account{Id: id})
    return err
}


type A_B_data struct{
    a_b float64 `xorm:"A_B"`
}


func getDatas(engine *xorm.Engine) ( []A_B_data, error){
    var a_b  []A_B_data
    has,err := engine.Table("product_information_table").Select("A_B").Get(&a_b)
    fmt.Println(has)
    if err !=nil{
        err = fmt.Errorf("GetMailBoxAddress:%v",err)
        return nil,err
    }


    return a_b,nil
}

type ProductInformationTable struct {
    Id      int64           
    
    
    DomSupplier         string      `xorm:"DomSupplier"`
    DpSupplier          string      `xorm:"DpSupplier"`
    ProductCn           int32       `xorm:"ProductCn"`
    LRstationDifference string      `xorm:"LRstationDifference"`

    A_B                 float64     `xorm:"A_B"`
    B_D                 float64     `xorm:"B_D"`
    E_F                 float64     `xorm:"E_F"`
    G_H                 float64     `xorm:"G_H"`


    Result              bool        `xorm:"Result"`
    Angle               float64     `xorm:"Angle"`
    SizeA               float64     `xorm:"SizeA"`
    SizeB               float64     `xorm:"SizeB"`
    SizeC               float64     `xorm:"SizeC"`
    SizeD               float64     `xorm:"SizeD"`
    SizeE               float64     `xorm:"SizeE"`
    SizeF               float64     `xorm:"SizeF"`
    SizeG               float64     `xorm:"SizeG"`
    SizeH               float64     `xorm:"SizeH"`

    Ctime               time.Time   `xorm:"created"`//`xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP created" json:"ctime"`
}


type SysParamTable struct {
    Id                  int64
    //================== param tolerance ========================
    LowerTolerance      float64     `xorm:"LowerTolerance"`
    UpperTolerance      float64     `xorm:"UpperTolerance"`

    //==================            =============================
}

func main(){
    x, err := xorm.NewEngine("mysql", "root:luxshare123@/mdata?charset=utf8")
    if err != nil {
        log.Fatalf("Fail to create engine: %v\n", err)
    }

    var paramTb SysParamTable

    x.Sync2(paramTb)

    x.Id(1).Get(&paramTb)

    fmt.Println("lower: ",paramTb.LowerTolerance,"       upper: ",paramTb.UpperTolerance)


/*
    var info []ProductInformationTable
    er := x.Table("product_information_table").Select("A_B,B_D,E_F,G_H").Where("ctime between '2019-08-05 15:02:45' and '2019-08-06 15:10:10'").Find(&info)
    
    if er != nil {
        fmt.Println(er)
    } else {
        fmt.Println(info)
    }
    */
    //SELECT id ,name FROM `student` WHERE (id in (SELECT id FROM `studentinfo` WHERE (status = 2)))
    //sql := "SELECT * FROM `product_information_table` where ctime between '2019-08-05 15:08:45' and '2019-08-05 15:10:10';"
    sql := "SELECT count(*) FROM `product_information_table` where ctime between '2019-08-05 15:02:45' and '2019-08-06 15:10:10' and Result=1;"
    results, err := x.Query(sql)
    fmt.Println(results[0]["count(*)"])

    //var A_B []float64
    //sql = "select 'A_B' from `product_information_table` where ctime between '2019-08-05 15:02:45' and '2019-08-06 15:10:10'"
    //results, err = x.Query(sql)
    //x.Query(sql).Find(&A_B)
   // x.Table("product_information_table").Select("A_B").Query().Find(&A_B)
    //fmt.Println(results[0]["A_B"])
    //A_B,e := getDatas(x)
    //if e != nil {
    //    fmt.Println(e)
    //} else {
    //    fmt.Println(A_B)
    //}
    

    count, err := x.SQL("SELECT count(*) FROM `product_information_table` where ctime between '2019-08-05 15:02:45' and '2019-08-06 15:10:10' and Result=0").Count()
    if err != nil{
        fmt.Println(err)
    } else {
        fmt.Println(count)
    }
}