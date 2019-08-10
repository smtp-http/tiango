package dataanalysis

import (
	"fmt"
	"github.com/gookit/event"
)

///////////////////////////////////////////////////////////////////////////////////////
//                                      events                                       //
///////////////////////////////////////////////////////////////////////////////////////




type YieldEarlyWarningEvent struct {
	event.BasicEvent
	ThresholdValue    	float64        	//阈值
	
	SubscriberList 		[]string  		//事件的订阅者，也就是要通知到的人，邮箱地址
}




func init() {
	fmt.Println("__________$%^&*()_+++++++++++++++")

}




