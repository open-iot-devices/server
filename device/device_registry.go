package device

import (
	"fmt"
	"sync"
)

var devicesByID map[uint64]*Device
var deviceLock sync.RWMutex

// FindDeviceByID looks up device by id.
// Returns Device or nil if not found
func FindDeviceByID(id uint64) *Device {
	deviceLock.RLock()
	defer deviceLock.RUnlock()

	return devicesByID[id]
}

// AddDevice adds new device into registry
func AddDevice(device *Device) error {
	deviceLock.Lock()
	defer deviceLock.Unlock()

	if _, ok := devicesByID[device.ID]; ok {
		return fmt.Errorf("Device with ID %x already exists in registry", device.ID)
	}

	devicesByID[device.ID] = device

	return nil
}

// ReplaceAllDevicesWith deletes all existing devices and
// adds new devices provided array
func ReplaceAllDevicesWith(devices []*Device) {
	deviceLock.Lock()
	defer deviceLock.Unlock()

	devicesByID = make(map[uint64]*Device)
	for _, dev := range devices {
		devicesByID[dev.ID] = dev
	}
}

// DeleteAllDevices deletes all registered devices
func DeleteAllDevices() {
	ReplaceAllDevicesWith(nil)
}

// DeleteDeviceByID deletes device from registry
func DeleteDeviceByID(id uint64) error {
	deviceLock.Lock()
	defer deviceLock.Unlock()

	if _, ok := devicesByID[id]; !ok {
		return fmt.Errorf("Device with ID %x not found", id)
	}

	delete(devicesByID, id)

	return nil
}

// GetAllDevices returns all registered devices in array
func GetAllDevices() []*Device {
	var index int
	res := make([]*Device, len(devicesByID))

	deviceLock.RLock()
	defer deviceLock.RUnlock()

	for _, dev := range devicesByID {
		res[index] = dev
		index++
	}

	return res
}

func init() {
	devicesByID = make(map[uint64]*Device)
}