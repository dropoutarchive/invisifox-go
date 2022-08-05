package invisifox

import "net/http"

type CaptchaClient struct {
	APIKey     string
	HTTPClient *http.Client
}

type BalanceResponse struct {
	CaptchaBalance float64 `json:"captchaBalance"`
	ProxyBalance   float64 `json:"proxyBalance"`
}

type SolveCaptchaResponse struct {
	Status string `json:"status"`
	TaskID string `json:"taskId"`
}

type SolutionResponse struct {
	Solution string `json:"solution"`
	Status   string `json:"status"`
	TaskID   string `json:"taskId"`
}
