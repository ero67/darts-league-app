package middleware

// import (
// 		"fmt"
// 	"log"
// 	"time"
// 	"github.com/gin-gonic/gin"
// )

// // Logger middleware for request logging
// func Logger() gin.HandlerFunc {
// 	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
// 		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
// 			param.ClientIP,
// 			param.TimeStamp.Format(time.RFC1123),
// 			param.Method,
// 			param.Path,
// 			param.Request.Proto,
// 			param.StatusCode,
// 			param.Latency,
// 			param.Request.UserAgent(),
// 			param.ErrorMessage,
// 		)
// 	})
// }

// // Recovery middleware
// func Recovery() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				log.Printf("Panic recovered: %v", err)
// 				c.JSON(500, gin.H{
// 					"success": false,
// 					"error":   "Internal server error",
// 				})
// 				c.Abort()
// 			}
// 		}()
// 		c.Next()
// 	}
// }