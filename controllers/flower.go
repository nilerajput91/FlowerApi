package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/flowerapi/model"

	"github.com/flowerapi/db"

	"github.com/gin-gonic/gin"
)

// Create new flower details
func Create(c *gin.Context) {
	var buffer bytes.Buffer
	id := c.PostForm("id")
	name := c.PostForm("name")
	category := c.PostForm("category")
	price := c.PostForm("price")
	photo := c.PostForm("photo")
	descriptions := c.PostForm("descriptions")

	stmt, err := db.Init().Prepare("insert into flower(id,name,category,price,photo,descriptions)values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(id, name, category, price, photo, descriptions)
	// Fastest way to append strings
	buffer.WriteString(name)
	buffer.WriteString(" ")
	defer stmt.Close()
	flowername := buffer.String()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully created:", flowername),
	})

}

// GetAllFlowers
func GetAllFlowers(c *gin.Context) {
	var (
		flower  model.Flower
		flowers []model.Flower
	)
	rows, err := db.Init().Query("select * from flower	;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&flower.ID, &flower.Name, &flower.Category, &flower.Price, &flower.Photo, &flower.Description)
		flowers = append(flowers, flower)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": flowers,
	})
}

// GetFlower details
func GetFlower(c *gin.Context) {
	var (
		flower model.Flower
		result gin.H
	)
	id := c.Param("id")
	err := db.Init().QueryRow("select * from flower where id=?; ", id).Scan(&flower.ID, &flower.Name, &flower.Category, &flower.Price, &flower.Photo, &flower.Description)
	if err != nil {
		// if no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": flower,
		}
	}
	c.JSON(http.StatusOK, result)
}

// DeleteFlower
func DeleteFlower(c *gin.Context) {
	id := c.Query("id")
	stmt, err := db.Init().Prepare("delete from flower where id =?;")
	if err != nil {
		fmt.Printf(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("successfully deleted flower:%s", id),
	})
}
