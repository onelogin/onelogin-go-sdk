package models

// AppParameters is the contract for parameters.
type AppParameters struct {
	ID                        *int32  `json:"id,omitempty"`
	Label                     *string `json:"label,omitempty"`
	UserAttributeMappings     *string `json:"user_attribute_mappings,omitempty"`
	UserAttributeMacros       *string `json:"user_attribute_macros,omitempty"`
	AttributesTransformations *string `json:"attributes_transformations,omitempty"`
	SkipIfBlank               *bool   `json:"skip_if_blank,omitempty"`
	Values                    *string `json:"values,omitempty,omitempty"`
	DefaultValues             *string `json:"default_values,omitempty"`
	ProvisionedEntitlements   *bool   `json:"provisioned_entitlements,omitempty"`
	SafeEntitlementsEnabled   *bool   `json:"safe_entitlements_enabled,omitempty"`
	IncludeInSamlAssertion    *bool   `json:"include_in_saml_assertion,omitempty"`
}
