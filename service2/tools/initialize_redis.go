// command/testdata.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"service2/models"
	"time"

	"github.com/go-redis/redis"
)

func ClientRedis() *redis.Client {
	ENV_REDIS_HOST := os.Getenv("REDIS_HOST")
	if len(ENV_REDIS_HOST) == 0 {
		ENV_REDIS_HOST = "localhost"
	}

	ENV_REDIS_PORT := os.Getenv("REDIS_PORT")
	if len(ENV_REDIS_PORT) == 0 {
		ENV_REDIS_PORT = "6379"
	}

	return redis.NewClient(&redis.Options{
		Addr: ENV_REDIS_HOST + ":" + ENV_REDIS_PORT,
		DB:   0,
	})
}

func main() {
	var userList []models.User
	userList = append(userList, models.User{ID: "3810", Name: "Mahatma Gandhi", Gender: "male", Born: "India"})
	userList = append(userList, models.User{ID: "8383", Name: "Liam Neeson", Gender: "male", Born: "United Kingdom"})
	userList = append(userList, models.User{ID: "4281", Name: "Chiaki Kuriyama", Gender: "female", Born: "Japan"})
	userList = append(userList, models.User{ID: "8779", Name: "Hande Erçel", Gender: "female", Born: "Turkey"})
	userList = append(userList, models.User{ID: "0473", Name: "Noah Beck", Gender: "male", Born: "United States"})
	userList = append(userList, models.User{ID: "0369", Name: "Han Jimin", Gender: "female", Born: "South Korea"})
	userList = append(userList, models.User{ID: "9948", Name: "Shekhar Mehta ", Gender: "male", Born: "Ugandan"})
	userList = append(userList, models.User{ID: "4715", Name: "Emma Stone", Gender: "female", Born: "United States"})
	userList = append(userList, models.User{ID: "7808", Name: "Penélope Cruz", Gender: "female", Born: "Spain"})
	userList = append(userList, models.User{ID: "1120", Name: "Coco Martin", Gender: "male", Born: "Philippines"})
	userList = append(userList, models.User{ID: "2993", Name: "Neymar da Silva Santos Júnior", Gender: "male", Born: "Brazil"})

	client := ClientRedis()
	client.FlushAll().Result()

	for _, v := range userList {
		profile := map[string]string{
			"Name":   v.Name,
			"Gender": v.Gender,
			"Born":   v.Born,
		}
		serialize, _ := json.Marshal(profile)

		key := v.ID
		//fmt.Println("key: ", key)
		err := client.Set(key, serialize, time.Hour*24).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	log.Println("Redis initialization was completed successfully.")
}
