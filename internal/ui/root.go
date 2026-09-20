// Package ui: bubbletea interface
package ui

import (
	"log"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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

type inputReceiver int

const (
	left inputReceiver = iota
	right
)

type pageModel interface {
	Update(msg tea.Msg) (pageModel, tea.Cmd)
	View() string
}

// MAIN BUBBLETEA INTERFACE
type rootModel struct {
	current                                page
	previous                               []page
	pages                                  map[page]pageModel
	queue                                  pageModel
	nowPlaying                             pageModel
	activePane                             inputReceiver
	listPaneH, listPaneW, npPaneH, npPaneW int
	ctrl                                   *controller.Controller
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

	found, newRoot, newCmd = m.checkUpdateMessage(msg)
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
	// lipgloss uses the entire pane width (including border)
	leftStyle := lipgloss.NewStyle().Width(m.listPaneW).Height(m.listPaneH).Border(lipgloss.RoundedBorder())
	rightStyle := lipgloss.NewStyle().Width(m.listPaneW).Height(m.listPaneH).Border(lipgloss.RoundedBorder())
	bottomStyle := lipgloss.NewStyle().Width(m.npPaneW).Height(m.npPaneH).Border(lipgloss.RoundedBorder())

	left := leftStyle.Render(m.pages[m.current].View())
	right := rightStyle.Render(m.queue.View())
	bottom := bottomStyle.Render(m.nowPlaying.View())

	top := lipgloss.JoinHorizontal(lipgloss.Top, left, right)
	full := lipgloss.JoinVertical(lipgloss.Top, top, bottom)
	v := tea.NewView(full)
	v.AltScreen = true
	return v
}

func (m rootModel) Init() tea.Cmd {
	return nil
}

// UPDATE() HELPERS
func (m rootModel) setWindowSize(msg tea.Msg) (bool, rootModel, tea.Cmd) {
	// list.SetSize() needs usable area (minus border width)
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		const border = 2
		m.listPaneW = msg.Width / 2
		m.listPaneH = msg.Height * 85 / 100
		m.npPaneW = msg.Width
		m.npPaneH = msg.Height * 15 / 100
		for key, p := range m.pages {
			if lp, ok := p.(listPageModel); ok {
				lp.list.SetSize(m.listPaneW-border, m.listPaneH-border)
				m.pages[key] = lp
			}
		}
		q := m.queue.(listPageModel)
		q.list.SetSize(m.listPaneW-border, m.listPaneH-border)
		m.queue = q
		np := m.nowPlaying.(nowPlayingModel)
		np.prog.SetWidth(m.npPaneW)
		m.nowPlaying = np
		return true, m, nil
	}
	return false, m, nil
}

func (m rootModel) globalKeyPresses(key tea.Msg) (bool, rootModel, tea.Cmd) {
	if key, ok := key.(tea.KeyPressMsg); ok && !m.isFiltering() {
		switch key.String() {
		case "ctrl+c":
			return true, m, tea.Quit
		case "esc":
			return true, m.goToPreviousPage(), nil
		case "1":
			m.activePane = left
			return true, m, nil
		case "2":
			m.activePane = right
			return true, m, nil
		case "space":
			return true, m, func() tea.Msg {
				if err := m.ctrl.TogglePlayPause(); err != nil {
					log.Println("UI: Pause Failed")
				}
				return nil
			}
		case "x":
			return true, m, func() tea.Msg {
				if err := m.ctrl.Stop(); err != nil {
					log.Println("UI: Stop Failed")
				}
				return nil
			}
		}

	}
	return false, m, nil
}

func (m rootModel) isFiltering() bool {
	for _, p := range m.pages {
		if lp, ok := p.(listPageModel); ok {
			if lp.list.FilterState() == list.Filtering {
				return true
			}
		}
	}
	return false
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
			if err := m.ctrl.PlaySong(msg.s); err != nil {
				log.Println("UI: Play Failed")
			}
			return nil
		}
	case addToQueueMsg:
		return true, m, func() tea.Msg {
			m.ctrl.QueueAddToEnd(msg.s)
			return nil
		}
	case playNextMsg:
		return true, m, func() tea.Msg {
			m.ctrl.QueueAddNext(msg.s)
			return nil
		}
	case removeFromQueueMsg:
		return true, m, func() tea.Msg {
			m.ctrl.QueueRemove(msg.index)
			return nil
		}
	case clearQueueMsg:
		return true, m, func() tea.Msg {
			m.ctrl.QueueClear()
			return nil
		}
	}
	return false, m, nil
}

func (m rootModel) checkUpdateMessage(msg tea.Msg) (bool, rootModel, tea.Cmd) {
	_, ok := msg.(updateMsg)
	if !ok {
		return false, m, nil
	}
	u := msg.(updateMsg).u
	switch u.Type {
	case controller.NowPlaying:
		np := m.nowPlaying.(nowPlayingModel).updateSong(u.NowPlaying)
		m.nowPlaying = np
		return true, m, nil
	case controller.QueueUpdate:
		lp := m.queue.(listPageModel)
		m.queue = lp.updateList(queueToListItems(u.Queue))
		return true, m, nil
	case controller.TimePosUpdate:
		//log.Printf("UI: %f\tNEW: %f", m.nowPlaying.(nowPlayingModel).timePos, u.TimePos)
		np := m.nowPlaying.(nowPlayingModel).updateTP(u.TimePos)
		m.nowPlaying = np
		return true, m, nil
	}
	return false, m, nil
}

func (m rootModel) delegateToSubpages(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.activePane == left {
		m.pages[m.current], cmd = m.pages[m.current].Update(msg)
	} else {
		m.queue, cmd = m.queue.Update(msg)
	}
	return m, cmd
}
