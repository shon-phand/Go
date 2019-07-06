package main

import (
	// "bytes"
	// "database/sql"

	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Email         string `json : "email"`
	Name          string `json : "name"`
	Password      string `-`
	Exp           string `json : "email"`
	PrevsComapany string `json : "prevsComapany"`
	Gender        string `json : "gender"`
	Id            int
}

func main() {

	db, err := sql.Open("mysql", "shon:Shon@2019@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("conneced to database")
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", homepage)

	r.GET("/v1/u/login", getLogin)
	r.GET("/v1/u/register", getRegistration)
	r.POST("/v1/u/register", postRegistration)
	r.POST("/v1/u/login", postLogin)
	r.GET("/v1/u/registered", getAllRegistered)
	r.GET("/v1/u/login/:id", getUserById)

	r.Run(":8082")

	// func getAllRegistered(c *gin.Context) {
	// 	row := db.QueryRow("select id, first_name, last_name from person;")
	// 	err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)
	// }

}

func searchUser(username interface{}) Data {

	var user Data

	db, err := sql.Open("mysql", "shon:Shon@2019@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	row := db.QueryRow("select * from user where email=?", username)
	err = row.Scan(&user.Name, &user.Email, &user.Password, &user.Exp, &user.PrevsComapany, &user.Gender, &user.Id)
	if err != nil {
		fmt.Print(err.Error())
	}
	return user

}

func getUserById(c *gin.Context) {

	db, err := sql.Open("mysql", "shon:Shon@2019@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	var user Data

	id := c.Param("id")

	row := db.QueryRow("select * from user where id=?", id)
	fmt.Println(row)
	err = row.Scan(&user.Name, &user.Email, &user.Password, &user.Exp, &user.PrevsComapany, &user.Gender, &user.Id)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(user)

	c.JSON(http.StatusOK, user)

}

func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		gin.H{
			"title": "Login Page",
		})
}

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html",
		gin.H{
			"title": "Login Page",
		})

}

func postLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	user := searchUser(username)

	if user.Email == username && user.Password == password {

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in",
			"username": username})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid credentials"})
	}
}

func getRegistration(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html",
		gin.H{
			"title": "Registration Page",
		})
}

func postRegistration(c *gin.Context) {

	db, err := sql.Open("mysql", "shon:Shon@2019@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	exp := c.PostForm("exp")
	company := c.PostForm("prevsComapany")
	gender := c.PostForm("gender")
	//fmt.Println(name, email, password, exp, company, gender)

	stmt, err := db.Prepare("insert into user (name,email,password,exp,presComapany,gender) values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, email, password, exp, company, gender)

	if err != nil {
		fmt.Print(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": name + " you are successfully registered.",
		})
	}
	defer stmt.Close()

}

func getAllRegistered(c *gin.Context) {

	var result gin.H
	var user Data
	var users []Data
	db, err := sql.Open("mysql", "shon:Shon@2019@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available

	rows, err := db.Query("select name,email,password,exp,presComapany,gender from user;")

	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {

		err = rows.Scan(&user.Name, &user.Email, &user.Password, &user.Exp, &user.PrevsComapany, &user.Gender)
		if err != nil {
			fmt.Print(err.Error())
		}
		users = append(users, user)
	}
	defer rows.Close()
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  len(users),
		}
	}
	c.JSON(http.StatusOK, result)
}
