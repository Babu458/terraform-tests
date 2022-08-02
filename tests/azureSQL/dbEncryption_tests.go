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

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyE(t, terraformOptions)
	// if _, err := terraform.InitAndApplyE(t, terraformOptions); err != nil {
	// 	policy_error = +1
	// 	t.Log(err)
	// 	terraform.Destroy(t, terraformOptions)
	// }

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	expectedEncryptStatusOne := terraform.Output(t, terraformOptions, "db_encryption_status_test_one")
	expectedEncryptStatusTwo := terraform.Output(t, terraformOptions, "db_encryption_status_test_two")
	expectedEncryptStatusThree := terraform.Output(t, terraformOptions, "db_encryption_status_test_three")

	assert.Equal(t, expectedEncryptStatusOne, "true")
	assert.Equal(t, expectedEncryptStatusTwo, "false")
	assert.Equal(t, expectedEncryptStatusThree, "true")

}