package cmd

const (
	defaultFoaasAddress   = "https://foaas.com"
	defaultFoaasTimeoutMs = 1000
	defaultMaxRequests    = 5
	defaultIntervalMs     = 10000
)

type flags struct {
	foaasAddress   string
	foaasTimeoutMs int
	maxRequests    int
	intervalMs     int
}
