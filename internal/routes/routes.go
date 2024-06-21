package routes

import (
	"log"
	"net/http"
	"simple-bank/internal"
	models "simple-bank/internal/Models"
	"simple-bank/internal/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db = database.LoadDatabase()

	r.POST("/signup", signup)
	r.POST("/updateprofile", internal.AuthenticateMiddleware, updateprofile)

	return r
}

func updateprofile(c *gin.Context) {
	// Retrieve the token from the request headers
	token := c.GetString("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "could not find the auth token the request",
		})
		return
	}

	var body models.UpdateNameModel
	error := c.BindJSON(&body)

	if error != nil && (body.FirstName == nil || body.MiddleName == nil || body.LastName == nil) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Can not decode Request body",
		})
		return
	}

	var user models.User

	// Find the user with the auth token
	result := db.Where("authentication_token = ?", token).First(&user)
	if result.RowsAffected == 1 {
		// set the new values to user values
		user.FirstName = body.FirstName
		user.MiddleName = body.MiddleName
		user.LastName = body.LastName

		// save the user
		result = db.Save(&user)
		if result.RowsAffected == 1 {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Saved the new values of user",
				"user":    user,
			})
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "could not find the user with the given auth token",
			})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Could not save the values of user",
		})
	}
}

func signup(c *gin.Context) {
	var body models.SignUpModel
	error := c.BindJSON(&body)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Can not decode Request body",
		})
	}

	var user models.User

	result := db.Where("email = ?", body.Email).First(&user)

	if result.RowsAffected == 0 {
		// Auth token for 1 days
		authToken, error := internal.CreateToken(body.Email, (24 * 1))
		if error != nil {
			log.Fatal("Can't create auth token")
		}

		// Refresh token for 30 days
		refreshToken, error := internal.CreateToken(body.Email, (24 * 30))
		if error != nil {
			log.Fatal("Can't create refresh token")
		}

		user = models.User{
			Email:               body.Email,
			Password:            body.Password,
			AuthenticationToken: authToken,
			RefreshToken:        refreshToken,
		}

		result = db.Create(&user)

		// Successfully added a user in DB
		if result.RowsAffected == 1 {
			c.JSON(http.StatusOK, gin.H{
				"code":         http.StatusOK,
				"message":      "Success",
				"authToken":    user.AuthenticationToken,
				"refreshToken": user.RefreshToken,
			})
		}
	} else { // if the rows affected is 0 then it means the user already exists
		c.JSON(http.StatusConflict, gin.H{
			"code":    http.StatusConflict,
			"message": "The user already exist, please login or try again with different user",
		})
	}
}
