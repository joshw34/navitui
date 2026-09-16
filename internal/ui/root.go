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
	artistsList
	artist
	album
)

type rootModel struct {
	current         page
	previous        []page
	mainMenu        mainModel
	artistsListPage artistsListModel
	artistPage      artistModel
	albumPage       albumModel
	ctrl            *controller.Controller
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.mainMenu.list.SetSize(msg.Width, msg.Height)
		m.artistsListPage.list.SetSize(msg.Width, msg.Height)
		m.artistPage.list.SetSize(msg.Width, msg.Height)
		m.albumPage.list.SetSize(msg.Width, msg.Height)
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

	case getArtistsListMsg:
		m.previous = append(m.previous, m.current)
		m.current = artistsList
		return m, loadArtistsList(m.ctrl)

	case artistsListLoadedMsg:
		if msg.err != nil {
			return m, nil
		}
		m.artistsListPage = m.artistsListPage.buildList(msg.data)
		return m, nil

	case getArtistMsg:
		m.previous = append(m.previous, m.current)
		m.current = artist
		return m, loadArtist(m.ctrl, msg.artistID)

	case artistLoadedMsg:
		if msg.err != nil {
			return m, nil
		}
		m.artistPage = m.artistPage.buildList(msg.data)
		return m, nil

	case getAlbumMsg:
		m.previous = append(m.previous, m.current)
		m.current = album
		return m, loadAlbum(m.ctrl, msg.albumID)

	case albumLoadedMsg:
		if msg.err != nil {
			return m, nil
		}
		m.albumPage = m.albumPage.buildList(msg.data)
		return m, nil

	}

	var cmd tea.Cmd
	switch m.current {
	case main:
		m.mainMenu, cmd = m.mainMenu.Update(msg)
	case artistsList:
		m.artistsListPage, cmd = m.artistsListPage.Update(msg)
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
	case artistsList:
		s = m.artistsListPage.View()
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
