package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type listPageModel struct {
	list       list.Model
	onKeypress func(tea.KeyPressMsg, list.Item, int) tea.Cmd
}

func (l listPageModel) updateList(items []list.Item) listPageModel {
	l.list.SetItems(items)
	return l
}

func (l listPageModel) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	// Check for page-specific keypress
	if key, ok := msg.(tea.KeyPressMsg); ok && l.list.FilterState() != list.Filtering {
		if it := l.list.SelectedItem(); it != nil {
			listCmd := l.onKeypress(key, it, l.list.Index())
			if listCmd != nil {
				return l, listCmd
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
