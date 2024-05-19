package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync/atomic"
	"github.com/gin-gonic/gin"
)

type SharkAttack struct {
	Date     string `json:"date"`
	Country  string `json:"country"`
	Name     string `json:"name"`
	Activity string `json:"activity"`
	Age      string `json:"age"`
	Injury   string `json:"injury"`
}

var (
	posts atomic.Value
)

func main() {
	filename, err := os.Open("global-shark-attack.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := filename.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := io.ReadAll(filename)
	if err != nil {
		log.Fatal(err)
	}

	var records []SharkAttack
	err = json.Unmarshal(data, &records)
	if err != nil {
		log.Fatal(err)
	}

	var initialPosts []SharkAttack
	for i := 0; i < 10; i++ {
		initialPosts = append(initialPosts, records[rand.Intn(len(records))])
	}

	posts.Store(initialPosts)

	r := gin.Default()

	r.GET("/posts", getPosts)
	r.POST("/posts", postPost)
	r.GET("/posts/:id", getPostByID)
	r.DELETE("/posts/:id", deletePostByID)

	fmt.Println("Server is listening at port 8080")
	log.Fatal(r.Run(":8080"))
}

func getPosts(c *gin.Context) {
	currentPosts := posts.Load().([]SharkAttack)
	c.JSON(http.StatusOK, currentPosts)
}

func postPost(c *gin.Context) {
	var post SharkAttack
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentPosts := posts.Load().([]SharkAttack)
	newPosts := append(currentPosts, post)
	posts.Store(newPosts)
	c.JSON(http.StatusCreated, post)
}

func getPostByID(c *gin.Context) {
	id := c.Param("id")
	var index int
	if _, err := fmt.Sscanf(id, "%d", &index); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	currentPosts := posts.Load().([]SharkAttack)
	if index < 0 || index >= len(currentPosts) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	post := currentPosts[index]
	c.JSON(http.StatusOK, post)
}

func deletePostByID(c *gin.Context) {
	id := c.Param("id")
	var index int
	if _, err := fmt.Sscanf(id, "%d", &index); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	currentPosts := posts.Load().([]SharkAttack)
	if index < 0 || index >= len(currentPosts) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	newPosts := append(currentPosts[:index], currentPosts[index+1:]...)
	posts.Store(newPosts)
	c.Status(http.StatusNoContent)
}
