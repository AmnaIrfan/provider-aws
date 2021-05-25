package route53resolver

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awsr53r "github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/google/go-cmp/cmp"

	"github.com/crossplane/provider-aws/apis/ec2/v1beta1"
	"github.com/crossplane/provider-aws/apis/route53resolver/v1alpha1"
	aws "github.com/crossplane/provider-aws/pkg/clients"
)

var (
	allocationID            = "allocation"
	associationID           = "association"
	testIPAddress           = "0.0.0.0"
	poolName                = "pool"
	domain                  = "vpc"
	instanceID              = "instance"
	networkBorderGroup      = "border"
	networkInterfaceID      = "network"
	networkInterfaceOwnerID = "owner"
	testKey                 = "key"
	testValue               = "value"
	ec2tag                  = ec2.Tag{Key: &testKey, Value: &testValue}
	v1beta1Tag              = v1beta1.Tag{Key: testKey, Value: testValue}
)

func TestGenerateResolverRuleAssociationObservation(t *testing.T) {
	cases := map[string]struct {
		in  awsr53r.ResolverRuleAssociation
		out v1alpha1.ResolverRuleAssociationObservation
	}{
		"AllFilled": {
			in: ec2.Address{
				AllocationId:            aws.String(allocationID),
				AssociationId:           aws.String(associationID),
				CustomerOwnedIp:         aws.String(testIPAddress),
				CustomerOwnedIpv4Pool:   aws.String(poolName),
				Domain:                  ec2.DomainType(domain),
				InstanceId:              aws.String(instanceID),
				NetworkBorderGroup:      aws.String(networkBorderGroup),
				NetworkInterfaceId:      aws.String(networkInterfaceID),
				NetworkInterfaceOwnerId: aws.String(networkInterfaceOwnerID),
				PrivateIpAddress:        aws.String(testIPAddress),
				PublicIp:                aws.String(testIPAddress),
				PublicIpv4Pool:          aws.String(poolName),
				Tags:                    []ec2.Tag{ec2tag},
			},
			out: v1beta1.AddressObservation{
				AllocationID:            allocationID,
				AssociationID:           associationID,
				CustomerOwnedIP:         testIPAddress,
				CustomerOwnedIPv4Pool:   poolName,
				InstanceID:              instanceID,
				NetworkBorderGroup:      networkBorderGroup,
				NetworkInterfaceID:      networkInterfaceID,
				NetworkInterfaceOwnerID: networkInterfaceOwnerID,
				PrivateIPAddress:        testIPAddress,
				PublicIP:                testIPAddress,
				PublicIPv4Pool:          poolName,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := GenerateAddressObservation(tc.in)
			if diff := cmp.Diff(tc.out, r); diff != "" {
				t.Errorf("GenerateAddressObservation(...): -want, +got:\n%s", diff)
			}
		})
	}
}
