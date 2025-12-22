package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Verse struct {
	Index      int    `json:"Index"`
	Testament  string `json:"Testament"`
	Book       string `json:"Book"`
	BookNumber int    `json:"BoookNumber"`
	Chapter    int    `json:"Chapter"`
	Verse      int    `json:"Verse"`
	Text       string `json:"Text"`
	Title      string `json:"Title"`
}

var bible []Verse

func loadBible() {
	file, err := os.ReadFile("data/Biblia_Reina_Valera_1960_Esp.json")
	if err != nil {
		log.Fatal("Error leyendo la biblia:", err)
	}

	if err := json.Unmarshal(file, &bible); err != nil {
		log.Fatal("Error parseando JSON:", err)
	}

	log.Println("Biblia cargada:", len(bible), "versículos")
}

func main() {
	loadBible()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Salud
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Biblia API activa",
			"verses": len(bible),
		})
	})

	// 📖 Obtener libros únicos
	r.GET("/books", func(c *gin.Context) {
		books := make(map[string]bool)
		var result []string

		for _, v := range bible {
			if !books[v.Book] {
				books[v.Book] = true
				result = append(result, v.Book)
			}
		}

		c.JSON(http.StatusOK, result)
	})

	// 📘 Capítulos de un libro
	r.GET("/books/:book/chapters", func(c *gin.Context) {
		book := c.Param("book")
		chapters := make(map[int]bool)
		var result []int

		for _, v := range bible {
			if v.Book == book && !chapters[v.Chapter] {
				chapters[v.Chapter] = true
				result = append(result, v.Chapter)
			}
		}

		c.JSON(http.StatusOK, result)
	})

	// 📖 Versículos por libro y capítulo
	r.GET("/books/:book/chapters/:chapter", func(c *gin.Context) {
		book := c.Param("book")
		chapter, _ := strconv.Atoi(c.Param("chapter"))

		var verses []Verse
		for _, v := range bible {
			if v.Book == book && v.Chapter == chapter {
				verses = append(verses, v)
			}
		}

		c.JSON(http.StatusOK, verses)
	})

	// 🔍 Versículo específico
	r.GET("/verse", func(c *gin.Context) {
		book := c.Query("book")
		chapter, _ := strconv.Atoi(c.Query("chapter"))
		verseNum, _ := strconv.Atoi(c.Query("verse"))

		for _, v := range bible {
			if v.Book == book && v.Chapter == chapter && v.Verse == verseNum {
				c.JSON(http.StatusOK, v)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Versículo no encontrado"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.Run("0.0.0.0:" + port)

}
