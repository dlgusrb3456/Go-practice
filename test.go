package main

import "fmt"

type Report interface {
	Report() string
}

type EmailReport struct {
	email string
}

type FaxReport struct {
	fax string
}

func (e *EmailReport) Report() string {
	return e.email
}

func (f *FaxReport) Report() string {
	return f.fax
}

type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct {
}

func (e *EmailSender) Send(r *Report) {
	fmt.Println("Email!")
	// 이메일 전송 구현
}

type FaxSender struct {
}

func (f *FaxSender) Send(r *Report) {
	// 팩스 전송 구현
	fmt.Println("Fax!")
}

type SendType string

func SendReport(r *Report, method SendType, receiver string) {
	switch method {
	case "Email":
		sender := EmailSender{}
		sender.Send(r)
	case "Fax":
		sender := FaxSender{}
		sender.Send(r)
	}
}

func main() {
	t_email := &EmailReport{}
	t_fax := &FaxReport{}
	var report Report
	report = t_email
	SendReport(&report, "Email", "test")
	report = t_fax
	SendReport(&report, "Fax", "test")
}
