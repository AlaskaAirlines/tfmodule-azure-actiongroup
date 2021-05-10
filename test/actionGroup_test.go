package test

import (
	"fmt"
	"strings"

	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const (
	// AzureSubscriptionID is an optional env variable supported by the `azurerm` Terraform provider to
	// designate a target Azure subscription ID
	AzureSubscriptionID = "ARM_SUBSCRIPTION_ID"

	// AzureResGroupName is an optional env variable custom to Terratest to designate a target Azure resource group
	AzureResGroupName = "tfmodulevalidation-test-group"
)

func TestTerraformEmail(t *testing.T) {
	t.Parallel()
	_random := strings.ToLower(random.UniqueId())

	expectedName := fmt.Sprintf("%s-test-actiongroup", _random)

	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/email/.",
		Vars: map[string]interface{}{
			"resource_group_name": AzureResGroupName,
			"appName":             _random,
			"environment":         "test",
			"shortName":           "blah",
			"enableEmail":         true,
			"emailName":           "TestName",
			"emailAddress":        "sample@test.com",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	actionGroupID := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(actionGroupID)
	assert.Contains(actionGroupID, expectedName)

	actionGroup := azure.GetActionGroupResource(t, expectedName, "", "")

	assert.NotNil(actionGroup)
	assert.Equal(1, len(*actionGroup.EmailReceivers))
	assert.Equal(0, len(*actionGroup.SmsReceivers))
	assert.Equal(0, len(*actionGroup.WebhookReceivers))
}

func TestTerraformEmailAndWebHook(t *testing.T) {
	t.Parallel()
	_random := strings.ToLower(random.UniqueId())

	expectedName := fmt.Sprintf("%s-test-actiongroup", _random)

	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/emailAndWebHook/.",
		Vars: map[string]interface{}{
			"resource_group_name": AzureResGroupName,
			"appName":             _random,
			"environment":         "test",
			"shortName":           "blah",
			"enableEmail":         true,
			"emailName":           "TestName",
			"emailAddress":        "sample@test.com",
			"enableWebHook":       true,
			"webhookName":         "webhookTestName",
			"webhookServiceUri":   "http://example.com/alert",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	actionGroupID := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(actionGroupID)
	assert.Contains(actionGroupID, expectedName)

	actionGroup := azure.GetActionGroupResource(t, expectedName, "", "")

	assert.NotNil(actionGroup)
	assert.Equal(1, len(*actionGroup.EmailReceivers))
	assert.Equal(0, len(*actionGroup.SmsReceivers))
	assert.Equal(1, len(*actionGroup.WebhookReceivers))
}

func TestTerraformSms(t *testing.T) {
	t.Parallel()
	_random := strings.ToLower(random.UniqueId())

	expectedName := fmt.Sprintf("%s-test-actiongroup", _random)

	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/sms/.",
		Vars: map[string]interface{}{
			"resource_group_name": AzureResGroupName,
			"appName":             _random,
			"environment":         "test",
			"shortName":           "blah",
			"enableSMS":           true,
			"smsName":             "TestName",
			"smsPhoneNumber":      5551234567,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	actionGroupID := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(actionGroupID)
	assert.Contains(actionGroupID, expectedName)

	actionGroup := azure.GetActionGroupResource(t, expectedName, "", "")

	assert.NotNil(actionGroup)
	assert.Equal(0, len(*actionGroup.EmailReceivers))
	assert.Equal(1, len(*actionGroup.SmsReceivers))
	assert.Equal(0, len(*actionGroup.WebhookReceivers))
}

func TestTerraformWebhook(t *testing.T) {
	t.Parallel()
	_random := strings.ToLower(random.UniqueId())

	expectedName := fmt.Sprintf("%s-test-actiongroup", _random)

	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/webhook/.",
		Vars: map[string]interface{}{
			"resource_group_name": AzureResGroupName,
			"appName":             _random,
			"environment":         "test",
			"shortName":           "blah",
			"enableWebHook":       true,
			"webhookName":         "webhookTestName",
			"webhookServiceUri":   "http://example.com/alert",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	actionGroupID := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(actionGroupID)
	assert.Contains(actionGroupID, expectedName)

	actionGroup := azure.GetActionGroupResource(t, expectedName, "", "")

	assert.NotNil(actionGroup)
	assert.Equal(0, len(*actionGroup.EmailReceivers))
	assert.Equal(0, len(*actionGroup.SmsReceivers))
	assert.Equal(1, len(*actionGroup.WebhookReceivers))
}
