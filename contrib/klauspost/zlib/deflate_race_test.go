//go:build race
// +build race

package zlib_test

import (
	"testing"

	"github.com/CAFxX/httpcompression/contrib/internal"
	"github.com/CAFxX/httpcompression/contrib/klauspost/zlib"
)

func TestDeflateRace(t *testing.T) {
	t.Parallel()
	c, _ := zlib.New(zlib.Options{})
	internal.RaceTestCompressionProvider(c, 100)
}
