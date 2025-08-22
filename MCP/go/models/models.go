package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// EpisodeMinimum represents the EpisodeMinimum schema from the OpenAPI specification
type EpisodeMinimum struct {
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Link string `json:"link,omitempty"` // Web link of this episode.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Title string `json:"title,omitempty"` // Episode name.
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
}

// PodcastAudienceResponse represents the PodcastAudienceResponse schema from the OpenAPI specification
type PodcastAudienceResponse struct {
	By_regions []map[string]interface{} `json:"by_regions,omitempty"`
}

// CustomAudio represents the CustomAudio schema from the OpenAPI specification
type CustomAudio struct {
	Thumbnail string `json:"thumbnail,omitempty"` // Low resolution image url of this custom audio.
	Title string `json:"title,omitempty"` // Custom audio title.
	Audio string `json:"audio,omitempty"` // Audio url, which can be played directly.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length in seconds.
	Image string `json:"image,omitempty"` // High resolution image url of this custom audio.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date (in milliseconds) of this custom audio. For now, it's the same as **added_at_ms** of this playlist item.
}

// GetCuratedPodcastsResponse represents the GetCuratedPodcastsResponse schema from the OpenAPI specification
type GetCuratedPodcastsResponse struct {
	Total int `json:"total"`
	Curated_lists []CuratedListSimple `json:"curated_lists"`
	Has_next bool `json:"has_next"`
	Has_previous bool `json:"has_previous"`
	Next_page_number int `json:"next_page_number"`
	Page_number int `json:"page_number"`
	Previous_page_number int `json:"previous_page_number"`
}

// GetEpisodesInBatchResponse represents the GetEpisodesInBatchResponse schema from the OpenAPI specification
type GetEpisodesInBatchResponse struct {
	Episodes []EpisodeSimple `json:"episodes"`
}

// CuratedListFull represents the CuratedListFull schema from the OpenAPI specification
type CuratedListFull struct {
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
	Description string `json:"description,omitempty"` // This curated list's description.
	Podcasts []PodcastSimple `json:"podcasts,omitempty"` // Complete meta data of all podcasts in this curated list.
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
	Title string `json:"title,omitempty"` // Curated list name.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
}

// Genre represents the Genre schema from the OpenAPI specification
type Genre struct {
	Name string `json:"name,omitempty"` // Genre name.
	Parent_id int `json:"parent_id,omitempty"` // Parent genre id.
	Id int `json:"id,omitempty"` // Genre id
}

// RelatedSearchesResponse represents the RelatedSearchesResponse schema from the OpenAPI specification
type RelatedSearchesResponse struct {
	Terms []string `json:"terms"` // Related search terms
}

// SpellCheckResponse represents the SpellCheckResponse schema from the OpenAPI specification
type SpellCheckResponse struct {
	Corrected_text_html string `json:"corrected_text_html"` // The corrected text for the entire search term (multiple words/tokens), where misspelled tokens are replaced with the correct texts and html tags <b><i>
	Tokens []map[string]interface{} `json:"tokens"` // The word in the text query string that is not spelled correctly
}

// CuratedListSimple represents the CuratedListSimple schema from the OpenAPI specification
type CuratedListSimple struct {
	Description string `json:"description,omitempty"` // This curated list's description.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
	Podcasts []PodcastMinimum `json:"podcasts,omitempty"` // Minimum meta data of up to 5 podcasts in this curated list.
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Title string `json:"title,omitempty"` // Curated list name.
}

// SubmitPodcastForm represents the SubmitPodcastForm schema from the OpenAPI specification
type SubmitPodcastForm struct {
	Rss string `json:"rss"` // A valid podcast rss url.
	Email string `json:"email,omitempty"` // A valid email address. If **email** is specified, then we'll notify this email address once the podcast is accepted.
}

// PodcastMinimumRss represents the PodcastMinimumRss schema from the OpenAPI specification
type PodcastMinimumRss struct {
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Title string `json:"title,omitempty"` // Podcast name.
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
}

// EpisodeFull represents the EpisodeFull schema from the OpenAPI specification
type EpisodeFull struct {
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Link string `json:"link,omitempty"` // Web link of this episode.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
	Title string `json:"title,omitempty"` // Episode name.
	Podcast PodcastSimple `json:"podcast,omitempty"`
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Transcript string `json:"transcript,omitempty"` // The transcript of this episode, in plain text (with the newline character \n). If there's not transcript, it is null. This field is available only in the PRO/ENTERPRISE plan.
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
}

// PodcastLookingForField represents the PodcastLookingForField schema from the OpenAPI specification
type PodcastLookingForField struct {
	Cohosts bool `json:"cohosts,omitempty"` // Whether this podcast is looking for cohosts.
	Cross_promotion bool `json:"cross_promotion,omitempty"` // Whether this podcast is looking for cross promotion opportunities with other podcasts.
	Guests bool `json:"guests,omitempty"` // Whether this podcast is looking for guests.
	Sponsors bool `json:"sponsors,omitempty"` // Whether this podcast is looking for sponsors.
}

// GetGenresResponse represents the GetGenresResponse schema from the OpenAPI specification
type GetGenresResponse struct {
	Genres []Genre `json:"genres"`
}

// DeletePodcastResponse represents the DeletePodcastResponse schema from the OpenAPI specification
type DeletePodcastResponse struct {
	Status string `json:"status"` // The status of this podcast deletion request.
}

// EpisodeSimple represents the EpisodeSimple schema from the OpenAPI specification
type EpisodeSimple struct {
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
	Title string `json:"title,omitempty"` // Episode name.
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Link string `json:"link,omitempty"` // Web link of this episode.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Podcast PodcastMinimum `json:"podcast,omitempty"`
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
}

// GetLanguagesResponse represents the GetLanguagesResponse schema from the OpenAPI specification
type GetLanguagesResponse struct {
	Languages []string `json:"languages"`
}

// GetEpisodeRecommendationsResponse represents the GetEpisodeRecommendationsResponse schema from the OpenAPI specification
type GetEpisodeRecommendationsResponse struct {
	Recommendations []EpisodeSimple `json:"recommendations"`
}

// TrendingSearchesResponse represents the TrendingSearchesResponse schema from the OpenAPI specification
type TrendingSearchesResponse struct {
	Terms []string `json:"terms"` // Trending search terms
}

// GetPodcastRecommendationsResponse represents the GetPodcastRecommendationsResponse schema from the OpenAPI specification
type GetPodcastRecommendationsResponse struct {
	Recommendations []PodcastSimple `json:"recommendations"`
}

// PodcastDomainResponse represents the PodcastDomainResponse schema from the OpenAPI specification
type PodcastDomainResponse struct {
	Page_number int `json:"page_number,omitempty"`
	Podcasts []PodcastSimple `json:"podcasts,omitempty"`
	Previous_page_number int `json:"previous_page_number,omitempty"`
	Has_next bool `json:"has_next,omitempty"`
	Has_previous bool `json:"has_previous,omitempty"`
	Next_page_number int `json:"next_page_number,omitempty"`
}

// GetPodcastsInBatchForm represents the GetPodcastsInBatchForm schema from the OpenAPI specification
type GetPodcastsInBatchForm struct {
	Itunes_ids string `json:"itunes_ids,omitempty"` // Comma-separated Apple Podcasts (iTunes) ids, e.g., 659155419
	Next_episode_pub_date int `json:"next_episode_pub_date,omitempty"` // For latest episodes pagination. It's the value of **next_episode_pub_date** from the response of last request. If not specified, just return latest 15 episodes.
	Rsses string `json:"rsses,omitempty"` // Comma-separated rss urls.
	Show_latest_episodes int `json:"show_latest_episodes,omitempty"` // Whether or not to fetch up to 15 latest episodes from these podcasts, sorted by pub_date. 1 is yes, and 0 is no.
	Spotify_ids string `json:"spotify_ids,omitempty"` // Comma-separated Spotify ids, e.g., 3DDfEsKDIDrTlnPOiG4ZF4
	Ids string `json:"ids,omitempty"` // Comma-separated list of podcast ids.
}

// PodcastExtraField represents the PodcastExtraField schema from the OpenAPI specification
type PodcastExtraField struct {
	Patreon_handle string `json:"patreon_handle,omitempty"` // Patreon username affiliated with this podcast
	Wechat_handle string `json:"wechat_handle,omitempty"` // WeChat username affiliated with this podcast
	Facebook_handle string `json:"facebook_handle,omitempty"` // Facebook username affiliated with this podcast
	Google_url string `json:"google_url,omitempty"` // Google Podcasts url for this podcast
	Linkedin_url string `json:"linkedin_url,omitempty"` // LinkedIn url affiliated with this podcast
	Twitter_handle string `json:"twitter_handle,omitempty"` // Twitter username affiliated with this podcast
	Url2 string `json:"url2,omitempty"` // Url affiliated with this podcast
	Amazon_music_url string `json:"amazon_music_url,omitempty"` // Amazon Music url for this podcast
	Spotify_url string `json:"spotify_url,omitempty"` // Spotify url for this podcast
	Url1 string `json:"url1,omitempty"` // Url affiliated with this podcast
	Instagram_handle string `json:"instagram_handle,omitempty"` // Instagram username affiliated with this podcast
	Url3 string `json:"url3,omitempty"` // Url affiliated with this podcast
	Youtube_url string `json:"youtube_url,omitempty"` // YouTube url affiliated with this podcast
}

// DeletedItem represents the DeletedItem schema from the OpenAPI specification
type DeletedItem struct {
	Status string `json:"status,omitempty"` // The status of this episode or podcast. For now, the only possible value is **deleted**.
	Title string `json:"title,omitempty"` // Episode title or podcast title.
	ErrorField string `json:"error,omitempty"` // Why this episode or podcast is deleted?
	Id string `json:"id,omitempty"` // Episode id or podcast id.
}

// EpisodeSearchResult represents the EpisodeSearchResult schema from the OpenAPI specification
type EpisodeSearchResult struct {
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of this episode's title
	Transcripts_highlighted []string `json:"transcripts_highlighted,omitempty"` // Up to 2 highlighted segments of the audio transcript of this episode.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Link string `json:"link,omitempty"` // Web link of this episode.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of this episode's description
	Title_original string `json:"title_original,omitempty"` // Plain text of this episode' title
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Description_original string `json:"description_original,omitempty"` // Plain text of this episode's description
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Podcast map[string]interface{} `json:"podcast,omitempty"` // The podcast that this episode belongs to.
}

// GetEpisodesInBatchForm represents the GetEpisodesInBatchForm schema from the OpenAPI specification
type GetEpisodesInBatchForm struct {
	Ids string `json:"ids"` // Comma-separated list of episode ids.
}

// SubmitPodcastResponse represents the SubmitPodcastResponse schema from the OpenAPI specification
type SubmitPodcastResponse struct {
	Podcast PodcastMinimum `json:"podcast"`
	Status string `json:"status"` // The status of this submission.
}

// PodcastMinimum represents the PodcastMinimum schema from the OpenAPI specification
type PodcastMinimum struct {
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Title string `json:"title,omitempty"` // Podcast name.
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
}

// BestPodcastsResponse represents the BestPodcastsResponse schema from the OpenAPI specification
type BestPodcastsResponse struct {
	Page_number int `json:"page_number"`
	Parent_id int `json:"parent_id"` // The id of parent genre.
	Total int `json:"total"`
	Id int `json:"id"` // The id of this genre
	Listennotes_url string `json:"listennotes_url"` // Url of the list of best podcasts on [ListenNotes.com](https://www.ListenNotes.com).
	Next_page_number int `json:"next_page_number"`
	Has_previous bool `json:"has_previous"`
	Name string `json:"name"` // This genre's name.
	Podcasts []PodcastSimple `json:"podcasts"`
	Previous_page_number int `json:"previous_page_number"`
	Has_next bool `json:"has_next"`
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	Next_offset int `json:"next_offset,omitempty"` // Pass this value to the **offset** parameter to do pagination of search results.
	Results []interface{} `json:"results,omitempty"` // A list of search results.
	Took float64 `json:"took,omitempty"` // The time it took to fetch these search results. In seconds.
	Total int `json:"total,omitempty"` // The total number of search results.
	Count int `json:"count,omitempty"` // The number of search results in this page.
}

// PlaylistsResponse represents the PlaylistsResponse schema from the OpenAPI specification
type PlaylistsResponse struct {
	Previous_page_number int `json:"previous_page_number,omitempty"`
	Total int `json:"total,omitempty"`
	Has_next bool `json:"has_next,omitempty"`
	Has_previous bool `json:"has_previous,omitempty"`
	Next_page_number int `json:"next_page_number,omitempty"`
	Page_number int `json:"page_number,omitempty"`
	Playlists []map[string]interface{} `json:"playlists,omitempty"`
}

// PlaylistItem represents the PlaylistItem schema from the OpenAPI specification
type PlaylistItem struct {
	TypeField string `json:"type,omitempty"` // The type of this playlist item. If a playlist is **episode_list**, then an item could be either **episode** or **custom_audio**. If it's **podcast_list**, then an item can only be **podcast**.
	Added_at_ms int `json:"added_at_ms,omitempty"` // Timestamp (in milliseconds) when this item is added.
	Data interface{} `json:"data,omitempty"`
	Id int `json:"id,omitempty"` // Playlist item id.
	Notes string `json:"notes,omitempty"` // Notes for this item.
}

// PodcastTypeaheadResult represents the PodcastTypeaheadResult schema from the OpenAPI specification
type PodcastTypeaheadResult struct {
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Publisher_highlighted string `json:"publisher_highlighted,omitempty"` // Highlighted segment of this podcast's publisher name.
	Publisher_original string `json:"publisher_original,omitempty"` // Plain text of this podcast's publisher name.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of podcast name.
	Title_original string `json:"title_original,omitempty"` // Plain text of podcast name.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
}

// GetRegionsResponse represents the GetRegionsResponse schema from the OpenAPI specification
type GetRegionsResponse struct {
	Regions map[string]interface{} `json:"regions"`
}

// PodcastSimple represents the PodcastSimple schema from the OpenAPI specification
type PodcastSimple struct {
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Genre_ids []int `json:"genre_ids,omitempty"`
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Country string `json:"country,omitempty"` // The country where this podcast is produced.
	Extra PodcastExtraField `json:"extra,omitempty"`
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Looking_for PodcastLookingForField `json:"looking_for,omitempty"`
	Is_claimed bool `json:"is_claimed,omitempty"` // Whether this podcast is claimed by its producer on [ListenNotes.com](https://www.ListenNotes.com).
	Language string `json:"language,omitempty"` // The language of this podcast. You can get all supported languages from `GET /languages`.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Title string `json:"title,omitempty"` // Podcast name.
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	TypeField string `json:"type,omitempty"` // The type of this podcast - episodic or serial.
}

// PodcastFull represents the PodcastFull schema from the OpenAPI specification
type PodcastFull struct {
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Extra PodcastExtraField `json:"extra,omitempty"`
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Title string `json:"title,omitempty"` // Podcast name.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Language string `json:"language,omitempty"` // The language of this podcast. You can get all supported languages from `GET /languages`.
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Country string `json:"country,omitempty"` // The country where this podcast is produced.
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Episodes []EpisodeMinimum `json:"episodes,omitempty"`
	Looking_for PodcastLookingForField `json:"looking_for,omitempty"`
	Is_claimed bool `json:"is_claimed,omitempty"` // Whether this podcast is claimed by its producer on [ListenNotes.com](https://www.ListenNotes.com).
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Genre_ids []int `json:"genre_ids,omitempty"`
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Next_episode_pub_date int `json:"next_episode_pub_date,omitempty"` // Passed to the **next_episode_pub_date** parameter of `GET /podcasts/{id}` to paginate through episodes of that podcast.
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	TypeField string `json:"type,omitempty"` // The type of this podcast - episodic or serial.
	Description string `json:"description,omitempty"` // Html of this episode's full description
}

// GetPodcastsInBatchResponse represents the GetPodcastsInBatchResponse schema from the OpenAPI specification
type GetPodcastsInBatchResponse struct {
	Latest_episodes []EpisodeSimple `json:"latest_episodes,omitempty"` // Up to 10 latest episodes from these podcasts, sorted by **pub_date**. This field shows up only when **show_latest_episodes** is 1.
	Podcasts []PodcastSimple `json:"podcasts"`
}

// CuratedListSearchResult represents the CuratedListSearchResult schema from the OpenAPI specification
type CuratedListSearchResult struct {
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Podcasts []PodcastMinimum `json:"podcasts,omitempty"` // Up to 5 podcasts in this curated list.
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of this curated list's description
	Description_original string `json:"description_original,omitempty"` // Plain text of this curated list's description
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of this curated list's title
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Title_original string `json:"title_original,omitempty"` // Plain text of this curated list's title
}

// TypeaheadResponse represents the TypeaheadResponse schema from the OpenAPI specification
type TypeaheadResponse struct {
	Genres []Genre `json:"genres,omitempty"` // Genre suggestions. It'll show up when the **show_genres** parameter is *1*.
	Podcasts []PodcastTypeaheadResult `json:"podcasts,omitempty"` // Podcast suggestions. It'll show up when the **show_podcasts** parameter is 1.
	Terms []string `json:"terms"` // Search term suggestions.
}

// PodcastSearchResult represents the PodcastSearchResult schema from the OpenAPI specification
type PodcastSearchResult struct {
	Description_original string `json:"description_original,omitempty"` // Plain text of podcast description
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Publisher_highlighted string `json:"publisher_highlighted,omitempty"` // Highlighted segment of this podcast's publisher name.
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of podcast description
	Publisher_original string `json:"publisher_original,omitempty"` // Plain text of this podcast's publisher name.
	Title_original string `json:"title_original,omitempty"` // Plain text of podcast name.
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of podcast name.
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Genre_ids []int `json:"genre_ids,omitempty"`
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
}

// PlaylistResponse represents the PlaylistResponse schema from the OpenAPI specification
type PlaylistResponse struct {
	Total int `json:"total,omitempty"` // Total number of items in this playlist.
	Description string `json:"description,omitempty"` // Playlist description.
	Id string `json:"id,omitempty"` // A 11-character playlist id, which can be used to further fetch detailed playlist metadata via `GET /playlists/{id}`.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this playlist on ListenNotes.com.
	Visibility string `json:"visibility,omitempty"` // Visibility of this playlist.
	Image string `json:"image,omitempty"` // High resolution image url of the playlist.
	Items []PlaylistItem `json:"items,omitempty"` // A list of playlist items.
	Name string `json:"name,omitempty"` // Playlist name.
	Thumbnail string `json:"thumbnail,omitempty"` // Low resolution image url of the playlist.
	Total_audio_length_sec int `json:"total_audio_length_sec,omitempty"` // Total audio length of all episodes in this playlist, in seconds. It will have a valid value only when type is **episode_list**. In other words, it will be 0 if type is **podcast_list**.
	TypeField string `json:"type,omitempty"` // The type of this playlist, which should be either **episode_list** or **podcast_list**.
	Last_timestamp_ms int `json:"last_timestamp_ms,omitempty"` // Passed to the **last_timestamp_ms** parameter of `GET /playlists/{id}` to paginate through items of that playlist.
}
