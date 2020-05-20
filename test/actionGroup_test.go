package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformEmail(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/Email/.",
	}
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndPlan(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(outputValue)
}

func TestTerraformEmailAndWebHook(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/EmailAndWebHook/.",
	}
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndPlan(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(outputValue)
}

func TestTerraformSms(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/sms/.",
	}
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndPlan(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(outputValue)
}

func TestTerraformWebhook(t *testing.T) {
	// Arrange
	terraformOptions := &terraform.Options{
		TerraformDir: "../example/webhook/.",
	}
	defer terraform.Destroy(t, terraformOptions)

	// Act
	terraform.InitAndPlan(t, terraformOptions)

	// Assert
	assert := assert.New(t)

	outputValue := terraform.Output(t, terraformOptions, "action_group_id")
	assert.NotNil(outputValue)
}
