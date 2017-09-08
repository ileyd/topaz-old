package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/ileyd/topaz/controllers"
	"github.com/ileyd/topaz/db"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	key1 := []byte("Xyp0tNYbk+sCX8AUBD17wCfRxr2TImBioBnQ1eGSHGySDV5r3t9lEjrL+0Rsz5sIw1UCvv/1qdoxjZg1eaeuGA==")
	key2 := []byte("Xyp0tNYbk+sCX8AUBD17wCfRxr2TImBioBnQ1eGSHGySDV5r3t9lEjrL+0Rsz5sIw1UCvv/1qdoxjZg1eaeuGA==")

	store := sessions.NewCookieStore(key1, key2)

	r.Use(sessions.Sessions("topaz-session", store))

	r.Use(CORSMiddleware())

	log.Println("DELTA1")

	db.Init()

	log.Println("DELTA2")

	api := r.Group("/api")
	{
		sonarrEvents := new(controllers.SonarrEventController)
		sEC := api.Group("/events")
		sEC.POST("/create", sonarrEvents.Create)
		sEC.GET("/all", sonarrEvents.All)
	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}
