// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

import (
	"kvm-agent/internal/config"
	"kvm-agent/internal/runner/handlers"

	"github.com/gin-gonic/gin"
)

// InitRouter init router.
func InitRouter() *gin.Engine {
	if config.GetString("app.debug") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.POST("/api/plug/run", handlers.RunPlugin)
	r.GET("/api/plug/param", handlers.GetPluginParam)
	r.POST("/api/plug/result", handlers.SendPluginResult)

	// Default task
	r.POST("/api/plug/ping", handlers.RunPingPlug)
	r.POST("/api/plug/fio", handlers.RunFIOPlug)
	r.POST("/api/plug/dbtest", handlers.RunDBTestPlug)
	r.POST("/api/plug/ptp4l", handlers.RunPTP4LPlug)

	return r
}
