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

func (s *StorageProxy) SqlQuery(sql string) ([]map[string][]byte,error) {
	return s.engine.Query(sql)
}

func (s *StorageProxy) GetCount(sql string)(int64,error){
	return s.engine.SQL(sql).Count()
}


func (s *StorageProxy) LoadSysParam(param *SysParam)error {
	var paramTb SysParamTable
	s.engine.Id(1).Get(&paramTb)

	param.Tolerance.LowerTolerance = paramTb.LowerTolerance
	param.Tolerance.UpperTolerance = paramTb.UpperTolerance

	return nil
}

func (s *StorageProxy) GetCpkCalculateData(strStart string,strEnd string) (error,[]ProductInformationTable){
	var info []ProductInformationTable
	whereInfo := fmt.Sprintf("ctime between '%s' and '%s'",strStart,strEnd)
	
    err := s.engine.Table("product_information_table").Select("A_B,B_D,E_F,G_H").Where(whereInfo).Find(&info)
    
    if err != nil {
        fmt.Println(err)
        return err,nil
    } 

    return nil,info
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

