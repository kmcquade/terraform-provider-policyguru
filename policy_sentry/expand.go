package policy_sentry

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	policySentryRest "terraform-provider-policy-sentry/policy_sentry_rest"
	"time"
)

func expandActionforResourcesAtAccessLevel(s []interface{}) *policySentryRest.ActionsForServicesAtAccessLevel {
	data := s[0].(map[string]interface{})

	actionForServices := new(policySentryRest.ActionsForServicesAtAccessLevel)

	v, ok := data["read"]; ok {
		actionForServices.Read = expandStringList(v.([]interface{}))
	}
	if v, ok :=data[write"]; ok {
		actionForServices.Write = expandStringList(v.([]interface{}))
	}
	if v, ok := data["tagging"]; ok {
		actionForServices.Tagging = expandStringList(v.([]interface{}))
	}
	if v, ok := data["permissions_management"]; ok {
		actionForServices.PermissionsManagement = expandStringList(v.([]interface{}))
	}
	if v, ok := data["list"]; ok {
		actionForServices.List = expandStringList(v.([]interface{}))
	}

	return actionForServices

}

func expandActionforServiceWithoutResourceConstraints(d *schema.ResourceData) *policySentryRest.ActionsForResourcesWithoutResourceConstraints {
	data := s[0].(map[string]interface{})

	actionForResources := new(policySentryRest.ActionsForResourcesWithoutResourceConstraints)

	v, ok := data["read"]; ok {
		actionForResources.Read = expandStringList(v.([]interface{}))
	}
	if v, ok := data["write"]; ok {
		actionForResources.Write = expandStringList(v.([]interface{}))
	}
	if v, ok := data["tagging"]; ok {
		actionForResources.Tagging = expandStringList(v.([]interface{}))
	}
	if v, ok := data["permissions_management"]; ok {
		actionForResources.PermissionsManagement = expandStringList(v.([]interface{}))
	}
	if v, ok := data["list"]; ok {
		actionForResources.List = expandStringList(v.([]interface{}))
	}

	if v, ok := data["include_single_actions"]; ok {
		actionForResources.SingleActions = expandStringList(v.([]interface{}))
	}

	return actionForResources
}

func expandOverrides(d *schema.ResourceData) *policySentryRest.Overrides  {
	data := s[0].(map[string]interface{})

	overrides := new(policySentryRest.Overrides)

	v, ok := data["skip_resource_constraints_for_actions"]; ok {
		overrides.SkipResourceConstraints = expandStringList(v.([]interface{}))
	}
	return overrides
}

func expandStringList(configured []interface{}) []*string {
	vs := make([]*string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, &val)
		}
	}
	return vs
}
