package main

// define Struct
type CxcxcPerson struct {
	Name string
	Age int

}

// define struct's method (no return)
func (cxcxc *CxcxcPerson) ChangeName(newName string){
	cxcxc.Name = newName
}

// define struct's method (has return)
func (cxcxc CxcxcPerson) EvalAge() string {

	if cxcxc.Age > 65 {
		return "已退休"
	}else {
		return "就業中"
	}
}

// can not change instance's attribute example
func (cxcxc CxcxcPerson) CanNotChangeAge(newAge int)  {

	cxcxc.Age=newAge
}

/*

func main(){

	// 生成一個實例，並打印出樣貌
	demoCxcxcPerson := CxcxcPerson{Name:"cxcxc",Age:99}
	fmt.Println(demoCxcxcPerson )

	// 使用實例的ChangeName，並打印，觀察結果是否有變更
	demoCxcxcPerson.ChangeName("lbh")
	fmt.Println(demoCxcxcPerson )

	// 使用實例的CanNotChangeAge，並打印，觀察結果是否有變更
	demoCxcxcPerson.CanNotChangeAge(80)
	fmt.Println(demoCxcxcPerson )

	// 調度方法，展示真的有傳回一個字串。
	fmt.Println(demoCxcxcPerson.EvalAge())

}

*/

