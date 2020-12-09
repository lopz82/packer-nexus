package image_test

import (
	"flag"
	"fmt"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/ssh"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
	"time"
)

var ami = flag.String("ami", "", "Ami to be tested")

func TestNexusRunningProperly(t *testing.T) {
	opts := &terraform.Options{
		TerraformDir: "./terraform",
		Vars: map[string]interface{}{
			"ami": *ami,
		},
		VarFiles: []string{"./test.tfvars"},
	}
	defer terraform.Destroy(t, opts)

	terraform.InitAndApply(t, opts)
	instanceIP := terraform.OutputRequired(t, opts, "public_ip")

	host := ssh.Host{
		Hostname:    instanceIP,
		SshUserName: "ubuntu",
		SshAgent:    true,
	}

	maxRetries := 10
	betweenRetries := 5 * time.Second
	commands := []string{
		"systemctl is-enabled --quiet nexus.service && echo Nexus repository is enabled",
		"systemctl is-active --quiet nexus.service && echo Nexus repository is active",
	}
	retry.DoWithRetry(t, "Checking if Nexus repository service is enabled and active", maxRetries, betweenRetries, func() (string, error) {
		for _, command := range commands {
			res, err := ssh.CheckSshCommandE(t, host, command)
			fmt.Printf("Output from '%s': %s", command, res)
			if err != nil {
				return "FAILED", err
			}
		}
		return "OK", nil
	})
	maxRetries = 12
	betweenRetries = 30 * time.Second
	commands = []string{
		"cat /opt/sonatype-work/nexus3/log/nexus.log | grep 'Started Sonatype Nexus OSS'",
		"curl -I http://localhost:8081",
	}
	retry.DoWithRetry(t, "Checking if Nexus repository is running", maxRetries, betweenRetries, func() (string, error) {
		for _, command := range commands {
			res, err := ssh.CheckSshCommandE(t, host, command)
			fmt.Printf("Output from '%s': %s", command, res)
			if err != nil {
				return "FAILED", err
			}
		}
		return "OK", nil
	})
}
