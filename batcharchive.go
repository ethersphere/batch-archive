package batcharchive

import (
	_ "embed"
)

//go:embed archive/export.ndjson.gzip
var payloadGZIP []byte

func GetBatchSnapshot() []byte {
	return payloadGZIP
}
