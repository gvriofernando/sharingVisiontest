package main

import (
	"database/sql"
	f "fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Posts struct {
	ID           int    `json:"id" binding:"required"`
	TITLE        string `json:"title" binding:"required"`
	CONTENT      string `json:"content" binding:"required"`
	CATEGORY     string `json:"category" binding:"required"`
	CREATED_DATE string `json:"created_date" binding:"required"`
	UPDATED_DATE string `json:"updated_date" binding:"required"`
	STATUS       string `json:"status" binding:"required"`
}

func checkInput(title string, content string, category string, status string) string {
	var errorMessage = ""

	if len(title) < 20 {
		errorMessage = "invalid format for TITLE"
	}

	if len(content) < 200 {
		errorMessage = "invalid format for CONTENT"
	}

	if len(category) < 3 {
		errorMessage = "invalid format for CATEGORY"
	}

	if status != "publish" && status != "draft" && status != "thrash" {
		errorMessage = "invalid format for STATUS"
	}

	return errorMessage
}

func main() {
	var err error
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_article?multiStatements=true")

	if err != nil {
		f.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		f.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	f.Println("Successful Connection to Database!")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))

	router.POST("/article/", func(c *gin.Context) {
		var posts Posts
		c.ShouldBindJSON(&posts)

		var errorInput string = ""

		errorInput = checkInput(posts.TITLE, posts.CONTENT, posts.CATEGORY, posts.STATUS)

		if errorInput != "" {
			c.JSON(http.StatusOK, gin.H{
				"message": errorInput,
			})
			return
		}

		insert, err := db.Query("INSERT INTO `posts` (`Title`,`Content`,`Category`,`Created_date`,`Updated_date`,`Status`) VALUES ('" + posts.TITLE + "', '" + posts.CONTENT + "', '" + posts.CATEGORY + "', current_timestamp(), current_timestamp(), '" + posts.STATUS + "');")

		if err != nil {
			f.Println("INSERT INTO `posts` (`Title`,`Content`,`Category`,`Created_date`,`Updated_date`,`Status`) VALUES ('" + posts.TITLE + "', '" + posts.CONTENT + "', '" + posts.CATEGORY + "', current_timestamp(), current_timestamp(), '" + posts.STATUS + "');")
			panic(err.Error())
		}
		defer insert.Close()

		c.JSON(http.StatusOK, gin.H{})
		return
	})

	router.GET("/articles/:limit/:offset", func(c *gin.Context) {
		limit := c.Param("limit")
		offset := c.Param("offset")

		var result gin.H
		rows, err := db.Query("SELECT * FROM `posts` LIMIT " + limit + " OFFSET " + offset + ";")
		// err = row.Scan(&posts.ID, &posts.TITLE, &posts.CONTENT, &posts.CATEGORY, &posts.CREATED_DATE, &posts.UPDATED_DATE, &posts.STATUS)
		if err != nil {
			result = gin.H{
				"message": "Error Query",
			}
		}

		defer rows.Close()

		var items []Posts
		for rows.Next() {
			var i Posts
			if err := rows.Scan(
				&i.ID,
				&i.TITLE,
				&i.CONTENT,
				&i.CATEGORY,
				&i.CREATED_DATE,
				&i.UPDATED_DATE,
				&i.STATUS,
			); err != nil {
				result = gin.H{
					"message": "Error Query",
				}
			}
			items = append(items, i)
		}

		if err != nil {
			// If no results send null
			result = gin.H{
				"message": "Data Not Found",
			}
		} else {
			result = gin.H{
				"result": items,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	router.GET("/article/:id", func(c *gin.Context) {
		id := c.Param("id")

		var posts Posts
		var result gin.H
		row := db.QueryRow("SELECT * FROM `posts` WHERE Id = " + id + ";")
		err = row.Scan(&posts.ID, &posts.TITLE, &posts.CONTENT, &posts.CATEGORY, &posts.CREATED_DATE, &posts.UPDATED_DATE, &posts.STATUS)
		if err != nil {
			// If no results send null
			result = gin.H{
				"message": "Data Not Found",
			}
		} else {
			result = gin.H{
				"result": posts,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	router.PATCH("/article/:id", func(c *gin.Context) {
		id := c.Param("id")

		var posts Posts

		c.ShouldBindJSON(&posts)

		var errorInput string

		errorInput = checkInput(posts.TITLE, posts.CONTENT, posts.CATEGORY, posts.STATUS)

		if errorInput != "" {
			c.JSON(http.StatusOK, gin.H{
				"message": errorInput,
			})
			return
		}

		updates, err := db.Query("UPDATE `posts` SET Title = '" + posts.TITLE + "', Content = '" + posts.CONTENT + "',Category = '" + posts.CATEGORY + "',Status = '" + posts.STATUS + "' WHERE Id = " + id + ";")

		if err != nil {
			panic(err.Error())
		}
		defer updates.Close()

		c.JSON(http.StatusOK, gin.H{})
	})

	router.DELETE("/article/:id", func(c *gin.Context) {
		id := c.Param("id")

		delete, err := db.Query("DELETE FROM `posts` WHERE id =" + id + ";")

		if err != nil {
			f.Println(";")
			panic(err.Error())
		}
		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
