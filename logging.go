package ensurecp

import (
	"encoding/json"
	"time"
)

type EnsurecpConfig struct {
	EnableLogging bool
}

type CopyLogOutput struct {
	FileList []CopyLog `json:"list"`
}

type CopyLog struct {
	Source      string `json:"source"`
	Destination string `json:"destionation"`
	Hash        string `json:"hash"`
	Timestamp   time.Time `json:"timestamp"`
	Size        int64   `json:"size"`
}

var localConfig = EnsurecpConfig{true}
var currentLog = []CopyLog{}

func SetLogging(mode bool) {
	localConfig.EnableLogging = mode;
}

func ClearLog() {
	currentLog = []CopyLog{}
}

func ExportLog() string {
	if len(currentLog) > 0 {
		out, err := json.MarshalIndent(CopyLogOutput{currentLog}, "", " ")

		if err != nil {
			return ""
		}

		return string(out)
	}

	return ""
}

