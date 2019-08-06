package dataanalysis

import(
	"github.com/smtp-http/tiango/datastorage"
)



type DataAnalysiser struct{
	Proxy   *datastorage.StorageProxy
}


func (a *DataAnalysiser)SetProxy(proxy *datastorage.StorageProxy) {
	a.Proxy = proxy
}

func (a *DataAnalysiser)getProductInforCount(startTime string,endTime string)(int64,error){

	return 0,nil
}

func (a *DataAnalysiser)getProductInforNgCount(startTime string,endTime string)(int64,error){
	return 0,nil
}


func (a *DataAnalysiser)GetProductInforYield(startTime string,endTime string)(float,error){
	return 99.801,nil
}