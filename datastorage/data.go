package datastorage

import (
	//"time"
	//"reflect"
	"time"
)

//===================================== Prodection info ===========================================


type ProductInformation struct {
	// common
	DomSupplier 		string 		`json:"DomSupplier"`
	DpSupplier 			string 		`json:"DpSupplier"`
	ProductCn 			int32 		`json:"ProductCn"`
	LRstationDifference string 		`json:"LRstationDifference"`

	A_B 				float64 	`json:"A_B"`
	B_D 				float64 	`json:"B_D"`
	E_F 				float64 	`json:"E_F"`
	G_H 				float64 	`json:"G_H"`


	Result  		 	bool 		`json:"Result"`
	Angle 				float64 	`json:"Angle"`
	SizeA 				float64 	`json:"SizeA"`
	SizeB 				float64 	`json:"SizeB"`
	SizeC 				float64 	`json:"SizeC"`
	SizeD 				float64 	`json:"SizeD"`
	SizeE 				float64 	`json:"SizeE"`
	SizeF 				float64 	`json:"SizeF"`
	SizeG 				float64 	`json:"SizeG"`
	SizeH				float64 	`json:"SizeH"`
}


type ProductInformationTable struct {
	Id 		int64 			
	
	
	DomSupplier 		string 		`xorm:"DomSupplier"`
	DpSupplier 			string 		`xorm:"DpSupplier"`
	ProductCn 			int32 		`xorm:"ProductCn"`
	LRstationDifference string 		`xorm:"LRstationDifference"`

	A_B 				float64 	`xorm:"A_B"`
	B_D 				float64 	`xorm:"B_D"`
	E_F 				float64 	`xorm:"E_F"`
	G_H 				float64 	`xorm:"G_H"`


	Result  		 	bool 		`xorm:"Result"`
	Angle 				float64 	`xorm:"Angle"`
	SizeA 				float64 	`xorm:"SizeA"`
	SizeB 				float64 	`xorm:"SizeB"`
	SizeC 				float64 	`xorm:"SizeC"`
	SizeD 				float64 	`xorm:"SizeD"`
	SizeE 				float64 	`xorm:"SizeE"`
	SizeF 				float64 	`xorm:"SizeF"`
	SizeG 				float64 	`xorm:"SizeG"`
	SizeH				float64 	`xorm:"SizeH"`

	Ctime 				time.Time 	`xorm:"created"`//`xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP created" json:"ctime"`
}



//=========================== DP size ===========================================



type DpSize struct {
	DomSupplier 		string 		`json:"DomSupplier"`
	DpSupplier 			string 		`json:"DpSupplier"`
	ProductCn 			int32 		`json:"ProductCn"`
	LRstationDifference string 		`json:"LRstationDifference"`

	Length 				float64 	`json:"Length"`
	Width 				float64 	`json:"Width"`
	LongSideAngle 		float64 	`json:"LongSideAngle"`
	ShortSideAngle 		float64 	`json:"ShortSideAngle"`
	Angle1 				float64 	`json:"Angle1"`
	Angle2 				float64 	`json:"Angle2"`
	Angle3 				float64 	`json:"Angle3"`
	Angle4 				float64 	`json:"Angle4"`
	DX 					float64 	`json:"DX"`
	DY 					float64 	`json:"DY"`
	DR 					float64 	`json:"DR"`

}

type DpSizeTable struct {
	Id 					int64 		

	Result 				string 		`xorm:"Result"`

	Position			string 		`xorm:"Position"`
	DomSupplier 		string 		`xorm:"DomSupplier"`
	DpSupplier 			string 		`xorm:"DpSupplier"`
	ProductCn 			int32 		`xorm:"ProductCn"`
	LRstationDifference string 		`xorm:"LRstationDifference"`

	Length 				float64 	`xorm:"Length"`
	Width 				float64 	`xorm:"Width"`
	LongSideAngle 		float64 	`xorm:"LongSideAngle"`
	ShortSideAngle 		float64 	`xorm:"ShortSideAngle"`
	Angle1 				float64 	`xorm:"Angle1"`
	Angle2 				float64 	`xorm:"Angle2"`
	Angle3 				float64 	`xorm:"Angle3"`
	Angle4 				float64 	`xorm:"Angle4"`
	DX 					float64 	`xorm:"DX"`
	DY 					float64 	`xorm:"DY"`
	DR 					float64 	`xorm:"DR"`

}

//=========================== Dom size ===========================================



type DomSize struct {
	DomSupplier 		string 		`json:"DomSupplier"`
	DpSupplier 			string 		`json:"DpSupplier"`
	ProductCn 			int32 		`json:"ProductCn"`
	LRstationDifference string 		`json:"LRstationDifference"`

	Length 				float64 	`json:"Length"`
	Width 				float64 	`json:"Width"`
	LongSideAngle 		float64 	`json:"LongSideAngle"`
	ShortSideAngle 		float64 	`json:"ShortSideAngle"`
	Angle1 				float64 	`json:"Angle1"`
	Angle2 				float64 	`json:"Angle2"`
	Angle3 				float64 	`json:"Angle3"`
	Angle4 				float64 	`json:"Angle4"`
	DX 					float64 	`json:"DX"`
	DY 					float64 	`json:"DY"`
	DR 					float64 	`json:"DR"`

}

type DomSizeTable struct {
	Id 					int64 		

	Result 				string 		`xorm:"Result"`

	Position			string 		`xorm:"Position"`
	DomSupplier 		string 		`xorm:"DomSupplier"`
	DpSupplier 			string 		`xorm:"DpSupplier"`
	ProductCn 			int32 		`xorm:"ProductCn"`
	LRstationDifference string 		`xorm:"LRstationDifference"`

	Length 				float64 	`xorm:"Length"`
	Width 				float64 	`xorm:"Width"`
	LongSideAngle 		float64 	`xorm:"LongSideAngle"`
	ShortSideAngle 		float64 	`xorm:"ShortSideAngle"`
	Angle1 				float64 	`xorm:"Angle1"`
	Angle2 				float64 	`xorm:"Angle2"`
	Angle3 				float64 	`xorm:"Angle3"`
	Angle4 				float64 	`xorm:"Angle4"`
	DX 					float64 	`xorm:"DX"`
	DY 					float64 	`xorm:"DY"`
	DR 					float64 	`xorm:"DR"`

}

//=========================== Environment ==========================================

type Environment struct {                          // from modbus
	Temperature 		float64 	
	Pressure 			float64 
}








//=================================================== param ============================================================

//============================ MaterialInputGuidance ================================
type ParamMaterialInputGuidance struct {
	PhotoDelay 					float64 	`json:"PhotoDelay"`
	CompensationX1 				float64 	`json:"CompensationX1"`
	CompensationY1 				float64 	`json:"CompensationY1"`
	CompensationR1 				float64 	`json:"CompensationR1"`
	CompensationX2 				float64 	`json:"CompensationX2"`
	CompensationY2 				float64 	`json:"CompensationY2"`
	CompensationR2				float64 	`json:"CompensationR2"`
	CompensationX3				float64 	`json:"CompensationX3"`
	CompensationY3				float64 	`json:"CompensationY3"`
	CompensationR3				float64 	`json:"CompensationR3"`
	CompensationX4				float64 	`json:"CompensationX4"`
	CompensationY4 				float64 	`json:"CompensationY4"`
	CompensationR4				float64 	`json:"CompensationR4"`

	MaterialInputReferenceX 	float64 	`json:"MaterialInputReferenceX"`
	MaterialInputReferenceY 	float64 	`json:"MaterialInputReferenceY"`
	MaterialInputReferenceZ 	float64 	`json:"MaterialInputReferenceZ"`
	MaterialInputReferenceR 	float64 	`json:"MaterialInputReferenceR"`

	FallingInitialSpeed 		float64 	`json:"FallingInitialSpeed"`
	FallingAcceleration 		float64 	`json:"FallingAcceleration"`
	FallingDeceleration 		float64 	`json:"FallingDeceleration"`
	FallingSpeed 				float64 	`json:"FallingSpeed"`

	PutOnTableInitialSpeed		float64 	`json:"PutOnTableInitialSpeed"`
	PutOnTableAcceleration 		float64 	`json:"PutOnTableAcceleration"`
	PutOnTableDeceleration 		float64 	`json:"PutOnTableDeceleration"`
	PutOnTableSpeed 			float64 	`json:"PutOnTableSpeed"`

	MaterialInputDelay 			float64 	`json:"MaterialInputDelay"`
}


type ParamMaterialInputGuidanceTable struct {
	Id 					int64 		

	PhotoDelay 					float64 	`xorm:"PhotoDelay"`
	CompensationX1 				float64 	`xorm:"CompensationX1"`
	CompensationY1 				float64 	`xorm:"CompensationY1"`
	CompensationR1 				float64 	`xorm:"CompensationR1"`
	CompensationX2 				float64 	`xorm:"CompensationX2"`
	CompensationY2 				float64 	`xorm:"CompensationY2"`
	CompensationR2				float64 	`xorm:"CompensationR2"`
	CompensationX3				float64 	`xorm:"CompensationX3"`
	CompensationY3				float64 	`xorm:"CompensationY3"`
	CompensationR3				float64 	`xorm:"CompensationR3"`
	CompensationX4				float64 	`xorm:"CompensationX4"`
	CompensationY4 				float64 	`xorm:"CompensationY4"`
	CompensationR4				float64 	`xorm:"CompensationR4"`

	MaterialInputReferenceX 	float64 	`xorm:"MaterialInputReferenceX"`
	MaterialInputReferenceY 	float64 	`xorm:"MaterialInputReferenceY"`
	MaterialInputReferenceZ 	float64 	`xorm:"MaterialInputReferenceZ"`
	MaterialInputReferenceR 	float64 	`xorm:"MaterialInputReferenceR"`

	FallingInitialSpeed 		float64 	`xorm:"FallingInitialSpeed"`
	FallingAcceleration 		float64 	`xorm:"FallingAcceleration"`
	FallingDeceleration 		float64 	`xorm:"FallingDeceleration"`
	FallingSpeed 				float64 	`xorm:"FallingSpeed"`

	PutOnTableInitialSpeed		float64 	`xorm:"PutOnTableInitialSpeed"`
	PutOnTableAcceleration 		float64 	`xorm:"PutOnTableAcceleration"`
	PutOnTableDeceleration 		float64 	`xorm:"PutOnTableDeceleration"`
	PutOnTableSpeed 			float64 	`xorm:"PutOnTableSpeed"`

	MaterialInputDelay 			float64 	`xorm:"MaterialInputDelay"`
}



//
type ParamSendMaterial struct {
	SendMaterialSpeed 	float64 	`json:"SendMaterialSpeed"`
	StopDelay 			float64 	`json:"StopDelay"`

	FitBenchmarkX 		float64 	`json:"FitBenchmarkX"`
	FitBenchmarkY 		float64 	`json:"FitBenchmarkY"`
	FitBenchmarkZ 		float64 	`json:"FitBenchmarkZ"`
	FitBenchmarkR 		float64 	`json:"FitBenchmarkR"`

	CameraHeight 		float64 	`json:"CameraHeight"`
	FitCompensationX 	float64 	`json:"FitCompensationX"`
	FitCompensationY 	float64 	`json:"FitCompensationY"`
	FitCompensationZ 	float64 	`json:"FitCompensationZ"`
	FitCompensationR 	float64 	`json:"FitCompensationR"`

	RemoveFitInitialSpeed 		float64 	`json:"RemoveFitInitialSpeed"`
	RemoveFitAcceleration 		float64 	`json:"RemoveFitAcceleration"`
	RemoveFitDeceleration 		float64 	`json:"RemoveFitDeceleration"`
	RemoveFitSpeed 				float64 	`json:"RemoveFitSpeed"`

	ReturnInitialSpeed 			float64 	`json:"ReturnInitialSpeed"`
	ReturnAcceleration 			float64 	`json:"ReturnAcceleration"`
	ReturnDeceleration 			float64 	`json:"ReturnDeceleration"`
	ReturnSpeed 				float64 	`json:"ReturnSpeed"`

	GotoPhotoInitialSpeed 		float64 	`json:"GotoPhotoInitialSpeed"`
	GotoPhotoAcceleration 		float64 	`json:"GotoPhotoAcceleration"`
	GotoPhotoDeceleration 		float64 	`json:"GotoPhotoDeceleration"`
	GotoPhotoSpeed 				float64 	`json:"GotoPhotoSpeed"`

	FitInitialSpeed 		float64 	`json:"FitInitialSpeed"`
	FitAcceleration 		float64 	`json:"FitAcceleration"`
	FitDeceleration 		float64 	`json:"FitDeceleration"`
	FitSpeed 				float64 	`json:"FitSpeed"`

	PutOnTableInitialSpeed		float64 	`json:"PutOnTableInitialSpeed"`
	PutOnTableAcceleration 		float64 	`json:"PutOnTableAcceleration"`
	PutOnTableDeceleration 		float64 	`json:"PutOnTableDeceleration"`
	PutOnTableSpeed 			float64 	`json:"PutOnTableSpeed"`

	BlowingHeight 				float64 	`json:"BlowingHeight"`
	AdjustmentTimes 			int32 		`json:"AdjustmentTimes"`
}

type ParamSendMaterialTable struct {
	Id 					int64 		

	SendMaterialSpeed 	float64 	`xorm:"SendMaterialSpeed"`
	StopDelay 			float64 	`xorm:"StopDelay"`

	FitBenchmarkX 		float64 	`xorm:"FitBenchmarkX"`
	FitBenchmarkY 		float64 	`xorm:"FitBenchmarkY"`
	FitBenchmarkZ 		float64 	`xorm:"FitBenchmarkZ"`
	FitBenchmarkR 		float64 	`xorm:"FitBenchmarkR"`

	CameraHeight 		float64 	`xorm:"CameraHeight"`
	FitCompensationX 	float64 	`xorm:"FitCompensationX"`
	FitCompensationY 	float64 	`xorm:"FitCompensationY"`
	FitCompensationZ 	float64 	`xorm:"FitCompensationZ"`
	FitCompensationR 	float64 	`xorm:"FitCompensationR"`

	RemoveFitInitialSpeed 		float64 	`xorm:"RemoveFitInitialSpeed"`
	RemoveFitAcceleration 		float64 	`xorm:"RemoveFitAcceleration"`
	RemoveFitDeceleration 		float64 	`xorm:"RemoveFitDeceleration"`
	RemoveFitSpeed 				float64 	`xorm:"RemoveFitSpeed"`

	ReturnInitialSpeed 			float64 	`xorm:"ReturnInitialSpeed"`
	ReturnAcceleration 			float64 	`xorm:"ReturnAcceleration"`
	ReturnDeceleration 			float64 	`xorm:"ReturnDeceleration"`
	ReturnSpeed 				float64 	`xorm:"ReturnSpeed"`

	GotoPhotoInitialSpeed 		float64 	`xorm:"GotoPhotoInitialSpeed"`
	GotoPhotoAcceleration 		float64 	`xorm:"GotoPhotoAcceleration"`
	GotoPhotoDeceleration 		float64 	`xorm:"GotoPhotoDeceleration"`
	GotoPhotoSpeed 				float64 	`xorm:"GotoPhotoSpeed"`

	FitInitialSpeed 		float64 	`xorm:"FitInitialSpeed"`
	FitAcceleration 		float64 	`xorm:"FitAcceleration"`
	FitDeceleration 		float64 	`xorm:"FitDeceleration"`
	FitSpeed 				float64 	`xorm:"FitSpeed"`

	PutOnTableInitialSpeed		float64 	`xorm:"PutOnTableInitialSpeed"`
	PutOnTableAcceleration 		float64 	`xorm:"PutOnTableAcceleration"`
	PutOnTableDeceleration 		float64 	`xorm:"PutOnTableDeceleration"`
	PutOnTableSpeed 			float64 	`xorm:"PutOnTableSpeed"`

	BlowingHeight 				float64 	`xorm:"BlowingHeight"`
	AdjustmentTimes 			int32 		`xorm:"AdjustmentTimes"`
}