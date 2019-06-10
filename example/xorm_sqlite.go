package main
import (
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-xorm/xorm"
    "log"
    "os"
    "time"
    "fmt"
)


//var engine *xorm.Engine

var hrengine *xorm.Engine

type User struct {
    Id   int64
    Name string  `xorm:"varchar(25) notnull unique 'usr_name'"`
    Age  uint32
    CreatedAt time.Time `xorm:"created"` //创建时间 自动
    UpdatedAt time.Time `xorm:"updated"`    //更新时间 自动
}



// 启动程序后就执行
func main() {
   // 创建 ORM 引擎与数据库
	var err error
	hrengine, err = xorm.NewEngine("sqlite3", "f:\\db\\test.db")
	if err != nil {
		log.Fatalf("Fail to create hrengine: %v\n", err)
	}

   // 同步结构体与数据表,和python的 migrate 一样同步数据库
	if err = hrengine.Sync2(new(User)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
   // 创建日志 可以选用
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}

	//engine.Logger = xorm.NewSimpleLogger(f)
	xorm.NewSimpleLogger(f)

	user := new(User)
	user.Name = "myname"
	u, err := hrengine.Id(20).Update(user)
	if err != nil{
		fmt.Println(err) 
	} else {
		fmt.Println(u)
	}


	pEveryOne := make([]*User, 0)
	err = hrengine.Find(&pEveryOne)

	if err != nil{
		fmt.Println(err) 
	} else {
		for _, u := range pEveryOne {
        	fmt.Printf("%v\n", u)
		}
	}
/*
	

	user := new(User) // 创建一个对象
	user.Name = "myname" // 给对象的字段赋值
	user.Age = 25
	u, err := hrengine.InsertOne(user) // 执行插入操作
	if err != nil{
		fmt.Println(err) 
	} else {
		fmt.Println(u)
	}

// INSERT INTO user (name) values (?)

	// ================= users =========================
	users := make([]*User, 0)
	usr := new(User)
	usr.Name = "name0"
	//users[0].Name = "name0"
	users = append(users,usr)
	affected, err := hrengine.Insert(&users)
	if err != nil{
		fmt.Println(err) 
	} else {
		fmt.Println(affected)
	}
*/
}

