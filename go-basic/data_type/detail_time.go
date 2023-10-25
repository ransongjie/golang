package data_type

import(
	"fmt"
	"time"
)

func RunTime(){
	test1()
}

//时间比较
func test1(){
	t1 := time.Now()
	t2 := time.Now()
	fmt.Println(t1.Equal(t2))// false
	fmt.Println(t1.After(t2))// false
	fmt.Println(t1.Before(t2))// true

	s1 := "2015-03-20 08:50:29"
	s2 := "2015-03-21 09:04:25"
	t1, err := time.Parse("2006-01-02 15:04:05", s1)
	t2, err = time.Parse("2006-01-02 15:04:05", s2)
	if err == nil && t1.Before(t2) {
		fmt.Println("true")// true
	}
}