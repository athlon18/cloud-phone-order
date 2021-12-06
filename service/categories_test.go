package service

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	data := CategoriesList(1,0)
	fmt.Println(data)
}

func Test1(t *testing.T) {
	//data := reflect.TypeOf(time.Now().Weekday())
	sd, _ := time.ParseDuration("24h")
	fmt.Println(int(time.Now().Add(sd).Weekday()))
}