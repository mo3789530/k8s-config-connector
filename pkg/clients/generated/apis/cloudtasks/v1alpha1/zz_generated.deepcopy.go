//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudTasksQueue) DeepCopyInto(out *CloudTasksQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudTasksQueue.
func (in *CloudTasksQueue) DeepCopy() *CloudTasksQueue {
	if in == nil {
		return nil
	}
	out := new(CloudTasksQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudTasksQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudTasksQueueList) DeepCopyInto(out *CloudTasksQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudTasksQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudTasksQueueList.
func (in *CloudTasksQueueList) DeepCopy() *CloudTasksQueueList {
	if in == nil {
		return nil
	}
	out := new(CloudTasksQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudTasksQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudTasksQueueSpec) DeepCopyInto(out *CloudTasksQueueSpec) {
	*out = *in
	if in.AppEngineRoutingOverride != nil {
		in, out := &in.AppEngineRoutingOverride, &out.AppEngineRoutingOverride
		*out = new(QueueAppEngineRoutingOverride)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.RateLimits != nil {
		in, out := &in.RateLimits, &out.RateLimits
		*out = new(QueueRateLimits)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RetryConfig != nil {
		in, out := &in.RetryConfig, &out.RetryConfig
		*out = new(QueueRetryConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.StackdriverLoggingConfig != nil {
		in, out := &in.StackdriverLoggingConfig, &out.StackdriverLoggingConfig
		*out = new(QueueStackdriverLoggingConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudTasksQueueSpec.
func (in *CloudTasksQueueSpec) DeepCopy() *CloudTasksQueueSpec {
	if in == nil {
		return nil
	}
	out := new(CloudTasksQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudTasksQueueStatus) DeepCopyInto(out *CloudTasksQueueStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudTasksQueueStatus.
func (in *CloudTasksQueueStatus) DeepCopy() *CloudTasksQueueStatus {
	if in == nil {
		return nil
	}
	out := new(CloudTasksQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueAppEngineRoutingOverride) DeepCopyInto(out *QueueAppEngineRoutingOverride) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueAppEngineRoutingOverride.
func (in *QueueAppEngineRoutingOverride) DeepCopy() *QueueAppEngineRoutingOverride {
	if in == nil {
		return nil
	}
	out := new(QueueAppEngineRoutingOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueRateLimits) DeepCopyInto(out *QueueRateLimits) {
	*out = *in
	if in.MaxBurstSize != nil {
		in, out := &in.MaxBurstSize, &out.MaxBurstSize
		*out = new(int)
		**out = **in
	}
	if in.MaxConcurrentDispatches != nil {
		in, out := &in.MaxConcurrentDispatches, &out.MaxConcurrentDispatches
		*out = new(int)
		**out = **in
	}
	if in.MaxDispatchesPerSecond != nil {
		in, out := &in.MaxDispatchesPerSecond, &out.MaxDispatchesPerSecond
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueRateLimits.
func (in *QueueRateLimits) DeepCopy() *QueueRateLimits {
	if in == nil {
		return nil
	}
	out := new(QueueRateLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueRetryConfig) DeepCopyInto(out *QueueRetryConfig) {
	*out = *in
	if in.MaxAttempts != nil {
		in, out := &in.MaxAttempts, &out.MaxAttempts
		*out = new(int)
		**out = **in
	}
	if in.MaxBackoff != nil {
		in, out := &in.MaxBackoff, &out.MaxBackoff
		*out = new(string)
		**out = **in
	}
	if in.MaxDoublings != nil {
		in, out := &in.MaxDoublings, &out.MaxDoublings
		*out = new(int)
		**out = **in
	}
	if in.MaxRetryDuration != nil {
		in, out := &in.MaxRetryDuration, &out.MaxRetryDuration
		*out = new(string)
		**out = **in
	}
	if in.MinBackoff != nil {
		in, out := &in.MinBackoff, &out.MinBackoff
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueRetryConfig.
func (in *QueueRetryConfig) DeepCopy() *QueueRetryConfig {
	if in == nil {
		return nil
	}
	out := new(QueueRetryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStackdriverLoggingConfig) DeepCopyInto(out *QueueStackdriverLoggingConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStackdriverLoggingConfig.
func (in *QueueStackdriverLoggingConfig) DeepCopy() *QueueStackdriverLoggingConfig {
	if in == nil {
		return nil
	}
	out := new(QueueStackdriverLoggingConfig)
	in.DeepCopyInto(out)
	return out
}