
package dataanalysis

import(
	"github.com/smtp-http/tiango/datastorage"
	"github.com/jasonlvhit/gocron"
	"github.com/gookit/event"
	"sync"
	"log"
	"fmt"
	"errors"
	"time"
)

	


type JobScheduler struct {
	
	BaseElements 		[]datastorage.JobBaseElement 

	//JobSet 				map[string]func(){}
	
}


func (s *JobScheduler) AddJob(elem *datastorage.JobBaseElement) error {
	if elem == nil {
		return errors.New("job elem is nil!")
	}

	return nil
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetTodayByeMinute(hour int,minute int) int64 {
	t1:=time.Now().Year()
	t2:=time.Now().Month()
	t3:=time.Now().Day()

	todayTime:=time.Date(t1,t2,t3,hour,minute,0,0,time.Local)

	return todayTime.Unix()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func JobDailySummary (){

	fmt.Println("====== day ======== ",time.Now())
	analysiser := GetDataAnalysiser()


	result,e := analysiser.GetProductInforYield(GetTodayByeMinute(8,30),GetTodayByeMinute(17,30))
    if e != nil {
        fmt.Println("+++  ",e)
        return
    } else {
        fmt.Println("--- result: ",result)
    }

    
    param := datastorage.GetSysParam()
    threshold := param.YieldThresholdValue

    if result < threshold {
    	fmt.Println("触发Yield过低通知事件")
    	//addrs := [2]string{"abc@124.com","def@shit.com"}
    	
		event.MustFire("LowYield", event.M{"Yield":result})
    }
}

func JobMonthlySummary(){

}

func JobMinutesSummary(){
	fmt.Println("====== minutes ========")
}




var scheduler *JobScheduler
var once_sch sync.Once
 
func GetJobScheduler() *JobScheduler {
    once_sch.Do(func() {

    	CreatEvents()

        scheduler = &JobScheduler{}

        proxy := datastorage.GetStorageProxy()
        err := proxy.GetEngine().Find(&scheduler.BaseElements)
        if err != nil {
        	log.Fatalf("fail to Find BaseElements : %v", err)
        	scheduler = nil
        	return 
        }

       
        go func(){
	        for _,v := range scheduler.BaseElements {

	        	fmt.Println(v.JobName)
	        	if v.CycleUnit == "DAY" {
	        		fmt.Println("===== day ")
	        		gocron.Every(v.TriggerCycle).Day().At(v.TriggerTimePoint).Do(JobDailySummary)    //Time point : example '10:30'
	        	} else if v.CycleUnit == "MONDAY" {

	        	} else if v.CycleUnit == "MINUTES" {
	        		fmt.Println("===== minutes ")
	        		gocron.Every(1).Minute().Do(JobMinutesSummary)
	        	}

	    	}
	    	<- gocron.Start()

    	}()
    })
    return scheduler
}