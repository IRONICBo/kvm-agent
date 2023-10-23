package request

type PluginInfo struct {
	PlugStateId    string `json:"plugStateId"`
	PlugRunCommand string `json:"plugRunCommand"`
	PlugRunParams  string `json:"plugRunParams"`
	InterfaceAddr  string `json:"interfaceAddr"`
	InterfacePath  string `json:"interfacePath"`
}
