// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeControllerManagers implements KubeControllerManagerInterface
type FakeKubeControllerManagers struct {
	Fake *FakeOperatorV1
}

var kubecontrollermanagersResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "kubecontrollermanagers"}

var kubecontrollermanagersKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "KubeControllerManager"}

// Get takes name of the kubeControllerManager, and returns the corresponding kubeControllerManager object, and an error if there is any.
func (c *FakeKubeControllerManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.KubeControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubecontrollermanagersResource, name), &operatorv1.KubeControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeControllerManager), err
}

// List takes label and field selectors, and returns the list of KubeControllerManagers that match those selectors.
func (c *FakeKubeControllerManagers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.KubeControllerManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubecontrollermanagersResource, kubecontrollermanagersKind, opts), &operatorv1.KubeControllerManagerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.KubeControllerManagerList{ListMeta: obj.(*operatorv1.KubeControllerManagerList).ListMeta}
	for _, item := range obj.(*operatorv1.KubeControllerManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeControllerManagers.
func (c *FakeKubeControllerManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubecontrollermanagersResource, opts))
}

// Create takes the representation of a kubeControllerManager and creates it.  Returns the server's representation of the kubeControllerManager, and an error, if there is any.
func (c *FakeKubeControllerManagers) Create(ctx context.Context, kubeControllerManager *operatorv1.KubeControllerManager, opts v1.CreateOptions) (result *operatorv1.KubeControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubecontrollermanagersResource, kubeControllerManager), &operatorv1.KubeControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeControllerManager), err
}

// Update takes the representation of a kubeControllerManager and updates it. Returns the server's representation of the kubeControllerManager, and an error, if there is any.
func (c *FakeKubeControllerManagers) Update(ctx context.Context, kubeControllerManager *operatorv1.KubeControllerManager, opts v1.UpdateOptions) (result *operatorv1.KubeControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubecontrollermanagersResource, kubeControllerManager), &operatorv1.KubeControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeControllerManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeControllerManagers) UpdateStatus(ctx context.Context, kubeControllerManager *operatorv1.KubeControllerManager, opts v1.UpdateOptions) (*operatorv1.KubeControllerManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubecontrollermanagersResource, "status", kubeControllerManager), &operatorv1.KubeControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeControllerManager), err
}

// Delete takes name of the kubeControllerManager and deletes it. Returns an error if one occurs.
func (c *FakeKubeControllerManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubecontrollermanagersResource, name, opts), &operatorv1.KubeControllerManager{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeControllerManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubecontrollermanagersResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.KubeControllerManagerList{})
	return err
}

// Patch applies the patch and returns the patched kubeControllerManager.
func (c *FakeKubeControllerManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.KubeControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubecontrollermanagersResource, name, pt, data, subresources...), &operatorv1.KubeControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeControllerManager), err
}
