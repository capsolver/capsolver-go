package capsolver_go

type CapSolver struct {
	ApiKey string
}

type solution struct {
	UserAgent string  `json:"userAgent,omitempty"`
	SecChUa   string  `json:"secChUa,omitempty"`
	Token     string  `json:"token,omitempty"`
	Cookie    string  `json:"cookie,omitempty"`
	Type      string  `json:"type,omitempty"`
	Objects   []int32 `json:"objects,omitempty"`

	// GeeTest
	Challenge     string `json:"challenge,omitempty"`
	Validate      string `json:"validate,omitempty"`
	SecCode       string `json:"seccode,omitempty"`
	CaptchaId     string `json:"captcha_id,omitempty"`
	CaptchaOutput string `json:"captcha_output,omitempty"`
	GenTime       string `json:"gen_time,omitempty"`
	LotNumber     string `json:"lot_number,omitempty"`
	PassToken     string `json:"pass_token,omitempty"`
	RiskType      string `json:"risk_Type,omitempty"`

	// reCaptcha
	CreateTime         int64  `json:"createTime,omitempty"`
	GRecaptchaResponse string `json:"gRecaptchaResponse,omitempty"`
	RecaptchaCaT       string `json:"recaptcha-ca-t,omitempty"`
	RecaptchaCaE       string `json:"recaptcha-ca-e,omitempty"`

	// cloudflare
	Cookies map[string]string `json:"cookies,omitempty"`

	// ImageToText
	Text    string   `json:"text,omitempty"`
	Answers []string `json:"answers,omitempty"`

	// ReCaptchaV2Classification
	Size int32 `json:"size,omitempty"`

	// AwsWafClassification
	Box      []float32 `json:"box,omitempty"`
	Distance int32     `json:"distance,omitempty"`

	// VisionEngine
	Angle int32              `json:"angle,omitempty"`
	Rects []map[string]int32 `json:"rects,omitempty"`
}

type CapSolverResponse struct {
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

type CapSolverTask struct {
	Type       string `json:"type"`
	WebsiteURL string `json:"websiteURL,omitempty"`
	WebsiteKey string `json:"websiteKey,omitempty"`
	Proxy      string `json:"proxy,omitempty"`
	UserAgent  string `json:"userAgent,omitempty"`

	// GeeTest
	Gt                        string `json:"gt,omitempty"`
	Challenge                 string `json:"challenge,omitempty"`
	CaptchaId                 string `json:"captchaId,omitempty"`
	GeetestApiServerSubdomain string `json:"geetestApiServerSubdomain,omitempty"`

	// reCaptcha
	PageAction        string             `json:"pageAction,omitempty"`
	EnterPrisePayload *enterPrisePayload `json:"enterprisePayload,omitempty"`
	IsInvisible       bool               `json:"isInvisible,omitempty"`
	IsSession         bool               `json:"isSession,omitempty"`
	ApiDomain         string             `json:"apiDomain,omitempty"`

	// DataDome
	CaptchaUrl string `json:"captchaUrl,omitempty"`

	// AWS WAF
	AwsKey           string `json:"awsKey,omitempty"`
	AwsIv            string `json:"awsIv,omitempty"`
	AwsContext       string `json:"awsContext,omitempty"`
	AwsChallengeJS   string `json:"awsChallengeJS,omitempty"`
	AwsApiJs         string `json:"awsApiJs,omitempty"`
	AwsProblemUrl    string `json:"awsProblemUrl,omitempty"`
	AwsApiKey        string `json:"awsApiKey,omitempty"`
	AwsExistingToken string `json:"awsExistingToken,omitempty"`

	// cloudflare
	Metadata map[string]string `json:"metadata"`
	Html     string            `json:"html"`

	Image    string   `json:"image,omitempty"`
	Images   []string `json:"images,omitempty"`
	Question string   `json:"question,omitempty"`

	// ImageToTextTask
	Body   string `json:"body,omitempty"`
	Module string `json:"module,omitempty"`

	// VisionEngine
	ImageBackground string `json:"imageBackground,omitempty"`
}

type enterPrisePayload struct {
	S string `json:"s,omitempty"`
}
