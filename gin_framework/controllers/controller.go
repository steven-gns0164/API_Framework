package controllers

import (
	m "gin_framework/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Error Connecting to DB")
		return
	}

	var student []m.Students
	queryResult := db.Last(&student)
	if queryResult.Error != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Error Retriving Data")
		return
	}

	sendSuccessResponseGIN(c, http.StatusOK, "Success")
}

func InsertNewStudents(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Error Connecting to DB")
		return
	}

	nimStr := c.Query("nim")
	nim, err := strconv.Atoi(nimStr)
	if err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Invalid Value for NIM")
		return
	}
	nama := c.Query("nama")

	student := m.Students{
		NIM:  nim,
		Nama: nama,
	}

	result := db.Create(&student)
	if result.Error != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Internal Server Error")
		return
	}

	sendSuccessResponseGIN(c, http.StatusOK, "Success")
}

func UpdateStudent(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Error Connecting to DB")
		return
	}

	nimStr := c.Query("nim")
	nim, err := strconv.Atoi(nimStr)
	if err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Invalid Value for NIM")
		return
	}
	nama := c.Query("nama")

	var student m.Students
	if err := db.Where("nim = ?", nim).First(&student).Error; err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Student Not Found")
		return
	}

	student.Nama = nama

	if err := db.Save(&student).Error; err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Internal Server Error")
		return
	}

	sendSuccessResponseGIN(c, http.StatusOK, "Success")
}

func DeleteStudent(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Error Connecting to DB")
		return
	}

	nimStr := c.Query("nim")
	nim, err := strconv.Atoi(nimStr)
	if err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Invalid Value for NIM")
		return
	}

	var student m.Students
	if err := db.Where("nim = ?", nim).First(&student).Error; err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Student Not Found")
		return
	}

	if err := db.Delete(&student).Error; err != nil {
		sendErrorResponseGIN(c, http.StatusOK, "Internal Server Error")
		return
	}

	sendSuccessResponseGIN(c, http.StatusOK, "Success")
}
