package vimeo

import (
	"time"
)

type PostResponse struct {
	URI      string `json:"uri"`
	Active   bool   `json:"active"`
	Type     string `json:"type"`
	BaseLink string `json:"base_link"`
	Sizes    []struct {
		Width              int    `json:"width"`
		Height             int    `json:"height"`
		Link               string `json:"link"`
		LinkWithPlayButton string `json:"link_with_play_button"`
	} `json:"sizes"`
	Link           string `json:"link"`
	ResourceKey    string `json:"resource_key"`
	DefaultPicture bool   `json:"default_picture"`
}

type UploadVideoResponse struct {
	URI            string      `json:"uri"`
	Name           string      `json:"name"`
	Description    interface{} `json:"description"`
	Type           string      `json:"type"`
	Link           string      `json:"link"`
	PlayerEmbedURL string      `json:"player_embed_url"`
	Duration       int         `json:"duration"`
	Width          int         `json:"width"`
	Language       interface{} `json:"language"`
	Height         int         `json:"height"`
	Embed          struct {
		HTML   string `json:"html"`
		Badges struct {
			Hdr  bool `json:"hdr"`
			Live struct {
				Streaming bool `json:"streaming"`
				Archived  bool `json:"archived"`
			} `json:"live"`
			StaffPick struct {
				Normal         bool `json:"normal"`
				BestOfTheMonth bool `json:"best_of_the_month"`
				BestOfTheYear  bool `json:"best_of_the_year"`
				Premiere       bool `json:"premiere"`
			} `json:"staff_pick"`
			Vod              bool `json:"vod"`
			WeekendChallenge bool `json:"weekend_challenge"`
		} `json:"badges"`
		Buttons struct {
			Watchlater bool `json:"watchlater"`
			Share      bool `json:"share"`
			Embed      bool `json:"embed"`
			Hd         bool `json:"hd"`
			Fullscreen bool `json:"fullscreen"`
			Scaling    bool `json:"scaling"`
			Like       bool `json:"like"`
		} `json:"buttons"`
		Logos struct {
			Vimeo  bool `json:"vimeo"`
			Custom struct {
				Active bool        `json:"active"`
				URL    interface{} `json:"url"`
				Link   interface{} `json:"link"`
				Sticky bool        `json:"sticky"`
			} `json:"custom"`
		} `json:"logos"`
		Title struct {
			Name     string `json:"name"`
			Owner    string `json:"owner"`
			Portrait string `json:"portrait"`
		} `json:"title"`
		EndScreen     []interface{} `json:"end_screen"`
		Playbar       bool          `json:"playbar"`
		Volume        bool          `json:"volume"`
		Color         string        `json:"color"`
		EventSchedule bool          `json:"event_schedule"`
		Interactive   bool          `json:"interactive"`
		URI           interface{}   `json:"uri"`
		Speed         bool          `json:"speed"`
	} `json:"embed"`
	CreatedTime        time.Time   `json:"created_time"`
	ModifiedTime       time.Time   `json:"modified_time"`
	ReleaseTime        time.Time   `json:"release_time"`
	ContentRating      []string    `json:"content_rating"`
	ContentRatingClass string      `json:"content_rating_class"`
	RatingModLocked    bool        `json:"rating_mod_locked"`
	License            interface{} `json:"license"`
	Privacy            struct {
		View     string `json:"view"`
		Embed    string `json:"embed"`
		Download bool   `json:"download"`
		Add      bool   `json:"add"`
		Comments string `json:"comments"`
	} `json:"privacy"`
	Pictures struct {
		URI      interface{} `json:"uri"`
		Active   bool        `json:"active"`
		Type     string      `json:"type"`
		BaseLink string      `json:"base_link"`
		Sizes    []struct {
			Width              int    `json:"width"`
			Height             int    `json:"height"`
			Link               string `json:"link"`
			LinkWithPlayButton string `json:"link_with_play_button"`
		} `json:"sizes"`
		ResourceKey    string `json:"resource_key"`
		DefaultPicture bool   `json:"default_picture"`
	} `json:"pictures"`
	Tags  []interface{} `json:"tags"`
	Stats struct {
		Plays int `json:"plays"`
	} `json:"stats"`
	Categories []interface{} `json:"categories"`
	Uploader   struct {
		Pictures struct {
			URI      interface{} `json:"uri"`
			Active   bool        `json:"active"`
			Type     string      `json:"type"`
			BaseLink string      `json:"base_link"`
			Sizes    []struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				Link   string `json:"link"`
			} `json:"sizes"`
			ResourceKey    string `json:"resource_key"`
			DefaultPicture bool   `json:"default_picture"`
		} `json:"pictures"`
	} `json:"uploader"`
	Metadata struct {
		Connections struct {
			Comments struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Total   int      `json:"total"`
			} `json:"comments"`
			Credits struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Total   int      `json:"total"`
			} `json:"credits"`
			Likes struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Total   int      `json:"total"`
			} `json:"likes"`
			Pictures struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Total   int      `json:"total"`
			} `json:"pictures"`
			Texttracks struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Total   int      `json:"total"`
			} `json:"texttracks"`
			Related         interface{} `json:"related"`
			Recommendations struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
			} `json:"recommendations"`
			Albums struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Total   int      `json:"total"`
			} `json:"albums"`
			AvailableAlbums struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Total   int      `json:"total"`
			} `json:"available_albums"`
			AvailableChannels struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Total   int      `json:"total"`
			} `json:"available_channels"`
			Versions struct {
				URI                     string      `json:"uri"`
				Options                 []string    `json:"options"`
				Total                   int         `json:"total"`
				LatestIncompleteVersion interface{} `json:"latest_incomplete_version"`
			} `json:"versions"`
		} `json:"connections"`
		Interactions struct {
			Watchlater struct {
				URI       string      `json:"uri"`
				Options   []string    `json:"options"`
				Added     bool        `json:"added"`
				AddedTime interface{} `json:"added_time"`
			} `json:"watchlater"`
			Report struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
				Reason  []string `json:"reason"`
			} `json:"report"`
			ViewTeamMembers struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
			} `json:"view_team_members"`
			Edit struct {
				URI           string        `json:"uri"`
				Options       []string      `json:"options"`
				BlockedFields []interface{} `json:"blocked_fields"`
			} `json:"edit"`
			EditContentRating struct {
				URI           string   `json:"uri"`
				Options       []string `json:"options"`
				ContentRating []string `json:"content_rating"`
			} `json:"edit_content_rating"`
			EditPrivacy struct {
				URI         string   `json:"uri"`
				Options     []string `json:"options"`
				ContentType string   `json:"content_type"`
				Properties  []struct {
					Name     string   `json:"name"`
					Required bool     `json:"required"`
					Options  []string `json:"options"`
				} `json:"properties"`
			} `json:"edit_privacy"`
			Delete struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
			} `json:"delete"`
			CanUpdatePrivacyToPublic struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
			} `json:"can_update_privacy_to_public"`
			Trim struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
			} `json:"trim"`
			Validate struct {
				URI     string   `json:"uri"`
				Options []string `json:"options"`
			} `json:"validate"`
		} `json:"interactions"`
		IsVimeoCreate  bool `json:"is_vimeo_create"`
		IsScreenRecord bool `json:"is_screen_record"`
	} `json:"metadata"`
	ManageLink string `json:"manage_link"`
	User       struct {
		URI          string `json:"uri"`
		Name         string `json:"name"`
		Link         string `json:"link"`
		Capabilities struct {
			HasLiveSubscription     bool `json:"hasLiveSubscription"`
			HasEnterpriseLihp       bool `json:"hasEnterpriseLihp"`
			HasSvvTimecodedComments bool `json:"hasSvvTimecodedComments"`
		} `json:"capabilities"`
		Location    string      `json:"location"`
		Gender      string      `json:"gender"`
		Bio         interface{} `json:"bio"`
		ShortBio    interface{} `json:"short_bio"`
		CreatedTime time.Time   `json:"created_time"`
		Pictures    struct {
			URI      interface{} `json:"uri"`
			Active   bool        `json:"active"`
			Type     string      `json:"type"`
			BaseLink string      `json:"base_link"`
			Sizes    []struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				Link   string `json:"link"`
			} `json:"sizes"`
			ResourceKey    string `json:"resource_key"`
			DefaultPicture bool   `json:"default_picture"`
		} `json:"pictures"`
		Websites []interface{} `json:"websites"`
		Metadata struct {
			Connections struct {
				Albums struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"albums"`
				Appearances struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"appearances"`
				Categories struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"categories"`
				Channels struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"channels"`
				Feed struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
				} `json:"feed"`
				Followers struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"followers"`
				Following struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"following"`
				Groups struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"groups"`
				Likes struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"likes"`
				Membership struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
				} `json:"membership"`
				ModeratedChannels struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"moderated_channels"`
				Portfolios struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"portfolios"`
				Videos struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"videos"`
				Watchlater struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"watchlater"`
				Shared struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"shared"`
				Pictures struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"pictures"`
				WatchedVideos struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"watched_videos"`
				FoldersRoot struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
				} `json:"folders_root"`
				Folders struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"folders"`
				Teams struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"teams"`
				PermissionPolicies struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"permission_policies"`
				Block struct {
					URI     string   `json:"uri"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
				} `json:"block"`
			} `json:"connections"`
		} `json:"metadata"`
		LocationDetails struct {
			FormattedAddress string      `json:"formatted_address"`
			Latitude         interface{} `json:"latitude"`
			Longitude        interface{} `json:"longitude"`
			City             interface{} `json:"city"`
			State            interface{} `json:"state"`
			Neighborhood     interface{} `json:"neighborhood"`
			SubLocality      interface{} `json:"sub_locality"`
			StateIsoCode     interface{} `json:"state_iso_code"`
			Country          interface{} `json:"country"`
			CountryIsoCode   interface{} `json:"country_iso_code"`
		} `json:"location_details"`
		Skills           []interface{} `json:"skills"`
		AvailableForHire bool          `json:"available_for_hire"`
		CanWorkRemotely  bool          `json:"can_work_remotely"`
		Preferences      struct {
			Videos struct {
				Rating  []string `json:"rating"`
				Privacy struct {
					View     string `json:"view"`
					Comments string `json:"comments"`
					Embed    string `json:"embed"`
					Download bool   `json:"download"`
					Add      bool   `json:"add"`
				} `json:"privacy"`
			} `json:"videos"`
		} `json:"preferences"`
		ContentFilter []string `json:"content_filter"`
		UploadQuota   struct {
			Space struct {
				Free    int64  `json:"free"`
				Max     int64  `json:"max"`
				Used    int    `json:"used"`
				Showing string `json:"showing"`
				Unit    string `json:"unit"`
			} `json:"space"`
			Periodic struct {
				Period    interface{} `json:"period"`
				Unit      interface{} `json:"unit"`
				Free      interface{} `json:"free"`
				Max       interface{} `json:"max"`
				Used      interface{} `json:"used"`
				ResetDate interface{} `json:"reset_date"`
			} `json:"periodic"`
			Lifetime struct {
				Unit string `json:"unit"`
				Free int64  `json:"free"`
				Max  int64  `json:"max"`
				Used int    `json:"used"`
			} `json:"lifetime"`
		} `json:"upload_quota"`
		ResourceKey string `json:"resource_key"`
		Account     string `json:"account"`
	} `json:"user"`
	ParentFolder            interface{} `json:"parent_folder"`
	LastUserActionEventDate time.Time   `json:"last_user_action_event_date"`
	ReviewPage              struct {
		Active bool   `json:"active"`
		Link   string `json:"link"`
	} `json:"review_page"`
	Files    []interface{} `json:"files"`
	Download []interface{} `json:"download"`
	Play     struct {
		Status string `json:"status"`
	} `json:"play"`
	App struct {
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"app"`
	Status      string `json:"status"`
	ResourceKey string `json:"resource_key"`
	Upload      struct {
		Status      string      `json:"status"`
		UploadLink  string      `json:"upload_link"`
		Form        interface{} `json:"form"`
		CompleteURI interface{} `json:"complete_uri"`
		Approach    string      `json:"approach"`
		Size        int         `json:"size"`
		RedirectURL interface{} `json:"redirect_url"`
		Link        interface{} `json:"link"`
	} `json:"upload"`
	Transcode struct {
		Status string `json:"status"`
	} `json:"transcode"`
	IsPlayable bool `json:"is_playable"`
	HasAudio   bool `json:"has_audio"`
}

type CreateLiveResponse struct {
	Album                    []string      `json:"album"`
	AutoCcEnabled            string        `json:"auto_cc_enabled"`
	AutoCcKeywords           string        `json:"auto_cc_keywords"`
	AutoCcLanguage           string        `json:"auto_cc_language"`
	AutoCcRemaining          string        `json:"auto_cc_remaining"`
	AutomaticallyTitleStream bool          `json:"automatically_title_stream"`
	ChatEnabled              bool          `json:"chat_enabled"`
	ContentRating            []interface{} `json:"content_rating"`
	CreatedTime              time.Time     `json:"created_time"`
	DashLink                 string        `json:"dash_link"`
	Embed                    struct {
		Autoplay             bool     `json:"autoplay"`
		AvailablePlayerLogos []string `json:"available_player_logos"`
		Byline               bool     `json:"byline"`
		ChatEmbedSource      string   `json:"chat_embed_source"`
		Color                string   `json:"color"`
		EmbedChat            string   `json:"embed_chat"`
		EmbedProperties      []string `json:"embed_properties"`
		EventSchedule        bool     `json:"event_schedule"`
		FullscreenButton     bool     `json:"fullscreen_button"`
		HideLiveLabel        bool     `json:"hide_live_label"`
		HideViewerCount      bool     `json:"hide_viewer_count"`
		HTML                 string   `json:"html"`
		LikeButton           bool     `json:"like_button"`
		Logos                struct {
			Custom struct {
				Active  bool   `json:"active"`
				Link    string `json:"link"`
				Sticky  bool   `json:"sticky"`
				URL     string `json:"url"`
				UseLink bool   `json:"use_link"`
			} `json:"custom"`
			Vimeo bool `json:"vimeo"`
		} `json:"logos"`
		Loop                   bool   `json:"loop"`
		Playbar                bool   `json:"playbar"`
		Playlist               bool   `json:"playlist"`
		Portrait               bool   `json:"portrait"`
		ResponsiveHTML         string `json:"responsiveHtml"`
		Schedule               bool   `json:"schedule"`
		ShowLatestArchivedClip bool   `json:"show_latest_archived_clip"`
		Title                  bool   `json:"title"`
		UseColor               string `json:"use_color"`
		Volume                 bool   `json:"volume"`
	} `json:"embed"`
	FromShowcase     string        `json:"from_showcase"`
	FromWebinar      string        `json:"from_webinar"`
	HeadClip         []interface{} `json:"head_clip"`
	Link             string        `json:"link"`
	LiveClips        []string      `json:"live_clips"`
	LiveDestinations []string      `json:"live_destinations"`
	LowLatency       bool          `json:"low_latency"`
	Metadata         struct {
		Connections struct {
			LiveVideo struct {
				Options []string `json:"options"`
				Status  string   `json:"status"`
				URI     string   `json:"uri"`
			} `json:"live_video"`
			Pictures struct {
				Options []string `json:"options"`
				URI     string   `json:"uri"`
			} `json:"pictures"`
			PreLiveVideo struct {
				Options []string `json:"options"`
				Status  string   `json:"status"`
				URI     string   `json:"uri"`
			} `json:"pre_live_video"`
			TeamMember struct {
				Options []string `json:"options"`
				URI     string   `json:"uri"`
			} `json:"team_member"`
			Videos struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"videos"`
		} `json:"connections"`
		Interactions struct {
			Activate struct {
				Options []string `json:"options"`
				URI     string   `json:"uri"`
			} `json:"activate"`
			Delete struct {
				Options []string `json:"options"`
				URI     string   `json:"uri"`
			} `json:"delete"`
			Edit struct {
				Options []string `json:"options"`
				URI     string   `json:"uri"`
			} `json:"edit"`
		} `json:"interactions"`
	} `json:"metadata"`
	NextOccurrenceTime time.Time `json:"next_occurrence_time"`
	ParentFolder       struct {
		AccessGrant             []interface{} `json:"access_grant"`
		Color                   string        `json:"color"`
		CreatedTime             time.Time     `json:"created_time"`
		CreatorURI              string        `json:"creator_uri"`
		FolderURI               string        `json:"folder_uri"`
		IsPinned                bool          `json:"is_pinned"`
		IsPrivateToUser         bool          `json:"is_private_to_user"`
		LastUserActionEventDate time.Time     `json:"last_user_action_event_date"`
		Link                    string        `json:"link"`
		Metadata                struct {
			Connections struct {
				AncestorPath []struct {
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"ancestor_path"`
				DataRetention struct {
					Policy struct {
						Options []string `json:"options"`
						Title   string   `json:"title"`
						URI     string   `json:"uri"`
					} `json:"policy"`
				} `json:"data_retention"`
				Folders struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"folders"`
				GroupFolderGrants struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"group_folder_grants"`
				Items struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"items"`
				ParentFolder struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"parent_folder"`
				TeamGroups struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"team_groups"`
				TeamMembers struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"team_members"`
				TeamPermissions struct {
					Options []string `json:"options"`
				} `json:"team_permissions"`
				UserFolderAccessGrants struct {
					FolderPermissionPolicies []struct {
						Name string `json:"name"`
						URI  string `json:"uri"`
					} `json:"folder_permission_policies"`
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"user_folder_access_grants"`
				Videos struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"videos"`
			} `json:"connections"`
			Interactions struct {
				AddSubfolder struct {
					CanAddSubfolders           string   `json:"can_add_subfolders"`
					ContentType                string   `json:"content_type"`
					Options                    []string `json:"options"`
					Properties                 []string `json:"properties"`
					SubfolderDepthLimitReached string   `json:"subfolder_depth_limit_reached"`
					URI                        string   `json:"uri"`
				} `json:"add_subfolder"`
				Delete struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"delete"`
				DeleteVideo struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"delete_video"`
				Edit struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"edit"`
				EditSettings struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"edit_settings"`
				Invite struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"invite"`
				View struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"view"`
			} `json:"interactions"`
		} `json:"metadata"`
		ModifiedTime time.Time `json:"modified_time"`
		Name         string    `json:"name"`
		PinnedOn     time.Time `json:"pinned_on"`
		Privacy      struct {
			View string `json:"view"`
		} `json:"privacy"`
		ResourceKey             string        `json:"resource_key"`
		SlackIncomingWebhooksID int           `json:"slack_incoming_webhooks_id"`
		SlackIntegrationChannel string        `json:"slack_integration_channel"`
		SlackLanguagePreference string        `json:"slack_language_preference"`
		SlackUserPreferences    []interface{} `json:"slack_user_preferences"`
		URI                     string        `json:"uri"`
		User                    struct {
			Account string `json:"account"`
			APIApps struct {
				Capabilities []string `json:"capabilities"`
			} `json:"api_apps"`
			AvailableForHire bool     `json:"available_for_hire"`
			BackgroundVideo  []string `json:"background_video"`
			Badge            struct {
				AltText string `json:"alt_text"`
				Text    string `json:"text"`
				Type    string `json:"type"`
				URL     string `json:"url"`
			} `json:"badge"`
			Bio             string        `json:"bio"`
			CanWorkRemotely bool          `json:"can_work_remotely"`
			Capabilities    []interface{} `json:"capabilities"`
			Clients         string        `json:"clients"`
			ContentFilter   []interface{} `json:"content_filter"`
			CreatedTime     time.Time     `json:"created_time"`
			Email           string        `json:"email"`
			EmailPrefs      []struct {
				Cast                     bool `json:"cast"`
				CollectionAdds           bool `json:"collection_adds"`
				CollectionClipDeletion   bool `json:"collection_clip_deletion"`
				Comments                 bool `json:"comments"`
				CreatePreviewReady       bool `json:"create_preview_ready"`
				DailyUpdate              bool `json:"daily_update"`
				Following                bool `json:"following"`
				JobShares                bool `json:"job_shares"`
				Like                     bool `json:"like"`
				Marketing                bool `json:"marketing"`
				Mentions                 bool `json:"mentions"`
				Messages                 bool `json:"messages"`
				Newsletter               bool `json:"newsletter"`
				OnlyFromFollowing        bool `json:"only_from_following"`
				Product                  bool `json:"product"`
				PurchaseVod              bool `json:"purchase_vod"`
				ReviewNotifications      bool `json:"review_notifications"`
				Shares                   bool `json:"shares"`
				Videos                   bool `json:"videos"`
				Vod                      bool `json:"vod"`
				VodEpisodeAdded          bool `json:"vod_episode_added"`
				VodFrequentPerformance   bool `json:"vod_frequent_performance"`
				VodInfrequentPerformance bool `json:"vod_infrequent_performance"`
			} `json:"email_prefs"`
			Emails                 []string      `json:"emails"`
			EntitlementUploadQuota []interface{} `json:"entitlement_upload_quota"`
			Gender                 string        `json:"gender"`
			HasInvalidEmail        string        `json:"has_invalid_email"`
			IsExpert               bool          `json:"is_expert"`
			LastActiveTime         time.Time     `json:"last_active_time"`
			Link                   string        `json:"link"`
			LiveQuota              struct {
				RemoteGuests int    `json:"remote_guests"`
				Status       string `json:"status"`
				Streams      struct {
					Maximum   int `json:"maximum"`
					Remaining int `json:"remaining"`
				} `json:"streams"`
				StudioSeats int `json:"studio_seats"`
				Time        struct {
					EventMaximum     int `json:"event_maximum"`
					MonthlyMaximum   int `json:"monthly_maximum"`
					MonthlyRemaining int `json:"monthly_remaining"`
				} `json:"time"`
			} `json:"live_quota"`
			Location        string `json:"location"`
			LocationDetails struct {
				City             string  `json:"city"`
				Country          string  `json:"country"`
				CountryIsoCode   string  `json:"country_iso_code"`
				FormattedAddress string  `json:"formatted_address"`
				Latitude         float64 `json:"latitude"`
				Longitude        float64 `json:"longitude"`
				Neighborhood     string  `json:"neighborhood"`
				State            string  `json:"state"`
				StateIsoCode     string  `json:"state_iso_code"`
				SubLocality      string  `json:"sub_locality"`
			} `json:"location_details"`
			Metadata struct {
				Connections struct {
					Albums struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"albums"`
					Appearances struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"appearances"`
					Block struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"block"`
					Categories struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"categories"`
					Channels struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"channels"`
					Clients struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"clients"`
					Collaborators struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"collaborators"`
					ConnectedApps struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"connected_apps"`
					Feed struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"feed"`
					Folders struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"folders"`
					FoldersRoot struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"folders_root"`
					Followers struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"followers"`
					Following struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"following"`
					Groups struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"groups"`
					JwtAuth struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"jwt_auth"`
					Likes struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"likes"`
					ManagedUser struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"managed_user"`
					Membership struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"membership"`
					ModeratedChannels struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"moderated_channels"`
					Notifications struct {
						NewTotal  int `json:"new_total"`
						Total     int `json:"total"`
						TotalNew  int `json:"total_new"`
						TypeCount struct {
							AccountExpirationWarning   int `json:"account_expiration_warning"`
							Comment                    int `json:"comment"`
							Credit                     int `json:"credit"`
							Follow                     int `json:"follow"`
							FollowedUserVideoAvailable int `json:"followed_user_video_available"`
							Like                       int `json:"like"`
							Mention                    int `json:"mention"`
							Reply                      int `json:"reply"`
							Share                      int `json:"share"`
							StorageWarning             int `json:"storage_warning"`
							VideoAvailable             int `json:"video_available"`
							VodPreorderAvailable       int `json:"vod_preorder_available"`
							VodPurchase                int `json:"vod_purchase"`
							VodRentalExpirationWarning int `json:"vod_rental_expiration_warning"`
						} `json:"type_count"`
						TypeUnseenCount struct {
							AccountExpirationWarning   int `json:"account_expiration_warning"`
							Comment                    int `json:"comment"`
							Credit                     int `json:"credit"`
							Follow                     int `json:"follow"`
							FollowedUserVideoAvailable int `json:"followed_user_video_available"`
							Like                       int `json:"like"`
							Mention                    int `json:"mention"`
							Reply                      int `json:"reply"`
							Share                      int `json:"share"`
							StorageWarning             int `json:"storage_warning"`
							VideoAvailable             int `json:"video_available"`
							VodPreorderAvailable       int `json:"vod_preorder_available"`
							VodPurchase                int `json:"vod_purchase"`
							VodRentalExpirationWarning int `json:"vod_rental_expiration_warning"`
						} `json:"type_unseen_count"`
						UnreadTotal int    `json:"unread_total"`
						URI         string `json:"uri"`
					} `json:"notifications"`
					PermissionPolicies struct {
						Count   string   `json:"count"`
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"permission_policies"`
					Pictures struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"pictures"`
					Portfolios struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"portfolios"`
					ProjectTypes struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"project_types"`
					RecommendedChannels struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"recommended_channels"`
					RecommendedUsers struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"recommended_users"`
					Shared struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"shared"`
					Skills struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"skills"`
					TeamMembers struct {
						InvitesRemaining int      `json:"invites_remaining"`
						MaxSeats         int      `json:"max_seats"`
						Options          []string `json:"options"`
						Total            int      `json:"total"`
						URI              string   `json:"uri"`
					} `json:"team_members"`
					Teams struct {
						Count   string   `json:"count"`
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"teams"`
					Videos struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"videos"`
					VimeoExperts struct {
						IsEnrolled bool     `json:"is_enrolled"`
						Options    []string `json:"options"`
						URI        string   `json:"uri"`
					} `json:"vimeo_experts"`
					Violations struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"violations"`
					WatchedVideos struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"watched_videos"`
					Watchlater struct {
						Options []string `json:"options"`
						Total   int      `json:"total"`
						URI     string   `json:"uri"`
					} `json:"watchlater"`
					WebexUploadFolder struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"webex_upload_folder"`
					ZoomUploadFolder struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"zoom_upload_folder"`
				} `json:"connections"`
				Interactions struct {
					AddPrivacyUser struct {
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"add_privacy_user"`
					Block struct {
						Added     bool      `json:"added"`
						AddedTime time.Time `json:"added_time"`
						Options   []string  `json:"options"`
						URI       string    `json:"uri"`
					} `json:"block"`
					ConnectedApps struct {
						AllScopes    []interface{} `json:"all_scopes"`
						IsConnected  bool          `json:"is_connected"`
						NeededScopes []interface{} `json:"needed_scopes"`
						Options      []string      `json:"options"`
						URI          string        `json:"uri"`
					} `json:"connected_apps"`
					Follow struct {
						Added   bool     `json:"added"`
						Options []string `json:"options"`
						URI     string   `json:"uri"`
					} `json:"follow"`
					Report struct {
						Options []string `json:"options"`
						Reason  []string `json:"reason"`
						URI     string   `json:"uri"`
					} `json:"report"`
				} `json:"interactions"`
				PublicVideos struct {
					Total int `json:"total"`
				} `json:"public_videos"`
			} `json:"metadata"`
			Name         string `json:"name"`
			PartnerAppID int    `json:"partner_app_id"`
			Payment      struct {
				Active bool `json:"active"`
				Cc     struct {
					ExpirationDate string `json:"expiration_date"`
					LastFour       string `json:"last_four"`
					PostalCode     string `json:"postal_code"`
					Type           string `json:"type"`
				} `json:"cc"`
				Type string `json:"type"`
			} `json:"payment"`
			Pictures struct {
				Active         bool   `json:"active"`
				BaseLink       string `json:"base_link"`
				DefaultPicture bool   `json:"default_picture"`
				Link           string `json:"link"`
				ResourceKey    string `json:"resource_key"`
				Sizes          []struct {
					Height             int    `json:"height"`
					Link               string `json:"link"`
					LinkWithPlayButton string `json:"link_with_play_button"`
					Width              int    `json:"width"`
				} `json:"sizes"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"pictures"`
			Preferences struct {
				Videos struct {
					Privacy struct {
						Password string `json:"password"`
					} `json:"privacy"`
				} `json:"videos"`
			} `json:"preferences"`
			PrimaryEmail       string `json:"primary_email"`
			ProfilePreferences struct {
				Layout                    string `json:"layout"`
				ProfileType               string `json:"profile_type"`
				ShouldAutoAddVideos       bool   `json:"should_auto_add_videos"`
				ShowAddVideoTip           bool   `json:"show_add_video_tip"`
				ShowClients               bool   `json:"show_clients"`
				ShowJoinVimeoExpertsModal bool   `json:"show_join_vimeo_experts_modal"`
				ShowProfileTypeTip        bool   `json:"show_profile_type_tip"`
				ShowRate                  bool   `json:"show_rate"`
			} `json:"profile_preferences"`
			ResourceKey string `json:"resource_key"`
			ShortBio    string `json:"short_bio"`
			Skills      struct {
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"skills"`
			TotalCollectionCount int `json:"total_collection_count"`
			UploadQuota          struct {
				Lifetime struct {
					Free int    `json:"free"`
					Max  int64  `json:"max"`
					Unit string `json:"unit"`
					Used int64  `json:"used"`
				} `json:"lifetime"`
				Periodic struct {
					Free      int    `json:"free"`
					Max       int64  `json:"max"`
					Period    string `json:"period"`
					ResetDate string `json:"reset_date"`
					Unit      string `json:"unit"`
					Used      int64  `json:"used"`
				} `json:"periodic"`
				Space struct {
					Free     int `json:"free"`
					Lifetime struct {
						Unit string `json:"unit"`
					} `json:"lifetime"`
					Max     int64  `json:"max"`
					Showing string `json:"showing"`
					Used    int64  `json:"used"`
				} `json:"space"`
			} `json:"upload_quota"`
			URI      string `json:"uri"`
			Verified bool   `json:"verified"`
			Websites []struct {
				Description string `json:"description"`
				Link        string `json:"link"`
				Name        string `json:"name"`
				Type        string `json:"type"`
				URI         string `json:"uri"`
			} `json:"websites"`
		} `json:"user"`
	} `json:"parent_folder"`
	Pictures struct {
		Active         bool   `json:"active"`
		BaseLink       string `json:"base_link"`
		DefaultPicture bool   `json:"default_picture"`
		Link           string `json:"link"`
		ResourceKey    string `json:"resource_key"`
		Sizes          []struct {
			Height             int    `json:"height"`
			Link               string `json:"link"`
			LinkWithPlayButton string `json:"link_with_play_button"`
			Width              int    `json:"width"`
		} `json:"sizes"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"pictures"`
	PlaylistSort      string        `json:"playlist_sort"`
	RtmpLink          string        `json:"rtmp_link"`
	RtmpsLink         string        `json:"rtmps_link"`
	Schedule          []interface{} `json:"schedule"`
	StreamDescription string        `json:"stream_description"`
	StreamKey         string        `json:"stream_key"`
	StreamPassword    string        `json:"stream_password"`
	StreamPrivacy     struct {
		Embed        string `json:"embed"`
		UnlistedHash string `json:"unlisted_hash"`
		View         string `json:"view"`
	} `json:"stream_privacy"`
	StreamTitle       string        `json:"stream_title"`
	StreamableClip    []interface{} `json:"streamable_clip"`
	TimeZone          string        `json:"time_zone"`
	Title             string        `json:"title"`
	UnlimitedAutoCc   string        `json:"unlimited_auto_cc"`
	UnlimitedDuration bool          `json:"unlimited_duration"`
	URI               string        `json:"uri"`
	User              struct {
		Account string `json:"account"`
		APIApps struct {
			Capabilities []string `json:"capabilities"`
		} `json:"api_apps"`
		AvailableForHire bool     `json:"available_for_hire"`
		BackgroundVideo  []string `json:"background_video"`
		Badge            struct {
			AltText string `json:"alt_text"`
			Text    string `json:"text"`
			Type    string `json:"type"`
			URL     string `json:"url"`
		} `json:"badge"`
		Bio             string        `json:"bio"`
		CanWorkRemotely bool          `json:"can_work_remotely"`
		Capabilities    []interface{} `json:"capabilities"`
		Clients         string        `json:"clients"`
		ContentFilter   []interface{} `json:"content_filter"`
		CreatedTime     time.Time     `json:"created_time"`
		Email           string        `json:"email"`
		EmailPrefs      []struct {
			Cast                     bool `json:"cast"`
			CollectionAdds           bool `json:"collection_adds"`
			CollectionClipDeletion   bool `json:"collection_clip_deletion"`
			Comments                 bool `json:"comments"`
			CreatePreviewReady       bool `json:"create_preview_ready"`
			DailyUpdate              bool `json:"daily_update"`
			Following                bool `json:"following"`
			JobShares                bool `json:"job_shares"`
			Like                     bool `json:"like"`
			Marketing                bool `json:"marketing"`
			Mentions                 bool `json:"mentions"`
			Messages                 bool `json:"messages"`
			Newsletter               bool `json:"newsletter"`
			OnlyFromFollowing        bool `json:"only_from_following"`
			Product                  bool `json:"product"`
			PurchaseVod              bool `json:"purchase_vod"`
			ReviewNotifications      bool `json:"review_notifications"`
			Shares                   bool `json:"shares"`
			Videos                   bool `json:"videos"`
			Vod                      bool `json:"vod"`
			VodEpisodeAdded          bool `json:"vod_episode_added"`
			VodFrequentPerformance   bool `json:"vod_frequent_performance"`
			VodInfrequentPerformance bool `json:"vod_infrequent_performance"`
		} `json:"email_prefs"`
		Emails                 []string      `json:"emails"`
		EntitlementUploadQuota []interface{} `json:"entitlement_upload_quota"`
		Gender                 string        `json:"gender"`
		HasInvalidEmail        string        `json:"has_invalid_email"`
		IsExpert               bool          `json:"is_expert"`
		LastActiveTime         time.Time     `json:"last_active_time"`
		Link                   string        `json:"link"`
		LiveQuota              struct {
			RemoteGuests int    `json:"remote_guests"`
			Status       string `json:"status"`
			Streams      struct {
				Maximum   int `json:"maximum"`
				Remaining int `json:"remaining"`
			} `json:"streams"`
			StudioSeats int `json:"studio_seats"`
			Time        struct {
				EventMaximum     int `json:"event_maximum"`
				MonthlyMaximum   int `json:"monthly_maximum"`
				MonthlyRemaining int `json:"monthly_remaining"`
			} `json:"time"`
		} `json:"live_quota"`
		Location        string `json:"location"`
		LocationDetails struct {
			City             string  `json:"city"`
			Country          string  `json:"country"`
			CountryIsoCode   string  `json:"country_iso_code"`
			FormattedAddress string  `json:"formatted_address"`
			Latitude         float64 `json:"latitude"`
			Longitude        float64 `json:"longitude"`
			Neighborhood     string  `json:"neighborhood"`
			State            string  `json:"state"`
			StateIsoCode     string  `json:"state_iso_code"`
			SubLocality      string  `json:"sub_locality"`
		} `json:"location_details"`
		Metadata struct {
			Connections struct {
				Albums struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"albums"`
				Appearances struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"appearances"`
				Block struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"block"`
				Categories struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"categories"`
				Channels struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"channels"`
				Clients struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"clients"`
				Collaborators struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"collaborators"`
				ConnectedApps struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"connected_apps"`
				Feed struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"feed"`
				Folders struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"folders"`
				FoldersRoot struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"folders_root"`
				Followers struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"followers"`
				Following struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"following"`
				Groups struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"groups"`
				JwtAuth struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"jwt_auth"`
				Likes struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"likes"`
				ManagedUser struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"managed_user"`
				Membership struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"membership"`
				ModeratedChannels struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"moderated_channels"`
				Notifications struct {
					NewTotal  int `json:"new_total"`
					Total     int `json:"total"`
					TotalNew  int `json:"total_new"`
					TypeCount struct {
						AccountExpirationWarning   int `json:"account_expiration_warning"`
						Comment                    int `json:"comment"`
						Credit                     int `json:"credit"`
						Follow                     int `json:"follow"`
						FollowedUserVideoAvailable int `json:"followed_user_video_available"`
						Like                       int `json:"like"`
						Mention                    int `json:"mention"`
						Reply                      int `json:"reply"`
						Share                      int `json:"share"`
						StorageWarning             int `json:"storage_warning"`
						VideoAvailable             int `json:"video_available"`
						VodPreorderAvailable       int `json:"vod_preorder_available"`
						VodPurchase                int `json:"vod_purchase"`
						VodRentalExpirationWarning int `json:"vod_rental_expiration_warning"`
					} `json:"type_count"`
					TypeUnseenCount struct {
						AccountExpirationWarning   int `json:"account_expiration_warning"`
						Comment                    int `json:"comment"`
						Credit                     int `json:"credit"`
						Follow                     int `json:"follow"`
						FollowedUserVideoAvailable int `json:"followed_user_video_available"`
						Like                       int `json:"like"`
						Mention                    int `json:"mention"`
						Reply                      int `json:"reply"`
						Share                      int `json:"share"`
						StorageWarning             int `json:"storage_warning"`
						VideoAvailable             int `json:"video_available"`
						VodPreorderAvailable       int `json:"vod_preorder_available"`
						VodPurchase                int `json:"vod_purchase"`
						VodRentalExpirationWarning int `json:"vod_rental_expiration_warning"`
					} `json:"type_unseen_count"`
					UnreadTotal int    `json:"unread_total"`
					URI         string `json:"uri"`
				} `json:"notifications"`
				PermissionPolicies struct {
					Count   string   `json:"count"`
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"permission_policies"`
				Pictures struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"pictures"`
				Portfolios struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"portfolios"`
				ProjectTypes struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"project_types"`
				RecommendedChannels struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"recommended_channels"`
				RecommendedUsers struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"recommended_users"`
				Shared struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"shared"`
				Skills struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"skills"`
				TeamMembers struct {
					InvitesRemaining int      `json:"invites_remaining"`
					MaxSeats         int      `json:"max_seats"`
					Options          []string `json:"options"`
					Total            int      `json:"total"`
					URI              string   `json:"uri"`
				} `json:"team_members"`
				Teams struct {
					Count   string   `json:"count"`
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"teams"`
				Videos struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"videos"`
				VimeoExperts struct {
					IsEnrolled bool     `json:"is_enrolled"`
					Options    []string `json:"options"`
					URI        string   `json:"uri"`
				} `json:"vimeo_experts"`
				Violations struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"violations"`
				WatchedVideos struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"watched_videos"`
				Watchlater struct {
					Options []string `json:"options"`
					Total   int      `json:"total"`
					URI     string   `json:"uri"`
				} `json:"watchlater"`
				WebexUploadFolder struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"webex_upload_folder"`
				ZoomUploadFolder struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"zoom_upload_folder"`
			} `json:"connections"`
			Interactions struct {
				AddPrivacyUser struct {
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"add_privacy_user"`
				Block struct {
					Added     bool      `json:"added"`
					AddedTime time.Time `json:"added_time"`
					Options   []string  `json:"options"`
					URI       string    `json:"uri"`
				} `json:"block"`
				ConnectedApps struct {
					AllScopes    []interface{} `json:"all_scopes"`
					IsConnected  bool          `json:"is_connected"`
					NeededScopes []interface{} `json:"needed_scopes"`
					Options      []string      `json:"options"`
					URI          string        `json:"uri"`
				} `json:"connected_apps"`
				Follow struct {
					Added   bool     `json:"added"`
					Options []string `json:"options"`
					URI     string   `json:"uri"`
				} `json:"follow"`
				Report struct {
					Options []string `json:"options"`
					Reason  []string `json:"reason"`
					URI     string   `json:"uri"`
				} `json:"report"`
			} `json:"interactions"`
			PublicVideos struct {
				Total int `json:"total"`
			} `json:"public_videos"`
		} `json:"metadata"`
		Name         string `json:"name"`
		PartnerAppID int    `json:"partner_app_id"`
		Payment      struct {
			Active bool `json:"active"`
			Cc     struct {
				ExpirationDate string `json:"expiration_date"`
				LastFour       string `json:"last_four"`
				PostalCode     string `json:"postal_code"`
				Type           string `json:"type"`
			} `json:"cc"`
			Type string `json:"type"`
		} `json:"payment"`
		Pictures struct {
			Active         bool   `json:"active"`
			BaseLink       string `json:"base_link"`
			DefaultPicture bool   `json:"default_picture"`
			Link           string `json:"link"`
			ResourceKey    string `json:"resource_key"`
			Sizes          []struct {
				Height             int    `json:"height"`
				Link               string `json:"link"`
				LinkWithPlayButton string `json:"link_with_play_button"`
				Width              int    `json:"width"`
			} `json:"sizes"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"pictures"`
		Preferences struct {
			Videos struct {
				Privacy struct {
					Password string `json:"password"`
				} `json:"privacy"`
			} `json:"videos"`
		} `json:"preferences"`
		PrimaryEmail       string `json:"primary_email"`
		ProfilePreferences struct {
			Layout                    string `json:"layout"`
			ProfileType               string `json:"profile_type"`
			ShouldAutoAddVideos       bool   `json:"should_auto_add_videos"`
			ShowAddVideoTip           bool   `json:"show_add_video_tip"`
			ShowClients               bool   `json:"show_clients"`
			ShowJoinVimeoExpertsModal bool   `json:"show_join_vimeo_experts_modal"`
			ShowProfileTypeTip        bool   `json:"show_profile_type_tip"`
			ShowRate                  bool   `json:"show_rate"`
		} `json:"profile_preferences"`
		ResourceKey string `json:"resource_key"`
		ShortBio    string `json:"short_bio"`
		Skills      struct {
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"skills"`
		TotalCollectionCount int `json:"total_collection_count"`
		UploadQuota          struct {
			Lifetime struct {
				Free int    `json:"free"`
				Max  int64  `json:"max"`
				Unit string `json:"unit"`
				Used int64  `json:"used"`
			} `json:"lifetime"`
			Periodic struct {
				Free      int    `json:"free"`
				Max       int64  `json:"max"`
				Period    string `json:"period"`
				ResetDate string `json:"reset_date"`
				Unit      string `json:"unit"`
				Used      int64  `json:"used"`
			} `json:"periodic"`
			Space struct {
				Free     int `json:"free"`
				Lifetime struct {
					Unit string `json:"unit"`
				} `json:"lifetime"`
				Max     int64  `json:"max"`
				Showing string `json:"showing"`
				Used    int64  `json:"used"`
			} `json:"space"`
		} `json:"upload_quota"`
		URI      string `json:"uri"`
		Verified bool   `json:"verified"`
		Websites []struct {
			Description string `json:"description"`
			Link        string `json:"link"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
		} `json:"websites"`
	} `json:"user"`
	Webinar []string `json:"webinar"`
}
