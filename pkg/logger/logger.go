package logger

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"gitlab.com/t9963/log-proxy-server/pkg/config"
)

type LogEntry struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	Stream map[string]string `json:"stream"`
	Values [][2]string       `json:"values"`
}

func SendLog(level, message string, additionalInfo map[string]string) error {
	stream := map[string]string{
		"logger":          additionalInfo["logger"],
		"service":         additionalInfo["service"],
		"application":     additionalInfo["application"],
		"logLevel":        level,
		"user_agent":      additionalInfo["user_agent"],
		"locationurlPath": additionalInfo["locationurlPath"],
		"userId":          additionalInfo["userId"],
	}

	for key, value := range additionalInfo {
		stream[key] = value
	}

	values := [][2]string{
		{fmt.Sprintf("%d000000", time.Now().UnixNano()/int64(time.Millisecond)), message},
	}

	logEntry := LogEntry{
		Streams: []Stream{
			{
				Stream: stream,
				Values: values,
			},
		},
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", config.BasicAuth).
		SetBody(logEntry).
		Post(config.LokiEndpoint)

	if err != nil {
		return fmt.Errorf("error sending log: %v", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("failed to send log: %s", resp.Status())
	}

	return nil
}
