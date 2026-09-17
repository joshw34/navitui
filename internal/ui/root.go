package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/controller"
)

type pageModel interface {
	Update(msg tea.Msg) (pageModel, tea.Cmd)
	View() string
}

type page int

const (
	main page = iota
	artists
	albums
	songs
)

type rootModel struct {
	current     page
	previous    []page
	mainMenu    mainModel
	artistsList artistsModel
	albumsList  albumsModel
	songsList   songsModel
	ctrl        *controller.Controller
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.mainMenu.list.SetSize(msg.Width, msg.Height)
		m.artistsList.list.SetSize(msg.Width, msg.Height)
		m.albumsList.list.SetSize(msg.Width, msg.Height)
		m.songsList.list.SetSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == "esc" {
			if len(m.previous) == 1 {
				m.current = m.previous[0]
				return m, nil
			}
			m.current = m.previous[len(m.previous)-1]
			m.previous = m.previous[:len(m.previous)-1]
		}

	case getAllArtistsMsg:
		m.previous = append(m.previous, m.current)
		m.current = artists
		return m, loadAllArtists(m.ctrl)

	case loadedAllArtistsMsg:
		if msg.err != nil {
			return m, nil
		}
		m.artistsList = m.artistsList.buildList(msg.data)
		return m, nil

	case getAlbumsByArtistMsg:
		m.previous = append(m.previous, m.current)
		m.current = albums
		return m, loadAlbumsByArtist(m.ctrl, msg.artistID)

	case getAllAlbumsMsg:
		m.previous = append(m.previous, m.current)
		m.current = albums
		return m, loadAllAlbums(m.ctrl)

	case loadedAlbumsMsg:
		if msg.err != nil {
			return m, nil
		}
		m.albumsList = m.albumsList.buildList(msg.data)
		return m, nil

	case getSongsByAlbumMsg:
		m.previous = append(m.previous, m.current)
		m.current = songs
		return m, loadSongsByAlbum(m.ctrl, msg.albumID)

	case getAllSongsMsg:
		m.previous = append(m.previous, m.current)
		m.current = songs
		return m, loadAllSongs(m.ctrl)

	case loadedSongsMsg:
		if msg.err != nil {
			return m, nil
		}
		m.songsList = m.songsList.buildList(msg.data)
		return m, nil

	case playSongMsg:
		return m, func() tea.Msg {
			_ = m.ctrl.Play(msg.songID)
			return nil
		}
	}

	var cmd tea.Cmd
	switch m.current {
	case main:
		m.mainMenu, cmd = m.mainMenu.Update(msg)
	case artists:
		m.artistsList, cmd = m.artistsList.Update(msg)
	case albums:
		m.albumsList, cmd = m.albumsList.Update(msg)
	case songs:
		m.songsList, cmd = m.songsList.Update(msg)
	}
	return m, cmd
}

func (m rootModel) View() tea.View {
	var s string
	switch m.current {
	case main:
		s = m.mainMenu.View()
	case artists:
		s = m.artistsList.View()
	case albums:
		s = m.albumsList.View()
	case songs:
		s = m.songsList.View()
	}
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

func (m rootModel) Init() tea.Cmd {
	return nil
}
