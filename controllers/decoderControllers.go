package controllers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/buabaj/voters-id-reader/utilities"
	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Welcome to the voters ID reader API",
	})
}

func ReadID(c *gin.Context) {
	// get the file from the request
	file, _ := c.FormFile("file")
	// save the file to the server
	c.SaveUploadedFile(file, file.Filename)
	// decode the file
	results, err := utilities.DecodeQR(file.Filename)
	os.Remove(file.Filename)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "error decoding the file",
		})
		return
	}
	// send the results back to the client
	surname, firstName, gender, dateOfBirth, pollingStationCode, dateOfRegistration, idNumber := results[1], results[2], results[3], results[4], results[5], results[6], results[7]
	dateValue, _ := time.Parse("20060102", dateOfBirth)
	dateOfBirth = dateValue.Format("2006-01-02")
	dateValue, _ = time.Parse("20060102", dateOfRegistration)
	dateOfRegistration = dateValue.Format("2006-01-02")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "ID successfully decoded",
		"data": gin.H{
			"surname":            strings.Title(strings.ToLower(surname)),
			"firstName":          strings.Title(strings.ToLower(firstName)),
			"gender":             strings.Title(strings.ToLower(gender)),
			"dateOfBirth":        dateOfBirth,
			"pollingStationCode": pollingStationCode,
			"dateOfRegistration": dateOfRegistration,
			"idNumber":           idNumber,
		},
	})

}
