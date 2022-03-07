package utils

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-micro.dev/v4/logger"
)

// 认证中心的token验证接口
const serverValidPath = "/cas/serviceValidate"

// Cas cas验证工具
type Cas struct {
	AuthServer string
}

// ValidToken 验证认证中心传过来的回调token
func (cas *Cas) ValidToken(service, token string) (AuthenticationSuccess, error) {
	// 1. 构造请求验证token
	urlPath := fmt.Sprintf("%s%s?ticket=%s&service=%s", cas.AuthServer, serverValidPath, token, service)
	logger.Debug(urlPath)
	resp, err := http.Get(urlPath)
	if err != nil {
		logger.Warn("Cas valid error: ", err)
		return AuthenticationSuccess{}, err
	}
	respData, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logger.Warn("Cas response body read error: ", err)
		return AuthenticationSuccess{}, err
	}
	info, err := cas.parseServerResponse(respData)
	if err != nil {
		logger.Warn("Cas response parse error: ", err, " with data: ", string(respData))
		return AuthenticationSuccess{}, err
	}
	return info, nil
}

// serviceResponse cas验证的返回值
type serviceResponse struct {
	XMLName xml.Name              `xml:"serviceResponse"`
	Success AuthenticationSuccess `xml:"authenticationSuccess"`
	Failure authenticationFailure `xml:"authenticationFailure"`
}

// AuthenticationSuccess 验证成功的字段
type AuthenticationSuccess struct {
	XMLName        xml.Name `xml:"authenticationSuccess"`
	User           string   `xml:"user"`
	Name           string   `xml:"name"`
	EmployeeNumber string   `xml:"employeeNumber"`
}

type authenticationFailure string

func (cas *Cas) parseServerResponse(data []byte) (AuthenticationSuccess, error) {
	var xmlResponse serviceResponse
	if err := xml.Unmarshal(data, &xmlResponse); err != nil {
		return AuthenticationSuccess{}, err
	}
	if xmlResponse.Failure != "" {
		return AuthenticationSuccess{}, errors.New(string(xmlResponse.Failure))
	}
	return xmlResponse.Success, nil
}
