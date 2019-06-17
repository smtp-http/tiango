package datastorage

import (
	//"time"
	//"reflect"
)
//================================================================================




type ProductInformation struct {
	// common
	DomSupplier 		string 	`json:"DomSupplier"`
	DpSupplier 			string 	`json:"DpSupplier"`
	ProductCn 			int32 	`json:"ProductCn"`
	LRstationDifference string 	`json:"LRstationDifference"`
	A_B 				float64 `json:"A_B"`
	B_D 				float64 `json:"B_D"`
	E_F 				float64 `json:"E_F"`
	G_H 				float64 `json:"G_H"`
	
	
	Result  		 	bool 	`json:"Result"`
}

type ProductInformationTable struct {
	Id 		int32 			`xorm:"id"`
	AminusB  float32 		`xorm:"AminusB"`
	BminusD  float32 		`xorm:"BminusD"`
	EminusF  float32 		`xorm:"EminusF"`
	GminusH  float32 		`xorm:"GminusH"`
	Result   bool       	`xorm:"Result"`
}