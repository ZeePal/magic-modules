package networksecurity_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigFullExample_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigFullExample_full(context),
			},
			{
				ResourceName:            "google_network_security_backend_authentication_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"terraform_labels", "labels"},
			},
			{
				Config: testAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigFullExample_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_network_security_backend_authentication_config.default", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_network_security_backend_authentication_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"terraform_labels", "labels"},
			},
		},
	})
}

func testAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigFullExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_certificate_manager_certificate" "certificate" {
  name     = "tf-test-my-certificate%{random_suffix}"
  location    = "global"
  self_managed {
    pem_certificate = file("test-fixtures/cert.pem")
    pem_private_key = file("test-fixtures/key.pem")
  }
  scope       = "CLIENT_AUTH"
}

resource "google_certificate_manager_trust_config" "trust_config" {
  name        = "tf-test-my-trust-config%{random_suffix}"
  description = "sample description for the trust config"
  location    = "global"

  trust_stores {
    trust_anchors { 
      pem_certificate = file("test-fixtures/cert.pem")
    }
    intermediate_cas { 
      pem_certificate = file("test-fixtures/cert.pem")
    }
  }
}

resource "google_network_security_backend_authentication_config" "default" {
  name               = "tf-test-my-backend-authentication-config%{random_suffix}"
  location           = "global"
  description        = "my description"
  well_known_roots   = "PUBLIC_ROOTS"
  client_certificate = google_certificate_manager_certificate.certificate.id
  trust_config       = google_certificate_manager_trust_config.trust_config.id
}
`, context)
}

func testAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigFullExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_certificate_manager_certificate" "certificate" {
  name     = "tf-test-my-certificate%{random_suffix}"
  location    = "global"
  self_managed {
    pem_certificate = file("test-fixtures/cert.pem")
    pem_private_key = file("test-fixtures/key.pem")
  }
  scope       = "CLIENT_AUTH"
}

resource "google_certificate_manager_trust_config" "trust_config" {
  name        = "tf-test-my-trust-config%{random_suffix}"
  description = "sample description for the trust config"
  location    = "global"

  trust_stores {
    trust_anchors { 
      pem_certificate = file("test-fixtures/cert.pem")
    }
    intermediate_cas { 
      pem_certificate = file("test-fixtures/cert.pem")
    }
  }
}

resource "google_network_security_backend_authentication_config" "default" {
  name     = "tf-test-my-backend-authentication-config%{random_suffix}"
  location           = "global"
  description        = "updated description"
  well_known_roots   = "NONE"
  client_certificate = google_certificate_manager_certificate.certificate.id
  trust_config       = google_certificate_manager_trust_config.trust_config.id
}
`, context)
}
