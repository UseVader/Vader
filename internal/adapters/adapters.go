package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/UseVader/vader/internal/constants"
	"github.com/UseVader/vader/internal/scripts"
)

type PreviewCode struct {
	OutputURL string `json:"outputURL"`
}

func VaderAdapter(outputs []scripts.ScriptOutput) (string, error) {
	url := constants.URL("outputs")

	jsonData, err := json.Marshal(outputs)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return "", nil
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error posting to Vader:", err)
		return "", nil
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", nil
	}

	var outputCode PreviewCode

	err = json.Unmarshal([]byte(string(body)), &outputCode)
	if err != nil {
		return "Failed to generate preview link", err
	}

	path := fmt.Sprintf("preview/%s", outputCode.OutputURL)
	return constants.FEURL(path), nil
}
