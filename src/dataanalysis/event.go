package dataanalysis

import (
	"fmt"
	"github.com/gookit/event"
	"github.com/smtp-http/tiango/datastorage"
	"github.com/smtp-http/tiango/utils"
	"log"
)

///////////////////////////////////////////////////////////////////////////////////////
//                                      events                                       //
///////////////////////////////////////////////////////////////////////////////////////




type YieldEarlyWarningEvent struct {
	event.BasicEvent
	EventName 			string
	ThresholdValue    	float64        	//阈值
	
	SubscriberList 		[]string  		//事件的订阅者，也就是要通知到的人，邮箱地址
}


func (y *YieldEarlyWarningEvent) HandleLowYield(e event.Event) error {
	fmt.Printf("handle event: %s   ---11111112\n", e.Name())
	data := e.Data()
    fmt.Printf(" %v\n",data["Yield"])
    proxy := datastorage.GetStorageProxy()
    mails := proxy.GetEventSubscriberMails("LowYield")

	//fmt.Println(mails)

	sender := utils.GetMailSender()

	body := fmt.Sprintf("Yield : %v , too low",data["Yield"])

	sender.SendMail(mails,"Yield low",body)

    return nil
}



func CreatEvents(){

	fmt.Println("\n\n------------------------","\n\n")
	var events []datastorage.Event



	proxy := datastorage.GetStorageProxy()


	err := proxy.GetEngine().Find(&events)
	if err != nil {
		log.Fatalf("fail to Find events : %v", err)
		scheduler = nil
		return 
	}

	for _,v := range events {
		fmt.Printf("+++++++++++++++++ ev name:%s\n",v.EventName)
		if v.EventName == "LowYield" {
			e := &YieldEarlyWarningEvent{}
			e.SetName("LowYield")
			event.AddEvent(e)

			event.On("LowYield", event.ListenerFunc(e.HandleLowYield))
		} else if  v.EventName == "CpkFault" {
			fmt.Printf("----build CpkFault event\n")
		} else {
			log.Fatalf("Fail to creat event, event name is wrong!")
		}
	}
}



func init() {
	fmt.Println("__________$%^&*()_+++++++++++++++")

}




