package main

type arguments struct {
	Path    string // Path to config file
	Version bool   // Show version information
}

type config struct {
	Timeout   int    `json:"timeout"`
	Interval  int    `json:"interval"`
	IPPrefix  string `json:"ip_prefix"`
	ReportURI string `json:"report_uri"`
	DeviceTag string `json:"device_tag"`
}
