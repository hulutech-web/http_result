# http_result

#### 一、安装
```go
go get -u github.com/hulutech-web/http_result

```
##### 1.1发布资源  
```go
go run . artisan vendor:publish --package=github.com:hulutech-web/http_result

```
##### 1.2 注册服务提供者:config/app.go
```go
import	"github.com:hulutech-web/http_result"

func init() {
"providers": []foundation.ServiceProvider{
	....
	&http_result.ServiceProvider{},
	
}
	
}

```
#### 二、使用

##### 2.1 使用说明:自定义默认返回
发布资源后，config/http_result.go中的配置文件中有默认的返回状态码和message，可自行修改
```go
config.Add("http_result", map[string]any{
		"Code":    500, //自定义修改默认状态码
		"Message": "获取成功",//自定义修改默认消息
})
```
#### 使用：ctx为goravel控制器中的默认(ctx http.Context) "github.com/goravel/framework/contracts/http"
##### 方式一：
一、成功返回：
```
import httpfacades "git@github.com:hulutech-web/http_result.git"
return httpfacades.NewResult(ctx).Success("", user)
```
二、失败返回：
```go
import httpfacades "git@github.com:hulutech-web/http_result.git"
httpfacades.NewResult(ctx).Error(500, "用户不存在", "no users find")
```
三、表单验证错误：
```go
httpfacades.NewResult(ctx).ValidError("验证失败", errors.All())
```