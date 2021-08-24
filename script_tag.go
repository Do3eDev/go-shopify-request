package go_shopify_request

import (
	"encoding/json"
	"github.com/Do3eDev/go-gcloud-engine"
	"net/http"
	"strconv"
)

func DeleteAllScriptTag(Env, SHOPIFY_API_VERSION string, scripts ScriptTags, accessToken string, domain string, r *http.Request) (stt int) {
	for j := 0; j < len(scripts.ScriptTags); j++ {
		urlDelete := "https://" + domain + "/admin" + SHOPIFY_API_VERSION + "/script_tags/" + strconv.FormatInt(scripts.ScriptTags[j].Id, 10) + ".json"
		stt1 := DeleteScriptTag(Env, urlDelete, accessToken, r)
		if stt != 200 {
			stt = stt1
		}
	}
	return
}

func AddScriptTag(Env, SHOPIFY_API_VERSION string, jsonStr []byte, domain string, accessToken string, r *http.Request) (status int, responseBody []byte, err error, Header map[string][]string) {
	scriptUrl := "https://" + domain + "/admin" + SHOPIFY_API_VERSION + "/script_tags.json"
	hder := map[string]string{"Content-Type": "application/json", "X-Shopify-Access-Token": accessToken, "Content-Length": strconv.Itoa(len(jsonStr))}
	return go_gcloud_engine.RequestCustomer(Env, "POST", scriptUrl, jsonStr, hder, r)
}

func DeleteScriptTag(Env string, urlDelete string, accessToken string, r *http.Request) (stt int) {
	hder := map[string]string{"Content-Type": "application/json", "X-Shopify-Access-Token": accessToken}
	stt, _, _, _ = go_gcloud_engine.RequestCustomer(Env, "DELETE", urlDelete, nil, hder, r)
	return
}

func ListAllScriptTag(Env, SHOPIFY_API_VERSION string, domain string, accessToken string, r *http.Request) (stt int, rs ScriptTags, err error) {
	scriptUrl := "https://" + domain + "/admin" + SHOPIFY_API_VERSION + "/script_tags.json"
	hder := map[string]string{"Content-Type": "application/json", "X-Shopify-Access-Token": accessToken}
	status, responseBody, err, _ := go_gcloud_engine.RequestCustomer(Env, "GET", scriptUrl, nil, hder, r)
	var scripts ScriptTags
	if status == 200 && err == nil {
		_ = json.Unmarshal(responseBody, &scripts)
	}
	return status, scripts, err
}
