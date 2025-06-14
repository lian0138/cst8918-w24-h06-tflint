package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "Removed ID"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Override the default terraform variables
		Vars: map[string]interface{}{
			"label_prefix": "lian0138",
		},
		// Increase retry attempts for Azure API operations
		MaxRetries:         3,
		TimeBetweenRetries: 10 * time.Second,
	}

	// Ensure cleanup happens even if test fails
	defer func() {
		// Check if resource group exists before attempting destroy
		resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
		if azure.ResourceGroupExists(t, resourceGroupName, subscriptionID) {
			t.Logf("Resource group %s exists, attempting to destroy", resourceGroupName)
			for i := 0; i < 3; i++ {
				_, err := terraform.DestroyE(t, terraformOptions)
				if err == nil {
					t.Logf("Successfully destroyed resources")
					break
				}
				t.Logf("Destroy attempt %d failed: %v", i+1, err)
				time.Sleep(10 * time.Second) // Wait before retrying
			}
		} else {
			t.Logf("Resource group %s does not exist, skipping destroy", resourceGroupName)
		}
	}()

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of output variable
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	nicName := terraform.Output(t, terraformOptions, "nic_name")

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID), "VM %s does not exist", vmName)

	// Confirm NIC exists and is connected to VM
	actualNicNames := azure.GetVirtualMachineNics(t, vmName, resourceGroupName, subscriptionID)
	assert.Equal(t, nicName, actualNicNames[0], "NIC %s is not attached to VM %s", nicName, vmName)

	// Confirm the VM is running the correct Ubuntu version
	vmImage := azure.GetVirtualMachineImage(t, vmName, resourceGroupName, subscriptionID)
	expectedOSPublisher := "Canonical"
	expectedOSVersion := "22_04-lts-gen2"
	assert.Equal(t, expectedOSPublisher, vmImage.Publisher, "VM image publisher is not %s", expectedOSPublisher)
	assert.Equal(t, expectedOSVersion, vmImage.SKU, "VM image SKU is not %s", expectedOSVersion)
}
