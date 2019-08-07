package dataanalysis

import(
	"fmt"
	"errors"
	"time"
	"github.com/smtp-http/tiango/datastorage"
	"errors"
)



type DataAnalysiser struct{
	Proxy   *datastorage.StorageProxy
}


func (a *DataAnalysiser)SetProxy(proxy *datastorage.StorageProxy) {
	a.Proxy = proxy
}

func (a *DataAnalysiser)getProductInforCount(startTime int64,endTime int64)(int64,error){    //startTime: Unix timestamp
	if a.Proxy == nil {
		return 0,errors.New("storage proxy is nil!")
	}

	strStart := time.Unix(startTime,0).Format("2006-01-02 15:04:05")
	strEnd := time.Unix(endTime,0).Format("2006-01-02 15:04:05")

	//strSql := "SELECT count(*) FROM `product_information_table` where ctime between '" + strStart "' and '" + strEnd + "2019-08-06 15:10:10'"
	var strSql string

	strSql = fmt.Sprintf("SELECT count(*) FROM `product_information_table` where ctime between '%s' and '%s'",strStart,strEnd)

	fmt.Println(strSql)

	count,err := a.Proxy.GetCount(strSql)
	if err != nil {
		return 0,err
	}

	return count,nil
}

func (a *DataAnalysiser)getProductInforOkCount(startTime int64,endTime int64)(int64,error){
	if a.Proxy == nil {
		return 0,errors.New("storage proxy is nil!")
	}

	strStart := time.Unix(startTime,0).Format("2006-01-02 15:04:05")
	strEnd := time.Unix(endTime,0).Format("2006-01-02 15:04:05")

	var strSql string

	strSql = fmt.Sprintf("SELECT count(*) FROM `product_information_table` where ctime between '%s' and '%s' and Result=1",strStart,strEnd)

	fmt.Println(strSql)

	count,err := a.Proxy.GetCount(strSql)
	if err != nil {
		return 0,err
	}

	return count,nil
}


func (a *DataAnalysiser)GetProductInforYield(startTime int64,endTime int64)(float64,error){
	count,err := a.getProductInforCount(startTime,endTime)
	if err != nil {
		return 0,err
	}

	var countOk int64

	countOk,err = a.getProductInforOkCount(startTime,endTime)
	if err != nil {
		return 0,err
	}
	
	result := float64(countOk)/float64(count)

	return result,nil
}


func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
		case 0:
			avg = 0
		default:
			for _,v := range xs {
				sum += v
			}
			avg = sum / float64(len(xs))
	}
	return
}


func getCpk(data []float64) (float64,error){
	if len(data) < 32 {
		return 0,errors.New("Insufficient data, at least 32 data required!")
	}


}

func (a *DataAnalysiser)GetProductCpk(startTime int64,endTime int64)(float64,error) {

	return 0,nil
}