package output

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mvisonneau/go-ebsnvme/pkg/ebsnvme"
)

func TestParseTypeFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Type
		err      bool
	}{
		{"valid type text", "text", TypeText, false},
		{"valid type json", "json", TypeJSON, false},
		{"invalid type", "invalid", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := ParseTypeFromString(tc.input)
			assert.Equal(t, tc.expected, result)

			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseFieldsFromStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected Fields
		err      bool
	}{
		{"valid fields", []string{"device-name", "device-path"}, Fields{FieldDeviceName, FieldDevicePath}, false},
		{"invalid fields", []string{"invalid-field"}, nil, true},
		{"mixed valid and invalid fields", []string{"device-name", "invalid-field"}, Fields{FieldDeviceName}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := ParseFieldsFromStringSlice(tc.input)
			assert.Equal(t, tc.expected, result)

			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFormatDeviceDetails(t *testing.T) {
	tests := []struct {
		name         string
		device       ebsnvme.Device
		outputType   Type
		outputFields []Field
		expected     string
	}{
		{
			name:         "default fields with text output",
			device:       ebsnvme.Device{Name: "xvdf", VolumeID: "vol-123456789"},
			outputType:   TypeText,
			outputFields: []Field{FieldVolumeID, FieldDeviceName},
			expected:     "vol-123456789\nxvdf",
		},
		{
			name:         "default fields with json output",
			device:       ebsnvme.Device{Name: "xvdf", VolumeID: "vol-123456789"},
			outputType:   TypeJSON,
			outputFields: []Field{FieldDeviceName, FieldVolumeID},
			expected:     "{\"device-name\":\"xvdf\",\"volume-id\":\"vol-123456789\"}",
		},
		{
			name:         "nvme related fields with text output",
			device:       ebsnvme.Device{NVMEPath: "/dev/nvme0n1p2"},
			outputType:   TypeText,
			outputFields: []Field{FieldDeviceNamespace, FieldDevicePartition},
			expected:     "1\n2",
		},
		{
			name:         "nvme related fields with json output",
			device:       ebsnvme.Device{NVMEPath: "/dev/nvme0n1p2"},
			outputType:   TypeJSON,
			outputFields: []Field{FieldDeviceNamespace, FieldDevicePartition},
			expected:     "{\"device-namespace\":\"1\",\"device-partition\":\"2\"}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, FormatDeviceDetails(tc.device, tc.outputType, tc.outputFields))
		})
	}
}
