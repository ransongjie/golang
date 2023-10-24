package oop

//方法的聚合体 》类型的聚合体
import (
    "fmt"
)

type Phone interface {
    call()
}

type HuaweiPhone struct {
}

func (huaweiPhone HuaweiPhone) call() {
    fmt.Println("I am Huawei, I can call you!")
}