package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestDbEncrypt(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/azureSQL",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	if _, err := terraform.InitAndApplyE(t, terraformOptions); err != nil {
		t.Log(err)
		expectedEncryptionStatus := terraform.Output(t, terraformOptions, "transparent_data_encryption_enabled")
	  
		assert.Equal(t, "false", expectedEncryptionStatus)
	}

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	expectedEncryptionStatus := terraform.Output(t, terraformOptions, "transparent_data_encryption_enabled")
	assert.Equal(t, "true", expectedEncryptionStatus)
}