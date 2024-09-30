package types

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDeviceDecode(t *testing.T) {
	testCases := []struct {
		in  interface{}
		out int
	}{
		{
			in:  "",
			out: -1,
		},
		{
			in:  "all",
			out: -1,
		},
		{
			in:  2,
			out: 2,
		},
	}

	for _, tc := range testCases {
		var deviceCount DeviceCount

		err := deviceCount.DecodeMapstructure(tc.in)
		assert.NilError(t, err)

		assert.Equal(t, deviceCount, DeviceCount(tc.out))
	}
}
