package models

type AppParameters struct {
	Id                        int32  `json:"id"`
	Label                     string `json:"label"`
	UserAttributeMappings     string `json:"user_attribute_mappings"`
	UserAttributeMacros       string `json:"user_attribute_macros"`
	AttributesTransformations string `json:"attributes_transformations"`
	SkipIfBlank               bool   `json:"skip_if_blank"`
	Values                    string `json:"values"`
	DefaultValues             string `json:"default_values"`
	ProvisionedEntitlements   bool   `json:"provisioned_entitlements"`
	SafeEntitlementsEnabled   bool   `json:"safe_entitlements_enabled"`
}
