package app

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/hakanyolat/go-todo-api/config"
	"github.com/jinzhu/gorm"
	"net/http"
)

type App struct {
	State        state
	router       *Router
	db           *gorm.DB
	config       config.Config
	services     []ServiceInterface
	models       []ModelInterface
	modelManager *ModelManager
	isMock       bool
}

func NewApp() *App {
	return &App{
		router:   nil,
		config:   config.Config{},
		db:       nil,
		services: nil,
		State:    State.Created,
	}
}

func (a *App) Configure(c config.Config) {
	fmt.Println("\u2023 Configuring application...")
	if a.State != State.Created {
		panic("Application already initialized.")
	}
	a.config = c
}

// Initialize the application
func (a *App) Init(tables []ModelInterface, services []ServiceInterface) {
	fmt.Println("\u2023 Initializing...")
	a.models = tables
	a.services = services
	a.router = NewRouter(a)

	db, err := gorm.Open(a.config.DB.Dialect, a.config.DB.GetDSN())
	if err != nil {
		panic("Database connection failed.")
	}

	a.db = db
	a.State = State.Ready

	if !a.isMock {
		a.modelManager = NewModelManager(a.db)
		a.migrateModels()
		a.seedModels()
	}

	a.provideServices()
}

func (a *App) migrateModels() {
	if a.State != State.Ready {
		panic("Application is not ready for database operations.")
	}

	fmt.Println("\u2023 Migrating models...")

	if len(a.models) == 0 {
		fmt.Println("  (No models found)")
		return
	}

	for _, t := range a.models {
		a.modelManager.Migrate(t)
		fmt.Println(fmt.Sprintf("  \u2713 (%T)", t))
	}

	fmt.Println(fmt.Sprintf("  %d model has been migrated.", len(a.models)))
}

func (a *App) seedModels() {
	if a.State != State.Ready {
		panic("Application is not ready for database operations.")
	}

	if len(a.models) == 0 {
		return
	}

	fmt.Println("\u2023 Seeding models...")
	for _, t := range a.models {
		seeded := a.modelManager.Seed(t)
		if seeded > 0 {
			fmt.Println(fmt.Sprintf("  \u2713 %d rows seeded for (%T)", seeded, t))
		}
	}
}

// Init service
func (a *App) provideServices() {
	if a.State != State.Ready {
		panic("Application is not ready for service providing. Please initialize the application.")
	}

	fmt.Println("\u2023 Initializing services...")
	for _, s := range a.services {
		s.Init(a.router, a.db)
		s.Provide()
		fmt.Println(fmt.Sprintf("  \u2713 (%T)", s))
	}
	fmt.Println(fmt.Sprintf("  %d service has been provided.", len(a.services)))
}

// Run application
func (a *App) Run() {
	host := a.config.TCP.GetHost()
	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	a.State = State.Running
	fmt.Println()
	fmt.Printf("\033[1;32m%s\033[0m", fmt.Sprintf("Server running at %s...", host))
	fmt.Println()
	panic(http.ListenAndServe(host, handlers.CORS(originsOk, headersOk, methodsOk)(a.router.mux)))
}

func (a *App) GetDB() *gorm.DB {
	return a.db
}

func (a *App) Mock() *App {
	a.isMock = true
	return a
}

func (a *App) IsMock() bool {
	return a.isMock
}