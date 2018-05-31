package configs

import (
	"os"
	"strconv"
)

var ServicePort, _ = strconv.Atoi(os.Getenv("SERVICE_PORT"))
var GSCredentialsFile = os.Getenv("GS_CREDENTIALS_FILE")
var GSBucketName = os.Getenv("GS_BUCKET_NAME")
