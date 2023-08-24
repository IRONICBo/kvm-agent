package utils

import (
	"fmt"
	"kvm-agent/internal/config"

	"github.com/Delta456/box-cli-maker/v2"
)

func KVMAgentBanner() {
	// logo
	logo := `
    __ ___    ____  ___      ___                    __ 
   / //_/ |  / /  |/  /     /   | ____ ____  ____  / /_
  / ,<  | | / / /|_/ /_____/ /| |/ __  / _ \/ __ \/ __/
 / /| | | |/ / /  / /_____/ ___ / /_/ /  __/ / / / /_  
/_/ |_| |___/_/  /_/     /_/  |_\__, /\___/_/ /_/\__/  
                               /____/                  

APP Mode:
- Version: %s
- Debug: %v
- Log file: %s

Agent Mode:
- UUID: %s
- Period: %d
- GZip: %v
`
	content := fmt.Sprintf(logo, config.Config.App.Version, config.Config.App.Debug, config.Config.App.LogFile, config.Config.Agent.UUID, config.Config.Agent.Period, config.Config.Agent.GZip)

	Box := box.New(box.Config{
		Px:       5,
		Py:       0,
		Type:     "Round",
		Color:    "Cyan",
		TitlePos: "Top",
	})
	Box.Print("KVM-Agent", content)
}
