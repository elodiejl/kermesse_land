package main

import (
	"back/database"
	"back/docs"
	"back/models"
	"back/repositories"
	"back/routes"
	"back/seeders"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/madkins23/gin-utils/pkg/ginzero"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
//
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer token
//
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	//DB connection
	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Connected to the database !")

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate the database: ", err)
	}
	// Migrate the schema in a controlled order
	if err := db.AutoMigrate(&models.Ticket{}); err != nil {
		log.Fatal("Failed to migrate tickets: ", err)
	}
	if err := db.AutoMigrate(&models.Prize{}); err != nil {
		log.Fatal("Failed to migrate prize: ", err)
	}
	if err := db.AutoMigrate(&models.Student{}); err != nil {
		log.Fatal("Failed to migrate student: ", err)
	}
	if err := db.AutoMigrate(&models.Activity{}); err != nil {
		log.Fatal("Failed to migrate activities: ", err)
	}
	if err := db.AutoMigrate(&models.ActivityParticipation{}); err != nil {
		log.Fatal("Failed to migrate participations: ", err)
	}
	if err := db.AutoMigrate(&models.Kermesse{}); err != nil {
		log.Fatal("Failed to migrate kermesses: ", err)
	}
	if err := db.AutoMigrate(&models.Organizer{}); err != nil {
		log.Fatal("Failed to migrate submissions: ", err)
	}
	log.Println("Migrated organisers!")
	if err := db.AutoMigrate(&models.Parent{}); err != nil {
		log.Fatal("Failed to migrate evaluations: ", err)
	}
	log.Println("Migrated parents!")
	if err := db.AutoMigrate(&models.Stand{}); err != nil {
		log.Fatal("Failed to migrate stands: ", err)
	}
	if err := db.AutoMigrate(&models.Message{}); err != nil {
		log.Fatal("Failed to migrate messages: ", err)
	}
	log.Println("Migrated messages!")
	if err := db.AutoMigrate(&models.Tombola{}); err != nil {
		log.Fatal("Failed to migrate tombolas: ", err)
	}
	if err := db.AutoMigrate(&models.Token{}); err != nil {
		log.Fatal("Failed to migrate tokens: ", err)
	}
	if err := db.AutoMigrate(&models.Transaction{}); err != nil {
		log.Fatal("Failed to migrate transactions: ", err)
	}
	log.Println("Database migrated!")

	// Password hashing
	/*password := "test1234"

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to hash the password: ", err)
		return
	}
	nouvelUtilisateur := models.User{LastName: "Dupont", FirstName: "Alice", Username: "aliced", Password: string(hashedPassword)}
	db.Create(&nouvelUtilisateur)*/

	// Context for services
	//ctx := context.Background()

	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(ginzero.Logger())

	// Middleware to add DB connection to context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		//AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, gin-zerolog example")
	})

	docs.SwaggerInfo.Title = "Kermesse Land API"
	docs.SwaggerInfo.Description = "API for the Kermesse Land project."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.GET("/chat/:teamId", controllers.HandleConnections)
	/*r.GET("/chat/messages/:teamId", func(c *gin.Context) {
		controllers.GetChatMessages(c)
	})*/
	//go controllers.HandleMessages()

	// Setup routes
	//ticketSer := services.NewTicketService(db)
	userRepo := repositories.NewUserRepository(db)
	//activityRepo := repositories.NewActivityRepositoryImpl(db)
	//activityParticipationRepo := repositories.NewActivityParticipationRepositoryImpl(db)
	kermesseRepo := repositories.NewKermesseRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	parentRepo := repositories.NewParentRepository(db)
	//prizeRepo := repositories.NewPrizeRepository(db)
	//standRepo := repositories.NewStandRepository(db)
	//ticketRepo := repositories.NewTicketRepository(db)
	//tokenRepo := repositories.NewTokenRepository(db)
	//tombolaRepo := repositories.NewTombolaRepository(db)
	//notificationRepo := repositories.NewNotificationRepository(db)
	//adminRepo := repositories.NewAdminRepository(db)

	routes.UserRoutes(r, db, userRepo)
	//routes.AdminRoutes(r, db, adminRepo)
	routes.KermesseRoutes(r, kermesseRepo)
	routes.TransactionRoutes(r, transactionRepo)
	routes.ParentRoutes(r, parentRepo)
	/*routes.SetupStepRouter(r, db, stepRepo, userRepo)
	routes.HackathonRoutes(r, db, submissionRepo, userRepo, stepRepo, teamRepo, hackathonRepo, participationRepo, messagingService)
	routes.SetupEvaluationRouter(r, db, evaluationRepo, userRepo, teamRepo, hackathonRepo, submissionRepo)
	routes.SetupSubmissionRouter(r, db, storageService, submissionRepo, userRepo, hackathonRepo, teamRepo, stepRepo)
	routes.SetupFeatureRouter(r, featureRepo, messagingService)
	routes.NotificationRoutes(r, db, notificationRepo, userRepo)
	routes.SetupSkillRouter(r, db, skillRepo, userRepo)*/

	if err := seeders.SeedUsers(db); err != nil {
		log.Fatal("Failed to seed users: ", err)
	}

	if err := seeders.SeedKermesses(db); err != nil {
		log.Fatal("Failed to seed kermesses: ", err)
	}

	if err := seeders.SeedParent(db); err != nil {
		log.Fatal("Failed to seed parents: ", err)
	}

	if err := seeders.SeedStudent(db); err != nil {
		log.Fatal("Failed to seed students: ", err)
	}

	if err := seeders.SeedStands(db); err != nil {
		log.Fatal("Failed to seed stands: ", err)
	}

	if err := seeders.SeedTombolas(db); err != nil {
		log.Fatal("Failed to seed students: ", err)
	}

	if err := seeders.SeedPrizes(db); err != nil {
		log.Fatal("Failed to seed students: ", err)
	}

	if err := seeders.SeedTickets(db); err != nil {
		log.Fatal("Failed to seed students: ", err)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
