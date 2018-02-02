/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

type Locale struct {

	Language string `json:"language,omitempty"`

	Country string `json:"country,omitempty"`

	DisplayCountry string `json:"displayCountry,omitempty"`

	DisplayLanguage string `json:"displayLanguage,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	DisplayScript string `json:"displayScript,omitempty"`

	DisplayVariant string `json:"displayVariant,omitempty"`

	ExtensionKeys []string `json:"extensionKeys,omitempty"`

	Iso3Country string `json:"iso3Country,omitempty"`

	Iso3Language string `json:"iso3Language,omitempty"`

	Script string `json:"script,omitempty"`

	UnicodeLocaleAttributes []string `json:"unicodeLocaleAttributes,omitempty"`

	UnicodeLocaleKeys []string `json:"unicodeLocaleKeys,omitempty"`

	Variant string `json:"variant,omitempty"`
}
