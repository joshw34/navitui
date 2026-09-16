package client

type jsonResponse struct {
	SubResp jsonSubsonicResponse `json:"subsonic-response"`
}

type jsonSubsonicResponse struct {
	Status  string                    `json:"status"`
	Artist  jsonArtist                `json:"artist"`
	Artists jsonArtistList            `json:"artists"`
	Album   jsonAlbum                 `json:"album"`
	Error   jsonSubsonicResponseError `json:"error"`
}

type jsonSubsonicResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type jsonArtistList struct {
	Indexes []struct {
		Artists []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			AlbumCount int    `json:"albumCount"`
		} `json:"artist"`
	} `json:"index"`
}

type jsonArtist struct {
	Albums []struct {
		ID       string `json:"id"`
		ArtistID string `json:"artistId"`
		Name     string `json:"name"`
		Genres   []struct {
			Name string `json:"name"`
		} `json:"genres"`
		Year      int `json:"year"`
		Duration  int `json:"duration"`
		SongCount int `json:"songCount"`
	} `json:"album"`
}

type jsonAlbum struct {
	Songs []struct {
		ID       string `json:"id"`
		ArtistID string `json:"artistId"`
		AlbumID  string `json:"albumId"`
		Title    string `json:"title"`
		FileType string `json:"suffix"`
		Track    int    `json:"track"`
		Year     int    `json:"year"`
		Duration int    `json:"duration"`
		Disc     int    `json:"discNumber"`
	} `json:"song"`
}
