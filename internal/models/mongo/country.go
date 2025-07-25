package mongo

type Country struct {
	ID            string `bson:"_id,omitempty" json:"id,omitempty"`
	CountryName   string `bson:"country_name" json:"country_name"`
	CountrySymbol string `bson:"country_symbol" json:"country_symbol"`
	CountryCode   string `bson:"country_code" json:"country_code"`
	DialCode      string `bson:"dial_code" json:"dial_code"`
}
