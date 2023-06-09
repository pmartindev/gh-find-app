package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cli/go-gh"
)

func _main() error {
	var org = ""
	var err = survey.AskOne(&survey.Input{
		Message: "Enter the name of the org you want to revoke access to:",
	}, &org)
	if err != nil {
		return err
	} else if org == "" {
		return fmt.Errorf("org cannot be empty")
	}
	appId := ""
	err = survey.AskOne(&survey.Password{
		Message: "Enter the github personal access token to be revoked:",
	}, &appId)
	if err != nil {
		return err
	} else if appId == "" {
		return fmt.Errorf("appId cannot be empty")
	}
	return nil
}

type AppInstallation struct {
	ID      int    `json:"id"`
	AppID   int    `json:"app_id"`
	AppSlug string `json:"app_slug"`
}

func getAppInstllations(org string) error {
	client, err := gh.RESTClient(nil)
	response := []AppInstallation{}
	client.Get(fmt.Sprintf("orgs/%s/installations", org), &response)

	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	// print response
	fmt.Println(response)
	return nil
}
func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "X %s", err.Error())
	}
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
