package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbEncryptTrue(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/azureSQL",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"DbEncryptionFalse.auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	// expectedEncryptStatusThree := terraform.Output(t, terraformOptions, "db_encryption_status_test_three")
}

func TestDbEncryptModify(t *testing.T) {
	t.Parallel()
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/azureSQL",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"DbEncryptionModifyTrue.auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyE(t, terraformOptions)

	terraformOptionsModify := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/azureSQL",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"DbEncryptionModifyFalse.auto.tfvars"},
	})

	terraform.InitAndApplyE(t, terraformOptionsModify)
}
