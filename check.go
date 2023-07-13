package capsolver_go

import (
	"errors"
	"fmt"
	"strings"
)

var (
	SUPPORT_CAPTCHA_TYPES = []string{

		"HCaptchaTask",
		"HCaptchaTaskProxyLess",
		"HCaptchaEnterpriseTask",
		"HCaptchaEnterpriseTaskProxyLess",
		"HCaptchaTurboTask",

		"FunCaptchaTask",
		"FunCaptchaTaskProxyLess",

		"GeeTestTask",
		"GeeTestTaskProxyLess",

		"ReCaptchaV2Task",
		"ReCaptchaV2TaskProxyLess",

		"ReCaptchaV2EnterpriseTaskProxyLess",
		"ReCaptchaV2EnterpriseTask",

		"ReCaptchaV3Task",
		"ReCaptchaV3TaskProxyLess",

		"ReCaptchaV3EnterpriseTask",
		"ReCaptchaV3EnterpriseTaskProxyLess",

		"MtCaptchaTask",
		"MtCaptchaTaskProxyLess",

		"DataDomeSliderTask",

		"AntiCloudflareTask",

		"AntiKasadaTask",

		"AntiAkamaiBMPTask",

		"ImageToTextTask",

		"HCaptchaClassification",

		"FunCaptchaClassification",

		"AwsWafClassification",
	}
)

const (
	websiteKey = "websiteKey"
	websiteURL = "websiteURL"
)

func checkParams(params map[string]any) error {

	captchaType, ok := params["type"].(string)
	exists := false
	for _, v := range SUPPORT_CAPTCHA_TYPES {
		if v == captchaType {
			exists = true
		}
	}
	if !exists {
		return errors.New("unSupported Type " + captchaType + "you need to pay attention to capitalization,current support types fellow\n" + formatTaskTypes())
	}
	if !ok {
		return formatParamError(captchaType, "type")
	}
	captchaType = strings.ToLower(captchaType)
	if strings.Contains(captchaType, "recaptcha") {
		if captchaType == "recaptchaclassification" {
		} else {
			return checkNormalCaptcha(params)
		}
	}
	if strings.Contains(captchaType, "hcaptcha") {
		if captchaType == "hcaptchaclassification" {
			if _, ok = params["queries"]; !ok {
				return formatParamError(captchaType, "queries")
			}
			if _, ok = params["question"]; !ok {
				return formatParamError(captchaType, "question")
			}
		} else {
			return checkNormalCaptcha(params)
		}
	}
	if strings.Contains(captchaType, "funcaptcha") {
		if captchaType == "funcaptchaclassification" {
			if _, ok = params["images"]; !ok {
				return formatParamError(captchaType, "images")
			}
		} else {
			if _, ok := params[websiteURL]; !ok {
				return formatParamError(captchaType, websiteURL)
			}
			if _, ok := params["websitePublicKey"]; !ok {
				return formatParamError(captchaType, "websitePublicKey")
			}
			return nil
		}
	}

	if strings.Contains(captchaType, "mtcaptcha") {
		return checkNormalCaptcha(params)
	}

	if strings.Contains(captchaType, "geetesttask") {
		if _, ok = params["gt"]; !ok {
			return formatParamError(captchaType, "gt")
		}
		if _, ok = params["challenge"]; !ok {
			return formatParamError(captchaType, "challenge")
		}
	}
	if strings.Contains(captchaType, "datadom") {
		if _, ok = params["proxy"]; !ok {
			return formatParamError(captchaType, "proxy")
		}
		if _, ok = params["useragent"]; !ok {
			return formatParamError(captchaType, "userAgent")
		}
	}
	if strings.Contains(captchaType, "anticloudflare") {
		if _, ok = params["metadata"]; !ok {
			return formatParamError(captchaType, "metadata")
		}
		if _, ok = params["proxy"]; !ok {
			return formatParamError(captchaType, "proxy")
		}
	}
	if strings.Contains(captchaType, "antikasada") {
		if _, ok = params["pageURL"]; !ok {
			return formatParamError(captchaType, "pageURL")
		}
		if _, ok = params["proxy"]; !ok {
			return formatParamError(captchaType, "proxy")
		}
	}
	if strings.Contains(captchaType, "antiakamaibmp") {
		if _, ok = params["packageName"]; !ok {
			return formatParamError(captchaType, "packageName")
		}
	}
	return nil
}

func checkNormalCaptcha(params map[string]any) error {
	captchaType, _ := params["type"].(string)
	if _, ok := params[websiteKey]; !ok {
		return formatParamError(captchaType, websiteKey)
	}
	if _, ok := params[websiteURL]; !ok {
		return formatParamError(captchaType, websiteURL)
	}
	return nil
}

func formatTaskTypes() string {

	t_list := []string{}
	for i, v := range SUPPORT_CAPTCHA_TYPES {
		t_list = append(t_list, fmt.Sprintf("%d,%s", i, v))
	}
	return strings.Join(t_list, "\n")
}

func formatParamError(captchaType, paramKey string) error {
	return errors.New(captchaType + fmt.Sprintf(" need %s param", paramKey))
}
