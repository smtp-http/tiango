
package dataanalysis

import(
	"github.com/smtp-http/tiango/datastorage"
	"github.com/jasonlvhit/gocron"
	"sync"
	"log"
	"fmt"
)

	


type JobScheduler struct {
	
	BaseElements 		[]datastorage.JobBaseElement 

	//JobSet 				map[string]func(){}
	
}

func JobDailySummary (){

}

func JobMonthlySummary(){

}

func JobMinutesSummary(){
	
}




var scheduler *JobScheduler
var once_sch sync.Once
 
func GetJobScheduler() *JobScheduler {
    once_sch.Do(func() {
        scheduler = &JobScheduler{}

        proxy := datastorage.GetStorageProxy()
        err := proxy.GetEngine().Find(&scheduler.BaseElements)
        if err != nil {
        	log.Fatalf("fail to Find BaseElements : %v", err)
        	scheduler = nil
        	return
        }

        for _,v := range scheduler.BaseElements {
        	//fmt.Println(v) //prints 0, 1, 2
        	fmt.Println(v.JobName)
        	if v.CycleUnit == "DAY" {
        		gocron.Every(v.TriggerCycle).Day().At(v.TriggerTimePoint).Do(JobDailySummary)    //Time point : example '10:30'
        	} else if v.CycleUnit == "MONDAY" {

        	} else if v.CycleUnit == "MINUTES" {
        		gocron.Every(2).Minutes().Do(JobMinutesSummary)
        	}

    	}
    })
    return scheduler
}