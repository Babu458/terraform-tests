package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestDbEncryptTrue(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/azureSQL",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"DbEncryptionTrue.auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyE(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	expectedEncryptStatus := terraform.Output(t, terraformOptions, "db_encryption_status_test_one")
	// expectedEncryptStatusThree := terraform.Output(t, terraformOptions, "db_encryption_status_test_three")

	assert.Equal(t, expectedEncryptStatus, "true")

}

func TestDbEncryptFalse(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/azureSQL",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"DbEncryptionFalse.auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyE(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	expectedEncryptStatusOne := terraform.Output(t, terraformOptions, "db_encryption_status_test_one")
	// expectedEncryptStatusThree := terraform.Output(t, terraformOptions, "db_encryption_status_test_three")

	assert.Equal(t, expectedEncryptStatusOne, "false")
}

func TestDbEncryptModify(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/azureSQL",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"DbEncryptionTrue.auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyE(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	expectedEncryptStatus := terraform.Output(t, terraformOptions, "db_encryption_status_test_one")
	// expectedEncryptStatusThree := terraform.Output(t, terraformOptions, "db_encryption_status_test_three")

	assert.Equal(t, expectedEncryptStatus, "true")

	terraformOptionsModify := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/azureSQL",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"DbEncryptionFalse.auto.tfvars"},
	})

	terraform.InitAndApplyE(t, terraformOptionsModify)

	assert.Equal(t, expectedEncryptStatus, "true")
}