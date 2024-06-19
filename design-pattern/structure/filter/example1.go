package main

import "fmt"

type Customer struct {
	name   string
	gender int  //1 male, 2 female
	single bool //true single, false non-single
}

type Creteria interface {
	filter(customers [](*Customer)) [](*Customer)
}

type IsMale struct{}
type IsFemale struct{}
type IsSingle struct{}

func (m *IsMale) filter(customers [](*Customer)) [](*Customer) {
	males := [](*Customer){}
	for i := 0; i < len(customers); i++ {
		if customers[i].gender == 1 {
			males = append(males, customers[i])
		}
	}
	return males
}

func (f *IsFemale) filter(customers [](*Customer)) [](*Customer) {
	females := [](*Customer){}
	for i := 0; i < len(customers); i++ {
		if customers[i].gender == 2 {
			females = append(females, customers[i])
		}
	}
	return females
}

func (s *IsSingle) filter(customers [](*Customer)) [](*Customer) {
	singles := [](*Customer){}
	for i := 0; i < len(customers); i++ {
		if customers[i].single {
			singles = append(singles, customers[i])
		}
	}
	return singles
}

type And struct {
	criteria1 Creteria
	criteria2 Creteria
}

type Or struct {
	criteria1 Creteria
	criteria2 Creteria
}

func (a *And) filter(customers [](*Customer)) [](*Customer) {
	//先用第1个标准过滤，再用第2个标准过滤
	customers1 := a.criteria1.filter(customers)
	customers2 := a.criteria2.filter(customers1)
	return customers2
}

func (a *Or) filter(customers [](*Customer)) [](*Customer) {
	//先分别使用第1个标准，第2个标准过滤，再求并集
	customers1 := a.criteria1.filter(customers)
	customers2 := a.criteria2.filter(customers)
	customers3 := [](*Customer){}
	customers3 = append(customers3, customers1...)
	customers3 = append(customers3, customers2...)
	return customers3
}

func main() {
	c1 := &Customer{"c1", 1, true}
	c2 := &Customer{"c2", 1, false}
	c3 := &Customer{"c3", 2, true}
	c4 := &Customer{"c4", 1, true}
	c5 := &Customer{"c5", 2, true}
	c6 := &Customer{"c6", 2, false}

	customers := [](*Customer){c1, c2, c3, c4, c5, c6}
	single := &IsSingle{}
	male := &IsMale{}
	female := &IsFemale{}

	show(single.filter(customers))
	and := &And{single, male}
	show(and.filter(customers))
	or := &Or{single, female}
	show(or.filter(customers))
}

func show(customers [](*Customer)) {
	for i := 0; i < len(customers); i++ {
		fmt.Print(customers[i])
	}
	fmt.Println()
}
