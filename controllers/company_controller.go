package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-microservice/db"
	"github.com/yourusername/go-microservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

// GetCompanies handles the GET /companies route
func GetCompanies(c *gin.Context) {
	var companies []models.Company
	cursor, err := db.CompanyCollection.Find(context.Background(), bson.D{{}}, options.Find())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var company models.Company
		if err := cursor.Decode(&company); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		companies = append(companies, company)
	}

	c.JSON(http.StatusOK, companies)
}

// CreateCompany handles the POST /companies route
func CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.CompanyCollection.InsertOne(context.Background(), company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// UpdateCompany handles the PUT /companies/:id route
func UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", company.Name}, {"address", company.Address}, {"website", company.Website}}}}

	result, err := db.CompanyCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated"})
}

// PatchCompany handles the PATCH /companies/:id route
func PatchCompany(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", company.Name}, {"address", company.Address}, {"website", company.Website}}}}

	result, err := db.CompanyCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company patched"})
}

// DeleteCompany handles the DELETE /companies/:id route
func DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	filter := bson.D{{"_id", id}}
	result, err := db.CompanyCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}
