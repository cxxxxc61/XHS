package main

import "unicode/utf8"


func Hello() {
	println("Hello, go!")
	println(utf8.RuneCountInString("你好"))
}