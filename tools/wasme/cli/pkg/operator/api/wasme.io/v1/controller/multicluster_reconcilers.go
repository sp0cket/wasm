// Code generated by skv2. DO NOT EDIT.

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	wasme_io_v1 "github.com/solo-io/wasm/tools/wasme/cli/pkg/operator/api/wasme.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the FilterDeployment Resource across clusters.
// implemented by the user
type MulticlusterFilterDeploymentReconciler interface {
	ReconcileFilterDeployment(clusterName string, obj *wasme_io_v1.FilterDeployment) (reconcile.Result, error)
}

// Reconcile deletion events for the FilterDeployment Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFilterDeploymentDeletionReconciler interface {
	ReconcileFilterDeploymentDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFilterDeploymentReconcilerFuncs struct {
	OnReconcileFilterDeployment         func(clusterName string, obj *wasme_io_v1.FilterDeployment) (reconcile.Result, error)
	OnReconcileFilterDeploymentDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFilterDeploymentReconcilerFuncs) ReconcileFilterDeployment(clusterName string, obj *wasme_io_v1.FilterDeployment) (reconcile.Result, error) {
	if f.OnReconcileFilterDeployment == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFilterDeployment(clusterName, obj)
}

func (f *MulticlusterFilterDeploymentReconcilerFuncs) ReconcileFilterDeploymentDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFilterDeploymentDeletion == nil {
		return nil
	}
	return f.OnReconcileFilterDeploymentDeletion(clusterName, req)
}

type MulticlusterFilterDeploymentReconcileLoop interface {
	// AddMulticlusterFilterDeploymentReconciler adds a MulticlusterFilterDeploymentReconciler to the MulticlusterFilterDeploymentReconcileLoop.
	AddMulticlusterFilterDeploymentReconciler(ctx context.Context, rec MulticlusterFilterDeploymentReconciler, predicates ...predicate.Predicate)
}

type multiclusterFilterDeploymentReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFilterDeploymentReconcileLoop) AddMulticlusterFilterDeploymentReconciler(ctx context.Context, rec MulticlusterFilterDeploymentReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFilterDeploymentMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFilterDeploymentReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterFilterDeploymentReconcileLoop {
	return &multiclusterFilterDeploymentReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &wasme_io_v1.FilterDeployment{})}
}

type genericFilterDeploymentMulticlusterReconciler struct {
	reconciler MulticlusterFilterDeploymentReconciler
}

func (g genericFilterDeploymentMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFilterDeploymentDeletionReconciler); ok {
		return deletionReconciler.ReconcileFilterDeploymentDeletion(cluster, req)
	}
	return nil
}

func (g genericFilterDeploymentMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*wasme_io_v1.FilterDeployment)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FilterDeployment handler received event for %T", object)
	}
	return g.reconciler.ReconcileFilterDeployment(cluster, obj)
}
