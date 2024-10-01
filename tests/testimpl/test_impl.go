package testimpl

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	armAppService "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestAppServicePlan(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionId) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %e\n", err)
	}

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud.AzurePublic,
		},
	}

	appServicePlanClient, err := armAppService.NewPlansClient(subscriptionId, credential, &options)
	if err != nil {
		t.Fatalf("Error getting App Service Plan client: %v", err)
	}

	t.Run("doesAppServicePlanExist", func(t *testing.T) {
		appServicePlanName := terraform.Output(t, ctx.TerratestTerraformOptions(), "name")
		resourceGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")

		appServicePlan, err := appServicePlanClient.Get(context.Background(), resourceGroupName, appServicePlanName, nil)
		if err != nil {
			t.Fatalf("Error getting app service plan: %v", err)
		}

		assert.Equal(t, appServicePlanName, *appServicePlan.Name)
	})
}
