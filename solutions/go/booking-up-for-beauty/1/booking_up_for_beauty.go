package booking

import (
	"fmt"
	"log"
	"time"
)

// REF: Mon Jan 2 15:04:05 -0700 MST 2006

func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func HasPassed(date string) bool {

	t, err := time.Parse("January 2, 2006 15:04:05", date)

	if err != nil {
		log.Fatal(err)
	}

	return t.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {

	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)

	if err != nil {
		log.Fatal(err)
	}

	return t.Hour() >= 12 && t.Hour() <= 18

}

func Description(date string) string {

	t, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		log.Fatal(err)
	}

	t_f := t.Format("Monday, January 2, 2006")

	return fmt.Sprintf("You have an appointment on %s, at %d:%d.", t_f, t.Hour(), t.Minute())

}

func AnniversaryDate() time.Time {

	date_string := fmt.Sprintf("September 15 %d", time.Now().Year())

	t, err := time.Parse("January 2 2006", date_string)

	if err != nil {
		log.Fatal(err)
	}

	return t

}
