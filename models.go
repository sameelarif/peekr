package main

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
)

func (b BotMitigationService) String() string {
	return [...]string{"Akamai", "AkamaiPixel", "Cloudflare", "Imperva", "PerimeterX", "Reblaze", "Radware", "DataDome", "Kasada"}[b]
}