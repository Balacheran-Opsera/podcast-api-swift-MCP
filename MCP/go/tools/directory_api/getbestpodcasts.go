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

func GetbestpodcastsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["genre_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("genre_id=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["region"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("region=%v", val))
		}
		if val, ok := args["publisher_region"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("publisher_region=%v", val))
		}
		if val, ok := args["language"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("language=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["safe_mode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("safe_mode=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/best_podcasts%s", cfg.BaseURL, queryString)
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
		var result models.BestPodcastsResponse
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

func CreateGetbestpodcastsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_best_podcasts",
		mcp.WithDescription("Fetch a list of best podcasts by genre"),
		mcp.WithString("X-ListenAPI-Key", mcp.Required(), mcp.Description("Get API Key on listennotes.com/api")),
		mcp.WithString("genre_id", mcp.Description("You can get the id from `GET /genres`. If not specified, it'll be the overall best podcasts, which can be considered as a special genre.")),
		mcp.WithNumber("page", mcp.Description("Page number of those podcasts in this genre.")),
		mcp.WithString("region", mcp.Description("Filter best podcasts by country/region.\nPlease note that podcasts that are \"best\" in a country/region may not be produced in that country/region.\nFor example, a podcast from the US may be very popular in Canada.\nYou can get the supported country codes (e.g., us, jp, gb...) from `GET /regions`.\nIf not specified, you'll get \"best podcasts\" in United States.\n")),
		mcp.WithString("publisher_region", mcp.Description("Filter best podcasts by the publisher's country/region.\nThis is to narrow down the results to include \"best podcasts\" produced in a specific country/region.\nYou can get the supported country codes (e.g., us, jp, gb...) from `GET /regions`.\nIf not specified, you'll get \"best podcasts\" produced in any country/region.\nIf you want to get a country/region's \"best podcasts\" that are also produced in that country/region,\nthen you need to specify both **region** and **publisher_region**,\ne.g., `region=jp` and `publisher_region=jp`.\n")),
		mcp.WithString("language", mcp.Description("Filter best podcasts by language.\nYou can get a list of supported languages (e.g., English, Chinese, Japanese...) from `GET /languages`.\nIf not specified, you'll get \"best podcasts\" in any language.\n")),
		mcp.WithString("sort", mcp.Description("How do you want to sort these podcasts?\nIf you'd like to sort by popularity, please use **listen_score**.\n")),
		mcp.WithNumber("safe_mode", mcp.Description("Whether or not to exclude podcasts with explicit language. 1 is yes, and 0 is no.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetbestpodcastsHandler(cfg),
	}
}
