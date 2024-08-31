package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nikitsingh/forky/backend/internal/config"
	"github.com/nikitsingh/forky/backend/internal/handler"
	"github.com/nikitsingh/forky/backend/internal/repo"
	"github.com/nikitsingh/forky/backend/internal/service"
)

type Router struct {
	router *gin.Engine
	db     *sqlx.DB
}

func NewRouter(db *sqlx.DB) *Router {
	return &Router{
		router: gin.Default(),
		db:     db,
	}
}

func (r *Router) SetupRouter() {
	if config.Envs.ENV == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	v1 := r.router.Group("/api/v1")

	// User section
	userRepo := repo.NewUserRepo(r.db)
	userService := service.NewUserService(userRepo)
	// User section end

	// Auth section
	authRepo := repo.NewAuthRepo(r.db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService, userService)

	v1.POST("/auth/magic-link", authHandler.CreateMagicLink)
	v1.POST("/auth/magic-link/verify", authHandler.VerifyMagicLink)
	// Auth section end
}
