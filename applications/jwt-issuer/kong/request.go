package kong

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	timeoutEnvName        = "KONG_REQUEST_TIMEOUT"
	kongAdmBaseURLEnvName = "KONG_ADMIN_BASE_URL"
	defaultTimeout        = 0
)

var (
	httpClient = &http.Client{
		Timeout: time.Second * getTimeout(),
	}
)

func getTimeout() time.Duration {
	timeout := defaultTimeout
	fromEnv := os.Getenv(timeoutEnvName)
	if fromEnv != "" {
		envTimeout, err := strconv.Atoi(fromEnv)
		fatal(err)

		timeout = envTimeout
	}

	return time.Duration(timeout) * time.Second
}

func getBaseURL() (string, error) {
	baseURL := os.Getenv(kongAdmBaseURLEnvName)
	if baseURL != "" {
		return baseURL, nil
	}

	return "", fmt.Errorf("Missing or empty env variable: %s", kongAdmBaseURLEnvName)
}

func createRequest(endpoint string) (int, []byte) {
	response, err := httpClient.Get(endpoint)
	fatal(err)

	body, _ := ioutil.ReadAll(response.Body)
	statusCode := response.StatusCode

	return statusCode, body
}

func fatal(err error) {
	if err != nil {
		log.Fatalf("Kong error : %s", err)
	}
}
