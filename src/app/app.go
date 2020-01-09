package app

import (
	"log"

	"github.com/mrKitikat/Vk-Parser-Service/src/app/models"
	parser "github.com/mrKitikat/Vk-Parser-Service/src/app/vk-parser"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// Main App
type App struct {
	Router   *fasthttprouter.Router // fasthttp Router
	VkParser *parser.VkParser       // Vk Parser is an engine
}

// NewApp creates the Vk Parser App.
func NewApp(conf *models.Config) *App {

	a := App{
		Router: fasthttprouter.New(),
	}

	// Creates Vk Parser
	a.VkParser = parser.NewVkParser(conf)
	// Sets routers
	a.setRouters()

	return &a
}

// Sets routers for api methods:
// 1. "/getProfiles/{id}" returns profiles from intersection of given id profile.
// 2. "/getlikes/{id}" returns posts, which user({id}) liked. (UNDER CONSTRUCTION)
func (a *App) setRouters() {

	// Routing for handling GET Profile from intersection.
	a.Router.POST("/getProfiles", a.VkParser.GetProfilesHandler)

}

// Run App
func (a *App) Run(host string) {
	log.Fatal(fasthttp.ListenAndServe(host, a.Router.Handler))
}

func filter(handlefunction fasthttprouter.Handle) fasthttprouter.Handle {
	return fasthttprouter.Handle(func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
		ctx.SetContentType("application/json")
		handlefunction(ctx, ps)
	})
}
