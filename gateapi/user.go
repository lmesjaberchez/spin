/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type User struct {

	AccountNonExpired bool `json:"accountNonExpired,omitempty"`

	AccountNonLocked bool `json:"accountNonLocked,omitempty"`

	AllowedAccounts []string `json:"allowedAccounts,omitempty"`

	Authorities []GrantedAuthority `json:"authorities,omitempty"`

	CredentialsNonExpired bool `json:"credentialsNonExpired,omitempty"`

	Email string `json:"email,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Roles []string `json:"roles,omitempty"`

	Username string `json:"username,omitempty"`
}
