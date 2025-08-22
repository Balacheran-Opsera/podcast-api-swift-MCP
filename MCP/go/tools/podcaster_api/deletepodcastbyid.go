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

func DeletepodcastbyidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["reason"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("reason=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/podcasts/%s%s", cfg.BaseURL, id, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result models.DeletePodcastResponse
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

func CreateDeletepodcastbyidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_podcasts_id",
		mcp.WithDescription("Request to delete a podcast"),
		mcp.WithString("X-ListenAPI-Key", mcp.Required(), mcp.Description("Get API Key on listennotes.com/api")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Podcast id. You can get podcast id from using other endpoints, e.g., `GET /search`, `GET /best_podcasts`...")),
		mcp.WithString("reason", mcp.Description("The reason why this podcast should be deleted, e.g., copyright violation, the podcaster wants to delete it... You can put \"testing\" here to indicate that you are testing this endpoint, so we will not actually delete the podcast.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletepodcastbyidHandler(cfg),
	}
}
