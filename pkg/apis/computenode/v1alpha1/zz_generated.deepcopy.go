// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeOpenStack) DeepCopyInto(out *ComputeNodeOpenStack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeOpenStack.
func (in *ComputeNodeOpenStack) DeepCopy() *ComputeNodeOpenStack {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeOpenStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeNodeOpenStack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeOpenStackList) DeepCopyInto(out *ComputeNodeOpenStackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputeNodeOpenStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeOpenStackList.
func (in *ComputeNodeOpenStackList) DeepCopy() *ComputeNodeOpenStackList {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeOpenStackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeNodeOpenStackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeOpenStackSpec) DeepCopyInto(out *ComputeNodeOpenStackSpec) {
	*out = *in
	if in.InfraDaemonSets != nil {
		in, out := &in.InfraDaemonSets, &out.InfraDaemonSets
		*out = make([]InfraDaemonSet, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeOpenStackSpec.
func (in *ComputeNodeOpenStackSpec) DeepCopy() *ComputeNodeOpenStackSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeOpenStackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeOpenStackStatus) DeepCopyInto(out *ComputeNodeOpenStackStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeOpenStackStatus.
func (in *ComputeNodeOpenStackStatus) DeepCopy() *ComputeNodeOpenStackStatus {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeOpenStackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraDaemonSet) DeepCopyInto(out *InfraDaemonSet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraDaemonSet.
func (in *InfraDaemonSet) DeepCopy() *InfraDaemonSet {
	if in == nil {
		return nil
	}
	out := new(InfraDaemonSet)
	in.DeepCopyInto(out)
	return out
}
