package util

var gbMultiple int64 = 1024 * 1024 * 1024

func BytesFromGb(sizeGb int) int64 {
	return int64(sizeGb) * gbMultiple
}

func GbFromBytes(sizeBytes int64) float64 {
	return float64(sizeBytes) / float64(gbMultiple)
}
