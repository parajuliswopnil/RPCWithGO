package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type User struct {
	Name string
	Password string
}

type API int

var UserList []User

func (r *API) AddUser(userdata User, reply *User) error {
	log.Printf("registering user")
	UserList = append(UserList, userdata)

	*reply = userdata
	return nil
}

func (r API) LoginUser(userdata User, reply *string) error {
	log.Println("logging in user")
	authStatus := "success"
	auth := authenticateUser(userdata)
	if auth == "illegal"{
		log.Fatal("Couldnot authenticate the user")
	}
	*reply = authStatus;
	return nil
}

func authenticateUser(userdata User) string{
	log.Println("authenticating")
	authenticate := "illegal"
	for _, users := range UserList {
		log.Println(users.Name)
		log.Println(users.Password)
		if users.Name == userdata.Name && users.Password == userdata.Password{
			authenticate = "legal"
			break
		}
	}
	return authenticate
}

func main(){
	api := new(API)
	err := rpc.Register(api)

	if err != nil{
		log.Fatal("error in registration")
	}

	rpc.HandleHTTP();
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("error in listening")
	}

	log.Printf("Serving on localhost:4040")

	err = http.Serve(listener, nil)

	if err != nil{
		log.Fatal("error in serving")
	}
}

