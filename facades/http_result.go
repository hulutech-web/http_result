package facades

import (
	"github.com/hulutech-web/http_result"
	"github.com/hulutech-web/http_result/contracts"
	"log"
)

func HttpResult() contracts.HttpResult {
	instance, err := http_result.App.Make(http_result.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.HttpResult)
}
