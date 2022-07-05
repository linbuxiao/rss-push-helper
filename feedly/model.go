package feedly

type Collection struct {
	Customizable bool   `json:"customizable"`
	Id           string `json:"id"`
	NumFeeds     int    `json:"numFeeds"`
	Feeds        []struct {
		ContentType         string   `json:"contentType"`
		Description         string   `json:"description"`
		Language            string   `json:"language"`
		FeedId              string   `json:"feedId"`
		Title               string   `json:"title"`
		Topics              []string `json:"topics"`
		Updated             int64    `json:"updated"`
		Website             string   `json:"website"`
		Subscribers         int      `json:"subscribers"`
		Velocity            float64  `json:"velocity"`
		IconUrl             string   `json:"iconUrl"`
		Partial             bool     `json:"partial"`
		EstimatedEngagement int      `json:"estimatedEngagement"`
		VisualUrl           string   `json:"visualUrl"`
	} `json:"feeds"`
	Label      string `json:"label"`
	Created    int64  `json:"created"`
	Enterprise bool   `json:"enterprise"`
}

type GetStreamIDsReponse struct {
	Ids          []string `json:"ids"`
	Continuation string   `json:"continuation"`
}

type GetEntriesResponse struct {
	Id         string `json:"id"`
	Unread     bool   `json:"unread"`
	Categories []struct {
		Id    string `json:"id"`
		Label string `json:"label"`
	} `json:"categories"`
	Tags []struct {
		Id    string `json:"id"`
		Label string `json:"label"`
	} `json:"tags"`
	Title     string   `json:"title"`
	Keywords  []string `json:"keywords"`
	Published int64    `json:"published"`
	Updated   int64    `json:"updated"`
	Crawled   int64    `json:"crawled"`
	Alternate []struct {
		Href string `json:"href"`
		Type string `json:"type"`
	} `json:"alternate"`
	Content struct {
		Direction string `json:"direction"`
		Content   string `json:"content"`
	} `json:"content"`
	Author string `json:"author"`
	Origin struct {
		StreamId string `json:"streamId"`
		Title    string `json:"title"`
		HtmlUrl  string `json:"htmlUrl"`
	} `json:"origin"`
	Engagement     int     `json:"engagement"`
	EngagementRate float64 `json:"engagementRate,omitempty"`
}
