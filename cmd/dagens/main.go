package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/antonlindstrom/dagens"
)

func main() {
	var namesOnly = flag.Bool("names", false, "Only print todays names")
	flag.Parse()

	resp, err := dagens.Date(time.Now())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Kunde inte hämta datum: %s\n", err)
		return
	}

	for _, date := range resp.Days {
		names := strings.Join(date.Names, ", ")
		if *namesOnly {
			if len(date.Names) > 0 {
				fmt.Printf("%s\n", names)
			} else {
				fmt.Println("Ingen namnsdag idag.")
			}

			continue
		}

		fmt.Printf("Idag är det %s (%s).\n", strings.ToLower(date.Weekday), date.Date)
		if len(date.Names) > 0 {
			fmt.Printf("- %s.\n", names)
		}

		if date.IsDayOff.Bool() {
			fmt.Println("- Arbetsfri dag.")
		} else {
			fmt.Println("- Arbetsdag.")
		}

		if date.IsRedDay.Bool() {
			fmt.Printf("- Röd dag (%s).\n", date.Holiday)
		}
	}
}
