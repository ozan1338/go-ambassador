package main

// import (
// 	"context"
// 	"encoding/json"
// 	"go-ambassador/src/database"
// 	"go-ambassador/src/models"
// 	"go-ambassador/src/services"

// 	"github.com/go-redis/redis/v8"
// )

// func main() {
// 	database.Connect()
// 	database.SetupRedis()
// 	services.Setup()

// 	ctx := context.Background()

// 	resp, err := services.UserService.Get("users","")

// 	if err != nil {
// 		panic(err)
// 	}

// 	var users []models.User

// 	json.NewDecoder(resp.Body).Decode(&users)

// 	for _, user := range users {
// 		if user.IsAmbassador {

// 			ambassador := models.Ambassador(user)
// 			ambassador.CalculateRevenue(database.DB)

// 			database.Cache.ZAdd(ctx, "rankings", &redis.Z{
// 				Score:  ambassador.Revenue,
// 				Member: user.Name(),
// 			})
// 		}
// 	}
// }
