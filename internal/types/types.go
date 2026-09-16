// Package types: shared types used my multiple packages
package types

type Artist struct {
	ID         string
	Name       string
	AlbumCount int
}

type Album struct {
	ID        string
	ArtistID  string
	Name      string
	Genres    []string
	Year      int
	Duration  int
	SongCount int
}

type Song struct {
	ID       string
	ArtistID string
	AlbumID  string
	Title    string
	FileType string
	Track    int
	Year     int
	Duration int
	Disc     int
}
