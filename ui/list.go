package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hasanhakkaev/chgdns"
	"github.com/muesli/termenv"
)

// ListModel is the UI in which the user can select which forks should be
// deleted if any, and see details on each oqf them.
type ListModel struct {
	dnsServers chgdns.DnsServers
	cursor     int
	selected   map[int]struct{}
}

// NewListModel creates a new ListModel with the required fields.
func NewListModel(dnsServers chgdns.DnsServers) ListModel {
	return ListModel{
		dnsServers: dnsServers,
		selected:   map[int]struct{}{},
	}
}

// Init initializes the CLI.
func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.dnsServers.DnsServers)-1 {
				m.cursor++
			}
		/*case "a":
			for i := range m.dnsServers.DnsServers {
				m.selected[i] = struct{}{}
			}
		case "n":
			for i := range m.selected {
				delete(m.selected, i)
			}*/
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
			/*dm := NewDeletingModel(m.client, deleteable, m)
			return dm, dm.Init()*/
		}
	}
	return m, nil
}

func (m ListModel) View() string {
	s := boldSecondaryForeground("Which of these DNS servers do you want to use?\n\n")

	for i, dnsServer := range m.dnsServers.DnsServers {
		line := dnsServer.Name
		if _, ok := m.selected[i]; ok {
			line = iconSelected + " " + line
		} else {
			line = faint(iconNotSelected + " " + line)
		}
		line += "\n"

		if m.cursor == i {
			nl := ""
			if i > 0 {
				nl = "\n"
			}
			line = nl + boldPrimaryForeground(line) + viewDnsServerDetails(m.dnsServers.DnsServers[i])
		}

		s += line
	}

	return s + helpView([]helpOption{
		{"up/down", "navigate", false},
		{"space", "toggle selection", false},
		/*{"d", "delete selected", true},
		{"a", "select all", false},
		{"n", "deselect all", false},*/
		{"q/esc", "quit", false},
	})
}

func viewDnsServerDetails(dnsServer chgdns.DnsServer) string {
	var details []string
	if dnsServer.Name != "" {
		// Don't display the name for second time
		//details = append(details, fmt.Sprint(dnsServer.Name))
		details = append(details, fmt.Sprint(dnsServer.PrimaryDNS))
		details = append(details, fmt.Sprint(dnsServer.SecondaryDNS))
	}

	if len(details) == 0 {
		return ""
	}

	var s string
	for _, d := range details {
		s += "    * " + d + "\n"
	}
	s += "\n"
	return termenv.String(s).Faint().Italic().String()
}
