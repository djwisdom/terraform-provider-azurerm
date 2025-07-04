// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pluginsdk

import (
	"context"
	"log"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type IDValidationFunc func(id string) error

type ImporterFunc = func(ctx context.Context, d *ResourceData, meta interface{}) ([]*ResourceData, error)

// ImporterValidatingResourceId validates the ID provided at import time is valid
// using the validateFunc.
func ImporterValidatingResourceId(validateFunc IDValidationFunc) *schema.ResourceImporter {
	thenFunc := func(ctx context.Context, d *ResourceData, meta interface{}) ([]*ResourceData, error) {
		return []*ResourceData{d}, nil
	}
	return ImporterValidatingResourceIdThen(validateFunc, thenFunc)
}

// ImporterValidatingResourceIdThen validates the ID provided at import time is valid
// using the validateFunc then runs the 'thenFunc', allowing the import to be customised.
func ImporterValidatingResourceIdThen(validateFunc IDValidationFunc, thenFunc ImporterFunc) *schema.ResourceImporter {
	return &schema.ResourceImporter{
		StateContext: func(ctx context.Context, d *ResourceData, meta interface{}) ([]*ResourceData, error) {
			log.Printf("[DEBUG] Importing Resource - parsing %q", d.Id())

			if _, ok := ctx.Deadline(); !ok {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, d.Timeout(schema.TimeoutRead))
				defer cancel()
			}

			if err := validateFunc(d.Id()); err != nil {
				// NOTE: we're intentionally not wrapping this error, since it's prefixed with `parsing %q:`
				return []*ResourceData{d}, err
			}

			return thenFunc(ctx, d, meta)
		},
	}
}

// ImporterValidatingIdentity validates the ID provided at import time is valid or that the resource identity data provided in the import block is valid
// based on the expected resource ID type.
func ImporterValidatingIdentity(id resourceids.ResourceId, idType ...ResourceTypeForIdentity) *schema.ResourceImporter {
	thenFunc := func(ctx context.Context, d *ResourceData, meta interface{}) ([]*ResourceData, error) {
		return []*ResourceData{d}, nil
	}

	return ImporterValidatingIdentityThen(id, thenFunc, idType...)
}

// ImporterValidatingIdentityThen validates the ID provided at import time is valid or that the resource identity data provided in the import block is valid
// based on the expected resource ID type, then runs the 'thenFunc', allowing the import to be customised.
func ImporterValidatingIdentityThen(id resourceids.ResourceId, thenFunc ImporterFunc, idType ...ResourceTypeForIdentity) *schema.ResourceImporter {
	return &schema.ResourceImporter{
		StateContext: func(ctx context.Context, d *ResourceData, meta interface{}) ([]*ResourceData, error) {
			log.Printf("[DEBUG] Importing Resource - parsing %q", d.Id())

			if _, ok := ctx.Deadline(); !ok {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, d.Timeout(schema.TimeoutRead))
				defer cancel()
			}

			if d.Id() != "" {
				parser := resourceids.NewParserFromResourceIdType(id)
				if _, err := parser.Parse(d.Id(), false); err != nil {
					// NOTE: we're intentionally not wrapping this error, since it's prefixed with `parsing %q:`
					return []*ResourceData{d}, err
				}
				return thenFunc(ctx, d, meta)
			}

			if err := ValidateResourceIdentityData(d, id, idType...); err != nil {
				return nil, err
			}

			return thenFunc(ctx, d, meta)
		},
	}
}
