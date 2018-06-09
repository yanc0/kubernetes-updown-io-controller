package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"


func (in *Check) DeepCopyInto(out *Check) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = CheckSpec{
		URL: in.Spec.URL,
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
