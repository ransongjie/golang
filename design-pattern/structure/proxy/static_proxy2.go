package main

import "fmt"

// static proxy

type Wedding interface {
	wed()
}

type BrideGroom struct{}

type WedCompany struct {
	brideGroom *BrideGroom
	// brideGroom BrideGroom
}

func (bg *BrideGroom) wed() {
	fmt.Println("新郎新娘结婚")
}

func (wc *WedCompany) wed() {
	fmt.Println("婚庆公司代理布置婚礼现场")
	wc.brideGroom.wed()
	fmt.Println("婚庆公司代理收拾婚礼现场")
}

func NewProxy() *WedCompany {
	return &WedCompany{&BrideGroom{}}
}

func main() {
	proxy := NewProxy()
	proxy.wed()
}
