package main

func init() {
	println("init func()")
}

func main() {
	// var count int64
	// var name string
	// var arr []string{
	// 	"1",
	// 	"2",
	// }

	// var slice []string //จะไม่สามารถ append ได้
	// slice := []string{}
	// slice = append(slice, "aa")
	// fmt.Printf("%+v", slice)

	// mapVal := map[string]string{ //element ต้องเป็น string เท่านั้น
	// 	"key": "Hello world",
	// }

	// mapVal := map[string]interface{}{ //type:interface ใช้ค่าอะไรก็ได้
	// 	"key":    "Hello world",
	// 	"number": 5,
	// }

	// mapVal := map[string]interface{}{ //type:interface ใช้ค่าอะไรก็ได้
	// 	"key":    "Hello world",
	// 	"number": 5,
	// 	"nestedMap": map[string]interface{}{
	// 		"childKey": 2,
	// 	},
	// }

	// type valMap struct{ // กำหนดโครงสร้าง struct ของตัวแปร
	// 	Key string
	// 	Number int
	// 	NestedMap map[string]string
	// }

	// fmt.Printf("%+v", mapVal)

	// printText("World", 12345)
	// result := Prinln(0)
	// fmt.Prinln(result)
}

// func printText(text string, number int) {
// 	fmt.Println("Hello", text, number)
// }

// func Prinln(a ...string) { // for range
// 	for _, e := range a { // _ คือการ ignore ตัว param ที่เราไม่ต้องการใช้
// 		print(e + " ")
// 	}
// }

// func Prinln(num int, a ...string) { // for range
// 	for _, e := range a { // _ คือการ ignore ตัว param ที่เราไม่ต้องการใช้
// 		print(e + " ")
// 	}
// }

// func Prinln(num int, a ...string) {
// 	for _, e := range a {
// 		print(e + " ")
// 	}
// 	return "val"
// }

// func Hello(name, surname string) {
// 	if i == 0 {
// 		fmt.Println(false)
// 	}
// }

// func Hello(name, surname string) {
// 	if b := i + 5; b < 5 {
// 		fmt.Println(false)
// 	}
// }
