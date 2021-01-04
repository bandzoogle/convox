package manifest

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

type BalancerAnnotations map[string]string

func (b Balancer) AnnotationsMap() map[string]string {
	annotations := map[string]string{}

	for k, value := range b.Annotations {
		annotations[k] = value
	}

	// add idle timeout and LB type if not specified already
	addIfMissing(annotations, "service.beta.kubernetes.io/aws-load-balancer-connection-idle-timeout", "3600")
	addIfMissing(annotations, "service.beta.kubernetes.io/aws-load-balancer-type", "nlb")

	return annotations
}

func addIfMissing(m map[string]string, k string, v string) {
	_, exists := m[k]
	if !exists {
		m[k] = v
	}
}
