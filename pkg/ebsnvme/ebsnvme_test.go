package ebsnvme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVolumeID(t *testing.T) {
	i := &nvmeIdentifyController{}
	copy(i.sn[:], "vol12345            ")
	assert.Equal(t, "vol-12345", i.getVolumeID())

	copy(i.sn[:], "vol-67890            ")
	assert.Equal(t, "vol-67890", i.getVolumeID())
}

func TestGetDevicePath(t *testing.T) {
	i := &nvmeIdentifyController{}
	copy(i.vs.bdev[:], "foo                             ")
	assert.Equal(t, "/dev/foo", i.getDevicePath())

	copy(i.vs.bdev[:], "/dev/foo                        ")
	assert.Equal(t, "/dev/foo", i.getDevicePath())

	copy(i.vs.bdev[:], "foobar                          ")
	assert.Equal(t, "/dev/foobar", i.getDevicePath())
}

func TestGetDeviceName(t *testing.T) {
	i := &nvmeIdentifyController{}
	copy(i.vs.bdev[:], "foo                             ")
	assert.Equal(t, "foo", i.getDeviceName())

	copy(i.vs.bdev[:], "/dev/foo                        ")
	assert.Equal(t, "foo", i.getDeviceName())

	copy(i.vs.bdev[:], "/dev/foobar")
	assert.Equal(t, "foobar", i.getDeviceName())
}
