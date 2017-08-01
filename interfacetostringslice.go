package main

import (
	"fmt"
)

func main() {
	type interfaceToStringSliceStruct struct {
		interfaceToStringSlice interface{}
	}

	interfaceToStringSliceStrct := interfaceToStringSliceStruct{}

	stringSlice := []string{"a", "b"}

	interfaceToStringSliceStrct.interfaceToStringSlice = stringSlice

	result := convertInterfaceToSliceString(interfaceToStringSliceStrct.interfaceToStringSlice)

	fmt.Println("interfaceToStringSliceStrct.interfaceToStringSlice -> ", interfaceToStringSliceStrct.interfaceToStringSlice)
	fmt.Println("result -> ", result)
	fmt.Println("result[0] -> ", result[0])
	fmt.Println("result[1] -> ", result[1])
}

func convertInterfaceToSliceString (interfaceToConvert interface{}) []string{
	a := interfaceToConvert.([]string)

	fmt.Println("a -> ", a)

	return a
}
