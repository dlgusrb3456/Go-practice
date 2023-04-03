// package main

// import (
// 	"fmt"
// 	"log"
// )

// type Report interface {
// 	Report() string
// }

// type EmailReport struct {
// 	email string
// }

// type FaxReport struct {
// 	fax string
// }

// func (e *EmailReport) Report() string {
// 	return e.email
// }

// func (f *FaxReport) Report() string {
// 	return f.fax
// }

// type ReportSender interface {
// 	Send(r *Report)
// }

// type EmailSender struct {
// }

// func (e *EmailSender) Send(r *Report) {
// 	fmt.Println("Email!")
// 	// 이메일 전송 구현
// }

// type FaxSender struct {
// }

// func (f *FaxSender) Send(r *Report) {
// 	// 팩스 전송 구현
// 	fmt.Println("Fax!")
// }

// type SendType string

// func SendReport(r *Report, method SendType, receiver string) {
// 	switch method {
// 	case "Email":
// 		sender := EmailSender{}
// 		sender.Send(r)
// 	case "Fax":
// 		sender := FaxSender{}
// 		sender.Send(r)
// 	}
// }
// func generateMessage(message string) {
// 	// Buffered Channel of type Boolean
// 	done := make(chan bool, 1)
// 	go printMessage(done, message)

// 	// Waiting to receive value from channel
// 	//<-done
// }

// func printMessage(done chan bool, message string) {
// 	defer func() {
// 		// Sending value to channel
// 		done <- true
// 	}()
// 	IsTimeEnabled := false
// 	// Inside Logic (Dont frighten xD)
// 	if IsTimeEnabled {
// 		log.Println(message)
// 		return
// 	}
// 	fmt.Println(message)
// }

// func main() {
// 	// t_email := &EmailReport{}
// 	// t_fax := &FaxReport{}
// 	// var report Report
// 	// report = t_email
// 	// SendReport(&report, "Email", "test")
// 	// report = t_fax
// 	// SendReport(&report, "Fax", "test")
// 	generateMessage("i have a boy")
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// // Handler for /example request
// func example(w http.ResponseWriter, req *http.Request) {

// 	fmt.Println("example handler started")

// 	// Accessing the context of the request
// 	context := req.Context()

// 	select {

// 	// Simulating some work by the server
// 	// Waits 10 seconds and then responds with "example\n"
// 	case <-time.After(10 * time.Second):
// 		fmt.Fprintf(w, "example\n")

// 	// Handling request cancellation
// 	case <-context.Done():
// 		err := context.Err()
// 		fmt.Println("server:", err)
// 	}

// 	fmt.Println("example handler ended")
// }

// func main() {

// 	http.HandleFunc("/example", example)
// 	http.ListenAndServe(":5000", nil)
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func longRunningTask(ctx context.Context, timeToRun time.Duration) {
// 	fmt.Println("start time: ", time.Now())
// 	select {
// 	case <-time.After(timeToRun):
// 		fmt.Println("completed before context deadline passed")
// 	case <-ctx.Done():
// 		fmt.Println("bailed because context deadline passed")
// 	}
// 	fmt.Println("end time: ", time.Now())
// }

// const duration = 5 * time.Second

// func main() {
// 	ctx := context.Background()

// 	// this will bail because the function runs longer than the context's deadline allows
// 	ctx1, _ := context.WithDeadline(ctx, time.Now().Add(duration))
// 	longRunningTask(ctx1, 10*time.Second)

// 	// this will complete because the function completes before the context's deadline arrives
// 	ctx2, _ := context.WithDeadline(ctx, time.Now().Add(duration))
// 	longRunningTask(ctx2, 1*time.Second)
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func longRunningTask(ctx context.Context, timeToRun time.Duration) {
// 	select {
// 	case <-time.After(timeToRun):
// 		fmt.Println("completed before context timed out")
// 	case <-ctx.Done():
// 		fmt.Println("bailed because context timed out")
// 	}
// }

// const timeout = 5 * time.Second

// func main() {
// 	ctx := context.Background()

// 	// this will bail because the function takes longer than the context allows
// 	// ctx1, _ := context.WithTimeout(ctx, timeout)
// 	// longRunningTask(ctx1, 10*time.Second)

// 	// this will complete because the function completes before the context times out
// 	ctx2, _ := context.WithTimeout(ctx, timeout)
// 	longRunningTask(ctx2, 1*time.Second)
// }

package main

import (
	"context"
	"fmt"
)

func tryAnotherKeyType(ctx context.Context, keyToConvert string) {
	type keyType2 string

	k := keyType2(keyToConvert)
	if v := ctx.Value(k); v != nil {
		fmt.Println("found a value for key type 2:", v)
	} else {
		fmt.Println("no value for key type 2")
	}
}

func main() {
	keyString := "foo"

	type keyType1 string

	k := keyType1(keyString)
	ctx := context.WithValue(context.Background(), k, "bar")

	if v := ctx.Value(k); v != nil {
		fmt.Println("found a value for key type 1:", v)
	} else {
		fmt.Println("no value for key type 1")
	}

	if v := ctx.Value(k); v != nil {
		fmt.Println("found a value for key type 1:", v)
	} else {
		fmt.Println("no value for key type 1")
	}

	tryAnotherKeyType(ctx, keyString)
}
