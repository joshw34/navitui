package navidrome

type jsonResponse struct {
	SubResp jsonSubsonicResponse `json:"subsonic-response"`
}

type jsonSubsonicResponse struct {
	Status       string                    `json:"status"`
	ArtistAlbums jsonArtistAlbums          `json:"artist"`
	AllArtists   jsonAllArtists            `json:"artists"`
	AlbumSongs   jsonAlbumSongs            `json:"album"`
	AllAlbums    jsonAllAlbums             `json:"albumList"`
	Search2      jsonSearch2               `json:"searchResult2"`
	Error        jsonSubsonicResponseError `json:"error"`
}

type jsonSubsonicResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type jsonAllArtists struct {
	Indexes []struct {
		Artists []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			AlbumCount int    `json:"albumCount"`
		} `json:"artist"`
	} `json:"index"`
}

type jsonArtistAlbums struct {
	Albums []jsonAlbum `json:"album"`
}

type jsonAllAlbums struct {
	Albums []jsonAlbum `json:"album"`
}

type jsonAlbumSongs struct {
	Songs []jsonSong `json:"song"`
}

type jsonSearch2 struct {
	Songs []jsonSong `json:"song"`
}

type jsonAlbum struct {
	ID       string `json:"id"`
	ArtistID string `json:"artistId"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Genres   []struct {
		Name string `json:"name"`
	} `json:"genres"`
	Year      int `json:"year"`
	Duration  int `json:"duration"`
	SongCount int `json:"songCount"`
}

type jsonSong struct {
	ID       string `json:"id"`
	ArtistID string `json:"artistId"`
	AlbumID  string `json:"albumId"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	FileType string `json:"suffix"`
	Track    int    `json:"track"`
	Year     int    `json:"year"`
	Duration int    `json:"duration"`
	Disc     int    `json:"discNumber"`
}
