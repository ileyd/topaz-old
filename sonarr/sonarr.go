package sonarr

var Client SonarrClient

func InitSonarrClient(address, apiKey string) error {
	Client, err := NewSonarrClient(address, apiKey)
	return err
}

func GetClient() {
	return Client
}
