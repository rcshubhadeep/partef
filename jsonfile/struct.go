package jsonfile

// TestDesc - The inner strcutrue to hold all the test cases
type TestDesc struct {
	EndPoint   string                 `json:"end_point"`
	Method     string                 `json:"method"`
	Data       map[string]interface{} `json:"data"`
	ReturnCode int                    `json:"return_code"`
	ReturnData map[string]interface{} `json:"return_data"`
	Concurency int                    `json:"concurreny"`
}

// APIDesc - the main structure
type APIDesc struct {
	Name    string     `json:"name"`
	Version string     `json:"version"`
	Tests   []TestDesc `json:"tests"`
}
