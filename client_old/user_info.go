/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

type UserInfo struct {

	// Specification links
	Schemas []string `json:"schemas,omitempty"`

	// Internal identifier of a user
	Id string `json:"id,omitempty"`

	// External identifier of a user
	ExternalId string `json:"externalId,omitempty"`

	// User metadata
	Meta *MetaInfo `json:"meta,omitempty"`

	// User mailbox
	UserName string `json:"userName,omitempty"`

	// User name
	Name *NameInfo `json:"name,omitempty"`

	// Status of a user
	Active bool `json:"active,omitempty"`

	// User email addresses
	Emails []EmailInfo `json:"emails,omitempty"`

	// User phone numbers
	PhoneNumbers []PhoneNumberInfo `json:"phoneNumbers,omitempty"`

	// User addresses
	Addresses []AddressInfo `json:"addresses,omitempty"`
}