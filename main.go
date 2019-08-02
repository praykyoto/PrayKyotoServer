package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

type Flower struct {
	gorm.Model
	Num uint
}

var (
	db *gorm.DB
)

const allowDomain string = "https://praykyoto.github.io"

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "/db/kyoto.db")
	if err != nil {
		panic("fail to connect database")
	}
	initTable()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	flower := r.Group("/api/flower")
	{
		flower.POST("/", addFlower)
		flower.GET("/", fetchFlower)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func initTable() {
	db.AutoMigrate(&Flower{})
	if !db.HasTable(&Flower{}) {
		db.CreateTable(&Flower{})
		db.Create(&Flower{Num: 0})
	}
}

func addFlower(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", allowDomain)
	c.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Header("Access-Control-Max-Age", "1728000")
	if origin != allowDomain {
		return
	}
	var flower Flower
	db.First(&flower, 1)
	db.Model(&flower).Update("Num", flower.Num+1)
	if err := db.Save(&flower).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func fetchFlower(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", allowDomain)
	c.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Header("Access-Control-Max-Age", "1728000")
	fmt.Print(origin)
	if origin != allowDomain {
		return
	}
	var flower Flower
	db.First(&flower, 1)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"num":    flower.Num,
	})
}
