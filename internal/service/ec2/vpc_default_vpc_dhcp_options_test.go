package ec2_test

import (
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)

func TestAccVPCDefaultVPCDHCPOptions_serial(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"basic": testAccDefaultVPCDHCPOptions_basic,
		"owner": testAccDefaultVPCDHCPOptions_owner,
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccDefaultVPCDHCPOptions_basic(t *testing.T) {
	var d ec2.DhcpOptions
	resourceName := "aws_default_vpc_dhcp_options.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, ec2.EndpointsID),
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      acctest.CheckDestroyNoop,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCDefaultVPCDHCPOptionsConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDHCPOptionsExists(resourceName, &d),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`dhcp-options/dopt-.+`)),
					resource.TestCheckResourceAttr(resourceName, "domain_name", tfec2.RegionalPrivateDNSSuffix(acctest.Region())),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "AmazonProvidedDNS"),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", "Default DHCP Option Set"),
				),
			},
		},
	})
}

func testAccDefaultVPCDHCPOptions_owner(t *testing.T) {
	var d ec2.DhcpOptions
	resourceName := "aws_default_vpc_dhcp_options.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, ec2.EndpointsID),
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      acctest.CheckDestroyNoop,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCDefaultVPCDHCPOptionsConfig_owner,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDHCPOptionsExists(resourceName, &d),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`dhcp-options/dopt-.+`)),
					resource.TestCheckResourceAttr(resourceName, "domain_name", tfec2.RegionalPrivateDNSSuffix(acctest.Region())),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "AmazonProvidedDNS"),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", "Default DHCP Option Set"),
				),
			},
		},
	})
}

const testAccVPCDefaultVPCDHCPOptionsConfig_basic = `
resource "aws_default_vpc_dhcp_options" "test" {
  tags = {
    Name = "Default DHCP Option Set"
  }
}
`

const testAccVPCDefaultVPCDHCPOptionsConfig_owner = `
data "aws_caller_identity" "current" {}

resource "aws_default_vpc_dhcp_options" "test" {
  owner_id = data.aws_caller_identity.current.account_id

  tags = {
    Name = "Default DHCP Option Set"
  }
}
`
