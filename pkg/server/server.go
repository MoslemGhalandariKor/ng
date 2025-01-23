package server

import (
	"fmt"
	"nextgen/internals/gintemplrenderer"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	router.LoadHTMLGlob("templates/**/**/**/*")
	router.Static("static", "static")
	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	return &Server{
		Router: router,
	}
}
func (s *Server) StartServer() {
	(&Routes{}).webRouter(s)
	(&Routes{}).authRouter(s)
	(&Routes{}).dashboardRouter(s)
	(&Routes{}).blogRouter(s)
	(&Routes{}).aboutusRouter(s)
	(&Routes{}).accountsRouter(s)
	
	err := s.Router.Run()
	if err != nil {
		fmt.Println("Error starting the server")
	}
}
