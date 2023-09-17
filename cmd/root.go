package cmd

import (
	"os"
	"time"

	commonLogger "github.com/hibare/GoCommon/v2/pkg/logger"
	"github.com/hibare/go-docker-healthcheck/internal/constants"
	"github.com/hibare/go-docker-healthcheck/internal/healthcheck"
	"github.com/hibare/go-docker-healthcheck/internal/version"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	url         string
	statusCodes []int
	debugMode   bool
	timeout     time.Duration
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "healthcheck",
	Short:   "Healthcheck CLI for docker images build from scratch",
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		if debugMode {
			commonLogger.SetLoggingLevel(commonLogger.LogLevelDebug)
			log.Debug().Msg("Debug mode enabled")
		}

		log.Debug().Msgf("URL: %s", url)
		log.Debug().Msgf("Status codes: %v", statusCodes)
		log.Debug().Msgf("Timeout %s", timeout)
		log.Debug().Msg("Running healthcheck")

		if healthcheck.Check(url, statusCodes, timeout) {
			log.Debug().Msg("Healthcheck successful")
		} else {
			log.Debug().Msg("Healthcheck failed")
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(commonLogger.InitLogger)
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "healthcheck URL")
	rootCmd.MarkPersistentFlagRequired("url")
	rootCmd.PersistentFlags().IntSliceVarP(&statusCodes, "status-code", "s", constants.DefaultSuccessStatusCodes, "success status codes")
	rootCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", constants.DefaultTimeout, "timeout")
	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "debug mode")
}
