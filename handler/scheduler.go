package handler

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

// enable_scheduler: true
// count down : "@every 25s"
// time on day : "TZ=Asia/Bangkok 59 23 * * *" #time_zone MM:HH
func StartScheduler() {
	_enableScheduler := false
	if _enableScheduler {
		log.Infoln("Start Scheduler")
		startCheckInBroadcastScheduler()
		startAlertScheduler()
	}
}

func startCheckInBroadcastScheduler() {
	_timeScheduler := "TZ=Asia/Bangkok 11 14 * * *" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, broadcastJob)
	c.Start()
	msg := fmt.Sprintf("Schedule Boibot Send CheckIn Broadcast running at: %s", c.Entries()[0].Next.String())
	log.Infoln("StartCheckInBroadcastScheduler : ", msg)
}

func startAlertScheduler() {
	_timeScheduler := "@every 10s" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, alertCheckInJob)
	c.Start()
	msg := fmt.Sprintf("Schedule Alert finish check in step running at: %s", c.Entries()[0].Next.String())
	log.Infoln("StartAlertScheduler : ", msg)
}

func broadcastJob() {
	log.Infoln("===============================")
	log.Infoln("Check In Broadcast Scheduler : ", time.Now())
	log.Infoln("===============================")

}

func alertCheckInJob() {
	log.Infoln("##############################")
	log.Infoln("Alert Check In Scheduler : ", time.Now())
	log.Infoln("##############################")

}
