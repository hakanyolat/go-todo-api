package app

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/hakanyolat/go-todo-api/config"
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/jinzhu/gorm"
	"net/http"
)

type App struct {
	router   *Router
	config   *config.Config
	db       *gorm.DB
	registry []ServiceInterface
	State    int
}

func NewApp() *App {
	return &App{
		router:   nil,
		config:   nil,
		db:       nil,
		registry: nil,
		State:    ApplicationState.Idled,
	}
}

func (a *App) GetDB() *gorm.DB {
	return a.db
}

// Initialize the application
func (a *App) Init(c *config.Config) {
	fmt.Println("Configuring application...")

	if a.State != ApplicationState.Idled {
		panic("Application already initialized.")
	}

	a.config = c
	db, err := gorm.Open(a.config.DB.Dialect, a.config.DB.GetDSN())

	if err != nil {
		panic("Database connection failed.")
	}

	a.db = model.Migrate(db)
	a.router = NewRouter(a)
	a.State = ApplicationState.Ready
	fmt.Println("Registering services...")
}

// Register service
func (a *App) Register(values ...ServiceInterface) {
	for _, s := range values {
		if a.State != ApplicationState.Ready {
			panic("Application is not ready for service providing. Please initialize the application first.")
		}

		a.registry = append(a.registry, s)
		s.Register(a.router, a.db)
		s.Provide()
		fmt.Println(fmt.Sprintf("+ (%T)", s))
	}
}

// Run application
func (a *App) Run() {
	fmt.Println(fmt.Sprintf("%d service has been registered.", len(a.registry)))
	host := a.config.TCP.GetHost()
	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	a.State = ApplicationState.Running

	fmt.Println(fmt.Sprintf("Server running on %s...", host))
	panic(http.ListenAndServe(host, handlers.CORS(originsOk, headersOk, methodsOk)(a.router.mux)))
}
