package dnsimple

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"domain_token"},
				DefaultFunc:   schema.EnvDefaultFunc("DNSIMPLE_EMAIL", nil),
				Description:   "A registered DNSimple email address.",
			},

			"token": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"domain_token"},
				DefaultFunc:   schema.EnvDefaultFunc("DNSIMPLE_TOKEN", nil),
				Description:   "The token key for API operations.",
			},

			"domain_token": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"email", "token"},
				DefaultFunc:   schema.EnvDefaultFunc("DNSIMPLE_DOMAIN_TOKEN", nil),
				Description:   "The domain token key for API operations.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"dnsimple_record": resourceDNSimpleRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Email:       d.Get("email").(string),
		Token:       d.Get("token").(string),
		DomainToken: d.Get("domain_token").(string),
	}

	return config.Client()
}
