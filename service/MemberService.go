package service

import (
	"encoding/json"
	"fmt"
	"ginShop/tool"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) SendCode(phone string) bool {

	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	config := tool.GetConfig().Sms

	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		fmt.Println("newclientwithkey error")
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone

	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)

	fmt.Println(response)

	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	//TODO: 2021/9/21:09 finish send codes
	if response.Code == "OK" {
		return true
	}
	return false
}

func NewClientWithAccessKey(id string, key string, secret string) (interface{}, interface{}) {

	return nil, nil
}
