package main

import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"

	"time"
)

/*

啟用docker-compose，將資料庫環境開啟
docker-compose up -d

下載套件，更新go.mod
require github.com/jinzhu/gorm v1.9.12 // indirect

引用相關資料庫套件
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

創建客戶端
https://gorm.io/docs/connecting_to_the_database.html

設計表格
https://gorm.io/docs/models.html

創造資料庫
創造表格
插入資料
讀取資料
	精準
	多筆
	複合條件
更新資料
刪除資料


*/


// 設計表格
// https://gorm.io/docs/models.html#Declaring-Models
type CxcxcMember struct{
	Name string
	Age int
	Birthday time.Time
	Email string `gorm:"primary_key"`
	Num  int     `gorm:"AUTO_INCREMENT"`
	IgnoreMe     string     `gorm:"-"` // ignore this field
	Family string  `gorm:"column:custom_column_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}


/*

func main() {

	// 先行用資料庫視覺化軟體，建立DB，cxcxc
	// 創造客戶端
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=cxcxc password=changeme sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// =======================================================



	// 創造表格
	//db.CreateTable(&CxcxcMember{})

	// 要求使用自定義表格名
	//db.Table("persons").CreateTable(&CxcxcMember{})



	// =======================================================

	// 新建多筆資料

	//var cxcxcMember1 CxcxcMember = CxcxcMember{Name:"cxcxc1",Age:18, Birthday:time.Now(), Email:"service1@cxcxc.io", IgnoreMe:"cxcxc",Family:"cxcxc-123"}
	//var cxcxcMember2 CxcxcMember = CxcxcMember{Name:"cxcxc2",Age:19, Birthday:time.Now(), Email:"service2@cxcxc.io", IgnoreMe:"cxcxc",Family:"cxcxc-123"}
	//var cxcxcMember3 CxcxcMember = CxcxcMember{Name:"cxcxc3",Age:20, Birthday:time.Now(), Email:"service3@cxcxc.io", IgnoreMe:"cxcxc",Family:"cxcxc-123"}
	//var cxcxcMember4 CxcxcMember = CxcxcMember{Name:"cxcxc4",Age:21, Birthday:time.Now(), Email:"service4@cxcxc.io", IgnoreMe:"cxcxc",Family:"cxcxc-123"}
	//var cxcxcMember5 CxcxcMember = CxcxcMember{Name:"cxcxc5",Age:22, Birthday:time.Now(), Email:"service5@cxcxc.io", IgnoreMe:"cxcxc",Family:"cxcxc-123"}
	//var cxcxcMember6 CxcxcMember = CxcxcMember{Name:"cxcxc6",Age:23, Birthday:time.Now(), Email:"service6@cxcxc.io", IgnoreMe:"cxcxc",Family:"cxcxc-123"}
	//
	////插入資料
	//if db.NewRecord(&cxcxcMember1) ==true {
	//	db.Create(&cxcxcMember1)
	//}
	//
	////插入多筆資料
	//db.Create(&cxcxcMember2)
	//db.Create(&cxcxcMember3)
	//db.Create(&cxcxcMember4)
	//db.Create(&cxcxcMember5)
	//db.Create(&cxcxcMember6)



	// =======================================================



	//// 按照Primary key作查詢，並將找回的內容，塞回物件內
	//var cxcxcMemberFind CxcxcMember = CxcxcMember{ Email:"service2@cxcxc.io", IgnoreMe:"cxcxc",Family:"cxcxc-123"}
	//db.Where(cxcxcMemberFind).Find(&cxcxcMemberFind)
	//fmt.Println(cxcxcMemberFind)
	//
	//// 查詢多筆資料，並塞回Array中
	//var moreCxcxcMember []CxcxcMember
	//db.Where(CxcxcMember{Family:"cxcxc-123"}).Find(&moreCxcxcMember)
	//fmt.Println(moreCxcxcMember)



	// =======================================================

	// 更新資料
	// 更新時，會依照PK去找 資料列
	// save or update
	// update 僅變更部分欄位
	// save 將所有欄位全部回填
	// https://gorm.io/docs/update.html

	// db.Model(cxcxcMemberFind).Update(CxcxcMember{ Name:"cxcxc98"})

	// =======================================================

	// 刪除資料
	// 若無設計軟刪除欄位，則會直接刪除
	// 若無指定PK，則會刪除所有資料


	//WARNING When deleting a record, you need to ensure its primary field has value,
	//and GORM will use the primary key to delete the record, if the primary key field is blank,
	//GORM will delete all records for the model


	// 有PK時，刪除資料正常

	// db.Delete(&cxcxcMemberFind)

	// 不指定PK，導致資料全部被刪除
	// db.Delete(&CxcxcMember{})


}

*/