package http_result

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/hulutech-web/http_result/commands"
)

const Binding = "http_result"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app
	app.Bind(Binding, func(app foundation.Application) (any, error) {
		config := app.MakeConfig()
		config.Add("http_result", map[string]any{
			"Code":    200,
			"Message": "返回成功",
		})
		return NewResult(nil), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Commands([]console.Command{
		commands.NewPublishHttpResult(),
	})
	app.Publishes("github.com/hulutech-web/http_result", map[string]string{
		"config/http_result.go": app.ConfigPath("http_result.go"),
	})
}
