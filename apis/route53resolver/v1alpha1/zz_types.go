/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type Filter struct {
	Name *string `json:"name,omitempty"`

	Values []*string `json:"values,omitempty"`
}

type IPAddressResponse struct {
	CreationTime *string `json:"creationTime,omitempty"`

	IP *string `json:"ip,omitempty"`

	IPID *string `json:"ipID,omitempty"`

	ModificationTime *string `json:"modificationTime,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`
}

type IPAddressUpdate struct {
	IP *string `json:"ip,omitempty"`

	IPID *string `json:"ipID,omitempty"`
}

type ResolverDNSSEcConfig struct {
	ID *string `json:"id,omitempty"`

	OwnerID *string `json:"ownerID,omitempty"`

	ResourceID *string `json:"resourceID,omitempty"`
}

type ResolverEndpoint_SDK struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *string `json:"creationTime,omitempty"`

	CreatorRequestID *string `json:"creatorRequestID,omitempty"`

	Direction *string `json:"direction,omitempty"`

	HostVPCID *string `json:"hostVPCID,omitempty"`

	ID *string `json:"id,omitempty"`

	IPAddressCount *int64 `json:"ipAddressCount,omitempty"`

	ModificationTime *string `json:"modificationTime,omitempty"`

	Name *string `json:"name,omitempty"`

	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`
}

type ResolverQueryLogConfig struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *string `json:"creationTime,omitempty"`

	CreatorRequestID *string `json:"creatorRequestID,omitempty"`

	ID *string `json:"id,omitempty"`

	OwnerID *string `json:"ownerID,omitempty"`

	ShareStatus *string `json:"shareStatus,omitempty"`
}

type ResolverQueryLogConfigAssociation struct {
	CreationTime *string `json:"creationTime,omitempty"`

	ID *string `json:"id,omitempty"`

	ResolverQueryLogConfigID *string `json:"resolverQueryLogConfigID,omitempty"`

	ResourceID *string `json:"resourceID,omitempty"`
}

type ResolverRuleConfig struct {
	Name *string `json:"name,omitempty"`

	ResolverEndpointID *string `json:"resolverEndpointID,omitempty"`

	TargetIPs []*TargetAddress `json:"targetIPs,omitempty"`
}

type ResolverRule_SDK struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *string `json:"creationTime,omitempty"`

	CreatorRequestID *string `json:"creatorRequestID,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	ID *string `json:"id,omitempty"`

	ModificationTime *string `json:"modificationTime,omitempty"`

	Name *string `json:"name,omitempty"`

	OwnerID *string `json:"ownerID,omitempty"`

	ResolverEndpointID *string `json:"resolverEndpointID,omitempty"`

	RuleType *string `json:"ruleType,omitempty"`

	ShareStatus *string `json:"shareStatus,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`

	TargetIPs []*TargetAddress `json:"targetIPs,omitempty"`
}

type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

type TargetAddress struct {
	IP *string `json:"ip,omitempty"`

	Port *int64 `json:"port,omitempty"`
}
