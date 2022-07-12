package models

type YuGiOhAPIResult struct {
	Cards []YuGiOhAPICard `json:"data"`
}

//Yu Gi Oh Card data
type YuGiOhAPICard struct {
	ID            int64                 `json:"id"`
	Name          string                `json:"name"`
	Type          string                `json:"type"`
	Description   string                `json:"desc"`
	Attack        int                   `json:"atk"`
	Defense       int                   `json:"def"`
	Level         int                   `json:"level"`
	Race          string                `json:"race"`
	Attribute     string                `json:"attribute"`
	Archetype     string                `json:"archetype"`
	LinkValue     int                   `json:"linkval"`
	LinkMarkers   []string              `json:"linkmarkers"`
	PendulumScale int                   `json:"scale"`
	CardSets      []YuGiOhAPICardSet    `json:"card_sets"`
	CardImages    []YuGiOhAPICardImage  `json:"card_images"`
	CardPrices    []YuGiOhAPICardPrices `json:"card_prices"`
}

type YuGiOhAPICardSet struct {
	Name       string  `json:"set_name"`
	Code       string  `json:"set_code"`
	Rarity     string  `json:"set_rarity"`
	RarityCode string  `json:"set_rarity_code"`
	PriceUSD   float32 `json:"set_price,string"`
}

type YuGiOhAPICardImage struct {
	ID       int    `json:"id"`
	URL      string `json:"image_url"`
	URLSmall string `json:"image_url_small"`
}

type YuGiOhAPICardPrices struct {
	CardMarket   float32 `json:"cardmarket_price,string"`
	TCGPlayer    float32 `json:"tcgplayer_price,string"`
	Ebay         float32 `json:"ebay_price,string"`
	Amazonconst  float32 `json:"amazon_price,string"`
	CoolstuffInc float32 `json:"coolstuffinc_price,string"`
}
