package handler

import (
	"fmt"
	"go-restful-app/models"
	"go-restful-app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) { //This function registers all API routes with the provided Gin engine (r).
	fmt.Println("hello2")
	r.GET("/fixlets", getAllFixlets)
	r.GET("/fixlets/:id", getFixletByID)
	r.POST("/fixlets", createFixlet)
	r.PUT("/fixlets/:id", updateFixlet)
	r.DELETE("/fixlets/:id", deleteFixlet)
}

func getAllFixlets(c *gin.Context) { //Fetches all fixlet records from the service layer.
	fmt.Println("hello1")
	records, err := service.GetAllRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} //If an error occurs, it responds with a 500 Internal Server Error..If successful, it responds with a 200 OK
	c.JSON(http.StatusOK, records)
}

func getFixletByID(c *gin.Context) {
	id := c.Param("id")
	record, err := service.GetRecordByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func createFixlet(c *gin.Context) {
	var record models.Record
	if err := c.BindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newRecord, err := service.CreateRecord(record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newRecord)
}

func updateFixlet(c *gin.Context) {
	id := c.Param("id")
	var record models.Record
	if err := c.BindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedRecord, err := service.UpdateRecord(id, record)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRecord)
}

func deleteFixlet(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteRecord(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fixlet deleted"})
}
