package route53resolver

import (
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/crossplane/provider-aws/apis/route53resolver/v1alpha1"
)

const (
	// ResolverRuleAssociation not found
	ResolverRuleAssociationNotFound = "NoResourceFoundException"
)

// ResolverRuleAssociationClient is the external client used for ResolverRuleAssociation Custom Resource
type ResolverRuleAssociationClient interface {
	GetResolverRuleAssociationRequest(*route53resolver.GetResolverRuleAssociationInput) route53resolver.GetResolverRuleAssociationRequest
	AssociateResolverRuleRequest(*route53resolver.AssociateResolverRuleInput) route53resolver.AssociateResolverRuleRequest
	DisassociateResolverRuleRequest(*route53resolver.DisassociateResolverRuleInput) route53resolver.DisassociateResolverRuleRequest
}

// IsResolverRuleAssociationNotFoundErr returns true when the error is due to the ResolverRuleAssociation not existing
func IsResolverRuleAssociationNotFoundErr(err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		if awsErr.Code() == ResolverRuleAssociationNotFound {
			return true
		}
	}
	return false
}

// GenerateRoute53ResolverObservation is used to produce v1alpha1.ResolverRuleAssociationObservation from the aws ResolverRuleAssociation resource
func GenerateRoute53ResolverObservation(resolverruleassociation route53resolver.ResolverRuleAssociation) v1alpha1.ResolverRuleAssociationObservation {
	o := v1alpha1.ResolverRuleAssociationObservation{
		ID:            resolverruleassociation.Id,
		RuleID:        resolverruleassociation.ResolverRuleId,
		VPCID:         resolverruleassociation.VPCId,
		StatusMessage: resolverruleassociation.StatusMessage,
		Status:        aws.String(string(resolverruleassociation.Status)),
	}

	return o
}

// LateInitializeResolverRuleAssociation fills the empty fields in *vv1alpha1.ResolverRuleAssociationParameters with
// the values seen in route53resolver.ResolverRuleAssociation
func LateInitializeResolverRuleAssociation(in *v1alpha1.ResolverRuleAssociationParameters, a *route53resolver.ResolverRuleAssociation) { // nolint:gocyclo
	if a == nil {
		return
	}
	in.Name = a.Name
}
