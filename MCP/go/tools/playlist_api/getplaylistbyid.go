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

func GetplaylistbyidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["last_timestamp_ms"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_timestamp_ms=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/playlists/%s%s", cfg.BaseURL, id, queryString)
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
		var result models.PlaylistResponse
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

func CreateGetplaylistbyidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_playlists_id",
		mcp.WithDescription("Fetch a playlist's info and items (i.e., episodes or podcasts)."),
		mcp.WithString("X-ListenAPI-Key", mcp.Required(), mcp.Description("Get API Key on listennotes.com/api")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Playlist id (always 11 characters, e.g., m1pe7z60bsw).\nYou can get the podcast id from the url of a playlist, e.g.,\nm1pe7z60bsw is the playlist id of listennotes.com/listen/podcasts-about-podcasting-m1pe7z60bsw\n")),
		mcp.WithString("type", mcp.Description("The type of this playlist, which should be either **episode_list** or **podcast_list**.\n")),
		mcp.WithNumber("last_timestamp_ms", mcp.Description("For playlist items pagination.\nIt's the value of **last_timestamp_ms** from the response of last request.\nIf it's 0 or not specified, just return the latest or the oldest 20 items,\ndepending on the value of the **sort** parameter.\n")),
		mcp.WithString("sort", mcp.Description("How do you want to sort playlist items?\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetplaylistbyidHandler(cfg),
	}
}
