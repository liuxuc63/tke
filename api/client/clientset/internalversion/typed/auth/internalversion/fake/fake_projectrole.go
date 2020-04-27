/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	auth "tkestack.io/tke/api/auth"
)

// FakeProjectRoles implements ProjectRoleInterface
type FakeProjectRoles struct {
	Fake *FakeAuth
}

var projectrolesResource = schema.GroupVersionResource{Group: "auth.tkestack.io", Version: "", Resource: "projectroles"}

var projectrolesKind = schema.GroupVersionKind{Group: "auth.tkestack.io", Version: "", Kind: "ProjectRole"}

// Get takes name of the projectRole, and returns the corresponding projectRole object, and an error if there is any.
func (c *FakeProjectRoles) Get(name string, options v1.GetOptions) (result *auth.ProjectRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(projectrolesResource, name), &auth.ProjectRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.ProjectRole), err
}

// List takes label and field selectors, and returns the list of ProjectRoles that match those selectors.
func (c *FakeProjectRoles) List(opts v1.ListOptions) (result *auth.ProjectRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(projectrolesResource, projectrolesKind, opts), &auth.ProjectRoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &auth.ProjectRoleList{ListMeta: obj.(*auth.ProjectRoleList).ListMeta}
	for _, item := range obj.(*auth.ProjectRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectRoles.
func (c *FakeProjectRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(projectrolesResource, opts))
}

// Create takes the representation of a projectRole and creates it.  Returns the server's representation of the projectRole, and an error, if there is any.
func (c *FakeProjectRoles) Create(projectRole *auth.ProjectRole) (result *auth.ProjectRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(projectrolesResource, projectRole), &auth.ProjectRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.ProjectRole), err
}

// Update takes the representation of a projectRole and updates it. Returns the server's representation of the projectRole, and an error, if there is any.
func (c *FakeProjectRoles) Update(projectRole *auth.ProjectRole) (result *auth.ProjectRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(projectrolesResource, projectRole), &auth.ProjectRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.ProjectRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjectRoles) UpdateStatus(projectRole *auth.ProjectRole) (*auth.ProjectRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(projectrolesResource, "status", projectRole), &auth.ProjectRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.ProjectRole), err
}

// Delete takes name of the projectRole and deletes it. Returns an error if one occurs.
func (c *FakeProjectRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(projectrolesResource, name), &auth.ProjectRole{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(projectrolesResource, listOptions)

	_, err := c.Fake.Invokes(action, &auth.ProjectRoleList{})
	return err
}

// Patch applies the patch and returns the patched projectRole.
func (c *FakeProjectRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.ProjectRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(projectrolesResource, name, pt, data, subresources...), &auth.ProjectRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.ProjectRole), err
}