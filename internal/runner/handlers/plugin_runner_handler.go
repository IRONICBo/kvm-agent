package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kvm-agent/internal/log"
	"kvm-agent/internal/runner/models/request"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	PluginStateSuccess = "1"
	PluginStateFail    = "2"
	ScriptPath         = "/root/script"
)

func SendResult(info request.PluginInfo, result string) error {
	client := resty.New()

	resp, err := client.R().
		// SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"stateCode": 1, "statePlugId": info.PlugStateId, "stateResponse": result}).
		Post(fmt.Sprintf("http://%s%s", info.InterfaceAddr, info.InterfacePath))
	if err != nil {
		log.Errorf("SendResult", "client.R().Post error: %v", err)
		return err
	}

	log.Debugf("SendResult", "resp: %+v", resp, fmt.Sprintf("http://%s%s", info.InterfaceAddr, info.InterfacePath))
	fmt.Println("resp:", resp, info.PlugStateId, fmt.Sprintf("http://%s%s", info.InterfaceAddr, info.InterfacePath))

	return nil
}

func RunPlugin(c *gin.Context) {
	var pluginInfo request.PluginInfo
	err := c.ShouldBindJSON(&pluginInfo)
	if err != nil {
		fmt.Println("c.ShouldBindJSON error:", err)
		log.Errorf("RunPlugin", "c.ShouldBindJSON error: %v", err)

		return
	}

	fmt.Println("pluginInfo:", pluginInfo)
	log.Debugf("RunPlugin", "pluginInfo: %+v", pluginInfo)

	// parse json params
	// [{"key":"host","value":"123"}]
	var params []string
	type Param struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	var paramList []Param
	err = json.Unmarshal([]byte(pluginInfo.PlugRunParams), &paramList)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		log.Errorf("RunPlugin", "json.Unmarshal error: %v", err)

		return
	}

	for _, param := range paramList {
		// params += fmt.Sprintf("%s %s ", param.Key, param.Value)
		params = append(params, fmt.Sprintf("--%s", param.Key), param.Value)
	}
	log.Debugf("RunPlugin", "params: %s", params)

	// run script
	err = os.Chdir(ScriptPath)
	if err != nil {
		log.Errorf("RunPlugin", "os.Chdir error: %v", err)

		return
	}
	// set chmod 777
	err = os.Chmod(ScriptPath, 0777)
	if err != nil {
		log.Errorf("RunPlugin", "os.Chmod error: %v", err)

		return
	}

	cmd := exec.Command(pluginInfo.PlugRunCommand, params...)
	cmd.Dir = ScriptPath
	fmt.Println("cmd:", cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("运行命令失败:", err.Error(), stderr.String())
		log.Error(err.Error(), stderr.String())
	} else {
		fmt.Println("运行命令成功:", out.String())
		log.Info(out.String())
	}
	// if err != nil {
	// 	fmt.Println("运行命令失败:", err)
	// 	return
	// }

	result := out.String()
	log.Debugf("RunPlugin", "result: %s", result)

	// send result
	err = SendResult(pluginInfo, result)
	if err != nil {
		log.Errorf("RunPlugin", "SendResult error: %v", err)

		return
	}
}
