package main

//import "fmt"

func Defer(){
	for i := 0; i < 5; i++ {
		defer func(){
			println(i)
		}()
	}
}

func Defer1(){
	for i := 0; i < 5; i++ {
		defer func(val int){
			println(val)
		}(i)
	}
}

func Defer2(){
	for i := 0; i < 5; i++ {
		j:=i
		defer func(){
			println(j)
		}()
	}
}