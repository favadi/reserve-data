package common

import (
	"os"
)

const (
	// modeEnv is the name environment variable that set the running mode of core.
	// See below constants for list of available modes.
	modeEnv = "KYBER_ENV"

	// devMode is default running mode. It is used for development.
	devMode = "dev"
	// productionMode is the running mode in staging environment.
	stagingMode = "staging"
	// productionMode is the running mode in production environment.
	productionMode = "production"
	// mainnetMode is the same as production mode, just for backward compatibility.
	mainnetMode = "mainnet"
	// kovanMode is the running mode for testing kovan network.
	kovanMode = "kovan"
	// kovanMode is the running mode for testing ropsten network.
	ropstenMode = "ropsten"
	// simulationMode is running mode in simulation.
	simulationMode = "simulation"
	// simulationMode is running mode for analytic development.
	analyticDevMode = "analytic_dev"
)

var validModes = map[string]struct{}{
	devMode:         {},
	stagingMode:     {},
	productionMode:  {},
	mainnetMode:     {},
	kovanMode:       {},
	ropstenMode:     {},
	simulationMode:  {},
	analyticDevMode: {},
}

// RunningMode returns the current running mode of application.
func RunningMode() string {
	mode, ok := os.LookupEnv(modeEnv)
	if !ok {
		return devMode
	}
	_, valid := validModes[mode]
	if !valid {
		return devMode
	}
	return mode
}
