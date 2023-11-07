package main

import (
	"net/http"
	"strings"

	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Abbey Road", Artist: "The Beatles", Price: 19.99},
	{ID: "5", Title: "Back in Black", Artist: "AC/DC", Price: 9.99},
	{ID: "6", Title: "Led Zeppelin IV", Artist: "Led Zeppelin", Price: 29.99},
	{ID: "7", Title: "Blood on the Tracks", Artist: "Bob Dylan", Price: 15.99},
	{ID: "8", Title: "Born to Run", Artist: "Bruce Springsteen", Price: 12.99},
	{ID: "9", Title: "Graceland", Artist: "Paul Simon", Price: 14.99},
	{ID: "10", Title: "Imagine", Artist: "John Lennon", Price: 17.99},
	{ID: "11", Title: "Rumours", Artist: "Fleetwood Mac", Price: 11.99},
	{ID: "12", Title: "Hotel California", Artist: "Eagles", Price: 12.99},
	{ID: "13", Title: "The Wall", Artist: "Pink Floyd", Price: 18.99},
	{ID: "14", Title: "Thriller", Artist: "Michael Jackson", Price: 16.99},
	{ID: "15", Title: "The Joshua Tree", Artist: "U2", Price: 14.99},
	{ID: "16", Title: "Licensed to Ill ", Artist: "Beastie Boys", Price: 13.99},
	{ID: "17", Title: "Appetite for Destruction", Artist: "Guns N' Roses", Price: 11.99},
	{ID: "18", Title: "Legend", Artist: "Bob Marley", Price: 17.99},
	{ID: "19", Title: "Nevermind", Artist: "Nirvana", Price: 13.99},
	{ID: "20", Title: "Pet Sounds", Artist: "The Beach Boys", Price: 15.99},
	{ID: "21", Title: "Back to Black", Artist: "Amy Winehouse", Price: 12.99},
	{ID: "22", Title: "The Chronic", Artist: "Dr. Dre", Price: 14.99},
	{ID: "23", Title: "American Idiot", Artist: "Green Day", Price: 11.99},
	{ID: "24", Title: "Born in the U.S.A.", Artist: "Bruce Springsteen", Price: 13.99},
	{ID: "25", Title: "Graduation", Artist: "Kanye West", Price: 15.99},
	{ID: "26", Title: "Dirt", Artist: "Alice in Chains", Price: 10.99},
	{ID: "27", Title: "The Miseducation of Lauryn Hill", Artist: "Lauryn Hill", Price: 14.99},
	{ID: "28", Title: "The Velvet Underground", Artist: "The Velvet Underground", Price: 18.99},
	{ID: "29", Title: "Are You Experienced", Artist: "The Jimi Hendrix Experience", Price: 19.99},
	{ID: "30", Title: "Dead Cities", Artist: "The Future Sound of London", Price: 11.99},
	{ID: "31", Title: "Rock and Roll", Artist: "John Lennon", Price: 12.99},
}

// getAlbums responds with a list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
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
