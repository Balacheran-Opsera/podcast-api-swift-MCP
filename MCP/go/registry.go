package main

import (
	"github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/config"
	"github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/models"
	tools_directory_api "github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/tools/directory_api"
	tools_search_api "github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/tools/search_api"
	tools_podcaster_api "github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/tools/podcaster_api"
	tools_playlist_api "github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/tools/playlist_api"
	tools_insights_api "github.com/listen-api-podcast-search-directory-and-insights-api/mcp-server/tools/insights_api"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_directory_api.CreateGetgenresTool(cfg),
		tools_directory_api.CreateGetpodcastrecommendationsTool(cfg),
		tools_search_api.CreateGettrendingsearchesTool(cfg),
		tools_podcaster_api.CreateDeletepodcastbyidTool(cfg),
		tools_directory_api.CreateGetpodcastbyidTool(cfg),
		tools_playlist_api.CreateGetplaylistsTool(cfg),
		tools_directory_api.CreateGetregionsTool(cfg),
		tools_search_api.CreateGetrelatedsearchesTool(cfg),
		tools_directory_api.CreateGetlanguagesTool(cfg),
		tools_insights_api.CreateGetpodcastaudienceTool(cfg),
		tools_search_api.CreateSearchTool(cfg),
		tools_search_api.CreateTypeaheadTool(cfg),
		tools_directory_api.CreateJustlistenTool(cfg),
		tools_insights_api.CreateGetpodcastsbydomainnameTool(cfg),
		tools_search_api.CreateSpellcheckTool(cfg),
		tools_directory_api.CreateGetbestpodcastsTool(cfg),
		tools_directory_api.CreateGetepisoderecommendationsTool(cfg),
		tools_directory_api.CreateGetcuratedpodcastbyidTool(cfg),
		tools_directory_api.CreateGetepisodebyidTool(cfg),
		tools_directory_api.CreateGetcuratedpodcastsTool(cfg),
		tools_playlist_api.CreateGetplaylistbyidTool(cfg),
	}
}
