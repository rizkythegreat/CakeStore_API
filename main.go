package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Cake struct represents the cake data
type Cake struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      float64    `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// DBConfig struct represents the database configuration
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var db *sql.DB

func main() {
	config := DBConfig{
		Username: "root",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Database: "cake_store",
	}

	db = connectDB(config)
	defer db.Close()

	router := gin.Default()

	router.GET("/cakes", getCakes)
	router.GET("/cakes/:id", getCake)
	router.POST("/cakes", addCake)
	router.PUT("/cakes/:id", updateCake)
	router.DELETE("/cakes/:id", deleteCake)

	log.Fatal(router.Run(":8000"))
}

// Connect to the database
func connectDB(config DBConfig) *sql.DB {
	connectionString := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Get a list of cakes
func getCakes(c *gin.Context) {
	query := "SELECT * FROM cakes ORDER BY rating DESC, title ASC"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	cakes := []Cake{}
	for rows.Next() {
		var cake Cake
		var createdAt, updatedAt, deletedAt []uint8

		err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		cake.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
		cake.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		if len(deletedAt) > 0 {
			deletedAtTime, _ := time.Parse("2006-01-02 15:04:05", string(deletedAt))
			cake.DeletedAt = &deletedAtTime
		}

		cakes = append(cakes, cake)
	}

	c.JSON(http.StatusOK, cakes)
}

// Get the detail of a cake
func getCake(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cake ID"})
		return
	}

	query := "SELECT * FROM cakes WHERE id = ?"
	row := db.QueryRow(query, id)

	var cake Cake
	var createdAt, updatedAt, deletedAt []uint8

	err = row.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cake not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	cake.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
	cake.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if len(deletedAt) > 0 {
		deletedAtTime, _ := time.Parse("2006-01-02 15:04:05", string(deletedAt))
		cake.DeletedAt = &deletedAtTime
	}

	c.JSON(http.StatusOK, cake)
}

// Add a new cake
func addCake(c *gin.Context) {
	var cake Cake
	if err := c.ShouldBindJSON(&cake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO cakes (title, description, rating, image) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, cake.Title, cake.Description, cake.Rating, cake.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// Update an existing cake
func updateCake(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cake ID"})
		return
	}

	var cake Cake
	if err := c.ShouldBindJSON(&cake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE cakes SET title = ?, description = ?, rating = ?, image = ? WHERE id = ?"
	_, err = db.Exec(query, cake.Title, cake.Description, cake.Rating, cake.Image, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// Delete a cake
func deleteCake(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cake ID"})
		return
	}

	query := "DELETE FROM cakes WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
