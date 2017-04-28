package vault

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func mountResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Path of the mount",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the mount",
			},
			"mount_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of mount",
			},
		},
	}
}
