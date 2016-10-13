package api

/************************************************
  generated by IDE. for [VPCRouterAPI]
************************************************/

import (
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *VPCRouterAPI) Reset() *VPCRouterAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *VPCRouterAPI) Offset(offset int) *VPCRouterAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *VPCRouterAPI) Limit(limit int) *VPCRouterAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *VPCRouterAPI) Include(key string) *VPCRouterAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *VPCRouterAPI) Exclude(key string) *VPCRouterAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *VPCRouterAPI) FilterBy(key string, value interface{}) *VPCRouterAPI {
	api.filterBy(key, value, false)
	return api
}

// func (api *VPCRouterAPI) FilterMultiBy(key string, value interface{}) *VPCRouterAPI {
// 	api.filterBy(key, value, true)
// 	return api
// }

// WithNameLike 名称条件
func (api *VPCRouterAPI) WithNameLike(name string) *VPCRouterAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *VPCRouterAPI) WithTag(tag string) *VPCRouterAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *VPCRouterAPI) WithTags(tags []string) *VPCRouterAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *VPCRouterAPI) WithSizeGib(size int) *VPCRouterAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *VPCRouterAPI) WithSharedScope() *VPCRouterAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *VPCRouterAPI) WithUserScope() *VPCRouterAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *VPCRouterAPI) SortBy(key string, reverse bool) *VPCRouterAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *VPCRouterAPI) SortByName(reverse bool) *VPCRouterAPI {
	api.sortByName(reverse)
	return api
}

// func (api *VPCRouterAPI) SortBySize(reverse bool) *VPCRouterAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// func (api *VPCRouterAPI) New() *sacloud.VPCRouter {
// 	return &sacloud.VPCRouter{}
// }

// func (api *VPCRouterAPI) Create(value *sacloud.VPCRouter) (*sacloud.VPCRouter, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.create(api.createRequest(value), res)
// 	})
// }

// func (api *VPCRouterAPI) Read(id string) (*sacloud.VPCRouter, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.read(id, nil, res)
// 	})
// }

// func (api *VPCRouterAPI) Update(id string, value *sacloud.VPCRouter) (*sacloud.VPCRouter, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.update(id, api.createRequest(value), res)
// 	})
// }

// func (api *VPCRouterAPI) Delete(id string) (*sacloud.VPCRouter, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.delete(id, nil, res)
// 	})
// }

/************************************************
  Inner functions
************************************************/

func (api *VPCRouterAPI) setStateValue(setFunc func(*sacloud.Request)) *VPCRouterAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

//func (api *VPCRouterAPI) request(f func(*sacloud.Response) error) (*sacloud.VPCRouter, error) {
//	res := &sacloud.Response{}
//	err := f(res)
//	if err != nil {
//		return nil, err
//	}
//	return res.VPCRouter, nil
//}
//
//func (api *VPCRouterAPI) createRequest(value *sacloud.VPCRouter) *sacloud.Request {
//	req := &sacloud.Request{}
//	req.VPCRouter = value
//	return req
//}
