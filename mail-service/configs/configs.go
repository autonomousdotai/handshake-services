package configs

import (
	"os"
	"strconv"
)

var ServicePort, _ = strconv.Atoi(os.Getenv("SERVICE_PORT"))
var SendgridApiKey = os.Getenv("SENDGRID_API_KEY")
