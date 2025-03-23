package scripts

type ParsedScript struct {
	Command         string `json:"command"`
	Title           string `json:"title"`
	ID              string `json:"id"`
	ScriptCommandID string `json:"scriptCommandId"`
	ScriptId        string `json:"scriptId"`
	ScriptName      string `json:"scriptName"`
}

type ScriptOutput struct {
	ScriptCommandID string `json:"scriptCommandId"`
	Output          string `json:"output"`
}

type ScriptExecutionMetric struct {
	Success int
	Failure int
}
