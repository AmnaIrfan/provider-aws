/*
Copyright 2019 The Crossplane Authors.

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
// +kubebuilder:object:generate=true
// +groupName=route53resolver.aws.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import "k8s.io/apimachinery/pkg/runtime/schema"

// Repository type metadata.
var (
	ResolverRuleAssociationKind             = "ResolverRuleAssociation"
	ResolverRuleAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: ResolverRuleAssociationKind}.String()
	ResolverRuleAssociationKindAPIVersion   = ResolverRuleAssociationKind + "." + GroupVersion.String()
	ResolverRuleAssociationGroupVersionKind = GroupVersion.WithKind(ResolverRuleAssociationKind)
)

func init() {
	SchemeBuilder.Register(&ResolverRuleAssociation{}, &ResolverRuleAssociationList{})
}
