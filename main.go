package main

import (
	"ajbates93/album-service/models"
	"net/http"
	"strings"

	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with a list of all albums as JSON
func getAlbums(c *gin.Context) {
	albums := models.GetAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	newAlbum := models.AddAlbum(c.BindJSON(&newAlbum))
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	album, err := models.GetAlbumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
	return
}

func getArtistBySearch(c *gin.Context) {
	query := c.DefaultQuery("name", "")

	if len(query) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no query parameter"})
		return
	}
	artists := []string{}
	for _, a := range albums {
		if strings.Contains(strings.ToLower(a.Artist), strings.ToLower(query)) {
			artists = append(artists, a.Artist)
		}
	}
	if len(artists) > 0 {
		c.IndentedJSON(http.StatusOK, artists)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "artist not found"})
	return
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()

	// Use custom logger middleware
	router.Use(LoggerMiddleware())

	// Use our custom authentication middleware for a specific group of routes
	authGroup := router.Group("/api")
	authGroup.Use(AuthMiddleware())

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.GET("/artists", getArtistBySearch)

	router.Run(":8080")
}
