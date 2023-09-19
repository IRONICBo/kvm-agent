package tasks

import (
	"context"
	"encoding/json"
	"io"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"kvm-agent/internal/monitor/metrics"
	"kvm-agent/internal/service"
	"kvm-agent/internal/utils"
	"strings"
	"time"

	"github.com/hpcloud/tail"
)

const WATCH_FILENAME = "/var/log/messages"

func StartGuestTriggerTask(config config.Agent, gzip bool) {

	tailConfig := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd},
		MustExist: false,
		Poll:      true,
	}

	svc := service.NewTriggerrService(context.Background())

	tails, err := tail.TailFile(WATCH_FILENAME, tailConfig)
	if err != nil {
		log.Errorf("tail %s failed, err:%v\n", WATCH_FILENAME, err)
	}

	for {
		msg, ok := <-tails.Lines // chan tail.Line
		if !ok {
			log.Errorf("[open failed] tail file close reopen, filename:%s\n",
				tails.Filename)
			time.Sleep(time.Second) // wait one second
			continue
		}

		// contains error message
		if !strings.Contains(msg.Text, "error") {
			continue
		}

		trigger := &metrics.TriggerInfo{
			UUID:      config.UUID,
			Timestamp: time.Now().Unix(),

			Key:   "Error",
			Value: msg.Text,
		}
		triggerString, err := json.Marshal(trigger)
		if err != nil {
			log.Errorf("StartGuestTriggerTask", "json.Marshal error: %v", err)
		}

		if gzip {
			triggerString, err = utils.CompressText(string(triggerString))
			if err != nil {
				log.Errorf("StartGuestTriggerTask", "utils.CompressText error: %v", err)
			}
		}

		_ = svc.GuestTriggerPush(config.UUID, string(triggerString))
		log.Infof("msg: %s", msg.Text)
	}
}
