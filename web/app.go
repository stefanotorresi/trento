package web

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/trento-project/trento/internal/consul"
	"github.com/trento-project/trento/web/services"
	"github.com/trento-project/trento/web/services/ara"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/trento-project/trento/docs" // docs is generated by Swag CLI, you have to import it.
)

//go:embed frontend/assets
var assetsFS embed.FS

//go:embed templates
var templatesFS embed.FS

const araAddrDefault = "127.0.0.1:8000"

type App struct {
	host string
	port int
	Dependencies
}

type Dependencies struct {
	consul               consul.Client
	engine               *gin.Engine
	store                cookie.Store
	checksService        services.ChecksService
	subscriptionsService services.SubscriptionsService
}

func DefaultDependencies() Dependencies {
	consulClient, _ := consul.DefaultClient()
	engine := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	araService := ara.NewAraService(araAddrDefault)
	checksService := services.NewChecksService(araService)
	subscriptionsService := services.NewSubscriptionsService(consulClient)

	return Dependencies{consulClient, engine, store, checksService, subscriptionsService}
}

func (d *Dependencies) SetAraAddr(araAddr string) {
	araService := ara.NewAraService(araAddr)
	d.checksService = services.NewChecksService(araService)
}

// shortcut to use default dependencies
func NewApp(host string, port int) (*App, error) {
	return NewAppWithDeps(host, port, DefaultDependencies())
}

// @title Trento API
// @version 1.0
// @description Trento API

// @contact.name Trento Project
// @contact.url https://www.trento-project.io
// @contact.email  trento-project@suse.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func NewAppWithDeps(host string, port int, deps Dependencies) (*App, error) {
	app := &App{
		Dependencies: deps,
		host:         host,
		port:         port,
	}

	InitAlerts()
	engine := deps.engine
	engine.HTMLRender = NewLayoutRender(templatesFS, "templates/*.tmpl")
	engine.Use(ErrorHandler)
	engine.Use(sessions.Sessions("session", deps.store))
	engine.StaticFS("/static", http.FS(assetsFS))
	engine.GET("/", HomeHandler)
	engine.GET("/about", NewAboutHandler(deps.subscriptionsService))
	engine.GET("/hosts", NewHostListHandler(deps.consul))
	engine.GET("/hosts/:name", NewHostHandler(deps.consul, deps.subscriptionsService))
	engine.GET("/catalog", NewChecksCatalogHandler(deps.checksService))
	engine.GET("/clusters", NewClusterListHandler(deps.consul, deps.checksService))
	engine.GET("/clusters/:id", NewClusterHandler(deps.consul, deps.checksService))
	engine.POST("/clusters/:id/settings", NewSaveClusterSettingsHandler(deps.consul))
	engine.GET("/sapsystems", NewSAPSystemListHandler(deps.consul))
	engine.GET("/sapsystems/:sid", NewSAPSystemHandler(deps.consul))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/ping", ApiPingHandler)

		apiGroup.GET("/tags", ApiListTag(deps.consul))
		apiGroup.POST("/hosts/:name/tags", ApiHostCreateTagHandler(deps.consul))
		apiGroup.DELETE("/hosts/:name/tags/:tag", ApiHostDeleteTagHandler(deps.consul))
		apiGroup.POST("/clusters/:id/tags", ApiClusterCreateTagHandler(deps.consul))
		apiGroup.DELETE("/clusters/:id/tags/:tag", ApiClusterDeleteTagHandler(deps.consul))
		apiGroup.POST("/sapsystems/:sid/tags", ApiSAPSystemCreateTagHandler(deps.consul))
		apiGroup.DELETE("/sapsystems/:sid/tags/:tag", ApiSAPSystemDeleteTagHandler(deps.consul))
	}

	return app, nil
}

func (a *App) Start() error {
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", a.host, a.port),
		Handler:        a,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	a.engine.ServeHTTP(w, req)
}
