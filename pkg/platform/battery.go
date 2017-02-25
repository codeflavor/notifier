package platform

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codeflavor/notifier/pkg/notice"
)

const (
	// NOTE: this shouldn't contain BAT0, since that could be different on laptops
	// with two batteries, or on systems that detect the batter to be BAT1.
	// Replace with regex.
	batteryInfo      = "/sys/class/power_supply/BAT0/capacity"
	defaultThreshold = 20
	appName          = "battery"
)

type battery struct {
	icon string
}

func (b *battery) Start() error {
	msg, err := checkBattThreshold()
	if err != nil {
		return err
	}

	if msg != "" {
		NotifyUser(msg)
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
		return fmt.Sprintf("Battery status below defined (%d) threshold", defaultThreshold), nil
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
