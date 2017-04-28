package vault

import (
	"fmt"
	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/vault/api"
	"testing"
)

func TestResourceMount(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers: testProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []r.TestStep{
			r.TestStep{
				Config: testResourceMount_initialConfig,
				Check:  testResourceMount_initialCheck,
			},
		},
	})
}

var testResourceMount_initialConfig = `
resource "vault_mount" "secrets" {
	path        = "secrets"
	description = "a place for secrets"
	mount_type  = "generic"
}
`

func testResourceMount_initialCheck(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["vault_mount.secrets"]
	if resourceState == nil {
		return fmt.Errorf("resource not found in state")
	}

	instanceState := resourceState.Primary
	if instanceState == nil {
		return fmt.Errorf("resource has no primary instance")
	}

	path := instanceState.ID
	if path != instanceState.Attributes["path"] {
		return fmt.Errorf("id doesn't match path")
	}

	client := testProvider.Meta().(*api.Client)
	mounts, err := client.Sys().ListMounts()
	if err != nil {
		return fmt.Errorf("error listing mounts: %s", err)
	}

	if got, want := mounts[path].Description, "a place for secrets"; got != want {
		return fmt.Errorf("mount description is %q; want %q", got, want)
	}

	return nil
}
