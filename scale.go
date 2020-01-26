package main
import (
  // "fmt"
  "github.com/gin-gonic/gin"
)

// func main(){
//     fmt.Println("Hello, world!");
// }

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  r.GET("/hello", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "World !",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080
}
