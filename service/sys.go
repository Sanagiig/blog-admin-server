package service

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"strings"
	"sync"
)

type initStore struct {
	Routes gin.RoutesInfo
}

var SysStore = initStore{}

func InitPermision() error {
	routes := SysStore.Routes
	wg := sync.WaitGroup{}
	for _, route := range routes {
		wg.Add(1)
		go func(r gin.RouteInfo) {
			permission := model.Permission{Value: r.Path}
			_ = FindPermission(&model.Permission{Value: r.Path})
			if permission.ID == "" {
				_ = CreatePermission(&model.Permission{
					Name:  strings.Replace(r.Path, "/api", "", 1),
					Code:  r.Path,
					Value: r.Path,
				})
			}
			wg.Done()
		}(route)
	}

	wg.Wait()
	return nil
}

func InitDictionary() error {
	wg := sync.WaitGroup{}
	typeDictionary := map[string]string{
		"permissionType": "权限类型",
		"tagType":        "标签类型",
		"categoryType":   "分类类型",
	}

	for key, val := range typeDictionary {
		wg.Add(1)
		go func(k string, v string) {
			dictionary := &model.Dictionary{Code: k}
			_ = FindDictionary(dictionary)
			if dictionary.ID == "" {
				_ = CreateDictionary(&model.Dictionary{Name: v, Code: k, Description: v})
			}
			wg.Done()
		}(key, val)
	}

	wg.Wait()
	return nil
}
