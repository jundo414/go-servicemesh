package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	pb "service2/proto"
	"time"

	"fmt"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

type Profile struct {
	Name   string `json:"Name"`
	Gender string `json:"Gender"`
	Born   string `json:"Born"`
}

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServer(server, &User{})
	server.Serve(listenPort)
}

type User struct{}

func (h *User) User(cts context.Context, req *pb.UserRequest) (*pb.UserProfile, error) {
	client := ClientRedis()
	_status := true

	switch req.Handling {
	/*
		case "LIST":
			_message := "[list user] completed successfully."
			res := pb.UserProfile{Status: _status, Message: _message}
			return &res, nil
	*/
	case "SHOW":
		_message := "[show user] completed successfully."
		key := req.Id
		val, err := client.Get(key).Result()

		jsonBytes := ([]byte)(val)
		data := new(Profile)

		if err == redis.Nil {
			_status = false
			_message := "[show user] key '" + key + "' does not exist."
			res := pb.UserProfile{Status: _status, Message: _message}
			return &res, nil
		} else if err != nil {
			_status = false
			_message := "[show user] failed while Redis(Get) proccess."
			res := pb.UserProfile{Status: _status, Message: _message}
			panic(err)
			return &res, nil
		} else if err := json.Unmarshal(jsonBytes, data); err != nil {
			_status = false
			_message := "[show user] failed while unmarshal proccess."
			res := pb.UserProfile{Status: _status, Message: _message}
			fmt.Println(err)
			return &res, nil
		} else {
			res := pb.UserProfile{
				Id:      req.Id,
				Name:    data.Name,
				Gender:  data.Gender,
				Born:    data.Born,
				Status:  _status,
				Message: _message,
			}
			return &res, nil
		}
	case "CREATE":
		_message := "[create user] completed successfully."
		key := req.Id
		_, err := client.Get(key).Result()
		if err == redis.Nil {
			profile := map[string]string{
				"Name":   req.Name,
				"Gender": req.Gender,
				"Born":   req.Born,
			}
			serialize, _ := json.Marshal(profile)

			//fmt.Println("key: ", key)
			err := client.Set(key, serialize, time.Hour*24).Err()
			if err != nil {
				_status = false
				_message = "[create user] failed while Redis(Set) proccess."
				res := pb.UserProfile{Status: _status, Message: _message}
				fmt.Println(err)
				return &res, nil
			}
			res := pb.UserProfile{
				Id:      req.Id,
				Name:    req.Name,
				Gender:  req.Gender,
				Born:    req.Born,
				Status:  _status,
				Message: _message,
			}
			return &res, nil
		} else if err != nil {
			_status = false
			_message = "[create user] failed while Redis(Set) proccess."
			res := pb.UserProfile{Status: _status, Message: _message}
			panic(err)
			return &res, nil
		} else {
			_status = false
			_message = "[create user] key '" + key + "' exist already. please try another keys."
			res := pb.UserProfile{Status: _status, Message: _message}
			return &res, nil
		}
	case "UPDATE":
		_message := "[update user] completed successfully."
		key := req.Id
		_, err := client.Get(key).Result()
		if err == redis.Nil {
			_status = false
			_message = "[update user] key '" + key + "' does not exist."
			res := pb.UserProfile{Status: _status, Message: _message}
			return &res, nil
		} else if err != nil {
			_status = false
			_message = "[update user] failed while Redis(Get) proccess."
			res := pb.UserProfile{Status: _status, Message: _message}
			panic(err)
			return &res, nil
		} else {
			profile := map[string]string{
				"Name":   req.Name,
				"Gender": req.Gender,
				"Born":   req.Born,
			}
			serialize, _ := json.Marshal(profile)

			//fmt.Println("key: ", key)
			err := client.Set(key, serialize, time.Hour*24).Err()
			if err != nil {
				_status = false
				_message = "[update user] failed while Redis(Set) proccess."
				res := pb.UserProfile{Status: _status, Message: _message}
				fmt.Println(err)
				return &res, nil
			}
			res := pb.UserProfile{
				Id:     req.Id,
				Name:   req.Name,
				Gender: req.Gender,
				Born:   req.Born,
			}
			return &res, nil
		}
	case "DELETE":
		_message := "[delete user] completed successfully."
		key := req.Id
		_, err := client.Get(key).Result()
		if err == redis.Nil {
			_status = false
			_message = "[delete user] key '" + key + "' does not exist."
			res := pb.UserProfile{Status: _status, Message: _message}
			return &res, nil
		} else if err != nil {
			_status = false
			_message = "[delete user] failed while Redis(Get) proccess."
			res := pb.UserProfile{Status: _status, Message: _message}
			panic(err)
			return &res, nil
		} else {
			err := client.Del(key).Err()
			if err != nil {
				_status = false
				_message = "[delete user] failed while Redis(Del) proccess."
				res := pb.UserProfile{Status: _status, Message: _message}
				fmt.Println(err)
				return &res, nil
			}
		}
		res := pb.UserProfile{Status: _status, Message: _message}
		return &res, nil
	}
	res := pb.UserProfile{}
	return &res, nil
}

func ClientRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}
