// main.go

package main

import (
	"go-gin-app/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// // Process the templates at the start so that they don't have to be loaded
	// // from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// router.GET("/", func(c *gin.Context) {

	//     // Call the HTML method of the Context to render a template
	//     c.HTML(
	//       // Set the HTTP status to 200 (OK)
	//       http.StatusOK,
	//       // Use the index.html template
	//       "index.html",
	//       // Pass the data that the page uses (in this case, 'title')
	//       gin.H{
	//         "title": "Home Page",
	//       },
	//     )

	//   })

	// Initialize the routes
	routes.InitializeRoutes(router)

	// Start serving the application
	router.Run(":8081")
}

