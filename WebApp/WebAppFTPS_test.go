package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestFtpstate(t *testing.T) {
	t.Parallel()
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/WebApp",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"WebAppFTPDeployFtpsonly.Auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyE(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	expectedFtpStatus := terraform.Output(t, terraformOptions, "WebApp_Ftp_status_test_one")
	// expectedFtpStatusThree := terraform.Output(t, terraformOptions, "WebApp_Ftp_status_test_three")

	assert.Equal(t, expectedFtpStatus, "ftp_State")

}

func TestDbFtpFalse(t *testing.T) {
	t.Parallel()
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/WebApp",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"WebAppFTPDeployDisabled.Auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	// expectedFtpStatusThree := terraform.Output(t, terraformOptions, "WebApp_Ftp_status_test_three")
}

func TestDbFtpFalse(t *testing.T) {
	t.Parallel()
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/WebApp",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"WebAppFTPDeployAllowed.Auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	// expectedFtpStatusThree := terraform.Output(t, terraformOptions, "WebApp_Ftp_status_test_three")
}

func TestDbFtpModify(t *testing.T) {
	t.Parallel()
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/WebApp",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"WebAppFTPDeployFtpsonly.Auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyE(t, terraformOptions)

	terraformOptionsModify := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/WebApp",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"WebAppFTPDeployDisabled.Auto.tfvars"},
	})

	terraform.InitAndApplyE(t, terraformOptionsModify)
}

func TestDbFtpModify(t *testing.T) {
	t.Parallel()
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/WebApp",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"WebAppFTPDeployFtpsonly.Auto.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyE(t, terraformOptions)

	terraformOptionsModify := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../modules/WebApp",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"WebAppFTPDeployAllowed.Auto.tfvars"},
	})

	terraform.InitAndApplyE(t, terraformOptionsModify)
}