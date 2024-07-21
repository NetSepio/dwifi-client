package wifi

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type DeviceInfo struct {
	MACAddress         string        `json:"macAddress"`
	IPAddress          string        `json:"ipAddress"`
	ConnectedAt        time.Time     `json:"connectedAt"`
	TotalConnectedTime time.Duration `json:"totalConnectedTime"`
	Connected          bool          `json:"connected"`
	LastChecked        time.Time     `json:"lastChecked"`
	DefaultGateway     string        `json:"defaultGateway"`
	Manufacturer       string        `json:"manufacturer"`
	InterfaceName      string        `json:"interfaceName"`
	HostSSID           string        `json:"hostSSID"`
}

type WiFiData struct {
	ID            uint         `json:"id"`
	Gateway       string       `json:"gateway"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	Status        []DeviceInfo `json:"status"`
	Password      string       `json:"password"`
	Location      string       `json:"location"`
	PricePerMin   string       `json:"price_per_min"`
	WalletAddress string       `json:"wallet_address"`
	ChainName     string       `json:"chain_name"`
}

func ConnectToWiFi(ssid, password string) error {
	profileName := "ErebrusDwifi"
	xmlProfile := fmt.Sprintf(`<?xml version="1.0"?>
		<WLANProfile xmlns="http://www.microsoft.com/networking/WLAN/profile/v1">
			<name>%s</name>
			<SSIDConfig>
				<SSID>
					<name>%s</name>
				</SSID>
			</SSIDConfig>
			<connectionType>ESS</connectionType>
			<connectionMode>auto</connectionMode>
			<MSM>
				<security>
					<authEncryption>
						<authentication>WPA2PSK</authentication>
						<encryption>AES</encryption>
						<useOneX>false</useOneX>
					</authEncryption>
					<sharedKey>
						<keyType>passPhrase</keyType>
						<protected>false</protected>
						<keyMaterial>%s</keyMaterial>
					</sharedKey>
				</security>
			</MSM>
		</WLANProfile>`, profileName, ssid, password)

	profilePath := "wifi_profile.xml"
	if err := os.WriteFile(profilePath, []byte(xmlProfile), 0644); err != nil {
		return fmt.Errorf("failed to create profile XML file: %w", err)
	}
	defer os.Remove(profilePath)

	cmd := exec.Command("netsh", "wlan", "add", "profile", fmt.Sprintf("filename=%s", profilePath))
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to add WiFi profile: %v, output: %s", err, output)
	}

	cmd = exec.Command("netsh", "wlan", "connect", fmt.Sprintf("name=%s", profileName), fmt.Sprintf("ssid=%s", ssid))
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to connect to WiFi: %v, output: %s", err, output)
	}

	return nil
}

func IsConnectedToWiFi(ssid string) (bool, error) {
	cmd := exec.Command("netsh", "wlan", "show", "interfaces")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("failed to check WiFi connection: %w", err)
	}

	return strings.Contains(strings.ToLower(string(output)), strings.ToLower(ssid)), nil
}

func ScanNearbyWiFi() ([]string, error) {
	cmd := exec.Command("netsh", "wlan", "show", "network", "mode=bssid")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to scan WiFi networks: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	var ssids []string
	for _, line := range lines {
		if strings.Contains(line, "SSID") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				ssid := strings.TrimSpace(parts[1])
				if ssid != "" {
					ssids = append(ssids, ssid)
				}
			}
		}
	}
	return ssids, nil
}

func DisconnectFromWiFi() error {
	cmd := exec.Command("netsh", "wlan", "disconnect")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to disconnect from WiFi: %v, output: %s", err, output)
	}
	return nil
}
