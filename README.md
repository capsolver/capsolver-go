# Capsolver
CapSolver official go library


## Supported CAPTCHA types:
- HCaptcha
- FunCaptcha
- Geetest
- ReCaptchaV2
- ReCaptchav3
- MtCaptcha
- Datadom
- Cloudflare
- Kasada
- Akamai BMP


## Installation

You don't need this source code unless you want to modify the package. If you just
want to use the package, just run:

```sh
go get github.com/capsolver/capsolver-go
```



## Usage

```bash
export CAPSOLVER_API_KEY='...'
```

Or set `capsolver apiKey` to its value:

```go
 capSolver := CapSolver{}
 //capSolver := CapSolver{apiKey:"..."} 
 s, err := capSolver.Solve(
            map[string]any{
            "type":       "ReCaptchaV2taskProxyLess",
            "websiteURL": "https://www.google.com/recaptcha/api2/demo",
            "websiteKey": "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
 })
 if err != nil {
    log.Fatal(err)
 }
 log.Println(s.Solution.GRecaptchaResponse)
	

 // recognition 
 b, err := os.ReadFile("queue-it.jpg")
 if err != nil {
    panic(err)
 }
 s, err := capSolver.Solve(
    map[string]any{
	"type": "ImageToTextTask",
        "module": "queueit",
        "body":   base64.StdEncoding.EncodeToString(b),
 })
 if err != nil {
    log.Fatal(err)
 }
 log.Println(s.Solution.Text)
		
  // get balance 
  b, err := capSolver.Balance()
  if err != nil {
    log.Fatal(err)
  }
  log.Println(b.Balance)
```


