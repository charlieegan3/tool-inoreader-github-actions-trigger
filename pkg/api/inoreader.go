package api

// InoreaderPost is the format of the data sent by Inoreader
type InoreaderPost struct {
	Items []struct {
		Alternate []struct {
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"alternate"`
		Annotations []interface{} `json:"annotations"`
		Author      string        `json:"author"`
		Canonical   []struct {
			Href string `json:"href"`
		} `json:"canonical"`
		Categories    []string      `json:"categories"`
		Comments      []interface{} `json:"comments"`
		CommentsNum   int64         `json:"commentsNum"`
		CrawlTimeMsec string        `json:"crawlTimeMsec"`
		ID            string        `json:"id"`
		LikingUsers   []interface{} `json:"likingUsers"`
		Origin        struct {
			HTMLURL  string `json:"htmlUrl"`
			StreamID string `json:"streamId"`
			Title    string `json:"title"`
		} `json:"origin"`
		Published int64 `json:"published"`
		Summary   struct {
			Content   string `json:"content"`
			Direction string `json:"direction"`
		} `json:"summary"`
		TimestampUsec string `json:"timestampUsec"`
		Title         string `json:"title"`
		Updated       int64  `json:"updated"`
	} `json:"items"`
	Rule struct {
		MatchesToday string `json:"matchesToday"`
		MatchesTotal string `json:"matchesTotal"`
		Name         string `json:"name"`
	} `json:"rule"`
}
