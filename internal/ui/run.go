package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/controller"
)

func StartUI(ctrl *controller.Controller) error {
	root := newRootModel(ctrl)
	p := tea.NewProgram(root)
	ctrl.SetUpdateHandler(func(u controller.Update) {
		p.Send(updateMsg{u})
	})
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func newRootModel(ctrl *controller.Controller) rootModel {
	return rootModel{
		current:  main,
		previous: []page{},
		pages: map[page]pageModel{
			main:    listPageModel{list: newMainList(), onKeypress: onKeypressMain},
			artists: listPageModel{list: newEmptyList("Artists"), onKeypress: onKeypressArtists},
			albums:  listPageModel{list: newEmptyList("Albums"), onKeypress: onKeypressAlbums},
			songs:   listPageModel{list: newEmptyList("Songs"), onKeypress: onKeypressSongs},
		},
		queue: listPageModel{list: newEmptyList("Queue"), onKeypress: onKeypressQueue},
		ctrl:  ctrl,
	}
}

func newEmptyList(title string) list.Model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = title
	return l
}

func newMainList() list.Model {
	l := newEmptyList("Main Menu")
	l.SetItems(mainOptionsToListItem())
	return l
}
