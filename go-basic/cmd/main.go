package main

import (
	"com.xcrj/golang/go-basic/concurrent"
	"com.xcrj/golang/go-basic/condition"
	"com.xcrj/golang/go-basic/data_type"
	"com.xcrj/golang/go-basic/err"
	"com.xcrj/golang/go-basic/generic"
	"com.xcrj/golang/go-basic/io"
	"com.xcrj/golang/go-basic/oop"
	"com.xcrj/golang/go-basic/operator"
	"com.xcrj/golang/go-basic/random"
	"com.xcrj/golang/go-basic/reflect"
	"com.xcrj/golang/go-basic/test"
)

func main() {
	// dataType()
	// operators()
	// conditions()
	// oops()
	// errs()
	tests()
	// generics()
	// concurrents()
	// reflects()
	// testRandom()
	// ioo()
}

func ioo() {
	io.RunIO()
}

func dataType() {
	// data_type.RunDataType()
	// data_type.RunDefaultValue()
	// data_type.RunTypeLen()
	// data_type.RunSliceDetail()
	// data_type.RunMapDetail()
	// data_type.RunEqual()
	data_type.RunTime()
}

func operators() {
	operator.RunOperator()
}

func conditions() {
	condition.RunCondition()
}

func oops() {
	oop.RunStruct()
}

func errs() {
	err.RunErr()
}

func tests() {
	test.RunTest()
}

func generics() {
	generic.RunGeneric()
}

func concurrents() {
	// concurrent.RunChan()
	// concurrent.RunContext()
	concurrent.RunAtomic()
}

func reflects() {
	reflect.RunReflect()
}

func testRandom() {
	random.RunRandom()
}
