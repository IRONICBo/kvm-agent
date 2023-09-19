package service

import (
	"log"
	"testing"
	"time"

	"github.com/hpcloud/tail"
)

func TestTriggerGetError(t *testing.T) {
	// Use fsnotify, This may have many Open and Close event, so may be need to update to monitor file stream.
	// Use file listening?
	// Use tail golang package
	filename := "./trigger.go"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile(filename, config)
	if err != nil {
		log.Printf("tail %s failed, err:%v\n", filename, err)
	}

	for {
		msg, ok := <-tails.Lines // chan tail.Line
		if !ok {
			log.Printf("[open failed] tail file close reopen, filename:%s\n",
				tails.Filename)
			time.Sleep(time.Second) // 读取出错等一秒
			continue
		}
		log.Println("msg:", msg.Text)
	}
}
