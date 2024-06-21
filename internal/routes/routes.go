package routes

import (
	"fmt"
	"log"
	"net/http"
	"simple-bank/internal"
	models "simple-bank/internal/Models"
	"simple-bank/internal/database"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db = database.LoadDatabase()

	r.POST("/signup", signup)
	r.POST("/updateprofile", authenticateMiddleware, updateprofile)

	return r
}

func authenticateMiddleware(c *gin.Context) {
	// Retrieve the token from the request headers
	token := c.GetHeader("Authorization")
	if token == "" {
		fmt.Println("Token missing in request")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Auth token is missing in the requst headers",
		})
		return
	}

	authToken := strings.Split(token, " ")[1]

	if authToken == "" {
		fmt.Println("Token missing in request")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Auth token is missing in the requst headers",
		})
		return
	}

	// Verify the token
	_, error := internal.VerifyToken(authToken)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Auth token is not valid",
		})
		return
	}
	c.Next()
}

func updateprofile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
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
		// Auth token for 7 days
		authToken, error := internal.CreateToken("viveksinghmehta.smr@gmail.com", (24 * 7))
		if error != nil {
			log.Fatal("Can't create token")
		}

		// Refresh token for 30 days
		refreshToken, error := internal.CreateToken("viveksinghmehta.smr@gmail.com", (24 * 30))
		if error != nil {
			log.Fatal("Can't create token")
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
