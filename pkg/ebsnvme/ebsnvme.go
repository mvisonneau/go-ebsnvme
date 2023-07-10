package ebsnvme

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

const (
	nvmeAdminIdentify         = 0x06
	nvmeIoctlAdminCmd uintptr = 0xC0484E41
	awsNvmeVolumeID           = 0x1D0F
	awsNvmeEbsMn              = "Amazon Elastic Block Store"
)

// Device represents a block device.
type Device struct {
	VolumeID string
	Name     string
}

type nvmeIdentifyController struct {
	vid       uint16
	ssvid     uint16
	sn        [20]byte
	mn        [40]byte
	fr        [8]byte
	rab       uint8
	ieee      []byte
	mic       uint8
	mdts      uint8
	reserved0 [(256 - 78)]uint8
	oacs      uint16
	acl       uint8
	aerl      uint8
	frmw      uint8
	lpa       uint8
	elpe      uint8
	npss      uint8
	avscc     uint8
	reserved1 [(512 - 265)]uint8
	sqes      uint8
	cqes      uint8
	reserved2 uint16
	nn        uint32
	oncs      uint16
	fuses     uint16
	fna       uint8
	vwc       uint8
	awun      uint16
	awupf     uint16
	nvscc     uint8
	reserved3 [(704 - 531)]uint8
	reserved4 [(2048 - 704)]uint8
	psd       [996]byte
	vs        struct {
		bdev      [32]byte
		reserved0 [(1024 - 32)]byte
	}
}

type nvmeAdminCommand struct {
	opcode    uint8
	flags     uint8
	cid       uint16
	nsid      uint32
	reserved0 uint64
	mptr      uint64
	addr      uintptr
	mlen      uint32
	alen      uint32
	cdw10     uint32
	cdw11     uint32
	cdw12     uint32
	cdw13     uint32
	cdw14     uint32
	cdw15     uint32
	reserved1 uint64
}

// ScanDevice returns a Device object based on its path.
func ScanDevice(device string) (d Device, e error) {
	f, err := open(device)
	if err != nil {
		e = err
		return
	}

	idCtrl := nvmeIdentifyController{}
	adminCmd := nvmeAdminCommand{
		opcode: nvmeAdminIdentify,
		addr:   uintptr(unsafe.Pointer(&idCtrl)),
		alen:   uint32(unsafe.Sizeof(idCtrl)),
		cdw10:  1,
	}

	err = ioctl(f, nvmeIoctlAdminCmd, uintptr(unsafe.Pointer(&adminCmd)))
	if err != nil {
		e = err
		return
	}

	if idCtrl.getVendorID() != awsNvmeVolumeID {
		e = fmt.Errorf("volume ID not matching an AWS EBS one")
		return
	}

	if idCtrl.getModuleNumber() != awsNvmeEbsMn {
		e = fmt.Errorf("module number not matching an AWS EBS one")
		return
	}

	d.VolumeID = idCtrl.getVolumeID()
	d.Name = idCtrl.getDeviceName()

	return
}

func open(device string) (uintptr, error) {
	f, err := syscall.Open(device, syscall.O_RDWR, 0o660)
	if err != nil {
		return 0, err
	}

	return uintptr(f), nil
}

func ioctl(fd, cmd, ptr uintptr) (err error) {
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr); errno != syscall.Errno(0x0) {
		err = errno
	}

	return
}

func (i *nvmeIdentifyController) getVolumeID() string {
	s := strings.TrimSpace(string(i.sn[:]))
	if s[3:4] != "-" {
		return "vol-" + s[3:]
	}

	return s
}

func (i *nvmeIdentifyController) getDeviceName() string {
	s := strings.TrimSpace(string(i.vs.bdev[:]))
	if len(s) < 5 || s[:5] != "/dev/" {
		return "/dev/" + s
	}

	return s
}

func (i *nvmeIdentifyController) getVendorID() uint16 {
	return i.vid
}

func (i *nvmeIdentifyController) getModuleNumber() string {
	return strings.TrimSpace(string(i.mn[:]))
}
