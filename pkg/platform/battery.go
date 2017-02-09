package platform

import (
	"fmt"
	"os"
	"strconv"
)

const (
	batteryInfo      = "/sys/class/power_supply/BAT0/capacity"
	defaultThreshold = 20
)

type battery struct {
}

func (b *battery) Start() error {
	err := checkBattThreshold()
	if err != nil {
		return err
	}
	return nil
}

func (b *battery) Stop() error {
	return nil
}

func (b *battery) Reload() error {
	return nil
}

func (b *battery) Load() error {
	return nil
}

func checkBattThreshold() (string, error) {
	threshold, err := getBatteryStatus()
	if err != nil {
		return "", err
	}
	t, err := strconv.ParseInt(threshold, 10, 32)
	if err != nil {
		return "", err
	}
	if t < defaultThreshold {
		return fmt.Sprintf("Battery status below defined default (%d) threshold", defaultThreshold), nil
	}
	return "", nil
}

func getBatteryStatus() (string, error) {
	fileHandler, err := os.Open(batteryInfo)
	if err != nil {
		return "", err
	}
	var readBytes []byte
	_, err = fileHandler.Read(readBytes)
	if err != nil {
		return "", err
	}
	return string(readBytes), nil
}
