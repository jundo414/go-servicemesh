package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"google.golang.org/grpc"

	pb "service1/proto"
)

func main() {
	http.HandleFunc("/users/", UsersHandler)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Gender string `json:"Gender"`
	Born   string `json:"Born"`
}

type ResponseAll struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
	Users   []User
}

type Response struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
	Users   *User
}

var (
	_users   []User
	_status  string
	_message string
)

type UserRequest struct {
	Name   string `json:"Name"`
	Gender string `json:"Gender"`
	Born   string `json:"Born"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		sub := strings.TrimPrefix(r.URL.Path, "/users")
		_, id := filepath.Split(sub)
		if id != "" {
			ShowUsers(w, r, id)
		}
		/*
			if id != "" {
				ShowUsers(w, r, id)
			} else {
				ListUsers(w, r)
			}
		*/
	case "POST":
		sub := strings.TrimPrefix(r.URL.Path, "/users")
		_, id := filepath.Split(sub)
		if id != "" {
			body := make([]byte, r.ContentLength)
			r.Body.Read(body)
			var userRequest UserRequest
			json.Unmarshal(body, &userRequest)

			CreateUsers(w, r, id, userRequest.Name, userRequest.Gender, userRequest.Born)
		}
	case "PUT":
		sub := strings.TrimPrefix(r.URL.Path, "/users")
		_, id := filepath.Split(sub)
		if id != "" {
			body := make([]byte, r.ContentLength)
			r.Body.Read(body)
			var userRequest UserRequest
			json.Unmarshal(body, &userRequest)

			UpdateUsers(w, r, id, userRequest.Name, userRequest.Gender, userRequest.Born)
		}
	case "DELETE":
		sub := strings.TrimPrefix(r.URL.Path, "/users")
		_, id := filepath.Split(sub)
		if id != "" {
			DeleteUsers(w, r, id)
		}
	}
}

/*
func ListUsers(w http.ResponseWriter, r *http.Request) {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer con.Close()
	client := pb.NewUserClient(con)
	req := &pb.UserRequest{
		Handling: "LIST",
	}
	res, _ := client.User(context.TODO(), req)

	_users := []User{
		User{res.Id, res.Name, res.Gender, res.Born},
	}
	if res.Status {
		_status = "Success"
	} else {
		_status = "Failed"
	}
	_message = res.Message
	response := ResponseAll{_status, _message, _users}
	json.NewEncoder(w).Encode(response)
}
*/
func CreateUsers(w http.ResponseWriter, r *http.Request, id string, name string, gender string, born string) {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer con.Close()
	client := pb.NewUserClient(con)
	req := &pb.UserRequest{
		Handling: "CREATE",
		Id:       id,
		Name:     name,
		Gender:   gender,
		Born:     born,
	}
	res, _ := client.User(context.TODO(), req)

	_users := []User{
		User{res.Id, res.Name, res.Gender, res.Born},
	}
	if res.Status {
		_status = "Success"
	} else {
		_status = "Failed"
	}
	_message = res.Message

	response := ResponseAll{_status, _message, _users}
	json.NewEncoder(w).Encode(response)
}

func ShowUsers(w http.ResponseWriter, r *http.Request, id string) {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer con.Close()
	client := pb.NewUserClient(con)
	req := &pb.UserRequest{
		Handling: "SHOW",
		Id:       id,
	}
	res, _ := client.User(context.TODO(), req)

	_users := []User{
		User{res.Id, res.Name, res.Gender, res.Born},
	}
	if res.Status {
		_status = "Success"
	} else {
		_status = "Failed"
	}
	_message = res.Message

	response := ResponseAll{_status, _message, _users}
	json.NewEncoder(w).Encode(response)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request, id string, name string, gender string, born string) {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer con.Close()
	client := pb.NewUserClient(con)
	req := &pb.UserRequest{
		Handling: "UPDATE",
		Id:       id,
		Name:     name,
		Gender:   gender,
		Born:     born,
	}
	res, _ := client.User(context.TODO(), req)

	_users := []User{
		User{res.Id, res.Name, res.Gender, res.Born},
	}
	if res.Status {
		_status = "Success"
	} else {
		_status = "Failed"
	}
	_message = res.Message

	response := ResponseAll{_status, _message, _users}
	json.NewEncoder(w).Encode(response)
}
func DeleteUsers(w http.ResponseWriter, r *http.Request, id string) {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer con.Close()
	client := pb.NewUserClient(con)
	req := &pb.UserRequest{
		Handling: "DELETE",
		Id:       id,
	}
	res, _ := client.User(context.TODO(), req)

	_users := []User{
		User{res.Id, res.Name, res.Gender, res.Born},
	}
	if res.Status {
		_status = "Success"
	} else {
		_status = "Failed"
	}
	_message = res.Message

	response := ResponseAll{_status, _message, _users}
	json.NewEncoder(w).Encode(response)
}
