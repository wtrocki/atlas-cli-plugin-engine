package main

import (
	"fmt"

	"github.com/danielgtaylor/openapi-cli-generator/cli"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/h2non/gentleman.v2"
)

var atlasApiTransformedSubcommand bool

// TODO use atlas-cli auth/config
// TODO requires custom template (ignore html/descriptions in commands)
// TODO use external command metadata

// AtlasApiTransformedCreateOrganization Create One Organization
func AtlasApiTransformedCreateOrganization(params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "create"
	if atlasApiTransformedSubcommand {
		handlerPath = "atlas-api-transformed " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = "https://cloud.mongodb.com"
	}

	url := server + "/api/atlas/v2/orgs"

	req := cli.Client.Post().URL(url)

	paramEnvelope := params.GetBool("envelope")
	if paramEnvelope != false {
		req = req.AddQuery("envelope", fmt.Sprintf("%v", paramEnvelope))
	}
	paramPretty := params.GetBool("pretty")
	if paramPretty != false {
		req = req.AddQuery("pretty", fmt.Sprintf("%v", paramPretty))
	}

	if body != "" {
		req = req.AddHeader("Content-Type", "application/vnd.atlas.2023-01-01+json").BodyString(body)
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// AtlasApiTransformedListOrganizations Return All Organizations
func AtlasApiTransformedListOrganizations(params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "listorganizations"
	if atlasApiTransformedSubcommand {
		handlerPath = "atlas-api-transformed " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = "https://cloud.mongodb.com"
	}

	url := server + "/api/atlas/v2/orgs"

	req := cli.Client.Get().URL(url)

	paramEnvelope := params.GetBool("envelope")
	if paramEnvelope != false {
		req = req.AddQuery("envelope", fmt.Sprintf("%v", paramEnvelope))
	}
	paramIncludecount := params.GetBool("includecount")
	if paramIncludecount != false {
		req = req.AddQuery("includeCount", fmt.Sprintf("%v", paramIncludecount))
	}
	paramItemsperpage := params.GetInt64("itemsperpage")
	if paramItemsperpage != 0 {
		req = req.AddQuery("itemsPerPage", fmt.Sprintf("%v", paramItemsperpage))
	}
	paramPagenum := params.GetInt64("pagenum")
	if paramPagenum != 0 {
		req = req.AddQuery("pageNum", fmt.Sprintf("%v", paramPagenum))
	}
	paramPretty := params.GetBool("pretty")
	if paramPretty != false {
		req = req.AddQuery("pretty", fmt.Sprintf("%v", paramPretty))
	}
	paramName := params.GetString("name")
	if paramName != "" {
		req = req.AddQuery("name", fmt.Sprintf("%v", paramName))
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

func atlasApiTransformedRegister(subcommand bool) {
	root := cli.Root

	if subcommand {
		root = &cobra.Command{
			Use:   "atlas-api-transformed",
			Short: "MongoDB Atlas Administration API",
			Long:  cli.Markdown("The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.\nTo learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/).\nThis OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources."),
		}
		atlasApiTransformedSubcommand = true
	} else {
		cli.Root.Short = "MongoDB Atlas Administration API"
		cli.Root.Long = cli.Markdown("The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.\nTo learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/).\nThis OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.")
	}

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "create",
			Short:   "Create One Organization",
			Long:    cli.Markdown("Creates one organization in MongoDB Cloud and links it to the requesting API Key's organization. To use this resource, the requesting API Key must have the Organization Owner role. The requesting API Key's organization must be a paying organization.\n## Request Schema (application/vnd.atlas.2023-01-01+json)\n\nproperties:\n  apiKey:\n    $ref: '#/components/schemas/ApiCreateApiKeyView'\n  name:\n    description: Human-readable label that identifies the organization.\n    maxLength: 64\n    minLength: 1\n    type: string\n  orgOwnerId:\n    description: Unique 24-hexadecimal digit string that identifies the Atlas user\n      that you want to assign the Organization Owner role. This user must be a member\n      of the same organization as the calling API key. This is required if you call\n      the Admin API endpoint directly, but not required when you call through the\n      Atlas CLI.\n    example: 32b6e34b3d91647abb20e7b8\n    maxLength: 24\n    minLength: 24\n    pattern: ^([a-f0-9]{24})$\n    type: string\nrequired:\n- name\ntype: object\n"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {
				body, err := cli.GetBody("application/vnd.atlas.2023-01-01+json", args[0:])
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get body")
				}

				_, decoded, err := AtlasApiTransformedCreateOrganization(params, body)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cmd.Flags().String("envelope", "", "")
		cmd.Flags().String("pretty", "", "")

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "list",
			Short:   "Return All Organizations",
			Long:    cli.Markdown("Returns all organizations to which you belong. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List."),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := AtlasApiTransformedListOrganizations(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cmd.Flags().String("envelope", "", "")
		cmd.Flags().String("includecount", "", "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
		cmd.Flags().Int64("itemsperpage", 0, "Number of items that the response returns per page.")
		cmd.Flags().Int64("pagenum", 0, "Number of the page that displays the current set of the total objects that the response returns.")
		cmd.Flags().String("pretty", "", "")
		cmd.Flags().String("name", "", "Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name.")

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()
	root.Execute()
}
