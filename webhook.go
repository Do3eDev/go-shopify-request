package go_shopify_request

import (
	"encoding/json"
	"fmt"
	"github.com/Do3eDev/go-gcloud-engine"
	"net/http"
	"strings"
)

func RegisterWebhookAppUninstallShopUpdate(Env, SHOPIFY_API_VERSION, apiUrl string, request *http.Request, shopDomain string, appCode string, accessToken string, topics []string) {
	aCode := strings.Split(appCode, "-")[0]
	header := map[string]string{"Content-Type": "application/json", "X-Shopify-Access-Token": accessToken}
	webHookUrl := "https://" + shopDomain + "/admin" + SHOPIFY_API_VERSION + "/webhooks.json"
	var _RegisterWebhook ShopRegisterWebhook
	_RegisterWebhook.Webhook.Format = "json"
	topics = append(topics, "app/uninstalled")
	topics = append(topics, "shop/update")
	for _, _topic := range topics {
		_RegisterWebhook.Webhook.Topic = _topic
		_RegisterWebhook.Webhook.Address = fmt.Sprint(apiUrl, "/webhook/", _topic, "/", aCode, "/", shopDomain)
		jsonByte, _ := json.Marshal(_RegisterWebhook)
		_, _, _, _ = go_gcloud_engine.RequestCustomer(Env, "POST", webHookUrl, jsonByte, header, request)
	}
}

func RegisterWebhookByEvents(Env, SHOPIFY_API_VERSION, apiUrl string, request *http.Request, shopDomain string, appCode string, accessToken string, events []string) {
	aCode := strings.Split(appCode, "-")[0]
	header := map[string]string{"Content-Type": "application/json", "X-Shopify-Access-Token": accessToken}
	webHookUrl := "https://" + shopDomain + "/admin" + SHOPIFY_API_VERSION + "/webhooks.json"
	var _RegisterWebhook ShopRegisterWebhook
	_RegisterWebhook.Webhook.Format = "json"
	var topics []string
	for _, ev1 := range events {
		topics = append(topics, ev1)
	}
	for _, _topic := range topics {
		_RegisterWebhook.Webhook.Topic = _topic
		_RegisterWebhook.Webhook.Address = fmt.Sprint(apiUrl, "/webhook/", _topic, "/", aCode, "/", shopDomain)
		jsonByte, _ := json.Marshal(_RegisterWebhook)
		_, _, _, _ = go_gcloud_engine.RequestCustomer(Env, "POST", webHookUrl, jsonByte, header, request)
	}
}

func ShopifyEventWebhookDelete(Env, SHOPIFY_API_VERSION string, request *http.Request, accessToken, shopDomain, event string) {
	webHookUrl := "https://" + shopDomain + "/admin" + SHOPIFY_API_VERSION + "/webhooks.json"
	urlGetTopic := webHookUrl + "?topic=" + event
	header := map[string]string{"Content-Type": "application/json", "X-Shopify-Access-Token": accessToken}
	stt, d1, err, _ := go_gcloud_engine.RequestCustomer(Env, "GET", urlGetTopic, nil, header, request)
	if stt == 200 && err == nil {
		var w1 struct {
			Webhooks []struct {
				Id    int64  `json:"id"`
				Topic string `json:"topic"`
			} `json:"webhooks"`
		}
		_ = json.Unmarshal(d1, &w1)
		for _, webhook := range w1.Webhooks {
			if webhook.Topic == event && webhook.Id > 0 {
				webhookId := fmt.Sprint(webhook.Id)
				deleteWebHookUrl := "https://" + shopDomain + "/admin" + SHOPIFY_API_VERSION + "/webhooks/" + webhookId + ".json"
				_, _, _, _ = go_gcloud_engine.RequestCustomer(Env, "DELETE", deleteWebHookUrl, nil, header, request)
			}
		}
	}
}
