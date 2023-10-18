package ebsnvme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceNamespaceAndPartition(t *testing.T) {
	testCases := []struct {
		name              string
		device            Device
		expectedNamespace string
		expectedPartition string
	}{
		{
			name:              "device basic",
			device:            Device{NVMEPath: "/dev/nvme0"},
			expectedNamespace: "",
			expectedPartition: "",
		},
		{
			name:              "device with namespace",
			device:            Device{NVMEPath: "/dev/nvme1n2"},
			expectedNamespace: "0",
			expectedPartition: "",
		},
		{
			name:              "device with namespace and partition",
			device:            Device{NVMEPath: "/dev/nvme3n4p5"},
			expectedNamespace: "4",
			expectedPartition: "5",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.expectedNamespace, tc.device.Namespace())
			assert.Equal(t, tc.expectedPartition, tc.device.Partition())
		})
	}
}
