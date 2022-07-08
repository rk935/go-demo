package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func start_cron() {

}

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* * * * * *", func() { fmt.Println("OK") })
	c.Start()
	select {}
}
