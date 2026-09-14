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
)

type rootModel struct {
	current     page
	mainMenu    mainModel
	artistsMenu artistsModel
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
		m.artistsMenu.data = msg.artists
		return m, nil
	}
	var cmd tea.Cmd
	switch m.current {
	case main:
		m.mainMenu, cmd = m.mainMenu.Update(msg)
	case artists:
		m.artistsMenu, cmd = m.artistsMenu.Update(msg)
	}
	return m, cmd
}

func (m rootModel) View() tea.View {
	var s string
	switch m.current {
	case main:
		s = m.mainMenu.View()
	case artists:
		s = m.artistsMenu.View()
	}
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

func (m rootModel) Init() tea.Cmd {
	return nil
}
