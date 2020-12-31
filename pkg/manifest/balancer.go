package manifest

import (
	"strings"
)

type Balancer struct {
	Name string `yaml:"-"`

	Ports       BalancerPorts       `yaml:"ports,omitempty"`
	Service     string              `yaml:"service,omitempty"`
	Whitelist   BalancerWhitelist   `yaml:"whitelist,omitempty"`
	Annotations BalancerAnnotations `yaml:"annotations,omitempty"`
}

type Balancers []Balancer

type BalancerPort struct {
	Source int `yaml:"-"`

	Protocol string `yaml:"protocol,omitempty"`
	Target   int    `yaml:"port,omitempty"`
}

type BalancerPorts []BalancerPort

type BalancerWhitelist []string

type BalancerAnnotations []string

func (b Balancer) AnnotationsMap() map[string]string {
	annotations := map[string]string{}

	for _, a := range b.Annotations {
		parts := strings.SplitN(a, "=", 2)
		annotations[parts[0]] = parts[1]
	}

	return annotations
}
