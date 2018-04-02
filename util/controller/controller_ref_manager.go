/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package controller

import (
	"fmt"
	"sync"

	"github.com/golang/glog"

	apps "k8s.io/api/apps/v1beta2"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
)

// GetControllerOf returns the controllerRef if controllee has a controller,
// otherwise returns nil.
func GetControllerOf(controllee metav1.Object) *metav1.OwnerReference {
	ownerRefs := controllee.GetOwnerReferences()
	for i := range ownerRefs {
		owner := &ownerRefs[i]
		if owner.Controller != nil && *owner.Controller == true {
			return owner
		}
	}
	return nil
}

// RecheckDeletionTimestamp returns a canAdopt() function to recheck deletion.
//
// The canAdopt() function calls getObject() to fetch the latest value,
// and denies adoption attempts if that object has a non-nil DeletionTimestamp.
func RecheckDeletionTimestamp(getObject func() (metav1.Object, error)) func() error {
	return func() error {
		obj, err := getObject()
		if err != nil {
			return fmt.Errorf("can't recheck DeletionTimestamp: %v", err)
		}
		if obj.GetDeletionTimestamp() != nil {
			return fmt.Errorf("%v/%v has just been deleted at %v", obj.GetNamespace(), obj.GetName(), obj.GetDeletionTimestamp())
		}
		return nil
	}
}

type baseControllerRefManager struct {
	controller metav1.Object
	selector   labels.Selector

	canAdoptErr  error
	canAdoptOnce sync.Once
	canAdoptFunc func() error
}

func (m *baseControllerRefManager) canAdopt() error {
	m.canAdoptOnce.Do(func() {
		if m.canAdoptFunc != nil {
			m.canAdoptErr = m.canAdoptFunc()
		}
	})
	return m.canAdoptErr
}

// claimObject tries to take ownership of an object for this controller.
//
// It will reconcile the following:
//   * Adopt orphans if the match function returns true.
//   * Release owned objects if the match function returns false.
//
// A non-nil error is returned if some form of reconciliation was attemped and
// failed. Usually, controllers should try again later in case reconciliation
// is still needed.
//
// If the error is nil, either the reconciliation succeeded, or no
// reconciliation was necessary. The returned boolean indicates whether you now
// own the object.
//
// No reconciliation will be attempted if the controller is being deleted.
func (m *baseControllerRefManager) claimObject(obj metav1.Object, match func(metav1.Object) bool, adopt, release func(metav1.Object) error) (bool, error) {
	controllerRef := GetControllerOf(obj)
	if controllerRef != nil {
		if controllerRef.UID != m.controller.GetUID() {
			// Owned by someone else. Ignore.
			return false, nil
		}
		if match(obj) {
			// We already own it and the selector matches.
			// Return true (successfully claimed) before checking deletion timestamp.
			// We're still allowed to claim things we already own while being deleted
			// because doing so requires taking no actions.
			return true, nil
		}
		// Owned by us but selector doesn't match.
		// Try to release, unless we're being deleted.
		if m.controller.GetDeletionTimestamp() != nil {
			return false, nil
		}
		if err := release(obj); err != nil {
			// If the pod no longer exists, ignore the error.
			if errors.IsNotFound(err) {
				return false, nil
			}
			// Either someone else released it, or there was a transient error.
			// The controller should requeue and try again if it's still stale.
			return false, err
		}
		// Successfully released.
		return false, nil
	}

	// It's an orphan.
	if m.controller.GetDeletionTimestamp() != nil || !match(obj) {
		// Ignore if we're being deleted or selector doesn't match.
		return false, nil
	}
	if obj.GetDeletionTimestamp() != nil {
		// Ignore if the object is being deleted
		return false, nil
	}
	// Selector matches. Try to adopt.
	if err := adopt(obj); err != nil {
		// If the pod no longer exists, ignore the error.
		if errors.IsNotFound(err) {
			return false, nil
		}
		// Either someone else claimed it first, or there was a transient error.
		// The controller should requeue and try again if it's still orphaned.
		return false, err
	}
	// Successfully adopted.
	return true, nil
}

// DaemonSetControllerRefManager is used to manage controllerRef of DaemontSet.
// Three methods are defined on this object 1: Classify 2: AdoptDaemonSet and
// 3: ReleaseDaemonSet which are used to classify the DaemonSet into appropriate
// categories and accordingly adopt or release them. See comments on these functions
// for more details.
type DaemonSetControllerRefManager struct {
	baseControllerRefManager
	controllerKind schema.GroupVersionKind
	client         kubernetes.Interface
}

func NewDaemonSetControllerRefManager(
	client kubernetes.Interface,
	controller metav1.Object,
	selector labels.Selector,
	controllerKind schema.GroupVersionKind,
	canAdopt func() error,
) *DaemonSetControllerRefManager {
	return &DaemonSetControllerRefManager{
		baseControllerRefManager: baseControllerRefManager{
			controller:   controller,
			selector:     selector,
			canAdoptFunc: canAdopt,
		},
		controllerKind: controllerKind,
		client:         client,
	}
}

// Claim tries to take ownership of a list of DaemonSets.
//
// It will reconcile the following:
//   * Adopt orphans if the selector matches.
//   * Release owned objects if the selector no longer matches.
//
// A non-nil error is returned if some form of reconciliation was attemped and
// failed. Usually, controllers should try again later in case reconciliation
// is still needed.
//
// If the error is nil, either the reconciliation succeeded, or no
// reconciliation was necessary. The list of DaemonSets that you now own is
// returned.
func (m *DaemonSetControllerRefManager) Claim(sets []*apps.DaemonSet) ([]*apps.DaemonSet, error) {
	var claimed []*apps.DaemonSet
	var errlist []error

	match := func(obj metav1.Object) bool {
		return m.selector.Matches(labels.Set(obj.GetLabels()))
	}
	adopt := func(obj metav1.Object) error {
		return m.Adopt(obj.(*apps.DaemonSet))
	}
	release := func(obj metav1.Object) error {
		return m.Release(obj.(*apps.DaemonSet))
	}

	for _, rs := range sets {
		ok, err := m.claimObject(rs, match, adopt, release)
		if err != nil {
			errlist = append(errlist, err)
			continue
		}
		if ok {
			claimed = append(claimed, rs)
		}
	}
	return claimed, utilerrors.NewAggregate(errlist)
}

// Adopt sends a patch to take control of the DaemonSet. It returns the error if
// the patching fails.
func (m *DaemonSetControllerRefManager) Adopt(ds *apps.DaemonSet) error {
	if err := m.canAdopt(); err != nil {
		return fmt.Errorf("can't adopt DaemontSet %v/%v (%v): %v", ds.Namespace, ds.Name, ds.UID, err)
	}
	// Note that ValidateOwnerReferences() will reject this patch if another
	// OwnerReference exists with controller=true.
	addControllerPatch := fmt.Sprintf(
		`{"metadata":{"ownerReferences":[{"apiVersion":"%s","kind":"%s","name":"%s","uid":"%s","controller":true,"blockOwnerDeletion":true}],"uid":"%s"}}`,
		m.controllerKind.GroupVersion(), m.controllerKind.Kind,
		m.controller.GetName(), m.controller.GetUID(), ds.UID)

	_, err := m.client.AppsV1beta2().DaemonSets(ds.Namespace).Patch(ds.Name, types.StrategicMergePatchType, []byte(addControllerPatch))
	return err
}

// Release sends a patch to free the DaemonSet from the control of the LoadBalancer controller.
// It returns the error if the patching fails. 404 and 422 errors are ignored.
func (m *DaemonSetControllerRefManager) Release(ds *apps.DaemonSet) error {
	glog.V(2).Infof("patching deamonset %s_%s to remove its controllerRef to %s/%s:%s",
		ds.Namespace, ds.Name, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.controller.GetName())
	deleteOwnerRefPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"$patch":"delete","uid":"%s"}],"uid":"%s"}}`, m.controller.GetUID(), ds.UID)
	_, err := m.client.AppsV1beta2().DaemonSets(ds.Namespace).Patch(ds.Name, types.StrategicMergePatchType, []byte(deleteOwnerRefPatch))
	if err != nil {
		if errors.IsNotFound(err) {
			// if DaemonSet no longer exists, ignore it
			return nil
		}
		if errors.IsInvalid(err) {
			// Invalid error will be returned in two cases: 1. the DaemonSet
			// has no owner reference, 2. the uid of the DaemonSet doesn't
			// match, which means the DaemonSet is deleted and then recreated.
			// In both cases, the error can be ignored.
			return nil
		}
	}
	return err
}

// DeploymentControllerRefManager is used to manage controllerRef of Deployment.
// Three methods are defined on this object 1: Classify 2: AdoptDeployment and
// 3: ReleaseDeployment which are used to classify the Deployment into appropriate
// categories and accordingly adopt or release them. See comments on these functions
// for more details.
type DeploymentControllerRefManager struct {
	baseControllerRefManager
	controllerKind schema.GroupVersionKind
	client         kubernetes.Interface
}

func NewDeploymentControllerRefManager(
	client kubernetes.Interface,
	controller metav1.Object,
	selector labels.Selector,
	controllerKind schema.GroupVersionKind,
	canAdopt func() error,
) *DeploymentControllerRefManager {
	return &DeploymentControllerRefManager{
		baseControllerRefManager: baseControllerRefManager{
			controller:   controller,
			selector:     selector,
			canAdoptFunc: canAdopt,
		},
		controllerKind: controllerKind,
		client:         client,
	}
}

// Claim tries to take ownership of a list of Deployments.
//
// It will reconcile the following:
//   * Adopt orphans if the selector matches.
//   * Release owned objects if the selector no longer matches.
//
// A non-nil error is returned if some form of reconciliation was attemped and
// failed. Usually, controllers should try again later in case reconciliation
// is still needed.
//
// If the error is nil, either the reconciliation succeeded, or no
// reconciliation was necessary. The list of Deployments that you now own is
// returned.
func (m *DeploymentControllerRefManager) Claim(sets []*apps.Deployment) ([]*apps.Deployment, error) {
	var claimed []*apps.Deployment
	var errlist []error

	match := func(obj metav1.Object) bool {
		return m.selector.Matches(labels.Set(obj.GetLabels()))
	}
	adopt := func(obj metav1.Object) error {
		return m.Adopt(obj.(*apps.Deployment))
	}
	release := func(obj metav1.Object) error {
		return m.Release(obj.(*apps.Deployment))
	}

	for _, rs := range sets {
		ok, err := m.claimObject(rs, match, adopt, release)
		if err != nil {
			errlist = append(errlist, err)
			continue
		}
		if ok {
			claimed = append(claimed, rs)
		}
	}
	return claimed, utilerrors.NewAggregate(errlist)
}

// Adopt sends a patch to take control of the Deployment. It returns the error if
// the patching fails.
func (m *DeploymentControllerRefManager) Adopt(d *apps.Deployment) error {
	if err := m.canAdopt(); err != nil {
		return fmt.Errorf("can't adopt Deployment %v/%v (%v): %v", d.Namespace, d.Name, d.UID, err)
	}
	// Note that ValidateOwnerReferences() will reject this patch if another
	// OwnerReference exists with controller=true.
	addControllerPath := fmt.Sprintf(
		`{"metadata":{"ownerReferences":[{"apiVersion":"%s","kind":"%s","name":"%s","uid":"%s","controller":true,"blockOwnerDeletion":true}],"uid":"%s"}}`,
		m.controllerKind.GroupVersion(), m.controllerKind.Kind,
		m.controller.GetName(), m.controller.GetUID(), d.UID)

	_, err := m.client.AppsV1beta2().Deployments(d.Namespace).Patch(d.Name, types.StrategicMergePatchType, []byte(addControllerPath))
	return err
}

// Release sends a patch to free the Deployment from the control of the LoadBalancer controller.
// It returns the error if the patching fails. 404 and 422 errors are ignored.
func (m *DeploymentControllerRefManager) Release(d *apps.Deployment) error {
	glog.V(2).Infof("patching deployment %s_%s to remove its controllerRef to %s/%s:%s",
		d.Namespace, d.Name, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.controller.GetName())
	deleteOwnerRefPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"$patch":"delete","uid":"%s"}],"uid":"%s"}}`, m.controller.GetUID(), d.UID)
	_, err := m.client.AppsV1beta2().Deployments(d.Namespace).Patch(d.Name, types.StrategicMergePatchType, []byte(deleteOwnerRefPatch))
	if err != nil {
		if errors.IsNotFound(err) {
			// if DaemonSet no longer exists, ignore it
			return nil
		}
		if errors.IsInvalid(err) {
			// Invalid error will be returned in two cases: 1. the DaemonSet
			// has no owner reference, 2. the uid of the DaemonSet doesn't
			// match, which means the DaemonSet is deleted and then recreated.
			// In both cases, the error can be ignored.
			return nil
		}
	}
	return err
}
