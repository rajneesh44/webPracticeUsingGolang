package main 

import (
	"fmt"
	"net"
	"io"
	"log"
)

func main(){
	li, er:= net.Listen("tcp", ":8080")
	if err!=nil{
		log.Panic(err)
	}
	defer li.Close()

	for{
		conn, err:= li.Accept()
		if err!=nil{
			log.Println(err)
		}

		io.WriteString(conn, "\n Hello from TCP\n")
		fmt.Fprintln(conn,"How is you rdays")
		fmt.Fprintf(conn, "%v", "well I hope")

		conn.Close()
	}
}
