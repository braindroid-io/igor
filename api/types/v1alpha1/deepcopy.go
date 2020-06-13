package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

func (in *WebSite) DeepCopyInto(out *WebSite) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = WebSiteSpec{
		Replicas: in.Spec.Replicas,
		Image: in.Spec.Image,
		Template: in.Spec.Template,
	}
}


func (in *WebSite) DeepCopyObject() runtime.Object {
	out := WebSite{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *WebSiteList) DeepCopyObject() runtime.Object {
	out := WebSiteList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]WebSite, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
