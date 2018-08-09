package main

import (
	"cms-util/monitor"
	"time"
)

func StartMonitor() {
	monitor := monitor.NewMonitor("localhost", "contentapi")
	monitor.AddRouteMonitor("/contentapi/comics/detail", "http://app.kngiths.mi.com/kngiths/contenetapi/comics/detail")

	time.Sleep(time.Minute * 10)
}
