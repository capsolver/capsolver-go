package capsolver_go

import "os"

var (
	ApiKey  = os.Getenv("CAPSOLVER_API_KEY")
	ApiHost = os.Getenv("CAPSOLVER_API_HOST")
)

const (
	STATUS_READY    = "ready"
	CREATE_TASK_URI = "/createTask"
	GET_TASK_URI    = "/getTaskResult"
	BALANCE_URI     = "/getBalance"
	TASK_TIMEOUT    = 45
)
