package main

import (
	"fmt"
	"math/rand"
)

func main() {

	debuglevel := 1

	if debuglevel > 10 {
		burgers()
	}

	if debuglevel < 10 {
		exampleSort()
	}

}

func exampleSort() {
	slice := generateSlice(20)
	fmt.Print("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Print("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size)
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func burgers() {

	logEntries := []string{`bread`, `bread`, `bread`, `ham`, `cheese`, `ham`, `ham`, `cheese`}

	type myStruct struct {
		counts map[string]int
	}

	var ms myStruct
	ms.counts = make(map[string]int)

	for _, log := range logEntries {
		fmt.Println("log:", log)

		ms.counts[log] = ms.counts[log] + 1

		counter := 0
		for k, v := range ms.counts {
			fmt.Println("k:", k, "v:", v)

			if ms.counts[k] > 0 {
				counter++
			}
		}

		fmt.Println("counter:", counter)
		if counter == 3 {
			fmt.Println("ready to make burger")

			for k, v := range ms.counts {
				fmt.Println("k:", k, "v:", v)

				if ms.counts[k] > 0 {
					ms.counts[k] = ms.counts[k] - 1
				}
			}

			fmt.Println("burger produced")

			for k, v := range ms.counts {
				fmt.Println("count after making burger = k:", k, "v:", v)
			}
		}
	}

	for k, v := range ms.counts {
		fmt.Println("end loop = k:", k, "v:", v)
	}

}

// for i := 0; i < 5; i++ {
// 	log.Println(echo.EchoRepeat("Hello, World!"))
// }

// const (
// 	//debugLevelCst = 11

// 	signalChannelSize = 10

// 	promListenCst           = ":9901"
// 	promPathCst             = "/metrics"
// 	promMaxRequestsInFlight = 10
// 	promEnableOpenMetrics   = true

// 	quantileError    = 0.05
// 	summaryVecMaxAge = 5 * time.Minute
// )

// var (
// 	// Passed by "go build -ldflags" for the show version
// 	commit string
// 	date   string

// 	pC = promauto.NewCounterVec(
// 		prometheus.CounterOpts{
// 			Subsystem: "counters",
// 			Name:      "webffmpeg",
// 			Help:      "webffmpeg counters",
// 		},
// 		[]string{"function", "variable", "type"},
// 	)
// 	pH = promauto.NewSummaryVec(
// 		prometheus.SummaryOpts{
// 			Subsystem: "histrograms",
// 			Name:      "webffmpeg",
// 			Help:      "webffmpeg historgrams",
// 			Objectives: map[float64]float64{
// 				0.1:  quantileError,
// 				0.5:  quantileError,
// 				0.99: quantileError,
// 			},
// 			MaxAge: summaryVecMaxAge,
// 		},
// 		[]string{"function", "variable", "type"},
// 	)
// )

// ctx, cancel := context.WithCancel(context.Background())
// defer cancel()

// version := flag.Bool("version", false, "version")

// promListen := flag.String("promListen", promListenCst, "Prometheus http listening socket")
// promPath := flag.String("promPath", promPathCst, "Prometheus http path. Default = /metrics")
// // curl -s http://[::1]:9111/metrics 2>&1 | grep -v "#"
// // curl -s http://127.0.0.1:9111/metrics 2>&1 | grep -v "#"

// flag.Parse()

// if *version {
// 	fmt.Println("commit:", commit, "\tdate(UTC):", date)
// 	os.Exit(0)
// }

// go initSignalHandler(cancel)

// go initPromHandler(ctx, *promPath, *promListen)

// exampleFunc()

// func exampleFunc() {

// 	startTime := time.Now()
// 	defer func() {
// 		pH.WithLabelValues("exampleFunc", "start", "complete").Observe(time.Since(startTime).Seconds())
// 	}()
// 	pC.WithLabelValues("exampleFunc", "start", "counter").Inc()

// }

// // initSignalHandler sets up signal handling for the process, and
// // will call cancel() when recieved
// func initSignalHandler(cancel context.CancelFunc) {
// 	c := make(chan os.Signal, signalChannelSize)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

// 	<-c
// 	log.Printf("Signal caught, closing application")
// 	cancel()
// 	os.Exit(0)
// }

// // initPromHandler starts the prom handler with error checking
// func initPromHandler(ctx context.Context, promPath string, promListen string) {
// 	// https: //pkg.go.dev/github.com/prometheus/client_golang/prometheus/promhttp?tab=doc#HandlerOpts
// 	http.Handle(promPath, promhttp.HandlerFor(
// 		prometheus.DefaultGatherer,
// 		promhttp.HandlerOpts{
// 			EnableOpenMetrics:   promEnableOpenMetrics,
// 			MaxRequestsInFlight: promMaxRequestsInFlight,
// 		},
// 	))
// 	go func() {
// 		err := http.ListenAndServe(promListen, nil)
// 		if err != nil {
// 			log.Fatal("prometheus error", err)
// 		}
// 	}()
// }
