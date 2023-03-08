package capsolver_go

type CapSolver struct {
	apiKey string
}

type solution struct {
	Object             []bool    `json:"objects,omitempty"`
	Box                []float32 `json:"box,omitempty"`
	ImageSizes         []int32   `json:"imageSize,omitempty"`
	Text               string    `json:"text,omitempty"`
	UserAgent          string    `json:"userAgent,omitempty"`
	ExpireTime         int64     `json:"expireTime,omitempty"`
	GRecaptchaResponse string    `json:"gRecaptchaResponse,omitempty"`
	Challenge          string    `json:"challenge,omitempty"`
	Validate           string    `json:"validate,omitempty"`
	CaptchaId          string    `json:"captcha-id,omitempty"`
	CaptchaOutput      string    `json:"captcha-output,omitempty"`
	GenTime            string    `json:"gen_time,omitempty"`
	LogNumber          string    `json:"log_number,omitempty"`
	PassToken          string    `json:"pass_token,omitempty"`
	RiskType           string    `json:"risk_Type,omitempty"`
	Token              string    `json:"token,omitempty"`
	Type               string    `json:"type,omitempty"`
}

type capSolverResponse struct {
	ErrorId          int32     `json:"errorId"`
	ErrorCode        string    `json:"errorCode"`
	ErrorDescription string    `json:"errorDescription,omitempty"`
	Status           string    `json:"status,omitempty"`
	Solution         *solution `json:"solution,omitempty"`
	TaskId           string    `json:"taskId,omitempty"`
	Balance          float32   `json:"balance,omitempty"`
	Packages         []string  `json:"packages,omitempty"`
}

type capSolverRequest struct {
	ClientKey string          `json:"ClientKey"`
	Task      *map[string]any `json:"task,omitempty"`
	TaskId    string          `json:"taskId,omitempty"`
}
type enterPrisePayload struct {
	S      string `json:"s,omitempty"`
	Rqdata string `json:"rqdata,omitempty"`
}
type cookies struct {
	Cookies []cookieItem `json:"cookies,omitempty"`
}
type cookieItem struct {
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}

type CapSolverTask struct {
	Type                      string             `json:"type"`
	WebsiteURL                string             `json:"websiteURL,omitempty"`
	WebsiteKey                string             `json:"websiteKey,omitempty"`
	Proxy                     string             `json:"proxy,omitempty"`
	EnterPrisePayload         *enterPrisePayload `json:"enterprisePayload,omitempty"`
	IsInvisible               bool               `json:"isInvisible,omitempty"`
	ApiDomain                 string             `json:"apiDomain,omitempty"`
	UserAgent                 string             `json:"userAgent,omitempty"`
	Cookies                   *cookies           `json:"cookies,omitempty"`
	Module                    string             `json:"module,omitempty"`
	Body                      string             `json:"body,omitempty"`
	Question                  string             `json:"question,omitempty"`
	Queries                   []string           `json:"Queries,omitempty"`
	PageAction                string             `json:"pageAction,omitempty"`
	MinScore                  float32            `json:"MinScore,omitempty"`
	Gt                        string             `json:"gt,omitempty"`
	Challenge                 string             `json:"challenge,omitempty"`
	GeetestApiServerSubdomain string             `json:"geetestApiServerSubdomain,omitempty"`
	CaptchaId                 string             `json:"captchaId,omitempty"`
	CaptchaUrl                string             `json:"captchaUrl,omitempty"`
	Metadata                  map[string]string  `json:"metadata"`
	Html                      string             `json:"html"`
}
