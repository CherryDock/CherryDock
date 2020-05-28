package monitoring

import (
	"math"
)

func networkStats(stats DockerStats) NetworkInfo {
	rx, rxUnit := convertNetworkStat(float64(stats.Networks.Eth0.RxBytes))
	tx, txUnit := convertNetworkStat(float64(stats.Networks.Eth0.TxBytes))

	return NetworkInfo{Rx{rx, rxUnit}, Tx{tx, txUnit}}
}

func convertNetworkStat(value float64) (float64, string) {
	var converted float64
	var unit string

	if value < math.Pow10(3) {
		converted = value
		unit = "b"
	} else if value < math.Pow10(6) {
		converted = value / math.Pow10(3)
		unit = "Kb"
	} else if value < math.Pow10(9) && value > math.Pow10(6) {
		converted = value / (math.Pow10(6))
		unit = "Mb"
	} else if value > math.Pow10(9) {
		converted = value / (math.Pow10(9))
		unit = "Gb"
	}
	return converted, unit
}
