package model

type DouYinVideo struct {
	StatusCode int `json:"status_code"`
	ItemList []ItemLists `json:"item_list"`
}

type ItemLists struct {
	AwemeId string `json:"aweme_id"`
	Desc string `json:"desc"`
	Music Music `json:"music"`
	Video Video `json:"video"`
}

type Music struct {
	Id int64 `json:"id"`
	PlayUrl MusicPlayUrl `json:"play_url"`
	CoverHd CoverHdList `json:"cover_hd"`
}

type MusicPlayUrl struct {
	Uri string `json:"uri"`
}

type CoverHdList struct {
	UrlList []string `json:"url_list"`
}

type Video struct {
	PlayAddr VideoPlayUrl `json:"play_addr"`
	Cover CoverUriList `json:"cover"`
	DynamicCover DynamicCoverList `json:"dynamic_cover"`
	OriginCover OriginCoverList `json:"origin_cover"`
}

type VideoPlayUrl struct {
	UrlList []string `json:"url_list"`
}

type CoverUriList struct {
	UrlList []string `json:"url_list"`
}

type DynamicCoverList struct {
	UrlList []string `json:"url_list"`
}

type OriginCoverList struct {
	UrlList []string `json:"url_list"`
}