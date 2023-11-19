package main

import "flag"

func (a *arguments) Read() {
	flag.StringVar(&a.Path, "config", "./config.json", "Path to config file")
	flag.BoolVar(&a.Version, "version", false, "Print version information")
	flag.Parse()
}
