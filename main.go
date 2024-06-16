package main

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/huh"
)

var (
	format_time = time.Kitchen
	format_date = time.DateOnly
)

type Input struct {
	format string
	time   string
	date   string
}

func (i Input) ToDiscordFormat() (string, error) {
	v, err := time.Parse(fmt.Sprintf("%v %v", format_date, format_time), fmt.Sprintf("%v %v", i.date, i.time))
	if err != nil {
		return "", err
	}
	unix := v.Unix()
	str := fmt.Sprintf("<t:%v:%v>\n", unix, i.format)
	return str, nil
}

func main() {
	v := new(Input)
	now := time.Now()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("time").
				Placeholder(now.Format(format_time)).
				Value(&v.time).
				// Validating fields is easy. The form will mark erroneous fields
				// and display error messages accordingly.
				Validate(func(str string) error {
					_, err := time.Parse(format_time, str)
					if err != nil {
						return err
					}
					return nil
				}),
			huh.NewInput().
				Title("date").
				Placeholder(now.Format(format_date)).
				Value(&v.date).
				// Validating fields is easy. The form will mark erroneous fields
				// and display error messages accordingly.
				Validate(func(str string) error {
					_, err := time.Parse(format_date, str)
					if err != nil {
						return err
					}
					return nil
				}),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose format").
				Options(
					huh.NewOption("short time", "t"),
					huh.NewOption("long time (with seconds)", "T"),
					huh.NewOption("short date", "d"),
					huh.NewOption("long date", "D"),
					huh.NewOption("date + time", "f"),
					huh.NewOption("day + date + time", "F"),
					huh.NewOption("relative", "R"),
				).
				Value(&v.format),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	str, err := v.ToDiscordFormat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
