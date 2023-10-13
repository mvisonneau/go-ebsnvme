package output

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mvisonneau/go-ebsnvme/pkg/ebsnvme"
)

func TestParseTypeFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected Type
		err      bool
	}{
		{"text", TypeText, false},
		{"json", TypeJSON, false},
		{"invalid", "", true},
	}

	for _, test := range tests {
		result, err := ParseTypeFromString(test.input)
		assert.Equal(t, test.expected, result)

		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestParseFieldsFromStringSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected Fields
		err      bool
	}{
		{[]string{"device-name", "device-path"}, Fields{FieldDeviceName, FieldDevicePath}, false},
		{[]string{"invalid-field"}, nil, true},
		{[]string{"device-name", "invalid-field"}, Fields{FieldDeviceName}, true},
	}

	for _, test := range tests {
		result, err := ParseFieldsFromStringSlice(test.input)
		assert.Equal(t, test.expected, result)

		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestFormatDeviceDetails(t *testing.T) {
	tests := []struct {
		device       ebsnvme.Device
		outputType   Type
		outputFields []Field
		expected     string
	}{
		{
			device:       ebsnvme.Device{Name: "xvdf", Path: "/dev/xvdf", VolumeID: "vol-123456789"},
			outputType:   TypeText,
			outputFields: []Field{FieldVolumeID, FieldDeviceName, FieldDevicePath},
			expected:     "vol-123456789\nxvdf\n/dev/xvdf",
		},
		{
			device:       ebsnvme.Device{Name: "xvdf", Path: "/dev/xvdf", VolumeID: "vol-123456789"},
			outputType:   TypeJSON,
			outputFields: []Field{FieldDeviceName, FieldDevicePath, FieldVolumeID},
			expected:     "{\"device-name\":\"xvdf\",\"device-path\":\"/dev/xvdf\",\"volume-id\":\"vol-123456789\"}",
		},
	}

	for _, test := range tests {
		result := FormatDeviceDetails(test.device, test.outputType, test.outputFields)
		assert.Equal(t, test.expected, result)
	}
}
