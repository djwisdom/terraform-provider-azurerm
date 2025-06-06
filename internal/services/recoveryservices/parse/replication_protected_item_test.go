// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = ReplicationProtectedItemId{}

func TestReplicationProtectedItemIDFormatter(t *testing.T) {
	actual := NewReplicationProtectedItemID("12345678-1234-9876-4563-123456789012", "group1", "vault1", "fabric1", "container1", "item1").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric1/replicationProtectionContainers/container1/replicationProtectedItems/item1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestReplicationProtectedItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReplicationProtectedItemId
	}{
		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/",
			Error: true,
		},

		{
			// missing VaultName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/",
			Error: true,
		},

		{
			// missing value for VaultName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/",
			Error: true,
		},

		{
			// missing ReplicationFabricName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/",
			Error: true,
		},

		{
			// missing value for ReplicationFabricName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/",
			Error: true,
		},

		{
			// missing ReplicationProtectionContainerName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric1/",
			Error: true,
		},

		{
			// missing value for ReplicationProtectionContainerName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric1/replicationProtectionContainers/",
			Error: true,
		},

		{
			// missing Name
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric1/replicationProtectionContainers/container1/",
			Error: true,
		},

		{
			// missing value for Name
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric1/replicationProtectionContainers/container1/replicationProtectedItems/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric1/replicationProtectionContainers/container1/replicationProtectedItems/item1",
			Expected: &ReplicationProtectedItemId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroup:                      "group1",
				VaultName:                          "vault1",
				ReplicationFabricName:              "fabric1",
				ReplicationProtectionContainerName: "container1",
				Name:                               "item1",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/RESOURCEGROUPS/GROUP1/PROVIDERS/MICROSOFT.RECOVERYSERVICES/VAULTS/VAULT1/REPLICATIONFABRICS/FABRIC1/REPLICATIONPROTECTIONCONTAINERS/CONTAINER1/REPLICATIONPROTECTEDITEMS/ITEM1",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ReplicationProtectedItemID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.VaultName != v.Expected.VaultName {
			t.Fatalf("Expected %q but got %q for VaultName", v.Expected.VaultName, actual.VaultName)
		}
		if actual.ReplicationFabricName != v.Expected.ReplicationFabricName {
			t.Fatalf("Expected %q but got %q for ReplicationFabricName", v.Expected.ReplicationFabricName, actual.ReplicationFabricName)
		}
		if actual.ReplicationProtectionContainerName != v.Expected.ReplicationProtectionContainerName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectionContainerName", v.Expected.ReplicationProtectionContainerName, actual.ReplicationProtectionContainerName)
		}
		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
	}
}
