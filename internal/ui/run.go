package ui

import (
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
		current:     main,
		mainMenu:    newMainModel(),
		artistsMenu: newArtistsModel(),
		ctrl:        ctrl,
	}
}

func newMainModel() mainModel {
	return mainModel{
		choices: []string{"artists"},
		cursor:  0,
	}
}

func newArtistsModel() artistsModel {
	return artistsModel{}
}
