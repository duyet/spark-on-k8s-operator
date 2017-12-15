// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ApplicationState).DeepCopyInto(out.(*ApplicationState))
			return nil
		}, InType: reflect.TypeOf(&ApplicationState{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Dependencies).DeepCopyInto(out.(*Dependencies))
			return nil
		}, InType: reflect.TypeOf(&Dependencies{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DriverInfo).DeepCopyInto(out.(*DriverInfo))
			return nil
		}, InType: reflect.TypeOf(&DriverInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DriverSpec).DeepCopyInto(out.(*DriverSpec))
			return nil
		}, InType: reflect.TypeOf(&DriverSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExecutorSpec).DeepCopyInto(out.(*ExecutorSpec))
			return nil
		}, InType: reflect.TypeOf(&ExecutorSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamePath).DeepCopyInto(out.(*NamePath))
			return nil
		}, InType: reflect.TypeOf(&NamePath{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SecretInfo).DeepCopyInto(out.(*SecretInfo))
			return nil
		}, InType: reflect.TypeOf(&SecretInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SparkApplication).DeepCopyInto(out.(*SparkApplication))
			return nil
		}, InType: reflect.TypeOf(&SparkApplication{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SparkApplicationList).DeepCopyInto(out.(*SparkApplicationList))
			return nil
		}, InType: reflect.TypeOf(&SparkApplicationList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SparkApplicationSpec).DeepCopyInto(out.(*SparkApplicationSpec))
			return nil
		}, InType: reflect.TypeOf(&SparkApplicationSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SparkApplicationStatus).DeepCopyInto(out.(*SparkApplicationStatus))
			return nil
		}, InType: reflect.TypeOf(&SparkApplicationStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationState) DeepCopyInto(out *ApplicationState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationState.
func (in *ApplicationState) DeepCopy() *ApplicationState {
	if in == nil {
		return nil
	}
	out := new(ApplicationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependencies) DeepCopyInto(out *Dependencies) {
	*out = *in
	if in.JarFiles != nil {
		in, out := &in.JarFiles, &out.JarFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PyFiles != nil {
		in, out := &in.PyFiles, &out.PyFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependencies.
func (in *Dependencies) DeepCopy() *Dependencies {
	if in == nil {
		return nil
	}
	out := new(Dependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverInfo) DeepCopyInto(out *DriverInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverInfo.
func (in *DriverInfo) DeepCopy() *DriverInfo {
	if in == nil {
		return nil
	}
	out := new(DriverInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverSpec) DeepCopyInto(out *DriverSpec) {
	*out = *in
	if in.PodName != nil {
		in, out := &in.PodName, &out.PodName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.DriverConfigMaps != nil {
		in, out := &in.DriverConfigMaps, &out.DriverConfigMaps
		*out = make([]NamePath, len(*in))
		copy(*out, *in)
	}
	if in.DriverSecrets != nil {
		in, out := &in.DriverSecrets, &out.DriverSecrets
		*out = make([]SecretInfo, len(*in))
		copy(*out, *in)
	}
	if in.DriverEnvVars != nil {
		in, out := &in.DriverEnvVars, &out.DriverEnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverSpec.
func (in *DriverSpec) DeepCopy() *DriverSpec {
	if in == nil {
		return nil
	}
	out := new(DriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutorSpec) DeepCopyInto(out *ExecutorSpec) {
	*out = *in
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.ExecutorConfigMaps != nil {
		in, out := &in.ExecutorConfigMaps, &out.ExecutorConfigMaps
		*out = make([]NamePath, len(*in))
		copy(*out, *in)
	}
	if in.ExecutorSecrets != nil {
		in, out := &in.ExecutorSecrets, &out.ExecutorSecrets
		*out = make([]SecretInfo, len(*in))
		copy(*out, *in)
	}
	if in.ExecutorEnvVars != nil {
		in, out := &in.ExecutorEnvVars, &out.ExecutorEnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutorSpec.
func (in *ExecutorSpec) DeepCopy() *ExecutorSpec {
	if in == nil {
		return nil
	}
	out := new(ExecutorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamePath) DeepCopyInto(out *NamePath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamePath.
func (in *NamePath) DeepCopy() *NamePath {
	if in == nil {
		return nil
	}
	out := new(NamePath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretInfo) DeepCopyInto(out *SecretInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretInfo.
func (in *SecretInfo) DeepCopy() *SecretInfo {
	if in == nil {
		return nil
	}
	out := new(SecretInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkApplication) DeepCopyInto(out *SparkApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkApplication.
func (in *SparkApplication) DeepCopy() *SparkApplication {
	if in == nil {
		return nil
	}
	out := new(SparkApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkApplicationList) DeepCopyInto(out *SparkApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SparkApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkApplicationList.
func (in *SparkApplicationList) DeepCopy() *SparkApplicationList {
	if in == nil {
		return nil
	}
	out := new(SparkApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkApplicationSpec) DeepCopyInto(out *SparkApplicationSpec) {
	*out = *in
	if in.MainClass != nil {
		in, out := &in.MainClass, &out.MainClass
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SparkConf != nil {
		in, out := &in.SparkConf, &out.SparkConf
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HadoopConf != nil {
		in, out := &in.HadoopConf, &out.HadoopConf
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SparkConfigMap != nil {
		in, out := &in.SparkConfigMap, &out.SparkConfigMap
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.HadoopConfigMap != nil {
		in, out := &in.HadoopConfigMap, &out.HadoopConfigMap
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	in.Driver.DeepCopyInto(&out.Driver)
	in.Executor.DeepCopyInto(&out.Executor)
	in.Deps.DeepCopyInto(&out.Deps)
	if in.LogsLocation != nil {
		in, out := &in.LogsLocation, &out.LogsLocation
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkApplicationSpec.
func (in *SparkApplicationSpec) DeepCopy() *SparkApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(SparkApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkApplicationStatus) DeepCopyInto(out *SparkApplicationStatus) {
	*out = *in
	out.DriverInfo = in.DriverInfo
	out.AppState = in.AppState
	if in.ExecutorState != nil {
		in, out := &in.ExecutorState, &out.ExecutorState
		*out = make(map[string]ExecutorState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkApplicationStatus.
func (in *SparkApplicationStatus) DeepCopy() *SparkApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(SparkApplicationStatus)
	in.DeepCopyInto(out)
	return out
}
