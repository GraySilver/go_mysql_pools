package test

import "src"

func main()  {
	StartHttpServer("0.0.0.0","8080","root:chenluyao@tcp(localhost:3306)/data?charset=utf8",5,5)
	}
