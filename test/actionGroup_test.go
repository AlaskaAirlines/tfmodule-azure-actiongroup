package test

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/resources/mgmt/insights"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// AzureSubscriptionID is an optional env variable supported by the `azurerm` Terraform provider to
	// designate a target Azure subscription ID
	AzureSubscriptionID = "ARM_SUBSCRIPTION_ID"

	// AzureResGroupName is an optional env variable custom to Terratest to designate a target Azure resource group
	AzureResGroupName = "tfmodulevalidation-test-group"
)

func TestTerraformEmail(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/email/.",
	}
	expectedName := "emailSample-test-actiongroup"
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(outputValue)
	assert.Contains(outputValue, expectedName)

	actionGroup := GetActionGroupsResource(t, expectedName)
	assert.Equal(1, len(*actionGroup.EmailReceivers))
	assert.Equal(0, len(*actionGroup.SmsReceivers))
	assert.Equal(0, len(*actionGroup.WebhookReceivers))
}

func TestTerraformEmailAndWebHook(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/emailAndWebHook/.",
	}
	expectedName := "emailWebHookSample-test-actiongroup"
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(outputValue)
	assert.Contains(outputValue, expectedName)

	actionGroup := GetActionGroupsResource(t, expectedName)
	assert.Equal(1, len(*actionGroup.EmailReceivers))
	assert.Equal(0, len(*actionGroup.SmsReceivers))
	assert.Equal(1, len(*actionGroup.WebhookReceivers))
}

func TestTerraformSms(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/sms/.",
	}
	expectedName := "smsSample-test-actiongroup"
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(outputValue)
	assert.Contains(outputValue, expectedName)

	actionGroup := GetActionGroupsResource(t, expectedName)
	assert.Equal(0, len(*actionGroup.EmailReceivers))
	assert.Equal(1, len(*actionGroup.SmsReceivers))
	assert.Equal(0, len(*actionGroup.WebhookReceivers))
}

func TestTerraformWebhook(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/webhook/.",
	}
	expectedName := "webhookSample-test-actiongroup"
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndApply(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(outputValue)
	assert.Contains(outputValue, expectedName)

	actionGroup := GetActionGroupsResource(t, expectedName)
	assert.Equal(0, len(*actionGroup.EmailReceivers))
	assert.Equal(0, len(*actionGroup.SmsReceivers))
	assert.Equal(1, len(*actionGroup.WebhookReceivers))
}

func GetActionGroupsResource(t *testing.T, ruleName string) *insights.ActionGroupResource {
	actionGroupResource, err := getActionGroupsResourceE(ruleName)
	require.NoError(t, err)

	return actionGroupResource
}

func getActionGroupsResourceE(ruleName string) (*insights.ActionGroupResource, error) {
	client, err := getActionGroupsClient()
	if err != nil {
		return nil, err
	}

	actionGroup, err := client.Get(context.Background(), AzureResGroupName, ruleName)

	return &actionGroup, nil
}

func getActionGroupsClient() (*insights.ActionGroupsClient, error) {
	subID := os.Getenv(AzureSubscriptionID)

	metricAlertsClient := insights.NewActionGroupsClient(subID)

	authorizer, err := azure.NewAuthorizer()
	if err != nil {
		return nil, err
	}

	metricAlertsClient.Authorizer = *authorizer

	return &metricAlertsClient, nil
}
