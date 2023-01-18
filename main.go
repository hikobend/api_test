package main

import "github.com/hikobend/qiita/controller"

func main() {
	r := controller.GetRouter()
	r.Run()
}
