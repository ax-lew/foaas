package cmd

import (
	"github.com/ax-lew/foaas/clock"
	"net/http"
	"time"

	"github.com/ax-lew/foaas/domain/service"
	"github.com/ax-lew/foaas/foaas"
	"github.com/ax-lew/foaas/logger"
	"github.com/ax-lew/foaas/ratelimiter"
	"github.com/ax-lew/foaas/server"
	"github.com/spf13/cobra"
)

func ServeCmd() *cobra.Command {
	var flags flags
	var cmd = &cobra.Command{
		Use:              "serve",
		Short:            "Run server",
		Long:             `Run server`,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {},
		Run: func(_ *cobra.Command, _ []string) {
			runServe(&flags)
		},
	}

	cmd.Flags().StringVar(&flags.foaasAddress, "foaas-address", defaultFoaasAddress,
		"foaas address")
	cmd.Flags().IntVar(&flags.foaasTimeoutMs, "foaas-timeout-ms", defaultFoaasTimeoutMs,
		"foaas request timeout")
	cmd.Flags().IntVar(&flags.maxRequests, "max-requests", defaultMaxRequests,
		"maximum allowed requests per time interval defined by interval_ms")
	cmd.Flags().IntVar(&flags.intervalMs, "interval-ms", defaultIntervalMs,
		"rate limiter time window")

	return cmd
}

func runServe(flags *flags) {
	logger.Initialize()

	rateLimiterConfig := &ratelimiter.Config{
		MaxRequests: flags.maxRequests,
		Interval:    time.Duration(flags.intervalMs) * time.Millisecond,
	}

	rateLimiter := ratelimiter.NewLocalRateLimiter(rateLimiterConfig, &clock.DefaultClock{})
	httpClient := &http.Client{Timeout: time.Duration(flags.foaasTimeoutMs) * time.Millisecond}
	foaasClient := foaas.NewClient(flags.foaasAddress, httpClient)
	service := service.NewDefaultService(foaasClient)

	server := server.NewServer(rateLimiter, service)
	server.Run()
}
