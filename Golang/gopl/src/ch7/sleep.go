package main

import (
	"flag"
	"fmt"
	"time"
)

type celsius float64
type fahrenheit float64

func (c celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func fToC(f fahrenheit) celsius {
	return celsius((f - 32) * 5 / 9)
}

type celsiusFlag struct {
	celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.celsius = celsius(value)
		return nil
	case "F", "°F":
		f.celsius = fToC(fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func newCelsiusFlag(name string, value celsius, usage string) *celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.celsius
}

var period = flag.Duration("period", 1*time.Second, "sleep period")
var temp = newCelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)
	fmt.Println()
}
