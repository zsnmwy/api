// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	autoscalingv1alpha1 "github.com/gocrane/api/autoscaling/v1alpha1"
	v2beta2 "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Analytics) DeepCopyInto(out *Analytics) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Analytics.
func (in *Analytics) DeepCopy() *Analytics {
	if in == nil {
		return nil
	}
	out := new(Analytics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Analytics) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsList) DeepCopyInto(out *AnalyticsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Analytics, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsList.
func (in *AnalyticsList) DeepCopy() *AnalyticsList {
	if in == nil {
		return nil
	}
	out := new(AnalyticsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyticsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsSpec) DeepCopyInto(out *AnalyticsSpec) {
	*out = *in
	if in.ResourceSelectors != nil {
		in, out := &in.ResourceSelectors, &out.ResourceSelectors
		*out = make([]ResourceSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CompletionStrategy.DeepCopyInto(&out.CompletionStrategy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsSpec.
func (in *AnalyticsSpec) DeepCopy() *AnalyticsSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyticsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsStatus) DeepCopyInto(out *AnalyticsStatus) {
	*out = *in
	if in.LastSuccessfulTime != nil {
		in, out := &in.LastSuccessfulTime, &out.LastSuccessfulTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Recommendations != nil {
		in, out := &in.Recommendations, &out.Recommendations
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsStatus.
func (in *AnalyticsStatus) DeepCopy() *AnalyticsStatus {
	if in == nil {
		return nil
	}
	out := new(AnalyticsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionStrategy) DeepCopyInto(out *CompletionStrategy) {
	*out = *in
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionStrategy.
func (in *CompletionStrategy) DeepCopy() *CompletionStrategy {
	if in == nil {
		return nil
	}
	out := new(CompletionStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRecommendation) DeepCopyInto(out *ContainerRecommendation) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRecommendation.
func (in *ContainerRecommendation) DeepCopy() *ContainerRecommendation {
	if in == nil {
		return nil
	}
	out := new(ContainerRecommendation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EffectiveHorizontalPodAutoscalerRecommendation) DeepCopyInto(out *EffectiveHorizontalPodAutoscalerRecommendation) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Prediction != nil {
		in, out := &in.Prediction, &out.Prediction
		*out = new(autoscalingv1alpha1.Prediction)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EffectiveHorizontalPodAutoscalerRecommendation.
func (in *EffectiveHorizontalPodAutoscalerRecommendation) DeepCopy() *EffectiveHorizontalPodAutoscalerRecommendation {
	if in == nil {
		return nil
	}
	out := new(EffectiveHorizontalPodAutoscalerRecommendation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Recommendation) DeepCopyInto(out *Recommendation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recommendation.
func (in *Recommendation) DeepCopy() *Recommendation {
	if in == nil {
		return nil
	}
	out := new(Recommendation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Recommendation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationList) DeepCopyInto(out *RecommendationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Recommendation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationList.
func (in *RecommendationList) DeepCopy() *RecommendationList {
	if in == nil {
		return nil
	}
	out := new(RecommendationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecommendationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationSpec) DeepCopyInto(out *RecommendationSpec) {
	*out = *in
	out.TargetRef = in.TargetRef
	in.CompletionStrategy.DeepCopyInto(&out.CompletionStrategy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationSpec.
func (in *RecommendationSpec) DeepCopy() *RecommendationSpec {
	if in == nil {
		return nil
	}
	out := new(RecommendationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationStatus) DeepCopyInto(out *RecommendationStatus) {
	*out = *in
	if in.EffectiveHPA != nil {
		in, out := &in.EffectiveHPA, &out.EffectiveHPA
		*out = new(EffectiveHorizontalPodAutoscalerRecommendation)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceRequest != nil {
		in, out := &in.ResourceRequest, &out.ResourceRequest
		*out = new(ResourceRequestRecommendation)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	if in.LastSuccessfulTime != nil {
		in, out := &in.LastSuccessfulTime, &out.LastSuccessfulTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationStatus.
func (in *RecommendationStatus) DeepCopy() *RecommendationStatus {
	if in == nil {
		return nil
	}
	out := new(RecommendationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequestRecommendation) DeepCopyInto(out *ResourceRequestRecommendation) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerRecommendation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequestRecommendation.
func (in *ResourceRequestRecommendation) DeepCopy() *ResourceRequestRecommendation {
	if in == nil {
		return nil
	}
	out := new(ResourceRequestRecommendation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSelector) DeepCopyInto(out *ResourceSelector) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSelector.
func (in *ResourceSelector) DeepCopy() *ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceSelector)
	in.DeepCopyInto(out)
	return out
}
