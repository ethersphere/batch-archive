package batcharchive

import (
	_ "embed"
)

//go:embed archive/export.ndjson
var payloadJSON []byte

//go:embed archive/export.ndjson.gzip
var payloadGZIP []byte

func GetBatchSnapshot(compress bool) []byte {
	if compress {
		return payloadGZIP
	}
	return payloadJSON
}
