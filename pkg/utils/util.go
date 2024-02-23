package utils

import (
	"os"
	"strconv"
)

type RestBody struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ConvertEnvToInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func ConvertEnvToBool(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func ConvertStrToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
