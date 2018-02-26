package conekta

type conektaParams struct {
	APIKey     string
	APIBase    string
	APIVersion string
	Locale     string
	Version    string
}

var defaultParams *conektaParams

// Init return conekta default params
func Init() *conektaParams {
	defaultParams = &conektaParams{
		APIBase:    "https://api.conekta.io",
		APIVersion: "2.0.0",
		Locale:     "es",
		Version:    "0.0.1",
	}
	return defaultParams
}

//GetParams returns defaultParams once this is initializated
func GetParams() *conektaParams {
	return defaultParams
}
