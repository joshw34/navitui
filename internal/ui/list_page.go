package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type listPageModel struct {
	list       list.Model
	onSelected func(list.Item) tea.Cmd
}

func (l listPageModel) updateList(items []list.Item) listPageModel {
	l.list.SetItems(items)
	return l
}

func (l listPageModel) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && l.list.FilterState() != list.Filtering {
		switch key.String() {
		case "enter":
			if it := l.list.SelectedItem(); it != nil {
				return l, l.onSelected(it)
			}
		}
	}

	// Pass keypress to the list
	var cmd tea.Cmd
	l.list, cmd = l.list.Update(msg)
	return l, cmd
}

func (l listPageModel) View() string {
	return l.list.View()
}
