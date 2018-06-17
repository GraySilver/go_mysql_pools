package main

import ("flag"
		"server"
)

func main()  {
	host := flag.String("h","127.0.0.1","http listen host.")
	port := flag.String("p","8080","http listen port.")
	connect := flag.String("c","","MySQL connect engine.")
	maxOpenConns := flag.Int("maxOpenConns",5,"MySQL maxOpenConns")
	maxIdleConns := flag.Int("maxIdleConns",5,"MySQL maxIdleConns")
	flag.Parse()
	server.StartHttpServer(*host,*port,*connect,*maxOpenConns,*maxIdleConns)

}