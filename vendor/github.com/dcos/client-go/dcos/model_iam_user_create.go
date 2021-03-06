/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type IamUserCreate struct {
	Description  string `json:"description,omitempty"`
	Password     string `json:"password,omitempty"`
	PublicKey    string `json:"public_key,omitempty"`
	ProviderType string `json:"provider_type,omitempty"`
	ProviderId   string `json:"provider_id,omitempty"`
	ClusterUrl   string `json:"cluster_url,omitempty"`
	CreatorUid   string `json:"creator_uid,omitempty"`
}
