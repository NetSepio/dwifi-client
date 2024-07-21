package utils

import (
	"fmt"
	"strings"
	"time"

	"chicken/wifi"

	"github.com/fatih/color"
)

func ErebrusBanner() {
	banner := `
	 ______          _                      _____               _____
	|  ____|        | |                    |  __ \          (_)|  __)(_)
	| |__   _ __ ___| |__  _ __ _   _ ___  | |  | |_      _  _ | |_   _
	|  __| | '__/ _ \ '_ \| '__| | | / __| | |  | \ \ /\ / /| ||  _| | |
	| |____| | |  __/ |_) | |  | |_| \__ \ | |__| |\ V  V / | || |   | |
	|______|_|  \___|_.__/|_|   \__,_|___/ |_____/  \_/\_/  |_||_|   |_|
																  
    `
	lines := strings.Split(banner, "\n")
	colors := []*color.Color{
		color.New(color.FgRed).Add(color.Bold),
		color.New(color.FgYellow).Add(color.Bold),
		color.New(color.FgGreen).Add(color.Bold),
		color.New(color.FgCyan).Add(color.Bold),
		color.New(color.FgBlue).Add(color.Bold),
		color.New(color.FgMagenta).Add(color.Bold),
	}

	for i, line := range lines {
		colors[i%len(colors)].Println(line)
	}

	fmt.Println()
	color.New(color.FgHiWhite).Add(color.Bold, color.Underline).Println("Welcome to Erebrus DWiFi!")
	fmt.Println()
}

func MonitorConnection(ssid string) {
	startTime := time.Now()
	disconnected := false

	for {
		connected, err := wifi.IsConnectedToWiFi(ssid)
		if err != nil {
			color.Red("Error checking WiFi connection: %v", err)
			return
		}

		if connected {
			if disconnected {
				color.Green("Reconnected to %s", ssid)
				disconnected = false
			}
			connectedTime := time.Since(startTime)
			color.Cyan("Connected to %s for %s", ssid, connectedTime.Round(time.Second))
		} else {
			if !disconnected {
				color.Yellow("Disconnected from %s", ssid)
				disconnected = true
			}
			color.Red("Not connected to %s. Waiting for reconnection...", ssid)
		}

		time.Sleep(5 * time.Second)
	}
}
