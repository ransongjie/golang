package main

import (
	"com.xcrj/golang/go-basic/data_type"
	"com.xcrj/golang/go-basic/operator"
	"com.xcrj/golang/go-basic/condition"
	"com.xcrj/golang/go-basic/oop"
	"com.xcrj/golang/go-basic/err"
	"com.xcrj/golang/go-basic/test"
)

func main() {
	dataType()
	// operators()
	// conditions()
	// oops()
	// errs()
	// tests()
}

func dataType() {
	// data_type.RunDataType()
	// data_type.RunDefaultValue()
	// data_type.RunTypeLen()
	// data_type.RunSliceDetail()
	// data_type.RunMapDetail()
	data_type.RunEqual()
	// data_type.RunTime()
}

func operators(){
	operator.RunOperator()
}

func conditions(){
	condition.RunCondition()
}

func oops(){
	oop.RunStruct()
}

func errs(){
	err.RunErr()
}

func tests(){
	test.RunTest()
}