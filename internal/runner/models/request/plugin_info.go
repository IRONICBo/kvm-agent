package request

type PluginInfo struct {
	ExecCommand      string  `json:"execCommand"`
	ExecNumber       int     `json:"execNumber"`
	ResponseUrl      string  `json:"responseUrl"`
	ExecResultIdList []int64 `json:"execResultIdList"`
	PlugType         int     `json:"plugType"`
	PlugResultType   int     `json:"plugResultType"`
	ExecParams       string  `json:"execParams"`
	PlugId           int64   `json:"plugId"`
	RecordId         int64   `json:"recordId"`
	PlugPath         string  `json:"plugPath"`
}

type HttpPluginParam struct {
	ExecResultId int64 `json:"execResultId"`
	PlugId       int64 `json:"plugId"`
}

type HttpPluginResult struct {
	ExecResultId       int64  `json:"execResultId"`
	PlugId             int64  `json:"plugId"`
	PlugResultText     string `json:"plugResult"`
	PlugResultFilePath string `json:"plugResultFilePath"`
}

type HttpPluginStop struct {
	PlugId      int64  `json:"plugId"`
	RecordId    int64  `json:"recordId"`
	ExecCommand string `json:"execCommand"`
}

// Deprecated: Use PluginInfo instead
// type PluginInfo struct {
// 	PlugStateId    string `json:"plugStateId"`
// 	PlugRunCommand string `json:"plugRunCommand"`
// 	PlugRunParams  string `json:"plugRunParams"`
// 	InterfaceAddr  string `json:"interfaceAddr"`
// 	InterfacePath  string `json:"interfacePath"`
// }
