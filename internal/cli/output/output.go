package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/mvisonneau/go-ebsnvme/pkg/ebsnvme"
)

type (
	Type  string
	Types []Type

	Field  string
	Fields []Field
)

const (
	// TypeText can be used to output RAW text.
	TypeText Type = "text"
	// TypeJSON can be used to output a JSON formatted string.
	TypeJSON Type = "json"

	// FieldDeviceName can be used to return local block device name.
	FieldDeviceName Field = "device-name"
	// FieldDeviceNamespace can be used to return the nvme namespace number of the device.
	FieldDeviceNamespace Field = "device-namespace"
	// FieldDevicePartition can be used to return the nvme partition number of the device.
	FieldDevicePartition Field = "device-partition"
	// FieldDevicePath can be used to return local block device path.
	FieldDevicePath Field = "device-path"
	// FieldVolumeID can be used to return EBS volume-id.
	FieldVolumeID Field = "volume-id"
)

func GetValidTypes() Types {
	return Types{
		TypeText,
		TypeJSON,
	}
}

func GetValidFields() Fields {
	return Fields{
		FieldDeviceName,
		FieldDeviceNamespace,
		FieldDevicePartition,
		FieldDevicePath,
		FieldVolumeID,
	}
}

func (t Type) String() string {
	return string(t)
}

func (f Field) String() string {
	return string(f)
}

func (fields Fields) StringSlice() (slice []string) {
	for _, f := range fields {
		slice = append(slice, string(f))
	}
	return
}

func ParseTypeFromString(t string) (parsedType Type, err error) {
	if !slices.Contains(GetValidTypes(), Type(t)) {
		err = fmt.Errorf("invalid type '%s'", t)
		return
	}

	parsedType = Type(t)
	return
}

func ParseFieldsFromStringSlice(fields []string) (parsedFields Fields, err error) {
	for _, field := range fields {
		parsedField := Field(field)
		if !slices.Contains(GetValidFields(), parsedField) {
			err = errors.Join(err, fmt.Errorf("invalid field '%s'", field))
			continue
		}
		parsedFields = append(parsedFields, Field(field))
	}

	return
}

func FormatDeviceDetails(d ebsnvme.Device, outputType Type, outputFields []Field) (s string) {
	output := map[string]string{}

	for _, field := range outputFields {
		var value string

		switch field {
		case FieldDeviceName:
			value = d.Name
		case FieldDeviceNamespace:
			value = d.Namespace()
		case FieldDevicePartition:
			value = d.Partition()
		case FieldDevicePath:
			value = d.Path()
		case FieldVolumeID:
			value = d.VolumeID
		default:
			continue
		}

		switch outputType {
		case TypeJSON:
			output[field.String()] = value
		case TypeText:
			s += value + "\n"
		}
	}

	switch outputType {
	case TypeJSON:
		jsonData, _ := json.Marshal(output)
		s = string(jsonData)
	case TypeText:
		s = strings.Trim(s, "\n")
	}

	return
}
