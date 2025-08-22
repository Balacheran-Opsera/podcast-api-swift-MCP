package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/config"
	"github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetpodcastsbydomainnameHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		domain_nameVal, ok := args["domain_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: domain_name"), nil
		}
		domain_name, ok := domain_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: domain_name"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/podcasts/domains/%s%s", cfg.BaseURL, domain_name, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			req.Header.Set("X-ListenAPI-Key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-ListenAPI-Key"]; ok {
			req.Header.Set("X-ListenAPI-Key", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.PodcastDomainResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetpodcastsbydomainnameTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_podcasts_domains_domain_name",
		mcp.WithDescription("Fetch podcasts by a publisher's domain name"),
		mcp.WithString("X-ListenAPI-Key", mcp.Required(), mcp.Description("Get API Key on listennotes.com/api")),
		mcp.WithString("domain_name", mcp.Required(), mcp.Description("A publisher's domain name, e.g., nytimes.com, wondery.com, npr.org...")),
		mcp.WithNumber("page", mcp.Description("Page number of the podcasts from this domain name")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetpodcastsbydomainnameHandler(cfg),
	}
}
