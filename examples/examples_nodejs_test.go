package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getTestOrg() string {
	return os.Getenv("PULUMI_TEST_OWNER")
}

func TestAccessTokenExample(t *testing.T) {
	cwd, _ := os.Getwd()
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Quick:       true,
		SkipRefresh: true,
		Dir:         path.Join(cwd, ".", "access-tokens"),
		Config: map[string]string{
			"orgName": getTestOrg(),
		},
		Dependencies: []string{
			"@pulumi/pulumiservice",
		},
	})
}

func TestStackTagsExample(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Quick:       true,
		SkipRefresh: true,
		Dir:         path.Join(cwd, ".", "ts-stack-tags"),
		Config: map[string]string{
			"orgName": getTestOrg(),
		},
		Dependencies: []string{
			"@pulumi/pulumiservice",
		},
	})
}
