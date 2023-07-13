package capsolver_go

import (
	"encoding/base64"
	"log"
	"os"
	"testing"
)

func TestTask(t *testing.T) {

	capSolver := CapSolver{}
	s, err := capSolver.Solve(
		map[string]any{
			"type":       "ReCaptchaV2TaskProxyLess",
			"websiteURL": "https://www.google.com/recaptcha/api2/demo",
			"websiteKey": "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution.GRecaptchaResponse)
}

func TestFunCaptcha(t *testing.T) {
	capSolver := CapSolver{}
	s, err := capSolver.Solve(
		map[string]any{
			"type":             "FunCaptchaTask",
			"websiteURL":       "",
			"websitePublicKey": "",
			"proxy":            "",
		})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution)

}
func TestHCaptcha(t *testing.T) {

	capSolver := CapSolver{}
	s, err := capSolver.Solve(
		map[string]any{
			"type":       "HCaptchaTurboTask",
			"websiteURL": "https://www.discord.com",
			"websiteKey": "4c672d35-0701-42b2-88c3-78380b0db560",
			"proxy":      "your proxy",
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution.GRecaptchaResponse)

}

func TestGeeTest(t *testing.T) {

	capSolver := CapSolver{}
	s, err := capSolver.Solve(
		map[string]any{
			"type": "GeeTestTaskProxyLess",
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution.GRecaptchaResponse)
}

func TestDataDom(t *testing.T) {

	capSolver := CapSolver{}
	s, err := capSolver.Solve(
		map[string]any{
			"type": "DataDomeSliderTask",
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution.GRecaptchaResponse)
}

func TestAntiCloudflareTask(t *testing.T) {

	capSolver := CapSolver{}
	s, err := capSolver.Solve(
		map[string]any{
			"type": "AntiCloudflareTask",
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution.GRecaptchaResponse)
}

func TestAntiKasadaTask(t *testing.T) {

	capSolver := CapSolver{}
	s, err := capSolver.Solve(
		map[string]any{
			"type": "AntiKasadaTask",
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution.GRecaptchaResponse)
}

func TestAntiAkamaiBMPTask(t *testing.T) {

	capSolver := CapSolver{}
	s, err := capSolver.Solve(
		map[string]any{
			"type": "AntiAkamaiBMPTask",
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution.GRecaptchaResponse)
}

func TestBalance(t *testing.T) {

	capSolver := CapSolver{}
	b, err := capSolver.Balance()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(b.Balance)
}

func TestRecognition(t *testing.T) {

	capSolver := CapSolver{}
	b, err := os.ReadFile("queue-it.jpg")
	if err != nil {
		panic(err)
	}
	s, err := capSolver.Solve(
		map[string]any{"type": "ImageToTextTask",
			"module": "queueit",
			"body":   base64.StdEncoding.EncodeToString(b),
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution.Text)
}

func TestHCaptchaClassfication(t *testing.T) {
	capSolver := CapSolver{}
	b, err := os.ReadFile("queue-it.jpg")
	if err != nil {
		panic(err)
	}
	s, err := capSolver.Solve(
		map[string]any{
			"type":     "HCaptchaClassification",
			"question": "Please click each image containing a truck",
			"queries": []string{
				base64.StdEncoding.EncodeToString(b),
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Solution)
}
