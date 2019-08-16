package dataanalysis

import(
	"fmt"
	"errors"
	"time"
	"github.com/smtp-http/tiango/datastorage"
	"github.com/grd/statistics"
	"sync"
)


type DataAnalysiser struct{
	Proxy   		*datastorage.StorageProxy
	
}

var analysiser *DataAnalysiser
var once_analysiser sync.Once
 
func GetDataAnalysiser() *DataAnalysiser {
    once_analysiser.Do(func() {
        analysiser = &DataAnalysiser{}
        analysiser.Proxy = datastorage.GetStorageProxy()
    })
    return analysiser
}


/*
func (a *DataAnalysiser)SetProxy(proxy *datastorage.StorageProxy) {
	a.Proxy = proxy
}*/

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


func average_ab(xs []datastorage.ProductInformationTable) (avg float64) {
	sum := 0.0
	switch len(xs) {
		case 0:
			avg = 0
		default:
			for _,v := range xs {

				sum += v.A_B
			}
			avg = sum / float64(len(xs))
	}
	return
}
func average_cd(xs []datastorage.ProductInformationTable) (avg float64) {
	sum := 0.0
	switch len(xs) {
		case 0:
			avg = 0
		default:
			for _,v := range xs {

				sum += v.B_D
			}
			avg = sum / float64(len(xs))
	}
	return
}

func average_ef(xs []datastorage.ProductInformationTable) (avg float64) {
	sum := 0.0
	switch len(xs) {
		case 0:
			avg = 0
		default:
			for _,v := range xs {

				sum += v.E_F
			}
			avg = sum / float64(len(xs))
	}
	return
}

func average_gh(xs []datastorage.ProductInformationTable) (avg float64) {
	sum := 0.0
	switch len(xs) {
		case 0:
			avg = 0
		default:
			for _,v := range xs {

				sum += v.G_H
			}
			avg = sum / float64(len(xs))
	}
	return
}

func variance(data statistics.Float64) float64 {
	return statistics.Variance(&data)
}



func getCpk(data statistics.Float64,mean float64)(float64,error){

	if data.Len() < 32 {
		return 0,errors.New("Insufficient data")
	}

	sysParam := datastorage.GetSysParam()

	varia := variance(data)

	if varia == 0 {
		return 0,errors.New("variance is zero!")
	}

	uperDiff := sysParam.Tolerance.UpperTolerance - mean
	lowerDiff := mean - sysParam.Tolerance.LowerTolerance
	fmt.Printf("++ mean: %f  uperDiff:%f lowerDiff:%f \n",mean,uperDiff,lowerDiff)
	if uperDiff > lowerDiff {

		return uperDiff/(3*varia),nil
	} else {

		return lowerDiff/(3*varia),nil
	}
}

func (a *DataAnalysiser)GetProductCpk(startTime int64,endTime int64)([]float64,error) {
	cpks := make([]float64,4)

	if a.Proxy == nil {
		fmt.Printf("storage proxy is nil!\n")
		return nil,errors.New("storage proxy is nil!")
	}

	strStart := time.Unix(startTime,0).Format("2006-01-02 15:04:05")
	strEnd := time.Unix(endTime,0).Format("2006-01-02 15:04:05")

	//strEnd = "2019-08-08 14:49:20"

	//fmt.Printf("start: %s      end: %s\n",strStart,strEnd)

    e,info := a.Proxy.GetCpkCalculateData(strStart,strEnd)
    if e != nil {
    	fmt.Println(e)
    	return nil,e
    }

    dataLen := len(info)



    if dataLen < 32 {
    	fmt.Println("Insufficient data")
    	return nil,errors.New("Insufficient data")
    }
//======================= A_B =============================

    ab_data := make(statistics.Float64, dataLen)


    i := 0
    for i < dataLen {
    	ab_data.SetValue(i,info[i].A_B)
    	i = i + 1
    }



    mean := average_ab(info)

    cpk,err := getCpk(ab_data,mean)
    if(err != nil) {
    	fmt.Printf("%v\n",err)
    	return nil,err
    }

    fmt.Printf("++++++++ cpk: %f\n",cpk)

    cpks[0] = cpk


//======================= c_d =============================
	bd_data := make(statistics.Float64, dataLen)

    i = 0
    for i < dataLen {
    	bd_data.SetValue(i,info[i].B_D)
    	i = i + 1
    }

    mean = average_cd(info)

    cpk,err = getCpk(bd_data,mean)
    if(err != nil) {
    	fmt.Printf("%v\n",err)
    	return nil,err
    }

    fmt.Printf("++++++++ cpk: %f\n",cpk)
    cpks[1] = cpk


//======================= e_f =============================
    ef_data := make(statistics.Float64, dataLen)

    i = 0
    for i < dataLen {
    	ef_data.SetValue(i,info[i].E_F)
    	i = i + 1
    }

    mean = average_ef(info)

    cpk,err = getCpk(ef_data,mean)
    if(err != nil) {
    	return nil,err
    }

    fmt.Printf("++++++++ cpk: %f\n",cpk)
    cpks[2] = cpk

//======================= g_h =============================
    gh_data := make(statistics.Float64, dataLen)

    i = 0
    for i < dataLen {
    	gh_data.SetValue(i,info[i].G_H)
    	i = i + 1
    }

    mean = average_gh(info)

    cpk,err = getCpk(gh_data,mean)
    if(err != nil) {
    	return nil,err
    }

    fmt.Printf("++++++++ cpk: %f\n",cpk)
    cpks[3] = cpk

	return cpks,nil
}




//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetStartTimeBefore (duration int32) int64 {

	if duration > 2 * 24 * 60 {
		duration = 2 * 24 * 60
	}

	tm := time.Now().Unix()
	tm = tm - int64(duration * 60)

	return tm
}

func (a *DataAnalysiser)GetConcentricRateStatisticalResult(duration int32) (*datastorage.ConcentricRateStatistical,error) {
	startTime := GetStartTimeBefore(duration)
	endTime := time.Now().Unix()

	strStart := time.Unix(startTime,0).Format("2006-01-02 15:04:05")
	strEnd := time.Unix(endTime,0).Format("2006-01-02 15:04:05")


	er,info := a.Proxy.GetCpkCalculateData(strStart,strEnd) 
	if er != nil {
		fmt.Println(er)
    	return nil,er
	}
   

    param := datastorage.GetSysParam()
    
    
    var crs datastorage.ConcentricRateStatistical
    length := len(info)
    crs.Count = length

    crs.AB_count = 0
    crs.CD_count = 0
    crs.EF_count = 0
    crs.GH_count = 0

	ab_data := make(statistics.Float64, length)
	cd_data := make(statistics.Float64, length)
	ef_data := make(statistics.Float64, length)
	gh_data := make(statistics.Float64, length)

    for k,v := range info {
    	if  v.A_B < param.Tolerance.LowerTolerance || v.A_B > param.Tolerance.UpperTolerance {
    		crs.AB_count += 1
    	} 

    	if v.B_D < param.Tolerance.LowerTolerance || v.B_D > param.Tolerance.UpperTolerance {
    		crs.CD_count += 1
    	} 

    	if v.E_F < param.Tolerance.LowerTolerance || v.E_F > param.Tolerance.UpperTolerance {
    		crs.EF_count += 1
    	} 

    	if v.G_H < param.Tolerance.LowerTolerance || v.G_H > param.Tolerance.UpperTolerance {
    		crs.GH_count += 1
    	} 

    	ab_data.SetValue(k,info[k].A_B)
    	cd_data.SetValue(k,info[k].B_D)
    	ef_data.SetValue(k,info[k].E_F)
    	gh_data.SetValue(k,info[k].G_H)

    }


    cpk,err := getCpk(ab_data,average_ab(info))
    if(err != nil) {
    	fmt.Println("get ab cpk err: ",err)
    	return nil,err
    }

    crs.AB_Cpk = cpk
    cpk,err = getCpk(cd_data,average_cd(info))
    if(err != nil) {
    	fmt.Println("get cd cpk err: ",err)
    	return nil,err
    }

    crs.CD_Cpk = cpk


	cpk,err = getCpk(ef_data,average_ef(info))
    if(err != nil) {
    	fmt.Println("get ef cpk err: ",err)
    	return nil,err
    }

    crs.EF_Cpk = cpk

    cpk,err = getCpk(gh_data,average_gh(info))
    if(err != nil) {
    	fmt.Println("get gh cpk err: ",err)
    	return nil,err
    }

    crs.GH_Cpk = cpk

    crs.Yield,err = a.GetProductInforYield(startTime,endTime)
    if err != nil{
    	fmt.Println("get yield err: ",err)
    	return nil,err
    }

    return &crs,nil
}

