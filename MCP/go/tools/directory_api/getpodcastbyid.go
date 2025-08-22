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

func GetpodcastbyidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["next_episode_pub_date"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_episode_pub_date=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/podcasts/%s%s", cfg.BaseURL, id, queryString)
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
		var result models.PodcastFull
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

func CreateGetpodcastbyidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_podcasts_id",
		mcp.WithDescription("Fetch detailed meta data and episodes for a podcast by id"),
		mcp.WithString("X-ListenAPI-Key", mcp.Required(), mcp.Description("Get API Key on listennotes.com/api")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Podcast id. You can get podcast id from using other endpoints, e.g., `GET /search`, `GET /best_podcasts`...")),
		mcp.WithNumber("next_episode_pub_date", mcp.Description("For episodes pagination. It's the value of **next_episode_pub_date** from the response of last request. If not specified, just return latest 10 episodes or oldest 10 episodes, depending on the value of the **sort** parameter.\n")),
		mcp.WithString("sort", mcp.Description("How do you want to sort the episodes of this podcast?\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetpodcastbyidHandler(cfg),
	}
}
