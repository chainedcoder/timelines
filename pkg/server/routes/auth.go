package routes

import (
	"chainedcoder/timelines/internal/handlers/auth"
	"chainedcoder/timelines/internal/orm"
	"chainedcoder/timelines/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Auth routes
func Auth(cfg *utils.ServerConfig, r *gin.Engine, orm *orm.ORM) error {
	provider := string(utils.ProjectContextKeys.ProviderCtxKey)
	// OAuth handlers
	g := r.Group(cfg.VersionedEndpoint("/auth"))
	g.GET("/:"+provider, auth.Begin())
	g.GET("/:"+provider+"/callback", auth.Callback(cfg, orm))
	// g.GET("/:"+provider+"/refresh", auth.Refresh(cfg, orm))
	return nil
}