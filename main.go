package main

import (
	"log"
	"net/http"
	"ngevent-api/auth"
	"ngevent-api/event"
	"ngevent-api/handler"
	"ngevent-api/helper"
	"ngevent-api/payment"
	"ngevent-api/transaction"
	"ngevent-api/user"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:halo123@tcp(127.0.0.1:3306)/db_ngevent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	userRepository := user.NewRepository(db)
	eventRepository := event.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	authService := auth.NewService()
	paymentService := payment.NewService(eventRepository)
	userService := user.NewService(userRepository)
	eventService := event.NewService(eventRepository)
	transactionService := transaction.NewService(transactionRepository, eventRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	eventHandler := handler.NewEventHandler(eventService, authService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("api/v1")

	routerUser := api.Group("/user")
	routerEvent := api.Group("/events")
	routerEventSingle := api.Group("/event")
	routerTransaction := api.Group("/transactions")

	routerUser.POST("/register", userHandler.RegisterNewUser)
	routerUser.POST("/login", userHandler.Login)
	routerUser.GET("/fetch", authMiddleware(authService, userService, "all"), userHandler.FetchUser)
	routerUser.POST("/update", authMiddleware(authService, userService, "all"), userHandler.UpdateDataUser)
	routerUser.POST("/change-password", authMiddleware(authService, userService, "all"), userHandler.UpdatePasswordUser)
	routerUser.POST("/change-avatar", authMiddleware(authService, userService, "all"), userHandler.UpdateAvatarUser)

	routerEvent.POST("", authMiddleware(authService, userService, "admin"), eventHandler.CreateNewEvent)
	routerEvent.GET("", eventHandler.GetAllEvent)
	routerEvent.GET(":id", eventHandler.GetEventById)
	routerEvent.POST(":id/images", authMiddleware(authService, userService, "admin"), eventHandler.CreateNewImageEvent)
	routerEvent.PUT(":id", authMiddleware(authService, userService, "admin"), eventHandler.UpdateDataEvent)
	routerEventSingle.PUT(":id_event/images/:id_image", authMiddleware(authService, userService, "admin"), eventHandler.MakeImagePrimary)
	routerEvent.DELETE(":id", authMiddleware(authService, userService, "admin"), eventHandler.DeleteEvent)
	routerEventSingle.DELETE(":id_event/images/:id_image", authMiddleware(authService, userService, "admin"), eventHandler.DeleteEventImage)

	routerTransaction.GET("", authMiddleware(authService, userService, "all"), transactionHandler.GetAllTransaction)
	routerTransaction.GET(":id", authMiddleware(authService, userService, "all"), transactionHandler.GetTransaction)
	routerTransaction.POST("", authMiddleware(authService, userService, "member"), transactionHandler.CreateNewTransaction)
	routerTransaction.POST("/notification", transactionHandler.GetNotification)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service, userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenString = tokenArray[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if user.UserType != userType && userType != "all" {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
