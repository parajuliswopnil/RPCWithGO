package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type User struct {
	Name string
	Password string
}

func main(){
	var reply User
	var authStat string
	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil{
		log.Fatal("error establishing connection")
	}

	user1 := User{Name: "swopnil", Password: "swop2233"}
	
	client.Call("API.AddUser", user1, &reply)
	fmt.Println(reply)

	user1 = User{Name: "swopnil", Password: "swop223"}

	client.Call("API.LoginUser", user1, &authStat)
	fmt.Println(authStat)
}


