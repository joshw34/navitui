package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/controller"
)

func StartUI(ctrl *controller.Controller) error {
	root := newRootModel(ctrl)
	p := tea.NewProgram(root)
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
			main:    listPageModel{list: newMainList(), onSelected: onSelectedMain},
			artists: listPageModel{list: newEmptyList("Artists"), onSelected: onSelectedArtists},
			albums:  listPageModel{list: newEmptyList("Albums"), onSelected: onSelectedAlbums},
			songs:   listPageModel{list: newEmptyList("Songs"), onSelected: onSelectedSongs},
		},
		ctrl: ctrl,
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
