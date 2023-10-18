package ebsnvme

import "regexp"

var nvmePathRegexp = regexp.MustCompile(`/dev/nvme(?P<Device>\d+)(n(?P<Namespace>\d+))?(p(?P<Partition>\d+))?`)

// Device represents a block device.
type Device struct {
	NVMEPath string

	VolumeID string
	Name     string
}

func (d *Device) Path() string {
	return "/dev/" + d.Name
}

func (d *Device) Namespace() string {
	return nvmePathRegexp.FindStringSubmatch(d.NVMEPath)[nvmePathRegexp.SubexpIndex("Namespace")]
}

func (d *Device) Partition() string {
	return nvmePathRegexp.FindStringSubmatch(d.NVMEPath)[nvmePathRegexp.SubexpIndex("Partition")]
}
