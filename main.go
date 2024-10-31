package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Is_seller bool   `json:"is_seller"`
}

type Locations struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Owner_id    int    `json:"owner_id"`
}

func main() {

	connStr := "user=postgres password=12345 dbname=LocalTour sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL UNIQUE, email VARCHAR(255) NOT NULL UNIQUE, password  VARCHAR(255) NOT NULL UNIQUE, is_seller BOOLEAN, CHECK (LENGTH(password) >= 5))")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS locations(id SERIAL PRIMARY KEY, title VARCHAR(255), description TEXT, owner_id INT)")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		message, err := createUser(db, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		message, err := updateUser(db, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": getUsers(db),
		})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": getUser(db, id),
		})
	})

	router.DELETE("/users", func(c *gin.Context) {
		id := c.Param("id")
		message, err := deleteUser(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.POST("/locations", func(c *gin.Context) {
		message, err := createLocation(db, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.PUT("/locations/:id", func(c *gin.Context) {
		message, err := updateLocation(db, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.GET("/locations", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": getLocations(db),
		})
	})

	router.GET("/locations/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": getLocation(db, id),
		})
	})

	router.DELETE("/locations", func(c *gin.Context) {
		id := c.Param("id")
		message, err := deleteLocations(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createUser(db *sql.DB, c *gin.Context) (string, error) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	is_seller := c.PostForm("is_seller")
	query := fmt.Sprintf(
		"INSERT INTO users (name, email, password, is_seller) VALUES ('%s', '%s', '%s','%s',);", name, email, password, is_seller,
	)
	_, err := db.Exec(query)
	if err != nil {
		return "", err
	}

	return "User created", nil
}

func updateUser(db *sql.DB, c *gin.Context) (string, error) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	is_seller := c.PostForm("is_seller")
	query := fmt.Sprintf(
		"UPDATE users SET name = '%s', email = '%s', password = '%s', is_seller = '%s' WHERE id = %s;", name, email, password, is_seller, id,
	)
	_, err := db.Exec(query)
	if err != nil {
		return "", err
	}

	return "User created", nil
}

func deleteUser(db *sql.DB, id string) (string, error) {
	_, err := db.Exec("DELETE FROM users WHERE id = " + id + ";")
	if err != nil {
		return "", err
	}

	return "User deleted", nil
}

func getUser(db *sql.DB, id string) []User {

	rows, err := db.Query("SELECT * FROM users WHERE id = " + id + ";")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	user := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Is_seller); err != nil {
			log.Fatal(err)
		}
		user = append(user, u)
	}
	return user
}

func getUsers(db *sql.DB) []User {

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Is_seller); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	return users
}

func updateLocation(db *sql.DB, c *gin.Context) (string, error) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	description := c.PostForm("description")
	owner_id := c.PostForm("owner_id")
	query := fmt.Sprintf(
		"UPDATE locations SET title = '%s', description = '%s', owner_id = '%s', WHERE id = %s;", title, description, owner_id, id,
	)
	_, err := db.Exec(query)
	if err != nil {
		return "", err
	}

	return "User created", nil
}

func getLocation(db *sql.DB, id string) []Locations {

	rows, err := db.Query("SELECT * FROM users WHERE id = " + id + ";")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	locations := []Locations{}
	for rows.Next() {
		var l Locations
		if err := rows.Scan(&l.Id, &l.Title, &l.Description, &l.Owner_id); err != nil {
			log.Fatal(err)
		}
		locations = append(locations, l)
	}
	return locations
}

func getLocations(db *sql.DB) []Locations {

	rows, err := db.Query("SELECT * FROM locations;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	locations := []Locations{}
	for rows.Next() {
		var l Locations
		if err := rows.Scan(&l.Id, &l.Title, &l.Description, &l.Owner_id); err != nil {
			log.Fatal(err)
		}
		locations = append(locations, l)
	}
	return locations
}

// dsf
func createLocation(db *sql.DB, c *gin.Context) (string, error) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	owner_id := c.PostForm("owner_id")
	query := fmt.Sprintf(
		"INSERT INTO locations (title, description, owner_id) VALUES ('%s', '%s', %s);", title, description, owner_id,
	)
	_, err := db.Exec(query)
	if err != nil {
		return "", err
	}

	return "Location created", nil
}

func deleteLocations(db *sql.DB, id string) (string, error) {
	_, err := db.Exec("DELETE FROM locations WHERE id = " + id + ";")
	if err != nil {
		return "", err
	}

	return "Location deleted", nil
}
