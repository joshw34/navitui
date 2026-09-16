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
		current:         main,
		previous:        []page{},
		mainMenu:        newMainModel(),
		artistsListPage: newArtistsListModel(),
		artistPage:      newArtistModel(),
		albumPage:       newAlbumModel(),
		ctrl:            ctrl,
	}
}

func newMainModel() mainModel {
	options := []mainItem{{option: "Artists", action: artistsList}}

	items := make([]list.Item, len(options))
	for i, opt := range options {
		items[i] = opt
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0) // pass items straight in
	l.Title = "Main Menu"

	return mainModel{list: l}
}

func newArtistsListModel() artistsListModel {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Artists"
	return artistsListModel{list: l}
}

func newArtistModel() artistModel {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Albums"
	return artistModel{list: l}
}

func newAlbumModel() albumModel {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Tracks"
	return albumModel{list: l}
}
