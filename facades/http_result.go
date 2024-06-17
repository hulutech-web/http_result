package facades

import (
	"goravel/packages/http_result"
	"goravel/packages/http_result/contracts"
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
