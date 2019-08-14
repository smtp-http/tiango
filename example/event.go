package main

import (
	"fmt"
	"github.com/gookit/event"
)

type MyEvent struct{
	event.BasicEvent
	customData string
}

func (e *MyEvent) CustomData() string {
    return e.customData
}

func (e *MyEvent) HandleData(s string) {
	fmt.Printf("+++++++ %s\n",s)
}

func (e *MyEvent) MyHandler(ev event.Event) error {
	fmt.Printf("handle event: %s   ---11111112\n", ev.Name())
	data := ev.Data()
    fmt.Printf("  %s   %s  %v\n",data["arg0"],data["arg1"],data["addr"])
    return nil
}

/*
func MyHandler(e event.Event) error {
	fmt.Printf("handle event: %s   ---11111112\n", e.Name())
        return nil
}

*/

func main() {
	/*
	// 注册事件监听器
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
        fmt.Printf("handle event: %s   ---1\n", e.Name())
        return nil
    }), event.Normal)
	
	// 注册多个监听器
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
        fmt.Printf("handle event: %s   ---2 ", e.Name())
        data := e.Data()
        fmt.Printf("  %s   %s  %v\n",data["arg0"],data["arg1"],data["addr"])
        return nil
    }), event.High)
	
	
	// 触发事件
	// 注意：第二个监听器的优先级更高，所以它会先被执行
	event.MustFire("evt1", event.M{"arg0": "val0", "arg1": "val1"})

	fmt.Printf("\n========================================================\n")

	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
        fmt.Printf("handle event: %s   ---3\n", e.Name())
        return nil
    }), event.High)

    addrs := [2]string{"abc@124.com","def@shit.com"}

    event.MustFire("evt1", event.M{"arg0": "val0", "arg1": "val1","addr":addrs})
*/
    e := &MyEvent{customData: "hello"}
	e.SetName("e1")
	event.AddEvent(e)

	fmt.Printf("----------: %s\n",e.Name())

	// add listener
	event.On("e1", event.ListenerFunc(func(e event.Event) error {
	   fmt.Printf("custom Data: %s\n", e.(*MyEvent).CustomData())
	   e.(*MyEvent).HandleData("shit")
	   return nil
	}))

	event.On("e1", event.ListenerFunc(e.MyHandler))

	// trigger
	//event.Fire("e1", nil)
	// OR
	// event.FireEvent(e)
	addrs := [2]string{"abc@124.com","def@shit.com"}
	event.MustFire("e1", event.M{"arg0": "val0", "arg1": "val1","addr":addrs})
}