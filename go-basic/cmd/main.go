package main

import (
	"com.xcrj/golang/go-basic/data_type"
	"com.xcrj/golang/go-basic/operator"
	"com.xcrj/golang/go-basic/condition"
	"com.xcrj/golang/go-basic/oop"
	"com.xcrj/golang/go-basic/err"
)

func main() {
	// dataType()
	// operators()
	// conditions()
	// oops()
	errs()
}

func dataType() {
	// data_type.RunDataType()
	// data_type.RunDefaultValue()
	// data_type.RunTypeLen()
	// data_type.RunSliceDetail()
	data_type.RunMapDetail()
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