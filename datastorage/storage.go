package datastorage


import (
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"errors"
	"fmt"
)

type DatabaseInfo struct {
	DriverName 		string 
	DataSourceName 	string
}


type StorageProxy struct {
	dbInfo DatabaseInfo
	engine *xorm.Engine
}


func (s *StorageProxy) SyncTable(table interface{}) error {
	if s.engine != nil {
		return s.engine.Sync2(table)
	} else {
		return errors.New("engin is nil!")
	}
	
}


func (s *StorageProxy) Insert(data interface{}) error {
	u, err := s.engine.InsertOne(data) // 执行插入操作
	if err != nil{
		fmt.Println(err) 
		return err
	} else {
		fmt.Println(u)
		return nil
	}
}

func (s *StorageProxy) Traversing() error {
	everyOne := make([]*ProductInformationTable, 0)
	err := s.engine.Find(&everyOne)

	if err != nil{
		fmt.Println(err) 
		return err
	} else {
		for _, u := range everyOne {
        	fmt.Printf("%v\n", u)
		}

		return nil
	}
}


func CreateStorageProxy(driverName string,dataSourceName string) *StorageProxy {
	var err error

	proxy := new(StorageProxy)
	proxy.engine, err = xorm.NewEngine(driverName,dataSourceName)
	if err != nil {
		fmt.Printf("Fail to create hrengine: %v\n", err)
		return nil
	}

	proxy.dbInfo.DriverName = driverName
	proxy.dbInfo.DataSourceName = dataSourceName

	//日志打印SQL
	//proxy.engine.ShowSQL(true)

	return proxy
}

