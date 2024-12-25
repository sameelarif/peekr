package models

type BotMitigationService int

const (
	Akamai BotMitigationService = iota
	AkamaiPixel
	Cloudflare
	Imperva
	PerimeterX
	Reblaze
	Radware
	DataDome
	Kasada
	Recaptcha
)

func (b BotMitigationService) String() string {
	return [...]string{"Akamai", "AkamaiPixel", "Cloudflare", "Imperva", "PerimeterX", "Reblaze", "Radware", "DataDome", "Kasada", "Recaptcha"}[b]
}
