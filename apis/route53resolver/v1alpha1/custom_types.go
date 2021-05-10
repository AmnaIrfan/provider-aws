package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

type CustomResolverEndpointParameters struct {
	// The ID of one or more security groups that you want to use to control access
	// to this VPC. The security group that you specify must include one or more
	// inbound rules (for inbound Resolver endpoints) or outbound rules (for outbound
	// Resolver endpoints). Inbound and outbound rules must allow TCP and UDP access.
	// For inbound access, open port 53. For outbound access, open the port that
	// you're using for DNS queries on your network.
	SecurityGroupIDs []string `json:"securityGroupIDs,omitempty"`
	// SecurityGroupIDRefs is a list of references to SecurityGroups used to set
	// the SecurityGroupIDs.
	// +optional
	SecurityGroupIDRefs []xpv1.Reference `json:"securityGroupIdRefs,omitempty"`

	// SecurityGroupIDsSelector selects references to SecurityGroupID used
	// to set the SecurityGroupIDs.
	// +optional
	SecurityGroupIDSelector *xpv1.Selector `json:"securityGroupIdSelector,omitempty"`

	IPAddresses []*IPAddressRequest `json:"ipAddresses"`
}

type CustomResolverRuleParameters struct {
}

type CustomResolverQueryLogConfigParameters struct {
}

type IPAddressRequest struct {
	IP *string `json:"ip,omitempty"`

	SubnetID *string `json:"subnetID,omitempty"`

	// SubnetIDRefs is a list of references to Subnets used to set
	// the SubnetIDs.
	SubnetIDRef *xpv1.Reference `json:"subnetIdRef,omitempty"`

	// SubnetIDSelector selects references to Subnets used
	// to set the SubnetIDs.
	SubnetIDSelector *xpv1.Selector `json:"subnetIdSelector,omitempty"`
}
