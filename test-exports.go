package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/costmanagement/mgmt/costmanagement"
	"github.com/pjirsa/azure-costmanagement-samples/export"
	"github.com/pjirsa/azure-costmanagement-samples/internal/config"
)

func main() {
	err := config.ParseEnvironment()
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()

	//scope := "subscriptions/" + config.SubscriptionID()
	scope := fmt.Sprintf("/providers/Microsoft.Management/managementGroups/%s", config.ManagementGroupID())

	ListExports(ctx, scope)
}

func ListExports(ctx context.Context, scope string) {
	resp, err := export.List(ctx, scope)

	if err != nil {
		fmt.Println("Error while getting list of exports", err)
		return
	}

	result, _ := json.Marshal(resp.Value)

	fmt.Println("List of configured exports:")
	fmt.Println(string(result))
}

func GetExport(ctx context.Context, scope, exportName string) {
	resp, err := export.Get(ctx, scope, exportName)

	if err != nil {
		fmt.Println("Error while getting export")
	}

	result, _ := json.Marshal(resp)

	fmt.Println("Export found")
	fmt.Println(string(result))
}

func DeleteExport(ctx context.Context, scope, exportName string) {
	_, err := export.Delete(ctx, scope, exportName)

	if err != nil {
		fmt.Println("Error while deleting export")
		return
	}

	fmt.Println("Deleted export")
}

func CreateExport(ctx context.Context, scope, exportName string) {
	resp, err := export.CreateOrUpdate(ctx, scope, exportName, CreateExportPreparer())

	if err != nil {
		fmt.Println("Error while creating export")
		return
	}

	result, _ := json.Marshal(resp)

	fmt.Println("Created new export")
	fmt.Println(result)
}

func CreateExportPreparer() (result costmanagement.Export) {

	return
}
