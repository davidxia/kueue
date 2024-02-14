/*
Copyright The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
	kueuev1beta1 "sigs.k8s.io/kueue/client-go/applyconfiguration/kueue/v1beta1"
)

// FakeWorkloadPriorityClasses implements WorkloadPriorityClassInterface
type FakeWorkloadPriorityClasses struct {
	Fake *FakeKueueV1beta1
}

var workloadpriorityclassesResource = v1beta1.SchemeGroupVersion.WithResource("workloadpriorityclasses")

var workloadpriorityclassesKind = v1beta1.SchemeGroupVersion.WithKind("WorkloadPriorityClass")

// Get takes name of the workloadPriorityClass, and returns the corresponding workloadPriorityClass object, and an error if there is any.
func (c *FakeWorkloadPriorityClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.WorkloadPriorityClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(workloadpriorityclassesResource, name), &v1beta1.WorkloadPriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadPriorityClass), err
}

// List takes label and field selectors, and returns the list of WorkloadPriorityClasses that match those selectors.
func (c *FakeWorkloadPriorityClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.WorkloadPriorityClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(workloadpriorityclassesResource, workloadpriorityclassesKind, opts), &v1beta1.WorkloadPriorityClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.WorkloadPriorityClassList{ListMeta: obj.(*v1beta1.WorkloadPriorityClassList).ListMeta}
	for _, item := range obj.(*v1beta1.WorkloadPriorityClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workloadPriorityClasses.
func (c *FakeWorkloadPriorityClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(workloadpriorityclassesResource, opts))
}

// Create takes the representation of a workloadPriorityClass and creates it.  Returns the server's representation of the workloadPriorityClass, and an error, if there is any.
func (c *FakeWorkloadPriorityClasses) Create(ctx context.Context, workloadPriorityClass *v1beta1.WorkloadPriorityClass, opts v1.CreateOptions) (result *v1beta1.WorkloadPriorityClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(workloadpriorityclassesResource, workloadPriorityClass), &v1beta1.WorkloadPriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadPriorityClass), err
}

// Update takes the representation of a workloadPriorityClass and updates it. Returns the server's representation of the workloadPriorityClass, and an error, if there is any.
func (c *FakeWorkloadPriorityClasses) Update(ctx context.Context, workloadPriorityClass *v1beta1.WorkloadPriorityClass, opts v1.UpdateOptions) (result *v1beta1.WorkloadPriorityClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(workloadpriorityclassesResource, workloadPriorityClass), &v1beta1.WorkloadPriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadPriorityClass), err
}

// Delete takes name of the workloadPriorityClass and deletes it. Returns an error if one occurs.
func (c *FakeWorkloadPriorityClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(workloadpriorityclassesResource, name, opts), &v1beta1.WorkloadPriorityClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkloadPriorityClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(workloadpriorityclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.WorkloadPriorityClassList{})
	return err
}

// Patch applies the patch and returns the patched workloadPriorityClass.
func (c *FakeWorkloadPriorityClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.WorkloadPriorityClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(workloadpriorityclassesResource, name, pt, data, subresources...), &v1beta1.WorkloadPriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadPriorityClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied workloadPriorityClass.
func (c *FakeWorkloadPriorityClasses) Apply(ctx context.Context, workloadPriorityClass *kueuev1beta1.WorkloadPriorityClassApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.WorkloadPriorityClass, err error) {
	if workloadPriorityClass == nil {
		return nil, fmt.Errorf("workloadPriorityClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(workloadPriorityClass)
	if err != nil {
		return nil, err
	}
	name := workloadPriorityClass.Name
	if name == nil {
		return nil, fmt.Errorf("workloadPriorityClass.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(workloadpriorityclassesResource, *name, types.ApplyPatchType, data), &v1beta1.WorkloadPriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadPriorityClass), err
}
