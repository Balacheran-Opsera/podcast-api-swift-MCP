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

func SearchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["q"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("q=%v", val))
		}
		if val, ok := args["sort_by_date"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by_date=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["len_min"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("len_min=%v", val))
		}
		if val, ok := args["len_max"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("len_max=%v", val))
		}
		if val, ok := args["episode_count_min"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("episode_count_min=%v", val))
		}
		if val, ok := args["episode_count_max"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("episode_count_max=%v", val))
		}
		if val, ok := args["update_freq_min"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("update_freq_min=%v", val))
		}
		if val, ok := args["update_freq_max"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("update_freq_max=%v", val))
		}
		if val, ok := args["genre_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("genre_ids=%v", val))
		}
		if val, ok := args["published_before"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("published_before=%v", val))
		}
		if val, ok := args["published_after"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("published_after=%v", val))
		}
		if val, ok := args["only_in"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("only_in=%v", val))
		}
		if val, ok := args["language"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("language=%v", val))
		}
		if val, ok := args["region"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("region=%v", val))
		}
		if val, ok := args["ocid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ocid=%v", val))
		}
		if val, ok := args["ncid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ncid=%v", val))
		}
		if val, ok := args["safe_mode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("safe_mode=%v", val))
		}
		if val, ok := args["unique_podcasts"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("unique_podcasts=%v", val))
		}
		if val, ok := args["page_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/search%s", cfg.BaseURL, queryString)
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
		var result models.SearchResponse
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

func CreateSearchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_search",
		mcp.WithDescription("Full-text search"),
		mcp.WithString("X-ListenAPI-Key", mcp.Required(), mcp.Description("Get API Key on listennotes.com/api")),
		mcp.WithString("q", mcp.Required(), mcp.Description("Search term, e.g., person, place, topic... You can use double quotes to do verbatim match, e.g., \"game of thrones\". Otherwise, it's fuzzy search.\n")),
		mcp.WithNumber("sort_by_date", mcp.Description("Sort by date or not? If 0, then sort by relevance. If 1, then sort by date.\n")),
		mcp.WithString("type", mcp.Description("What type of contents do you want to search for? \n")),
		mcp.WithNumber("offset", mcp.Description("Offset for search results, for pagination. You'll use **next_offset** from response for this parameter.\n")),
		mcp.WithNumber("len_min", mcp.Description("Minimum audio length in minutes. Applicable only when **type** parameter is **episode** or **podcast**.\nIf **type** parameter is **episode**, it's for audio length of an episode.\nIf **type** parameter is **podcast**, it's for average audio length of all episodes in a podcast.\n")),
		mcp.WithNumber("len_max", mcp.Description("Maximum audio length in minutes. Applicable only when **type** parameter is **episode** or **podcast**.\nIf **type** parameter is **episode**, it's for audio length of an episode.\nIf **type** parameter is **podcast**, it's for average audio length of all episodes in a podcast.\n")),
		mcp.WithNumber("episode_count_min", mcp.Description("Minimum number of episodes. Applicable only when type parameter is **podcast**.\n")),
		mcp.WithNumber("episode_count_max", mcp.Description("Maximum number of episodes. Applicable only when type parameter is **podcast**.\n")),
		mcp.WithNumber("update_freq_min", mcp.Description("Minimum update frequency in hours (how frequently does a podcast release a new episode). For example, if you want to find \"weekly\" podcasts, then you can set **update_freq_min**=144 hours (or 6 days) and **update_freq_max**=192 hours (or 8 days). Applicable only when type parameter is **podcast**.\n")),
		mcp.WithNumber("update_freq_max", mcp.Description("Maximum update frequency in hours (how frequently does a podcast release a new episode). For example, if you want to find \"weekly\" podcasts, then you can set **update_freq_min**=144 hours (or 6 days) and **update_freq_max**=192 hours (or 8 days). Applicable only when type parameter is **podcast**.\n")),
		mcp.WithString("genre_ids", mcp.Description("A comma-delimited string of a list of genre ids. If not specified, then all genres are included. You can find the id and the name of all genres from `GET /genres`. It works only when **type** is *episode* or *podcast*.\n")),
		mcp.WithNumber("published_before", mcp.Description("Only show episodes/podcasts/curated lists published before this timestamp (in milliseconds). If **published_before** & **published_after** are used at the same time, **published_before** should be bigger than **published_after**.\n")),
		mcp.WithNumber("published_after", mcp.Description("Only show episodes/podcasts/curated lists published after this timestamp (in milliseconds). If **published_before** & **published_after** are used at the same time, **published_before** should be bigger than **published_after**.\n")),
		mcp.WithString("only_in", mcp.Description("A comma-delimited string to search only in specific fields. Allowed values are title, description, author, and audio. If not specified, then search every fields.\n")),
		mcp.WithString("language", mcp.Description("Limit search results to a specific language. If not specified, it'll be any language. You can get a list of supported languages from `GET /languages`. It works only when **type** is *episode* or *podcast*.\n")),
		mcp.WithString("region", mcp.Description("Limit search results to a specific region (e.g., us, gb, in...). If not specified, it'll be any region. You can get the supported country codes from `GET /regions`. It works only when **type** is *episode* or *podcast*.\n")),
		mcp.WithString("ocid", mcp.Description("A comma-delimited string of podcast ids (up to 5 podcasts) - you can get a podcast id from the **podcast_id** field in response. This parameter is to limit search results from only a few specific podcasts. It works only when **type** is *episode*.\n")),
		mcp.WithString("ncid", mcp.Description("A comma-delimited string of podcast ids (up to 5 podcasts) - you can get a podcast id from the **podcast_id** field in response. This parameter is to exclude search results of a few specific podcasts. It works only when **type** is *episode*.\n")),
		mcp.WithNumber("safe_mode", mcp.Description("Whether or not to exclude podcasts/episodes with explicit language. 1 is yes and 0 is no. It works only when **type** is *episode* or *podcast*.\n")),
		mcp.WithNumber("unique_podcasts", mcp.Description("Whether or not to keep only one episode per podcast in search results. 1 is yes and 0 is no. It works only when **type** is *episode*.\n")),
		mcp.WithNumber("page_size", mcp.Description("The maximum number of search results per page. A valid value should be an integer between 1 and 10 (inclusive).\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SearchHandler(cfg),
	}
}
