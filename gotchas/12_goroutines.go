package main

import (
	"github.com/go-redis/redis"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

/*
	Gotcha #2: Never-Ending Goroutines

	Goroutines can be very useful,
	but they also pose the hazard of functions running in the Background and never terminating, if done stupidly.
*/

/*
	This example is a simple monitoring application
	to check the duration it took to connect to a Redis Instance,
	and execute a "PING" command.
	It then exports this value for scraping by Prometheus

	This example is intentionally bad programming,
	and could be programmed way better and more elegantly
*/
func neverendingGoroutine() {
	printTraceAndGap()

	/*
		This creates an anonymous function which executes checkthe function () and then sleeps for 5 Seconds
	*/
	go func() {
		for {
			check()
			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func check() {
	redisPingDuration.Set(redisSendPing())
}

func redisSendPing() (duratioMs float64) {
	printTraceAndGap()

	var duration time.Duration // we want to know how long the PING command took
	start := time.Now()

	// create a new client connection to the target Redis Server
	client := redis.NewClient(&redis.Options{
		Addr:     "docker-machine:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if checkError(err) {
		log.Panic("Could not connect to redis: ", err)
	} else {
		duration = time.Since(start)
	}

	duratioMs = float64(duration / time.Millisecond) // Prometheus wants a float for gauges, so we have to convert

	/*
		This is where the Gotcha happens.

		Remember that we created a redis.NewClient instance?
		Guess what: it does not end itself when we exit this function! (although the official documentation never explicity states this)
		We have to manually call client.Close() to destroy this instance.

		If the instance is never destroyed, garbage collection cannot collect it and memory cannot be reclaimed.
		If this program is left running for some weeks, one would see hundreds of thousands of active goroutines and gigabytes of memory usage.
	*/

	// client.Close()  // <- this would do the trick

	return
}

/*
	Prometheus metrics for better observability

*/

var (
	// redisPingDuration is a prometheus.Gauge variable which holds the amount of milliseconds it took to execute a PING command on the Redis server
	redisPingDuration = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "redis_ping_duration_milliseconds",
			Help: "The amount of time it took to execute a Redis PING",
		},
	)
)
