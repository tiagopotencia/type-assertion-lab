package main

import "fmt"

type structResult struct{
	name string
	age int
}

type struct2 struct {
	title string
	number int
}

func main() {
	var interfaceToBeConverted interface{}
	var interfaceToBeConvertedWithAnonymousStruct interface{}
	var interfaceToBeConvertedWithNil interface{} = nil
	var interfaceToBeConvertedWithNilConvertedToAPointer interface{}

	strc := structResult{name:"tp", age:23}

	interfaceToBeConverted = strc

	interfaceToBeConvertedWithAnonymousStruct = struct {
		name string
		age int
	}{
		name: "a",
		age: 2,
	}

	var structResultPointer *structResult
	interfaceToBeConvertedWithNilConvertedToAPointer = structResultPointer

	convertedInterface := convertInterfaceToStruct(interfaceToBeConverted)
	convertedInterfaceToStruct2 := convertInterfaceToStruct2(interfaceToBeConverted)
	convertedInterfaceToAonymousStruct := convertInterfaceToAnonymousStruct(interfaceToBeConvertedWithAnonymousStruct)
	convertedInterfaceWithNil := convertInterfaceToStructWithNil(interfaceToBeConvertedWithNil)
	convertedInterfaceToBeConvertedWithNilConvertedToAPointer := convertInterfaceToStructWithNilConvertedToAPointer(interfaceToBeConvertedWithNilConvertedToAPointer)

	fmt.Printf("interfaceToBeConverted type -> %T\n", interfaceToBeConverted)
	fmt.Println("interfaceToBeConverted -> ", interfaceToBeConverted)
	fmt.Printf("convertedInterface type -> %T\n", convertedInterface)
	fmt.Println("convertedInterface ->", convertedInterface)
	fmt.Printf("convertedInterfaceToStruct2 type -> %T\n", convertedInterfaceToStruct2)
	fmt.Println("convertedInterfaceToStruct2 ->", convertedInterfaceToStruct2)
	fmt.Printf("convertedInterfaceToAonymousStruct type -> %T\n", convertedInterfaceToAonymousStruct)
	fmt.Println("convertedInterfaceToAonymousStruct ->", convertedInterfaceToAonymousStruct)
	fmt.Printf("convertedInterfaceWithNil type -> %T\n", convertedInterfaceWithNil)
	fmt.Println("convertedInterfaceWithNil ->", convertedInterfaceWithNil)
	fmt.Printf("convertedInterfaceToBeConvertedWithNilConvertedToAPointer type -> %T\n", convertedInterfaceToBeConvertedWithNilConvertedToAPointer)
	fmt.Println("convertedInterfaceToBeConvertedWithNilConvertedToAPointer ->", convertedInterfaceToBeConvertedWithNilConvertedToAPointer)
}

func convertInterfaceToStruct(interfaceToBeConverted interface{}) structResult{
	strct := structResult{}
	strct = interfaceToBeConverted.(structResult)
	return strct
}

func convertInterfaceToStruct2(interfaceToBeConverted interface{}) struct2{
	strct := structResult{}
	strct = interfaceToBeConverted.(structResult)

	struct2 := struct2{number:strct.age, title: strct.name}

	return struct2
}

func convertInterfaceToAnonymousStruct(interfaceToBeConverted interface{}) structResult{
	strct := structResult{}
	strct = interfaceToBeConverted.(struct {
		name string
		age int
	})

	return strct
}

func convertInterfaceToStructWithNil(interfaceToBeConverted interface{}) *structResult{
	strct := structResult{}
	strct, ok := interfaceToBeConverted.(structResult)

	//The type assertion failed but it doesn't means that interfaceToBeConverted is nil
	if ok == false{
		if interfaceToBeConverted == nil{
			return nil
		} else {
			panic("I don't know what type to assert!! :(")
		}
	}

	return &strct
}

func convertInterfaceToStructWithNilConvertedToAPointer(interfaceToBeConverted interface{}) *structResult{
	var strct *structResult
	strct = interfaceToBeConverted.(*structResult)

	return strct
}