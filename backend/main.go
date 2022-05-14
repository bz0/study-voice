package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"io"
    "os"
	"log"
)

type Restaurant struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	  }))

	restaurant := Restaurant{
        Id: 3,
        Name: "サイゼリヤ",
    }

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, restaurant)
	})
	
	router.POST("/record", func(c *gin.Context) {
		dst, err := os.Create("test.webm")
		file, err := c.FormFile("data")
		
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		out, _ := os.Open(file.Filename)

		if _, err = io.Copy(dst, out); err != nil {
			log.Printf("err %v", err)
		}

		c.String(http.StatusOK, "Uploaded successfully.")
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}