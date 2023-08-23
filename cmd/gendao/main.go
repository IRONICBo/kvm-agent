package main

import (
	"flag"

	"kvm-agent/cmd/gendao/pkg"
	"kvm-agent/internal/models"
)

func main() {
	savePath := flag.String("path", "../../internal/dal/dao", "save path")
	flag.Parse()

	models := []interface{}{
		models.GuestInfo{},
	}

	for _, model := range models {
		pkg.NewDaoGenerator(model, *savePath).Generate().Flush()
	}
}
