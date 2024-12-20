// database/redis.go
package persistence

import (
	"context"
	"encoding/json"
	"time"

	datamodel "example.com/employee-mgmt/models"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis(addr string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func StoreEmployeeInCache(employee datamodel.Employee) error {
	// Store employee data in Redis
	employeeStr, err := json.Marshal(employee)
	if err != nil {
		return err
	}
	return RedisClient.Set(ctx, employee.Email, string(employeeStr), 5*time.Minute).Err() // Use Email as the key
}

func DeleteEmployeeInCache(email string) error {
	// Delete employee data in Redis
	return RedisClient.Del(ctx, email).Err() // Use Email as the key
}
