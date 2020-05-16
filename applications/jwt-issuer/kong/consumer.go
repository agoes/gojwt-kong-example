package kong

import (
	"encoding/json"
	"fmt"
)

const (
	consumerEndpointPath = "consumers"
)

// Consumer is unmarshaled JSON response from kong consumer jwt credentials
type Consumer struct {
	RsaPublicKey string `json:"rsa_public_key"`
	Key          string `json:"key"`
	Algorithm    string `json:"algorithm"`
}

// GetConsumerJwtCredentials will return kong consumer jwt information
func GetConsumerJwtCredentials(username string, key string) (int, *Consumer) {
	endpoint := fmt.Sprintf("%s/%s/jwt/%s", getEndpointPath(), username, key)
	statusCode, body := createRequest(endpoint)
	credentials := new(Consumer)
	json.Unmarshal(body, &credentials)

	return statusCode, credentials
}

func getEndpointPath() string {
	baseURL, err := getBaseURL()
	fatal(err)

	return fmt.Sprintf("%s/%s", baseURL, consumerEndpointPath)
}
