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
	artist
	album
)

type rootModel struct {
	current     page
	mainMenu    mainModel
	artistsPage artistsModel
	artistPage  artistModel
	albumPage   albumModel
	ctrl        *controller.Controller
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}

	case getArtistsMsg:
		m.current = artists
		return m, loadArtists(m.ctrl)

	case artistsLoadedMsg:
		if msg.err != nil {
			return m, nil
		}
		m.artistsPage.data = msg.data
		return m, nil

	case getArtistMsg:
		m.current = artist
		return m, loadArtist(m.ctrl, msg.artistID)

	case artistLoadedMsg:
		if msg.err != nil {
			return m, nil
		}
		m.artistPage.data = msg.data
		return m, nil

	case getAlbumMsg:
		m.current = album
		return m, loadAlbum(m.ctrl, msg.albumID)

	case albumLoadedMsg:
		if msg.err != nil {
			return m, nil
		}
		m.albumPage.data = msg.data
		return m, nil

	}

	var cmd tea.Cmd
	switch m.current {
	case main:
		m.mainMenu, cmd = m.mainMenu.Update(msg)
	case artists:
		m.artistsPage, cmd = m.artistsPage.Update(msg)
	case artist:
		m.artistPage, cmd = m.artistPage.Update(msg)
	case album:
		m.albumPage, cmd = m.albumPage.Update(msg)
	}
	return m, cmd
}

func (m rootModel) View() tea.View {
	var s string
	switch m.current {
	case main:
		s = m.mainMenu.View()
	case artists:
		s = m.artistsPage.View()
	case artist:
		s = m.artistPage.View()
	case album:
		s = m.albumPage.View()
	}
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

func (m rootModel) Init() tea.Cmd {
	return nil
}
