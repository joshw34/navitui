package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/controller"
)

// SUB-PAGES
type page int

const (
	main page = iota
	artists
	albums
	songs
)

type pageModel interface {
	Update(msg tea.Msg) (pageModel, tea.Cmd)
	View() string
}

// MAIN BUBBLETEA INTERFACE
type rootModel struct {
	current  page
	previous []page
	pages    map[page]pageModel
	ctrl     *controller.Controller
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	found, newRoot, newCmd := m.setWindowSize(msg)
	if found {
		return newRoot, newCmd
	}

	found, newRoot, newCmd = m.globalKeyPresses(msg)
	if found {
		return newRoot, newCmd
	}

	found, newRoot, newCmd = m.checkGetMessages(msg)
	if found {
		return newRoot, newCmd
	}

	found, newRoot, newCmd = m.checkLoadedMessages(msg)
	if found {
		return newRoot, newCmd
	}

	found, newRoot, newCmd = m.checkPlayerMessages(msg)
	if found {
		return newRoot, newCmd
	}

	return m.delegateToSubpages(msg)
}

func (m rootModel) View() tea.View {
	s := m.pages[m.current].View()
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

func (m rootModel) Init() tea.Cmd {
	m.pages[main].(listPageModel).Update(mainOptionsToListItem())
	return nil
}

// UPDATE() HELPERS
func (m rootModel) setWindowSize(msg tea.Msg) (bool, rootModel, tea.Cmd) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		for key, p := range m.pages {
			if lp, ok := p.(listPageModel); ok {
				lp.list.SetSize(msg.Width, msg.Height)
				m.pages[key] = lp
			}
		}
		return true, m, nil
	}
	return false, m, nil
}

func (m rootModel) globalKeyPresses(key tea.Msg) (bool, rootModel, tea.Cmd) {
	if key, ok := key.(tea.KeyPressMsg); ok {
		switch key.String() {
		case "ctrl+c":
			return true, m, tea.Quit
		case "esc":
			return true, m.goToPreviousPage(), nil
		}
	}
	return false, m, nil
}

func (m rootModel) goToPage(newPage page) rootModel {
	m.previous = append(m.previous, m.current)
	m.current = newPage
	return m
}

func (m rootModel) goToPreviousPage() rootModel {
	if len(m.previous) == 1 {
		m.current = m.previous[0]
	} else {
		m.current = m.previous[len(m.previous)-1]
		m.previous = m.previous[:len(m.previous)-1]
	}
	return m
}

func (m rootModel) checkGetMessages(msg tea.Msg) (bool, rootModel, tea.Cmd) {
	switch msg := msg.(type) {
	case getAllArtistsMsg:
		return true, m.goToPage(artists), loadAllArtists(m.ctrl)

	case getAllAlbumsMsg:
		return true, m.goToPage(albums), loadAllAlbums(m.ctrl)

	case getAllSongsMsg:
		return true, m.goToPage(songs), loadAllSongs(m.ctrl)

	case getAlbumsByArtistMsg:
		return true, m.goToPage(albums), loadAlbumsByArtist(m.ctrl, msg.artistID)

	case getSongsByAlbumMsg:
		return true, m.goToPage(songs), loadSongsByAlbum(m.ctrl, msg.albumID)
	}
	return false, m, nil
}

func (m rootModel) checkLoadedMessages(msg tea.Msg) (bool, rootModel, tea.Cmd) {
	switch msg := msg.(type) {
	case loadedAllArtistsMsg:
		if msg.err != nil {
			return true, m, nil
		}
		lp := m.pages[artists].(listPageModel)
		m.pages[artists] = lp.updateList(artistsToListItems(msg.data))
		return true, m, nil

	case loadedAlbumsMsg:
		if msg.err != nil {
			return true, m, nil
		}
		lp := m.pages[albums].(listPageModel)
		m.pages[albums] = lp.updateList(albumsToListItems(msg.data))
		return true, m, nil

	case loadedSongsMsg:
		if msg.err != nil {
			return true, m, nil
		}
		lp := m.pages[songs].(listPageModel)
		m.pages[songs] = lp.updateList(songsToListItems(msg.data))
		return true, m, nil
	}
	return false, m, nil
}

func (m rootModel) checkPlayerMessages(msg tea.Msg) (bool, rootModel, tea.Cmd) {
	switch msg := msg.(type) {
	case playSongMsg:
		return true, m, func() tea.Msg {
			_ = m.ctrl.Play(msg.songID)
			return nil
		}
	}
	return false, m, nil
}

func (m rootModel) delegateToSubpages(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.pages[m.current], cmd = m.pages[m.current].Update(msg)
	return m, cmd
}
