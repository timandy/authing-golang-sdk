package management

import (
	"crypto/tls"
	"encoding/json"
	"sync"
	"time"

	"github.com/Authing/authing-golang-sdk/v3/constant"
	"github.com/Authing/authing-golang-sdk/v3/dto"
	"github.com/Authing/authing-golang-sdk/v3/util"
	"github.com/Authing/authing-golang-sdk/v3/util/cache"
	"github.com/golang-jwt/jwt/v5"
	"github.com/valyala/fasthttp"
)

type JwtClaims struct {
	*jwt.RegisteredClaims
	//用户编号
	UID      string
	Username string
}

func GetAccessToken(client *ManagementClient) (string, error) {
	// 从缓存获取token
	cacheToken, b := cache.GetCache(constant.TokenCacheKeyPrefix + client.options.AccessKeyId)
	if b && cacheToken != nil {
		return cacheToken.(string), nil
	}
	// 从服务获取token，加锁
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	cacheToken, b = cache.GetCache(constant.TokenCacheKeyPrefix + client.options.AccessKeyId)
	if b && cacheToken != nil {
		return cacheToken.(string), nil
	}
	resp, err := QueryAccessToken(client)
	if err != nil {
		return "", err
	}

	if token, _ := jwt.Parse(resp.Data.AccessToken, nil); token != nil {
		userPoolId := token.Claims.(jwt.MapClaims)["scoped_userpool_id"]
		client.userPoolId = userPoolId.(string)
	}

	cache.SetCache(constant.TokenCacheKeyPrefix+client.options.AccessKeyId, resp.Data.AccessToken, time.Duration(resp.Data.ExpiresIn*int(time.Second)))
	return resp.Data.AccessToken, nil
}

func QueryAccessToken(client *ManagementClient) (*dto.GetManagementTokenRespDto, error) {
	variables := map[string]interface{}{
		"accessKeyId":     client.options.AccessKeyId,
		"accessKeySecret": client.options.AccessKeySecret,
	}

	b, err := client.SendHttpRequest("/api/v3/get-management-token", fasthttp.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	var r dto.GetManagementTokenRespDto
	if b != nil {
		json.Unmarshal(b, &r)
	}
	return &r, nil
}

func (client *ManagementClient) SendHttpRequest(url string, method string, reqDto interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	reqJsonBytes, err := json.Marshal(&reqDto)
	if err != nil {
		return nil, err
	}
	if method == fasthttp.MethodGet {
		variables := make(map[string]interface{})
		err = json.Unmarshal(reqJsonBytes, &variables)
		if err != nil {
			return nil, err
		}
		queryString := util.GetQueryString2(variables)
		if queryString != "" {
			url += "?" + queryString
		}
	}

	req.Header.SetMethod(method)
	req.SetRequestURI(client.options.Host + url)

	req.Header.Add("x-authing-app-tenant-id", client.options.TenantId)
	//req.Header.Add("x-authing-request-from", client.options.RequestFrom)
	req.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	req.Header.Add("x-authing-lang", client.options.Lang)
	if url != "/api/v3/get-management-token" {
		token, _ := GetAccessToken(client)
		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Add("x-authing-userpool-id", client.userPoolId)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	if method != fasthttp.MethodGet {
		req.SetBody(reqJsonBytes)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = client.httpClient.DoTimeout(req, resp, client.options.ReadTimeout)
	if err != nil {
		resultMap := make(map[string]interface{})
		if err == fasthttp.ErrTimeout {
			resultMap["statusCode"] = 504
			resultMap["message"] = "请求超时"
		} else {
			resultMap["statusCode"] = 500
			resultMap["message"] = err.Error()
		}
		b, err := json.Marshal(resultMap)
		if err != nil {
			return nil, err
		}
		return b, err
	}
	body := resp.Body()
	return body, err
}

func (client *ManagementClient) createHttpClient() *fasthttp.Client {
	options := client.options
	createClientFunc := options.CreateClientFunc
	if createClientFunc != nil {
		return createClientFunc(options)
	}
	return &fasthttp.Client{
		TLSConfig: &tls.Config{InsecureSkipVerify: options.InsecureSkipVerify},
	}
}
