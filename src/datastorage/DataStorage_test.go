package datastorage

import (
	"testing"
	//"time"
	"fmt"
	"log"
)


func Test_StorageProxy(t *testing.T) {
	proxy := CreateStorageProxy("sqlite3", "f:\\db\\test.db")
	fmt.Println(proxy)
	//fmt.Println(proxy.SyncTable(new(MaterialsTable)))
	if err := proxy.SyncTable(new(ProductInformationTable)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}

	pt := new(ProductInformationTable)

	//pt.Id = 10
	pt.Result = true
	pt.AminusB = 10.0

	proxy.Insert(pt)

	//everyOne := make([]*ProductInformationTable, 0)
	proxy.Traversing()
}