/*
Copyright 2019 Antrea Authors

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
	v1beta1 "github.com/vmware-tanzu/antrea/pkg/apis/networkpolicy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppliedToGroups implements AppliedToGroupInterface
type FakeAppliedToGroups struct {
	Fake *FakeNetworkpolicyV1beta1
}

var appliedtogroupsResource = schema.GroupVersionResource{Group: "networkpolicy.antrea.io", Version: "v1beta1", Resource: "appliedtogroups"}

var appliedtogroupsKind = schema.GroupVersionKind{Group: "networkpolicy.antrea.io", Version: "v1beta1", Kind: "AppliedToGroup"}

// Get takes name of the appliedToGroup, and returns the corresponding appliedToGroup object, and an error if there is any.
func (c *FakeAppliedToGroups) Get(name string, options v1.GetOptions) (result *v1beta1.AppliedToGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(appliedtogroupsResource, name), &v1beta1.AppliedToGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppliedToGroup), err
}

// List takes label and field selectors, and returns the list of AppliedToGroups that match those selectors.
func (c *FakeAppliedToGroups) List(opts v1.ListOptions) (result *v1beta1.AppliedToGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(appliedtogroupsResource, appliedtogroupsKind, opts), &v1beta1.AppliedToGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AppliedToGroupList{ListMeta: obj.(*v1beta1.AppliedToGroupList).ListMeta}
	for _, item := range obj.(*v1beta1.AppliedToGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appliedToGroups.
func (c *FakeAppliedToGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(appliedtogroupsResource, opts))
}
