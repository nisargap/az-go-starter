package routes

// "gopkg.in/mgo.v2/bson"
import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nisargap/az-go-starter/db"
	"github.com/nisargap/az-go-starter/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
	"time"
)

type NewUser struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

// Route handler
func SignUp(c *gin.Context) {
	var json NewUser
	const bcryptCost = 10 // cost for bcrypto algo higher == more secure
	const saltSize = 16
	const bcryptSize = 184
	const minPasswordLength = 6
	config := c.MustGet("config").(models.Config)
	if c.BindJSON(&json) == nil {
		// check if passwords match
		if json.Password != json.PasswordConfirm {
			c.JSON(http.StatusOK, gin.H{"status": false, "message": "Passwords do not match"})
			return
		}
		if len(json.Password) < minPasswordLength {
			c.JSON(http.StatusOK, gin.H{"status": false, "message": "Password must be greater than 6 characters"})
			return
		}
		if json.Password == "" || json.Username == "" || json.PasswordConfirm == "" {
			c.JSON(http.StatusOK, gin.H{"status": false, "message": "One or more fields are empty"})
			return
		}
		session := db.GetSession(c)
		defer session.Close()
		users := session.DB(config.DatabaseName).C("users")
		// Check if a user with that particular username exists
		num, numErr := users.Find(bson.M{"username": json.Username}).Count()
		if numErr != nil {
			c.JSON(http.StatusOK, gin.H{"status": false, "message": numErr})
			return
		}
		if num != 0 {
			c.JSON(http.StatusOK, gin.H{"status": false, "message": "User with that email already exists"})
			return
		}
		// generate the salt
		buf := make([]byte, saltSize, saltSize)
		_, err := io.ReadFull(rand.Reader, buf)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"status": false, "message": err})
			return
		}
		salt := string(buf[:])
		password := json.Password + salt
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"status": false, "message": err})
			return
		}
		userToInsert := models.User{
			Username:       json.Username,
			Salt:           salt,
			Password:       string(hashed[:]),
			Privilege:      "normal",
			DateRegistered: time.Now().UTC().String()}
		errInsert := users.Insert(&userToInsert)
		if errInsert != nil {
			c.JSON(http.StatusOK, gin.H{"status": false, "message": errInsert})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "User has been registered successfully"})
		return
	}
}
