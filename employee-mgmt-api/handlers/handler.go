// handlers/employee.go
package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	datamodel "example.com/employee-mgmt/models"
	"example.com/employee-mgmt/persistence"
	"example.com/employee-mgmt/utils"

	"github.com/gin-gonic/gin"
)

func UploadEmployeeData(c *gin.Context) {
	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Save the file to a temporary location
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	// Parse the Excel file
	employees, err := utils.ParseExcelFile(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Store each employee in MySQL and Redis
	for _, employee := range employees {
		if err := persistence.StoreEmployee(employee); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to store employee in MySQL"})
			return
		}
		if err := persistence.StoreEmployeeInCache(employee); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to store employee in Redis"})
			return
		}
	}

	// Respond with the parsed data
	c.JSON(http.StatusOK, gin.H{"employees": employees})

	// Optionally, you can delete the file after processing
	os.Remove(file.Filename)
}

// GetEmployee handles fetching employee data
func GetEmployee(c *gin.Context) {
	email := c.Param("email")

	// Check Redis cache first
	employeeDetail, err := persistence.RedisClient.Get(persistence.RedisClient.Context(), email).Result()

	var employee datamodel.Employee
	if err == nil {
		// If found in Redis, return the cached data
		err = json.Unmarshal([]byte(employeeDetail), &employee)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch employee in Redis"})
		}

		c.JSON(http.StatusOK, gin.H{"email": email, "employee": employee, "source": "cache"})
		return
	}

	// If not found in Redis, fetch from MySQL
	if err := persistence.DB.Where("email = ?", email).First(&employee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Store the fetched employee in Redis for future requests
	if err := persistence.StoreEmployeeInCache(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to store employee in Redis"})
		return
	}

	// Respond with the employee data
	// TODO: Store data in JSON format
	c.JSON(http.StatusOK, gin.H{"email": employee.Email, "employee": employee, "source": "database"})
}

// UpdateEmployee handles updating an employee record
func UpdateEmployee(c *gin.Context) {
	email := c.Param("email")
	var updatedEmployee datamodel.Employee

	// Bind the JSON body to the updatedEmployee struct
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update the employee in MySQL
	if err := persistence.DB.Model(&datamodel.Employee{}).Where("email = ?", email).Updates(updatedEmployee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update employee in MySQL"})
		return
	}

	// Update the employee in Redis
	if err := persistence.StoreEmployeeInCache(updatedEmployee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update employee in Redis"})
		return
	}

	// Respond with the updated employee data
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully", "employee": updatedEmployee})
}

// UpdateEmployee handles updating an employee record
func DeleteEmployee(c *gin.Context) {
	email := c.Param("email")

	// Update the employee in MySQL
	if err := persistence.DB.Model(&datamodel.Employee{}).Where("email = ?", email).Delete("email = ?").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete employee in MySQL"})
		return
	}

	// Update the employee in Redis
	if err := persistence.DeleteEmployeeInCache(email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete employee in Redis"})
		return
	}

	// Respond with the updated employee data
	c.JSON(http.StatusOK, gin.H{"message": "Employee delelted successfully", "employee-email": email})
}
