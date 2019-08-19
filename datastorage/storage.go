package datastorage


import (
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"errors"
	"fmt"
	"sync"
	"github.com/smtp-http/tiango/config"
	"log"
)

type DatabaseInfo struct {
	DriverName 		string 
	DataSourceName 	string
}


type StorageProxy struct {
	dbInfo DatabaseInfo
	engine *xorm.Engine
}

func init() {
	fmt.Println("__________$%^&*()_+++++++++++++++")

	loader := config.GetLoader()
	loader.Load("./config.json",config.GetConfig())

	var paramTable SysParamTable

    paramTable.LowerTolerance            =  -5.0
    paramTable.UpperTolerance             =   +5.1
    paramTable.MailAddr              =   "user@163.com"
    paramTable.Code    =   "123456"

    paramTable.SmtpServer                    =   "smtp@163.com"
    paramTable.YieldThresholdValue                    =   0.98

	engine, er := xorm.NewEngine(config.GetConfig().Database, config.GetConfig().DataSourceName)
    if er != nil {
        log.Fatalf("Fail to create engine: %v\n", er)
        return
    }

    engine.Sync2(&paramTable)

	has, err := engine.Table("sys_param_table").Exist() 
    if err == nil {
        if !has {
            _, e := engine.Insert(&paramTable) 
            if e != nil {
                fmt.Println("set param error, inster new record err: ",e)
            }
        } 
    } else {
    	fmt.Println("get exist err: ",err)
    }


}

func (s *StorageProxy) GetEngine() *xorm.Engine {
	return s.engine
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

func (s *StorageProxy) GetSysParam (id int64,data *SysParamTable)(bool,error){
	return s.engine.Id(id).Get(data)
}


func (s *StorageProxy) SqlQuery(sql string) ([]map[string][]byte,error) {
	return s.engine.Query(sql)
}

func (s *StorageProxy) GetCount(sql string)(int64,error){
	return s.engine.SQL(sql).Count()
}


func (s *StorageProxy) LoadSysParam(param *SysParam)error {
	var l sync.Mutex
	var paramTb SysParamTable
	s.engine.Id(1).Get(&paramTb)

	l.Lock()
	param.Tolerance.LowerTolerance = paramTb.LowerTolerance
	param.Tolerance.UpperTolerance = paramTb.UpperTolerance

	param.Mail.MailAddr = paramTb.MailAddr
	param.Mail.Code = paramTb.Code
	param.Mail.SmtpServer = paramTb.SmtpServer

	param.YieldThresholdValue = paramTb.YieldThresholdValue

	l.Unlock()

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


type SubscriberGroup struct {
    EventId int64 `xorm:"extends"`
    SubscriberId int64
}

func (s *SubscriberGroup) GetSubscriberId() int64 {
    return s.SubscriberId
}

func (s *StorageProxy) GetEventSubscriberMails(eventname string) []string {

	var subIds []int64 

	strNames := fmt.Sprintf("EventName = '%s'",eventname)

	err :=  s.engine.Table("event").Select("id").Where(strNames).Find(&subIds)
	if err != nil {
		fmt.Println("Get event  id err: ",err)
		return nil
	}

	fmt.Println(subIds)

	var submails []string

	strSubids := fmt.Sprintf("id in (select subscriber_id from ev_subscriber_table where event_id = %d)",subIds[0])

	err = s.engine.Table("subscriber").Select("SubscriberMail").Where(strSubids).Find(&submails)
	if err != nil {
		fmt.Println("Get subscriber mails err: ",err)
		return nil
	}

	return submails
}




var proxy *StorageProxy
var once_proxy sync.Once
 
func GetStorageProxy() *StorageProxy {

    once_proxy.Do(func() {
    	driverName := config.GetConfig().Database
    	dataSourceName := config.GetConfig().DataSourceName

    	fmt.Println("========== driverr name: ",driverName)
    	fmt.Println("========== dataSourceName: ",dataSourceName)

        proxy = &StorageProxy{}
        var err error
        proxy.engine, err = xorm.NewEngine(driverName,dataSourceName)
		if err != nil {
			fmt.Printf("Fail to create hrengine: %v\n", err)
			proxy = nil
			return
		}

		proxy.dbInfo.DriverName = driverName
		proxy.dbInfo.DataSourceName = dataSourceName

		if err := proxy.SyncTable(new(SysParamTable)); err != nil {
        	fmt.Printf("Fail to sync database SysParamTable: %v\n", err)
        	return
	    }

	    if err := proxy.SyncTable(new(ProductInformationTable)); err != nil {
	        fmt.Printf("Fail to sync database ProductInformationTable: %v\n", err)
	        return
	    }

	    if err := proxy.SyncTable(new(DpSizeTable)); err != nil {
	        fmt.Printf("Fail to sync database DpSizeTable: %v\n", err)
	        return
	    }

	    if err := proxy.SyncTable(new(DomSizeTable)); err != nil {
	        fmt.Printf("Fail to sync database DomSizeTable: %v\n", err)
	        return
	    }
	    
	    if err := proxy.SyncTable(new(ParamMaterialInputGuidanceTable)); err != nil {
	        fmt.Printf("Fail to sync database ParamMaterialInputGuidanceTable: %v\n", err)
	        return
	    }

	    if err := proxy.SyncTable(new(ParamSendMaterialTable)); err != nil {
	        fmt.Printf("Fail to sync database ParamSendMaterialTable: %v\n", err)
	        return
	    }

	    if err := proxy.SyncTable(new(JobBaseElement)); err != nil {
        	fmt.Printf("Fail to sync database JobBaseElement: %v\n", err)
        	return
    	}

		if err := proxy.SyncTable(new(Event)); err != nil {
			fmt.Printf("Fail to sync database EventTable: %v\n", err)
			return
		}

		if err := proxy.SyncTable(new(Subscriber)); err != nil {
			fmt.Printf("Fail to sync database SubscriberTable: %v\n", err)
			return
		}

		if err := proxy.SyncTable(new(EvSubscriberTable)); err != nil {
			fmt.Printf("Fail to sync database EvSubscriberTable: %v\n", err)
			return
		}

	    proxy.LoadSysParam(GetSysParam())

    })

    

    return proxy
}

//config.GetConfig().Database, config.GetConfig().DataSourceName
/*
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
*/
