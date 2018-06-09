package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CustomHeader struct {
	Key string `json:""`
	Value string `json:""`
}

type CheckSpec struct {
	URL string `json:"url"`
	Period int `json:"period"`
	ApdexT float64 `json:"apdexT"`
	Enabled bool `json:"enabled"`
	Published bool `json:"published"`
	Alias string `json:"alias"`
	StringMatch string `json:"stringMatch"`
	MuteUntil string `json:"muteUntil"`
	DisabledLocations []string `json:"disabledLocations"`
	CustomHeaders []CustomHeader `json:"customHeaders"`
}

type Check struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CheckSpec `json:"spec"`
}

func (c *Check) LoadDefaults() {
	if c.Spec.Period == 0 {
		c.Spec.Period = 60
	}
	if c.Spec.ApdexT == 0 {
		c.Spec.ApdexT = 0.5
	}
}

type CheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Check `json:"items"`
}
