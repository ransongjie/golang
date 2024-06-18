package main

import "fmt"

type EUSocket struct{}
type ChargerAdapter struct{}
type ChinaCharger struct{}

func (eus *EUSocket) charge() {
	fmt.Println("插入欧洲插座充电")
}

func (ca *ChargerAdapter) charge() {
	fmt.Println("我是充电适配器")
	(&EUSocket{}).charge()
}

func (cc *ChinaCharger) charge() {
	fmt.Println("我是中国充电器")
	(&ChargerAdapter{}).charge()
}

func main() {
	chinaCharger := &ChinaCharger{}
	chinaCharger.charge()
}
