package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type BirthDate time.Time

func Birthday(year,month,day int) BirthDate {
	return BirthDate(time.Date(year, time.Month(month), day,0,0,0,0,time.UTC))
}

func (d *BirthDate) Set(val string) error {
	t, err := time.Parse("02/01/2006", val)
	if err != nil {
		return err
	}
	*d = BirthDate(t)
	return nil
}

func (d *BirthDate) String() string {
	if d == nil {
		return ""
	}
	return (*time.Time)(d).Format("02/01/2006")
}

func (d BirthDate) day() int {
	return (time.Time)(d).Day()
}

func (d BirthDate) month() int {
	return int((time.Time)(d).Month())
}

func (d BirthDate) year() int {
	return (time.Time)(d).Year()
}

type options struct {
	cardID   string
	month    int
	year     int
	email    string
	birthday BirthDate
}


func main() {
	var opt options
	flag.StringVar(&opt.cardID, "id", "", "your TCL card number")
	flag.IntVar(&opt.month, "m", int(time.Now().Month()), "certificate for month")
	flag.IntVar(&opt.year, "y", time.Now().Year(), "certificate for year")
	flag.StringVar(&opt.email, "email", "", "your email address")
	flag.Var(&opt.birthday, "birthday", "your birthday (dd/MM/yyyy)")

	flag.Parse()

	var val validator
	val.requireString(opt.cardID, "id")
	val.requireString(opt.email, "email")
	val.requireDate(opt.birthday, "birthday")
	if len(val.errors) > 0 {
		for _, err := range val.errors {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
		}
		os.Exit(2)
	}

	fmt.Printf("retrieving certificate for card %s\n", opt.cardID)
	fmt.Printf("\tfor month %02d-%d\n", opt.month, opt.year)
	fmt.Printf("\tfor email %s\n", opt.email)

	api := NewApi(TCLBaseURL)

	err := api.RequestSubscriptionCertificate(opt.cardID, opt.email, opt.birthday, opt.month, opt.year)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		fmt.Printf("KO!\n")
	} else {
		fmt.Printf("OK!\n")
	}
}

type validator struct {
	errors []error
}

func (v *validator) add(err error) {
	v.errors = append(v.errors, err)
}

func (v *validator) requireString(val string, name string) {
	if val == "" {
		v.add(fmt.Errorf("%s is required", name))
	}
}

func (v *validator) requireDate(val BirthDate, name string) {
	var empty BirthDate
	if val == empty {
		v.add(fmt.Errorf("%s is required", name))
	}
}
