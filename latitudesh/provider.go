package latitudesh

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api "github.com/latitudesh/latitudesh-go"
)

const (
	userAgentForProvider = "Latitude-Terraform-Provider"
)

var currentVersion = "1.0.0-rc.1" // update varible when version updated

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LATITUDESH_AUTH_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"latitudesh_project":         resourceProject(),
			"latitudesh_server":          resourceServer(),
			"latitudesh_ssh_key":         resourceSSHKey(),
			"latitudesh_user_data":       resourceUserData(),
			"latitudesh_virtual_network": resourceVirtualNetwork(),
			"latitudesh_vlan_assignment": resourceVlanAssignment(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"latitudesh_plan":   dataSourcePlan(),
			"latitudesh_region": dataSourceRegion(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	authToken := d.Get("auth_token").(string)

	var diags diag.Diagnostics

	if authToken != "" {
		c := api.NewClientWithAuth("latitudesh", authToken, nil)
		c.UserAgent = fmt.Sprintf("%s/%s", userAgentForProvider, currentVersion)

		return c, diags
	}
	c := api.NewClientWithAuth("latitudesh", " ", nil)

	return c, diags
}
