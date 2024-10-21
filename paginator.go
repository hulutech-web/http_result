package http_result

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"math"
	"strconv"
	"strings"
)

type Meta struct {
	TotalPage   int   `json:"total_page"`
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type PageResult struct {
	Data  any   `json:"data"` // List of data
	Total int64 `json:"total"`
	Links Links `json:"links"`
	Meta  Meta  `json:"meta"`
}

// SearchByParams
// ?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx
func (h *HttpResult) SearchByParams(params map[string]string, conditionMap map[string]interface{}, excepts ...string) *HttpResult {
	for _, except := range excepts {
		delete(params, except)
	}
	query := facades.Orm().Query()

	// 再处理url查询
	h.Query = func(q orm.Query) orm.Query {
		//处理日期时间
		// 先处理过滤条件
		for key, val := range conditionMap {
			q = q.Where(key+" = ?", val).(orm.Query)
		}
		for key, value := range params {
			//如果key包含了[]符号

			if strings.Contains(key, "[]") || value == "" || key == "pageSize" || key == "total" || key == "currentPage" || key == "sort" || key == "order" {
				continue
			} else {
				q = q.Where(gorm.Expr(key+" LIKE ?", "%"+value+"%"))
			}
			//则表示是日期时间范围
			/**
			created_at[]: 2024-10-21 00:00:00
			created_at[]: 2024-10-21 23:59:59
			*/
			if strings.Contains(key, "[]") {
				key = strings.Replace(key, "[]", "", -1)
				if value == "" {
					continue
				}
				//按照，拆分value
				ranges := strings.Split(value, ",")
				if len(ranges) == 2 {
					q = q.Where(key+" BETWEEN ? AND ?", ranges[0], ranges[1])
				} else {
					continue
				}
			}
		}

		return q
	}(query)
	return h
}

func (r *HttpResult) ResultPagination(dest any, withes ...string) (http.Response, error) {
	request := r.Context.Request()
	pageSize := request.Query("pageSize", "10")
	pageSizeInt := cast.ToInt(pageSize)
	currentPage := request.Query("currentPage", "1")
	currentPageInt := cast.ToInt(currentPage)
	total := int64(0)
	for _, with := range withes {
		r.Query = r.Query.With(with)
	}
	r.Query.Paginate(currentPageInt, pageSizeInt, dest, &total)

	URL_PATH := r.Context.Request().Origin().URL.Path
	var proto = "http://"
	if request.Origin().TLS != nil {
		proto = "https://"
	}
	// Corrected links generation
	links := Links{
		First: proto + request.Origin().Host + URL_PATH + "?pageSize=" + pageSize + "&currentPage=1",
		Last:  proto + request.Origin().Host + URL_PATH + "?pageSize=" + pageSize + "&currentPage=" + strconv.Itoa(int(total)/pageSizeInt),
		Prev:  proto + request.Origin().Host + URL_PATH + "?pageSize=" + pageSize + "&currentPage=" + strconv.Itoa(currentPageInt-1),
		Next:  proto + request.Origin().Host + URL_PATH + "?pageSize=" + pageSize + "&currentPage=" + strconv.Itoa(currentPageInt+1),
	}

	// Corrected total page calculation
	totalPage := int(math.Ceil(float64(total) / float64(pageSizeInt)))

	meta := Meta{
		TotalPage:   totalPage,
		CurrentPage: currentPageInt,
		PerPage:     pageSizeInt,
		Total:       total,
	}

	pageResult := PageResult{
		Data:  dest,
		Total: total,
		Links: links,
		Meta:  meta,
	}

	// 返回构建好的分页结果
	return r.Context.Response().Success().Json(pageResult), nil
}

func (r *HttpResult) List(dest any) (http.Response, error) {
	err := r.Query.Find(dest)
	if err != nil {
		return nil, err
	}
	return r.Context.Response().Success().Json(dest), nil
}
