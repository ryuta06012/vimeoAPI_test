package vimeo

import (
	"time"
)

type ThumnailReponse struct {
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
