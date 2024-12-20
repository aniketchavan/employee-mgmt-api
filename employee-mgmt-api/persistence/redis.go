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

// Initialize the Redis Client
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
	// TODO: Make expiry time configurable - take value from config
	return RedisClient.Set(ctx, employee.Email, string(employeeStr), 5*time.Minute).Err()
}

func DeleteEmployeeInCache(email string) error {
	// Delete employee data in Redis
	return RedisClient.Del(ctx, email).Err()
}
