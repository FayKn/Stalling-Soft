package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "strconv"
)

type registrations struct {
	ID        int    `json:"id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	/*Email      string `json:"Email"`
	Phone      string `json:"Phone"`
	Adress     string `json:"Adress"`
	PostalCode string `json:"PostalCode"`
	City       string `json:"City"`
	Country    string `json:"Country"`
	IBAN       string `json:"IBAN"`*/
}

var registration = []registrations{
	{ID: 1, FirstName: "John", LastName: "Doe"},
	{ID: 2, FirstName: "Jane", LastName: "Doe"},
	{ID: 3, FirstName: "John", LastName: "Smith"},
}

func getRegistrations(c *gin.Context) {
	c.JSON(http.StatusOK, registration)
}

func getRegistrationByID(c *gin.Context) {
	idStr := c.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	for _, a := range registration {
		if a.ID == int(id) {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "registration not found"})
}

func registerRegistrations(c *gin.Context) {
	var newRegistrations registrations
	if err := c.BindJSON(&newRegistrations); err != nil {
		return
	}
	registration = append(registration, newRegistrations)
	c.IndentedJSON(http.StatusCreated, newRegistrations)
}

func main() {
	router := gin.Default()
	router.GET("/registrations", getRegistrations)
	router.GET("/registrations/:id", getRegistrationByID)
	router.POST("/registrations", registerRegistrations)

	router.Run(":8080")
}
