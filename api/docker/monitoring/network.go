package monitoring

import (
	"math"
)

type (
	NetworkInfo struct {
		In struct {
			Value float64
			Unit  string
		}
		Out struct {
			Value float64
			Unit  string
		}
	}
	Rx struct {
		Value float64
		Unit  string
	}
	Tx struct {
		Value float64
		Unit  string
	}
)

func networkStats(stats DockerStats) NetworkInfo {
	rx, rxUnit := convertNetworkStat(float64(stats.Networks.Eth0.RxBytes))
	tx, txUnit := convertNetworkStat(float64(stats.Networks.Eth0.TxBytes))

	return NetworkInfo{Rx{rx, rxUnit}, Tx{tx, txUnit}}
}

func convertNetworkStat(value float64) (float64, string) {
	var converted float64
	var unit string

	if value < math.Pow10(6) {
		converted = value / math.Pow10(3)
		unit = "Ko"
	} else if value < math.Pow10(9) && value > math.Pow10(6) {
		converted = value / (math.Pow10(6))
		unit = "Mo"
	} else if value > math.Pow10(9) {
		converted = value / (math.Pow10(9))
		unit = "Go"
	}
	return converted, unit
}
