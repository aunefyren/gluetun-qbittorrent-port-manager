package modules

type ConfigStruct struct {
	Timezone    string            `json:"timezone"`
	Version     string            `json:"version"`
	Environment string            `json:"environment"`
	LogLevel    string            `json:"log_level"`
	Interval    int               `json:"interval"`
	PortFile    string            `json:"port_file"`
	QBitTorrent QBitTorrentConfig `json:"qbittorrent"`
}

type QBitTorrentConfig struct {
	HTTPS    bool   `json:"https"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type QBitTorrentAppPreferences struct {
	ListenPort int `json:"listen_port"`
}
