package main

const (
	WEB_PREFIX = "/"
	GZIP_LEVEL = 9
)

type config struct {
	CORS   bool   `json:"cors"`
	Life   int    `json:"life"`
	Listen string `json:"listen"`
}

type arguments struct {
	Path    string // Path to config file
	Version bool   // Show version information
}

type serverOptions struct {
	Gzip       int
	CORS       bool
	WebPrefix  string
	APIPrefix  string
	DeviceLife int
	DeviceList []map[string]any
}

type httpHeader struct {
	Header string
	Value  string
}

type httpResponse struct {
	Time    string `json:"time"`
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
	Path    string `json:"path"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type clientRequest struct {
	IPAddress string `json:"ip_address" binding:"required"`
	DeviceTag string `json:"device_tag" binding:"required"`
}
