package config

type Config struct {
	ServerAddress string `json:"server_address,omitempty"`
	BaseURL       string `json:"base_url,omitempty"`
	UrlsPath      string `json:"file_path,omitempty"`
}
