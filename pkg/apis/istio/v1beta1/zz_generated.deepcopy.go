// +build !ignore_autogenerated

/*
Copyright 2019 Banzai Cloud.

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
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CitadelConfiguration) DeepCopyInto(out *CitadelConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CitadelConfiguration.
func (in *CitadelConfiguration) DeepCopy() *CitadelConfiguration {
	if in == nil {
		return nil
	}
	out := new(CitadelConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GalleyConfiguration) DeepCopyInto(out *GalleyConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GalleyConfiguration.
func (in *GalleyConfiguration) DeepCopy() *GalleyConfiguration {
	if in == nil {
		return nil
	}
	out := new(GalleyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfiguration) DeepCopyInto(out *GatewayConfiguration) {
	*out = *in
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceLabels != nil {
		in, out := &in.ServiceLabels, &out.ServiceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.SDS = in.SDS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfiguration.
func (in *GatewayConfiguration) DeepCopy() *GatewayConfiguration {
	if in == nil {
		return nil
	}
	out := new(GatewayConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySDSConfiguration) DeepCopyInto(out *GatewaySDSConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySDSConfiguration.
func (in *GatewaySDSConfiguration) DeepCopy() *GatewaySDSConfiguration {
	if in == nil {
		return nil
	}
	out := new(GatewaySDSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaysConfiguration) DeepCopyInto(out *GatewaysConfiguration) {
	*out = *in
	in.IngressConfig.DeepCopyInto(&out.IngressConfig)
	in.EgressConfig.DeepCopyInto(&out.EgressConfig)
	out.K8sIngress = in.K8sIngress
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaysConfiguration.
func (in *GatewaysConfiguration) DeepCopy() *GatewaysConfiguration {
	if in == nil {
		return nil
	}
	out := new(GatewaysConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Istio) DeepCopyInto(out *Istio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Istio.
func (in *Istio) DeepCopy() *Istio {
	if in == nil {
		return nil
	}
	out := new(Istio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Istio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioList) DeepCopyInto(out *IstioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Istio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioList.
func (in *IstioList) DeepCopy() *IstioList {
	if in == nil {
		return nil
	}
	out := new(IstioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IstioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioService) DeepCopyInto(out *IstioService) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioService.
func (in *IstioService) DeepCopy() *IstioService {
	if in == nil {
		return nil
	}
	out := new(IstioService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioSpec) DeepCopyInto(out *IstioSpec) {
	*out = *in
	if in.AutoInjectionNamespaces != nil {
		in, out := &in.AutoInjectionNamespaces, &out.AutoInjectionNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.SDS = in.SDS
	out.Pilot = in.Pilot
	out.Citadel = in.Citadel
	out.Galley = in.Galley
	in.Gateways.DeepCopyInto(&out.Gateways)
	out.Mixer = in.Mixer
	out.SidecarInjector = in.SidecarInjector
	out.NodeAgent = in.NodeAgent
	out.Proxy = in.Proxy
	out.ProxyInit = in.ProxyInit
	out.DefaultPodDisruptionBudget = in.DefaultPodDisruptionBudget
	out.OutboundTrafficPolicy = in.OutboundTrafficPolicy
	out.Tracing = in.Tracing
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioSpec.
func (in *IstioSpec) DeepCopy() *IstioSpec {
	if in == nil {
		return nil
	}
	out := new(IstioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioStatus) DeepCopyInto(out *IstioStatus) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string][]*unstructured.Unstructured, len(*in))
		for key, val := range *in {
			var outVal []*unstructured.Unstructured
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]*unstructured.Unstructured, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = (*in).DeepCopy()
					}
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioStatus.
func (in *IstioStatus) DeepCopy() *IstioStatus {
	if in == nil {
		return nil
	}
	out := new(IstioStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sIngressConfiguration) DeepCopyInto(out *K8sIngressConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sIngressConfiguration.
func (in *K8sIngressConfiguration) DeepCopy() *K8sIngressConfiguration {
	if in == nil {
		return nil
	}
	out := new(K8sIngressConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MixerConfiguration) DeepCopyInto(out *MixerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MixerConfiguration.
func (in *MixerConfiguration) DeepCopy() *MixerConfiguration {
	if in == nil {
		return nil
	}
	out := new(MixerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAgentConfiguration) DeepCopyInto(out *NodeAgentConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAgentConfiguration.
func (in *NodeAgentConfiguration) DeepCopy() *NodeAgentConfiguration {
	if in == nil {
		return nil
	}
	out := new(NodeAgentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutboundTrafficPolicyConfiguration) DeepCopyInto(out *OutboundTrafficPolicyConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutboundTrafficPolicyConfiguration.
func (in *OutboundTrafficPolicyConfiguration) DeepCopy() *OutboundTrafficPolicyConfiguration {
	if in == nil {
		return nil
	}
	out := new(OutboundTrafficPolicyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PDBConfiguration) DeepCopyInto(out *PDBConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PDBConfiguration.
func (in *PDBConfiguration) DeepCopy() *PDBConfiguration {
	if in == nil {
		return nil
	}
	out := new(PDBConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PilotConfiguration) DeepCopyInto(out *PilotConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PilotConfiguration.
func (in *PilotConfiguration) DeepCopy() *PilotConfiguration {
	if in == nil {
		return nil
	}
	out := new(PilotConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfiguration) DeepCopyInto(out *ProxyConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfiguration.
func (in *ProxyConfiguration) DeepCopy() *ProxyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ProxyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyInitConfiguration) DeepCopyInto(out *ProxyInitConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyInitConfiguration.
func (in *ProxyInitConfiguration) DeepCopy() *ProxyInitConfiguration {
	if in == nil {
		return nil
	}
	out := new(ProxyInitConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteIstio) DeepCopyInto(out *RemoteIstio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteIstio.
func (in *RemoteIstio) DeepCopy() *RemoteIstio {
	if in == nil {
		return nil
	}
	out := new(RemoteIstio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteIstio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteIstioList) DeepCopyInto(out *RemoteIstioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RemoteIstio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteIstioList.
func (in *RemoteIstioList) DeepCopy() *RemoteIstioList {
	if in == nil {
		return nil
	}
	out := new(RemoteIstioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteIstioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteIstioSpec) DeepCopyInto(out *RemoteIstioSpec) {
	*out = *in
	if in.EnabledServices != nil {
		in, out := &in.EnabledServices, &out.EnabledServices
		*out = make([]IstioService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoInjectionNamespaces != nil {
		in, out := &in.AutoInjectionNamespaces, &out.AutoInjectionNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Citadel = in.Citadel
	out.SidecarInjector = in.SidecarInjector
	out.Proxy = in.Proxy
	out.ProxyInit = in.ProxyInit
	in.signCert.DeepCopyInto(&out.signCert)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteIstioSpec.
func (in *RemoteIstioSpec) DeepCopy() *RemoteIstioSpec {
	if in == nil {
		return nil
	}
	out := new(RemoteIstioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteIstioStatus) DeepCopyInto(out *RemoteIstioStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteIstioStatus.
func (in *RemoteIstioStatus) DeepCopy() *RemoteIstioStatus {
	if in == nil {
		return nil
	}
	out := new(RemoteIstioStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SDSConfiguration) DeepCopyInto(out *SDSConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SDSConfiguration.
func (in *SDSConfiguration) DeepCopy() *SDSConfiguration {
	if in == nil {
		return nil
	}
	out := new(SDSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarInjectorConfiguration) DeepCopyInto(out *SidecarInjectorConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarInjectorConfiguration.
func (in *SidecarInjectorConfiguration) DeepCopy() *SidecarInjectorConfiguration {
	if in == nil {
		return nil
	}
	out := new(SidecarInjectorConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignCert) DeepCopyInto(out *SignCert) {
	*out = *in
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Root != nil {
		in, out := &in.Root, &out.Root
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Chain != nil {
		in, out := &in.Chain, &out.Chain
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignCert.
func (in *SignCert) DeepCopy() *SignCert {
	if in == nil {
		return nil
	}
	out := new(SignCert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfiguration) DeepCopyInto(out *TracingConfiguration) {
	*out = *in
	out.Zipkin = in.Zipkin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfiguration.
func (in *TracingConfiguration) DeepCopy() *TracingConfiguration {
	if in == nil {
		return nil
	}
	out := new(TracingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZipkinConfiguration) DeepCopyInto(out *ZipkinConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZipkinConfiguration.
func (in *ZipkinConfiguration) DeepCopy() *ZipkinConfiguration {
	if in == nil {
		return nil
	}
	out := new(ZipkinConfiguration)
	in.DeepCopyInto(out)
	return out
}
