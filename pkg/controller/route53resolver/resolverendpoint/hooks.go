package resolverendpoint

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/route53resolver"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	svcapitypes "github.com/crossplane/provider-aws/apis/route53resolver/v1alpha1"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

// SetupResolverEndpoint adds a controller that reconciles ResolverEndpoints
func SetupResolverEndpoint(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(svcapitypes.ResolverEndpointGroupKind)
	opts := []option{
		func(e *external) {
			e.preObserve = preObserve
			e.preCreate = preCreate
			e.postCreate = postCreate
			e.preDelete = preDelete
			e.preUpdate = preUpdate
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&svcapitypes.ResolverEndpoint{}).
		Complete(managed.NewReconciler(mgr,
			cpresource.ManagedKind(svcapitypes.ResolverEndpointGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithInitializers(managed.NewDefaultProviderConfig(mgr.GetClient())),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func preObserve(_ context.Context, cr *svcapitypes.ResolverEndpoint, obj *svcsdk.GetResolverEndpointInput) error {
	obj.ResolverEndpointId = aws.String(meta.GetExternalName(cr))
	return nil
}

func preCreate(_ context.Context, cr *svcapitypes.ResolverEndpoint, obj *svcsdk.CreateResolverEndpointInput) error {
	for _, sg := range cr.Spec.ForProvider.SecurityGroupIDs {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, aws.String(sg))
	}
	obj.IpAddresses = []*svcsdk.IpAddressRequest{}
	for _, ip := range cr.Spec.ForProvider.IPAddresses {
		if ip != nil {
			obj.IpAddresses = append(obj.IpAddresses, &svcsdk.IpAddressRequest{
				SubnetId: ip.SubnetID,
			})
		}
	}
	return nil
}

func postCreate(_ context.Context, cr *svcapitypes.ResolverEndpoint, obj *svcsdk.CreateResolverEndpointOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	meta.SetExternalName(cr, aws.StringValue(obj.ResolverEndpoint.Id))
	cre.ExternalNameAssigned = true
	return cre, err
}

func preDelete(_ context.Context, cr *svcapitypes.ResolverEndpoint, obj *svcsdk.DeleteResolverEndpointInput) (bool, error) {
	obj.ResolverEndpointId = aws.String(meta.GetExternalName(cr))
	return false, nil
}

func preUpdate(_ context.Context, cr *svcapitypes.ResolverEndpoint, obj *svcsdk.UpdateResolverEndpointInput) error {
	obj.ResolverEndpointId = aws.String(meta.GetExternalName(cr))
	return nil
}
