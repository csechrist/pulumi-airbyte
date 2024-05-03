package airbyte

import (
	"context"
	"errors"
	"strings"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func makeName(name string) string {
	newValue := strings.ReplaceAll(name, "airbyte_", "")
	newValue = strings.ReplaceAll(newValue, "_", " ")
	caser := cases.Title(language.English)
	newValue = caser.String(newValue)

	allvalues := strings.Split(newValue, " ")

	return strings.Join(allvalues, "")
}

func computeId(ctx context.Context, state resource.PropertyMap) (resource.ID, error) {
	if id, has := state["id"]; has {
		return resource.ID(id.StringValue()), nil
	}

	if id, has := state["source_id"]; has {
		return resource.ID(id.StringValue()), nil
	}

	if id, has := state["destination_id"]; has {
		return resource.ID(id.StringValue()), nil
	}

	if id, has := state["connection_id"]; has {
		return resource.ID(id.StringValue()), nil
	}

	if id, has := state["workspace_id"]; has {
		return resource.ID(id.StringValue()), nil
	}

	return resource.ID("no"), errors.New("computeId not implemented")
}

func airbyteResourceType(name string) *tfbridge.ResourceInfo {

	newName := makeName(name)

	return &tfbridge.ResourceInfo{
		Tok:       tfbridge.MakeResource(mainPkg, mainMod, newName),
		ComputeID: computeId,
	}
}
