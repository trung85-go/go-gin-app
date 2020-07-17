// handlers.user.go

package handlers

import (
	"math/rand"
	"net/http"
	"strconv"

	"go-gin-app/models"
	"go-gin-app/http"

	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	response.Render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func PerformLogin(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

    var sameSiteCookie http.SameSite;

	// Check if the username/password combination is valid
	if models.IsUserValid(username, password) {
		// If the username/password is valid set the token in a cookie
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.SetSameSite(sameSiteCookie)
		c.Set("is_logged_in", true)

		response.Render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func GenerateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func Logout(c *gin.Context) {

    var sameSiteCookie http.SameSite;

	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)
	c.SetSameSite(sameSiteCookie)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func ShowRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	response.Render(c, gin.H{
		"title": "Register"}, "register.html")
}

func Register(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

    var sameSiteCookie http.SameSite;

	if _, err := models.RegisterNewUser(username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.SetSameSite(sameSiteCookie)
		c.Set("is_logged_in", true)

		response.Render(c, gin.H{
			"title": "Successful registration & Login"}, "login-successful.html")

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})

	}
}
