package main

import (
	"GameEngine/internal/controller"
	"GameEngine/internal/logics"
	"GameEngine/internal/logics/game"
	"GameEngine/internal/logics/metadata"
	"GameEngine/internal/logics/ranking"
	"GameEngine/internal/logics/recommendation"
	"GameEngine/internal/logics/reservation"
	"GameEngine/internal/service"
	"fmt"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	s := g.Server()
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")

	service.RegisterAdminService(service.NewAdminService())
	service.RegisterFileEngine()
	service.RegisterGame(game.NewGame())
	service.RegisterMetadata(metadata.NewMetadata())
	service.RegisterRanking(ranking.NewRanking())
	service.RegisterRecommendation(recommendation.NewRecommendation())
	service.RegisterReservation(reservation.NewReservation())
	service.RegisterUserBehavior(logics.NewUserBehavier())

	s.Group("/api/v1/game-engine", func(group *ghttp.RouterGroup) {
		group.Middleware(CORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Middleware(Auth)
		// 游戏相关接口
		group.Bind(
			controller.GameController,
			controller.MetadataController,
			controller.RankingController,
			// controller.RecommendationController,
			controller.ReservationController,
			controller.UserBehavierController,
		)
	})

	s.Run()
}

func CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func Auth(r *ghttp.Request) {
	err := service.AdminService().Auth(r)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code":    http.StatusUnauthorized,
			"message": fmt.Sprintf("令牌校验失败: %s", err.Error()),
		})
		return
	}
	r.Middleware.Next()
}
