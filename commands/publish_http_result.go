package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type PublishHttpResult struct{}

func NewPublishHttpResult() *PublishHttpResult {
	return &PublishHttpResult{}
}

// Signature The name and signature of the console command.
func (receiver *PublishHttpResult) Signature() string {
	return "http_result:publish"
}

// Description The console command description.
func (receiver *PublishHttpResult) Description() string {
	return "发布http_result资源"
}

// Extend The console command extend.
func (receiver *PublishHttpResult) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *PublishHttpResult) Handle(ctx console.Context) error {
	return nil
}
