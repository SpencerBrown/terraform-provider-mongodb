package main

import (
	"github.com/SpencerBrown/terraform-provider-mongodb/mongodb"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return mongodb.Provider()
		},
	})
}
