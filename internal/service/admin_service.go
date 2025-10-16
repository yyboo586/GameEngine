package service

import (
	"GameEngine/internal/common"
	"GameEngine/internal/model"
	"context"
	"encoding/json"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type IAdminService interface {
	Auth(r *ghttp.Request) (err error)
}

var localAdminService IAdminService

func AdminService() IAdminService {
	if localAdminService == nil {
		panic("implement not found for interface IAdminService, forgot register?")
	}
	return localAdminService
}

func RegisterAdminService(i IAdminService) {
	localAdminService = i
}

type adminService struct {
	address string
	client  common.HTTPClient
}

func NewAdminService() IAdminService {
	return &adminService{
		address: g.Cfg().MustGet(context.Background(), "server.thirdService.adminService").String(),
		client:  common.NewHTTPClient(),
	}
}

func (a *adminService) Auth(r *ghttp.Request) (err error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil
	}

	url := a.address + "/api/v1/system/token/introspect"
	header := map[string]interface{}{
		"Authorization": token,
	}

	_, respBody, err := a.client.POST(r.Context(), url, header, nil)
	if err != nil {
		return err
	}

	var i httpResponse
	err = json.Unmarshal(respBody, &i)
	if err != nil {
		return err
	}

	if i.Code != 0 {
		return errors.New(i.Message)
	}

	var userInfo model.User
	userInfo.ID = int64(i.Data.(map[string]interface{})["user_id"].(float64))
	userInfo.Name = i.Data.(map[string]interface{})["user_name"].(string)

	r.SetCtxVar(model.UserInfoKey, userInfo)

	return
}

type httpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
