package builder

import "github.com/evanw/esbuild/pkg/api"

// MapSourceMap maps string to api.SourceMap
func MapSourceMap(sm string) api.SourceMap {
	switch sm {
	case "l":
		return api.SourceMapLinked
	case "in":
		return api.SourceMapInline
	default:
		return api.SourceMapNone
	}
}
