package models

import "fmt"

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

func GetAlbums() []album {
	return albums
}

func AddAlbum(newAlbum album) album {
	albums = append(albums, newAlbum)

	return newAlbum
}

func GetAlbumById(id string) (album, error) {
	for _, a := range albums {
		if a.ID == id {
			return a, nil
		}
	}

	var a album
	return a, fmt.Errorf("album with id %s not found", id)
}
