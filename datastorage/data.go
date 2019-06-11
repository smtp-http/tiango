package datastorage

import (
	"time"
)
//================================================================================

type Materials struct {
	MaterialNumber	int64
	Material 	string
	Supplier 	string
	Batch 		string
}

type MaterialsTable struct {
	Id 			int64
	MaterialNumber	int64 	`xorm:"material_number"`
	Material 	string 		`xorm:"varchar(25) notnull unique 'material_name'"`
	Supplier 	string 		`xorm:"supplier"`
	Batch 		string 		`xorm:"batch"`
	CreatedAt time.Time `xorm:"created"` //创建时间 自动	
	UpdatedAt time.Time `xorm:"updated"`    //更新时间 自动
}

//===================================================================================

type ProcessParameter struct {

}


type ProcessParameterTable struct{
	Id int64
}
