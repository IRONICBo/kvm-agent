package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kvm-agent/internal/log"
	"kvm-agent/internal/runner/models/request"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
)

const (
	PluginStateSuccess = "1"
	PluginStateFail    = "2"
	ScriptPath         = "/root/script"

	PluginResultTypeText = 1
	PluginResultTypeFile = 2
	PluginTypeCommand    = 1
	PluginTypeHTTP       = 2
)

// var pluginCache *expirable.LRU[int64, request.PluginInfo]

var pluginCache = expirable.NewLRU[int64, request.PluginInfo](1024, nil, time.Hour*24)

// For check plugin execId exist
var pluginExecCache = expirable.NewLRU[int64, request.PluginInfo](1024, nil, time.Hour*24)
var runningCache = expirable.NewLRU[int64, int](1024, nil, time.Hour*24)

// var pluginCache = expirable.NewLRU[int64, request.PluginInfo](60, nil, time.Millisecond*10)

// func init() {
// 	pluginCache = expirable.NewLRU[int64, request.PluginInfo](60, nil, time.Millisecond*10)
// }

func GetPluginParam(c *gin.Context) {
	// var httpPluginParam request.HttpPluginParam
	// err := c.ShouldBindJSON(&httpPluginParam)
	// if err != nil {
	// 	fmt.Println("c.ShouldBindJSON error:", err)
	// 	log.Errorf("GetPluginParam", "c.ShouldBindJSON error: %v", err)

	// 	return
	// }

	plugId, err := strconv.ParseInt(c.Query("plugId"), 10, 64)
	if err != nil {
		fmt.Println("Invalid plugId or missing parameter")
		log.Errorf("GetPluginParam", "Invalid plugId or missing parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "Invalid plugId or missing parameter"})
		return
	}

	execResultId, err := strconv.ParseInt(c.Query("execResultId"), 10, 64)
	if err != nil {
		fmt.Println("Invalid execResultId or missing parameter")
		log.Errorf("GetPluginParam", "Invalid execResultId or missing parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "Invalid execResultId or missing parameter"})
		return
	}

	httpPluginParam := request.HttpPluginParam{
		PlugId:       plugId,
		ExecResultId: execResultId,
	}

	fmt.Println(1111)
	for _, k := range pluginCache.Keys() {
		fmt.Println("pluginCache key:", k)
		v, ok := pluginCache.Get(k)
		if ok {
			fmt.Println("pluginCache value:", v)
		}
	}

	fmt.Printf("httpPluginParam: %#v\n", httpPluginParam)
	log.Debugf("GetPluginParam", "httpPluginParam: %+v", httpPluginParam)

	// get plugin info
	pluginInfo, ok := pluginCache.Get(httpPluginParam.PlugId)
	if !ok {
		fmt.Println("pluginInfo not found")
		log.Errorf("GetPluginParam", "pluginInfo not found")

		c.JSON(200, gin.H{
			"code":   -1,
			"msg":    "pluginInfo not found",
			"params": nil,
		})

		return
	}

	// send http result
	c.JSON(200, gin.H{
		"code":   0,
		"msg":    "success",
		"params": pluginInfo.ExecParams,
	})
}

func SendPluginResult(c *gin.Context) {
	var httpPluginResult request.HttpPluginResult
	err := c.ShouldBindJSON(&httpPluginResult)
	if err != nil {
		fmt.Println("c.ShouldBindJSON error:", err)
		log.Errorf("SendPluginResult", "c.ShouldBindJSON error: %v", err)

		return
	}

	fmt.Printf("httpPluginResult: %#v\n", httpPluginResult)
	log.Debugf("SendPluginResult", "httpPluginResult: %+v", httpPluginResult)

	// get plugin info
	pluginInfo, ok := pluginCache.Get(httpPluginResult.PlugId)
	if !ok {
		fmt.Println("pluginInfo not found")
		log.Errorf("SendPluginResult", "pluginInfo not found")

		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "pluginInfo send failed",
		})

		return
	}

	// send result
	switch pluginInfo.PlugResultType {
	case PluginResultTypeText:
		err = SendResult(pluginInfo, httpPluginResult.PlugResultText, 0, "")
		if err != nil {
			log.Errorf("SendPluginResult", "SendResult error: %v", err)

			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "pluginInfo send failed",
			})

			return
		}
	case PluginResultTypeFile:
		// err = SendResult(pluginInfo, "", 0, "/home/asklv/Projects/kvm-agent/internal/runner/handlers/plugin_runner_handler.go")
		err = SendResult(pluginInfo, "", 0, httpPluginResult.PlugResultFilePath)
		if err != nil {
			log.Errorf("SendPluginResult", "SendResult error: %v", err)

			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "pluginInfo send failed",
			})

			return
		}
	}

	// Try to remove plugin info from cache
	// pluginCache.Remove(httpPluginResult.PlugId)
	// Remove current execId
	runningCache.Remove(httpPluginResult.ExecResultId)

	// send http result
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func SendResult(info request.PluginInfo, result string, idx int, path string) error {
	client := resty.New()

	req := client.R().
		// SetHeader("Content-Type", "application/json").
		// SetHeader("Content-Type", "application/json").
		SetMultipartField("stateCode", "", "", strings.NewReader(fmt.Sprintf("%d", 1))).
		SetMultipartField("stateId", "", "", strings.NewReader(fmt.Sprintf("%d\n", info.ExecResultIdList[idx]))).
		SetMultipartField("stateResponse", "", "", strings.NewReader(result)).
		SetMultipartField("otherMessage", "", "", strings.NewReader(""))
		// SetMultipartFormData(map[string]string{
		// 	"stateCode":     1,
		// 	"stateId":       fmt.Sprintf("%d\n", info.ExecResultIdList[idx]),
		// 	"stateResponse": result,
		// 	"otherMessage":  "",
		// }).

	if path != "" {
		// get reader from file
		file, err := os.Open(path)
		if err != nil {
			log.Errorf("SendResult", "os.Open error: %v", err)
		}

		req.SetFileReader("files", file.Name(), file)
	}

	resp, err := req.Post(info.ResponseUrl)
	if err != nil {
		log.Errorf("SendResult", "client.R().Post error: %v", err)
		return err
	}

	log.Debugf("SendResult", "resp: %+v", resp, fmt.Sprintf("%s", info.ResponseUrl))
	fmt.Println("resp:", resp, info.ExecResultIdList[idx], fmt.Sprintf("%s", info.ResponseUrl))

	return nil
}

func StopPlugin(c *gin.Context) {
	var httpPluginStop request.HttpPluginStop
	err := c.ShouldBindJSON(&httpPluginStop)
	if err != nil {
		fmt.Println("c.ShouldBindJSON error:", err)
		log.Errorf("StopPlugin", "c.ShouldBindJSON error: %v", err)

		return
	}

	fmt.Printf("httpPluginStop: %#v\n", httpPluginStop)
	log.Debugf("StopPlugin", "httpPluginStop: %+v", httpPluginStop)

	// get plugin info
	runningPID, ok := runningCache.Get(httpPluginStop.RecordId)
	if !ok {
		fmt.Println("runningPID not found")
		log.Errorf("StopPlugin", "runningPID not found")

		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "runningPID not found",
		})
	}

	// Run kill
	kill := exec.Command("kill", "-9", fmt.Sprintf("%d", runningPID))
	err = kill.Run()
	if err != nil {
		fmt.Println("kill.Run error:", err)
		log.Errorf("StopPlugin", "kill.Run error: %v", err)

		return
	}

	// send http result
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})

	// Try to remove plugin info from cache
	pluginCache.Remove(httpPluginStop.PlugId)
	runningCache.Remove(httpPluginStop.RecordId)
}

func RunPlugin(c *gin.Context) {
	var pluginInfo request.PluginInfo
	err := c.ShouldBindJSON(&pluginInfo)
	if err != nil {
		fmt.Println("c.ShouldBindJSON error:", err)
		log.Errorf("RunPlugin", "c.ShouldBindJSON error: %v", err)

		return
	}

	fmt.Printf("pluginInfo: %#v\n", pluginInfo)
	log.Debugf("RunPlugin", "pluginInfo: %+v", pluginInfo)

	// Set plugin info to cache
	pluginCache.Add(pluginInfo.PlugId, pluginInfo)

	// fmt pluginCache
	for _, k := range pluginCache.Keys() {
		fmt.Println("pluginCache key:", k)
		v, ok := pluginCache.Get(k)
		if ok {
			fmt.Println("pluginCache value:", v)
		}
	}

	// parse json params
	// [{"key":"host","value":"123"}]
	var params []string
	type Param struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	var paramList []Param
	err = json.Unmarshal([]byte(pluginInfo.ExecParams), &paramList)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		log.Errorf("RunPlugin", "json.Unmarshal error: %v", err)

		return
	}

	for _, param := range paramList {
		// params += fmt.Sprintf("%s %s ", param.Key, param.Value)
		params = append(params, fmt.Sprintf("--%s", param.Key), param.Value)
	}

	// run script
	err = os.Chdir(pluginInfo.PlugPath)
	if err != nil {
		fmt.Println("os.Chdir error:", err)
		log.Errorf("RunPlugin", "os.Chdir error: %v", err)

		return
	}
	// set chmod 777
	// TODO: Need root permission
	// err = os.Chmod(pluginInfo.PlugPath, 0777)
	// if err != nil {
	// 	fmt.Println("os.Chmod error:", err)
	// 	log.Errorf("RunPlugin", "os.Chmod error: %v", err)

	// 	return
	// }

	for i := 0; i < pluginInfo.ExecNumber; i++ {
		var onceParams []string
		if pluginInfo.PlugType == PluginTypeCommand {
			onceParams = params
		}

		// Append id to params
		if pluginInfo.PlugType == PluginTypeHTTP {
			onceParams = append(onceParams, fmt.Sprintf("--%s", "pluginId"), fmt.Sprintf("%d", pluginInfo.PlugId))
			onceParams = append(onceParams, fmt.Sprintf("--%s", "execResultId"), fmt.Sprintf("%d", pluginInfo.ExecResultIdList[i]))
		}

		// Store this cache for concurrent exec
		pluginExecCache.Add(pluginInfo.ExecResultIdList[i], pluginInfo)

		fmt.Println("onceParams:", onceParams)
		log.Debugf("RunPlugin", "params: %s", onceParams)

		// set cmd current dir
		// cmd := exec.Command(pluginInfo.ExecCommand, params...)
		cmdParams := strings.Split(pluginInfo.ExecCommand, " ")
		cmdParams = append(cmdParams, onceParams...)
		var cmd *exec.Cmd
		if len(cmdParams) < 2 {
			cmd = exec.Command(cmdParams[0])
		} else {
			cmd = exec.Command(cmdParams[0], cmdParams[1:]...)
		}
		// cmd.Dir = ScriptPath
		cmd.Dir = pluginInfo.PlugPath
		fmt.Println("cmd:", cmd)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("Exec error:", err.Error(), stderr.String())
			log.Error(err.Error(), stderr.String())
		} else {
			fmt.Println("Exec success:", out.String())
			log.Info(out.String())
		}
		// if err != nil {
		// 	fmt.Println("运行命令失败:", err)
		// 	return
		// }

		// Store running pid
		runningCache.Add(pluginInfo.ExecResultIdList[i], cmd.Process.Pid)

		result := out.String()
		fmt.Println("RunPlugin result:", result)
		log.Debugf("RunPlugin", "result: %s", result)

		// Check if execId exist
		pluginInfo, ok := pluginExecCache.Get(pluginInfo.ExecResultIdList[i])
		if !ok {
			fmt.Println("pluginInfo not found or update already")
			log.Errorf("RunPlugin", "pluginInfo not found or update already")

			return
		}

		// send result
		switch pluginInfo.PlugType {
		case PluginTypeCommand:
			err = SendResult(pluginInfo, result, i, "")
			if err != nil {
				log.Errorf("RunPlugin", "SendResult error: %v", err)

				return
			}
		case PluginTypeHTTP:
			err = SendResult(pluginInfo, result, i, "")
			if err != nil {
				log.Errorf("RunPlugin", "SendResult error: %v", err)

				return
			}
		}
	}
}
