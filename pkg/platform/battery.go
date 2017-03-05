package platform

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/codeflavor/notifier/pkg/service"
)

const (
	// NOTE: this shouldn't contain BAT0, since that could be different on laptops
	// with two batteries, or on systems that detect the batter to be BAT1.
	// Replace with regex.
	batteryInfo      = "/sys/class/power_supply/BAT0/capacity"
	defaultThreshold = 20
	appName          = "battery"
)

// Battery holds information about the system battery if any.
type Battery struct {
	Icon     string
	Name     string
	Summary  string
	pollTime time.Duration
}

// Start enables battery check based on a defined duration.
func (b *Battery) Start() (string, error) {
	// adding this here for posterity
	b.pollTime = 6 * time.Second
	for {
		msg, err := checkBattThreshold()
		if err != nil {
			return "", err
		}
		if msg != "" {
			return msg, nil
		}
		time.Sleep(b.pollTime)
		continue
	}
}

// Stop stops the process permanently.
func (b *Battery) Stop() error {
	return nil
}

// Reload stops the process and reloads it.
func (b *Battery) Reload() error {
	return nil
}

// Load loads all the necessary properties of the process.
func (b *Battery) Load() error {
	return nil
}

// Info returns a list of messages about the process.
func (b *Battery) Info() (*service.Info, error) {
	return nil, nil
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
