package shim

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/provider"
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
)

// fix provider.Provider here to match whats in internal/provider
func New() tfpf.Provider {
	return provider.New("dev")()
}
