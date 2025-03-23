package scripts

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/UseVader/vader/internal/constants"
	"github.com/UseVader/vader/internal/runners"
)

func GetScript(scriptId string) ([]ParsedScript, error) {
	url := constants.URL("scripts/" + scriptId + "/parsed")
	res, err := http.Get(url)
	if err != nil {
		return []ParsedScript{}, err
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	var commands []ParsedScript

	for scanner.Scan() {
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &commands)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err)
			continue
		}
	}

	return commands, nil
}

func ExecuteScript(commands []ParsedScript) ([]ScriptOutput, ScriptExecutionMetric) {
	var outputs []ScriptOutput
	success, failure := 0, 0
	for _, command := range commands {

		fmt.Println("EXECUTING CMD: " + command.Title + "\n")
		output, err := runners.Bash(command.Command)

		if err != nil {
			log.Fatal(err)
			failure += 1
			return []ScriptOutput{}, ScriptExecutionMetric{Success: 0, Failure: 0}
		}

		success += 1
		outputs = append(outputs, ScriptOutput{
			ScriptCommandID: command.ScriptCommandID,
			Output:          output,
		})
	}
	return outputs, ScriptExecutionMetric{Success: success, Failure: failure}
}
