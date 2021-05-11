package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//GetResolverRuleAssociationWithContext(
//DisassociateResolverRuleWithContext
//AssociateResolverRuleWithContext

// +kubebuilder:object:root=true

// A ResolverRuleAssociation is a managed resource that represents the association between a Route53 Resolver Rule and VPC
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="ID",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="VPC",type="string",JSONPath=".spec.forProvider.vpcId"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResolverRuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResolverRuleAssociationSpec   `json:"spec"`
	Status ResolverRuleAssociationStatus `json:"status,omitempty"`
}

// A ResolverRuleAssociationSpec defines the desired state of a ResolverRuleAssociation
type ResolverRuleAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ResolverRuleAssociationParameters `json:"forProvider"`
}

//ResolverRuleAssociationParameters defines the desired state of a ResolverRuleAssociation
type ResolverRuleAssociationParameters struct {
	// ID of resolver rule association
	ID *string `json:"id"`
	// A name for the association that you're creating between a Resolver rule and
	// a VPC.
	// +optional
	Name *string `json:"name"`

	// The ID of the Resolver rule that you want to associate with the VPC.
	// +optional
	// +immutable
	ResolverRuleID *string `json:"resolverRuleId"`

	// VPCIDRef references a Resolver Rule to retrieve its resolverRuleId
	// +optional
	// +immutable
	ResolverRuleIDRef *xpv1.Reference `json:"resolverRuleIdRef,omitempty"`

	// ResolverRuleIDSelector selects a reference to a Resolver Rule to retrieve its resolverRuleId
	// +optional
	ResolverRuleIDSelector *xpv1.Selector `json:"resolverRuleIdSelector,omitempty"`

	// The ID of the VPC that you want to associate the Resolver rule with.
	// +optional
	// +immutable
	VPCID *string `json:"vpcId"`

	// VPCIDRef references a VPC to retrieve its vpcId
	// +optional
	// +immutable
	VPCIDRef *xpv1.Reference `json:"vpcIdRef,omitempty"`

	// VPCIDSelector selects a reference to a VPC to retrieve its vpcId
	// +optional
	VPCIDSelector *xpv1.Selector `json:"vpcIdSelector,omitempty"`
}

// A ResolverRuleAssociationStatus  represents the observed state of a ResolverRuleAssociation
type ResolverRuleAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ResolverRuleAssociationObservation `json:"atProvider"`
}

// ResolverRuleAssociationObservation keeps the state for the external resource
type ResolverRuleAssociationObservation struct {
	// The ID of the Resolver Rule Association
	ID *string `json:"id,omitempty"`

	// The ARN (Amazon Resource Name) for the Resolver Rule Association
	ARN *string `json:"arn,omitempty"`

	// The date and time that the Resolver Rule Association was created, in Unix time format and Coordinated Universal Time (UTC).
	CreationTime *string `json:"creationTime,omitempty"`

	// The ID of the VPC that the Resolver Rule is associated to
	VPCID *string `json:"VPCID,omitempty"`

	ModificationTime *string `json:"modificationTime,omitempty"`

	RuleID *string `json:"RuleID,omitempty"`

	Status *string `json:"status,omitempty"`

	// A detailed description of the status of the Resolver endpoint.
	StatusMessage *string `json:"statusMessage,omitempty"`
}

// +kubebuilder:object:root=true
// ResolverRuleAssociationList contains a list of ResolverRuleAssociations
type ResolverRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResolverRuleAssociation `json:"items"`
}
