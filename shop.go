package go_shopify_request

type ShopInfo struct {
	Shop Shop `json:"shop"`
}

type Shop struct {
	Id                               int64    `json:"id"`
	Name                             string   `json:"name"`
	Email                            string   `json:"email"`
	Domain                           string   `json:"domain"`
	Province                         string   `json:"province"`
	Country                          string   `json:"country"`
	Address1                         string   `json:"address1"`
	Zip                              string   `json:"zip"`
	City                             string   `json:"city"`
	Source                           string   `json:"source"`
	Phone                            string   `json:"phone"`
	Latitude                         float64  `json:"latitude"`
	Longitude                        float64  `json:"longitude"`
	PrimaryLocale                    string   `json:"primary_locale"`
	Address2                         string   `json:"address2"`
	CreatedAt                        string   `json:"created_at"`
	UpdatedAt                        string   `json:"updated_at"`
	CountryCode                      string   `json:"country_code"`
	CountryName                      string   `json:"country_name"`
	Currency                         string   `json:"currency"`
	CustomerEmail                    string   `json:"customer_email"`
	Timezone                         string   `json:"timezone"`
	IanaTimezone                     string   `json:"iana_timezone"`
	ShopOwner                        string   `json:"shop_owner"`
	MoneyFormat                      string   `json:"money_format"`
	MoneyWithCurrencyFormat          string   `json:"money_with_currency_format"`
	WeightUnit                       string   `json:"weight_unit"`
	ProvinceCode                     string   `json:"province_code"`
	TaxesIncluded                    bool     `json:"taxes_included"`
	CountyTaxes                      bool     `json:"county_taxes"`
	PlanDisplayName                  string   `json:"plan_display_name"`
	PlanName                         string   `json:"plan_name"`
	HasDiscounts                     bool     `json:"has_discounts"`
	HasGiftCards                     bool     `json:"has_gift_cards"`
	MyshopifyDomain                  string   `json:"myshopify_domain"`
	GoogleAppsDomain                 string   `json:"google_apps_domain"`
	GoogleAppsLoginEnabled           string   `json:"google_apps_login_enabled"`
	MoneyInEmailsFormat              string   `json:"money_in_emails_format"`
	MoneyWithCurrencyInEmailsFormat  string   `json:"money_with_currency_in_emails_format"`
	EligibleForPayments              bool     `json:"eligible_for_payments"`
	RequiresExtraPaymentsAgreement   bool     `json:"requires_extra_payments_agreement"`
	PasswordEnabled                  bool     `json:"password_enabled"`
	HasStorefront                    bool     `json:"has_storefront"`
	EligibleForCardReaderGiveaway    bool     `json:"eligible_for_card_reader_giveaway"`
	Finances                         bool     `json:"finances"`
	PrimaryLocationId                int64    `json:"primary_location_id"`
	CookieConsentLevel               string   `json:"cookie_consent_level"`
	VisitorTrackingConsentPreference string   `json:"visitor_tracking_consent_preference"`
	ForceSsl                         bool     `json:"force_ssl"`
	CheckoutApiSupported             bool     `json:"checkout_api_supported"`
	MultiLocationEnabled             bool     `json:"multi_location_enabled"`
	SetupRequired                    bool     `json:"setup_required"`
	PreLaunchEnabled                 bool     `json:"pre_launch_enabled"`
	EnabledPresentmentCurrencies     []string `json:"enabled_presentment_currencies"`
}

type ShopRegisterWebhook struct {
	Webhook RegisterWebhook `json:"webhook"`
}

type RegisterWebhook struct {
	Topic   string `json:"topic"`
	Address string `json:"address"`
	Format  string `json:"format"`
}

type Webhooks struct {
	Webhooks []Webhook `json:"webhooks"`
}

type Webhook struct {
	Id        int64  `json:"id"`
	Address   string `json:"address"`
	Topic     string `json:"topic"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Format    string `json:"format"`
	//Fields                     []interface{} `json:"fields"`
	//MetafieldNamespaces        []interface{} `json:"metafield_namespaces"`
	ApiVersion string `json:"api_version"`
	//PrivateMetafieldNamespaces []interface{} `json:"private_metafield_namespaces"`
}

type ScriptTags struct {
	ScriptTags []ScriptTag `json:"script_tags"`
}

type ScriptTag struct {
	Id           int64  `json:"id"`
	Src          string `json:"src"`
	Event        string `json:"event"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DisplayScope string `json:"display_scope"`
	Cache        bool   `json:"cache"`
}
