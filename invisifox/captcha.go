package invisifox

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func New(apikey string) *CaptchaClient {
	return &CaptchaClient{
		APIKey:     apikey,
		HTTPClient: &http.Client{},
	}
}

func (client *CaptchaClient) GetBalance() (*BalanceResponse, error) {
	query := url.Values{}
	query.Add("token", client.APIKey)

	req, err := http.NewRequest("GET", "https://api.invisifox.com/balance?"+query.Encode(), nil)
	HandleError(err)

	req.Header.Set("authority", "api.invisifox.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var reply BalanceResponse

	err = json.NewDecoder(res.Body).Decode(&reply)
	HandleError(err)

	return &reply, nil
}

func (client *CaptchaClient) SolveCaptcha(pageurl, sitekey, proxy, rqdata, useragent, cookies, invisible string) (*SolveCaptchaResponse, error) {
	query := url.Values{}
	query.Add("token", client.APIKey)
	query.Add("siteKey", sitekey)
	query.Add("pageurl", pageurl)
	query.Add("proxy", proxy)
	query.Add("useragent", useragent)
	query.Add("cookies", cookies)
	query.Add("invisible", invisible)

	req, err := http.NewRequest("GET", "https://api.invisifox.com/hcaptcha?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("authority", "api.invisifox.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")

	res, err := client.HTTPClient.Do(req)
	HandleError(err)

	var reply SolveCaptchaResponse

	err = json.NewDecoder(res.Body).Decode(&reply)
	HandleError(err)

	return &reply, nil
}

func (client *CaptchaClient) GetSolution(taskId string) (*SolutionResponse, error) {
	query := url.Values{}
	query.Add("token", client.APIKey)
	query.Add("taskId", taskId)

	req, err := http.NewRequest("GET", "https://api.invisifox.com/solution?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("authority", "api.invisifox.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var reply SolutionResponse

	err = json.NewDecoder(res.Body).Decode(&reply)
	HandleError(err)

	return &reply, nil
}
