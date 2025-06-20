// Code generated by Lark OpenAPI.

package larkbase

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"net/http"
)

type V2 struct {
	AppRole *appRole // app.role
}

func New(config *larkcore.Config) *V2 {
	return &V2{
		AppRole: &appRole{config: config},
	}
}

type appRole struct {
	config *larkcore.Config
}

// Create
//
// - 新增自定义角色
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=create&project=base&resource=app.role&version=v2
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/basev2/create_appRole.go
func (a *appRole) Create(ctx context.Context, req *CreateAppRoleReq, options ...larkcore.RequestOptionFunc) (*CreateAppRoleResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/base/v2/apps/:app_token/roles"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateAppRoleResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// List
//
// - 列出自定义角色
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=list&project=base&resource=app.role&version=v2
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/basev2/list_appRole.go
func (a *appRole) List(ctx context.Context, req *ListAppRoleReq, options ...larkcore.RequestOptionFunc) (*ListAppRoleResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/base/v2/apps/:app_token/roles"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListAppRoleResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *appRole) ListByIterator(ctx context.Context, req *ListAppRoleReq, options ...larkcore.RequestOptionFunc) (*ListAppRoleIterator, error) {
	return &ListAppRoleIterator{
		ctx:      ctx,
		req:      req,
		listFunc: a.List,
		options:  options,
		limit:    req.Limit}, nil
}

// Update
//
// - 更新自定义角色
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=update&project=base&resource=app.role&version=v2
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/basev2/update_appRole.go
func (a *appRole) Update(ctx context.Context, req *UpdateAppRoleReq, options ...larkcore.RequestOptionFunc) (*UpdateAppRoleResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/base/v2/apps/:app_token/roles/:role_id"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateAppRoleResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
