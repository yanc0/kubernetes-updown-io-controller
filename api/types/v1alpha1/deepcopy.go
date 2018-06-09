package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)


func (in *Check) DeepCopyInto(out *Check) {
	inDisabledLocations := make([]string, 0)
	for _, location := range in.Spec.DisabledLocations {
		inDisabledLocations = append(inDisabledLocations, location)
	}

	inCustomHeaders := make([]CustomHeader, 0)
	for _, customHeader := range inCustomHeaders {
		c := CustomHeader {
			Key: customHeader.Key,
			Value: customHeader.Value,
		}
		inCustomHeaders = append(inCustomHeaders, c)
	}

	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = CheckSpec{
		URL: in.Spec.URL,
		Period: in.Spec.Period,
		ApdexT: in.Spec.ApdexT,
		Enabled: in.Spec.Enabled,
		Published: in.Spec.Published,
		Alias: in.Spec.Alias,
		StringMatch: in.Spec.StringMatch,
		MuteUntil: in.Spec.MuteUntil,
		DisabledLocations: inDisabledLocations,
		CustomHeaders: inCustomHeaders,
	}
}

func (in *Check) DeepCopyObject() runtime.Object {
	out := Check{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *CheckList) DeepCopyObject() runtime.Object {
	out := CheckList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Check, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
