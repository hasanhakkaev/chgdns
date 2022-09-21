package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hasanhakkaev/chgdns"
)

// NewInitialModel creates a new InitialModel with required fields.
func NewInitialModel(dnsServers chgdns.DnsServers) InitialModel {
	s := spinner.New()
	s.Spinner = spinner.MiniDot

	return InitialModel{
		dnsServers: dnsServers,
		spinner:    s,
		loading:    true,
	}
}

// InitialModel is the UI when the CLI starts, basically loading the DNS servers.
type InitialModel struct {
	dnsServers chgdns.DnsServers
	err        error
	spinner    spinner.Model
	loading    bool
}

func (m InitialModel) Init() tea.Cmd {
	return tea.Batch(getDnsServers(), spinner.Tick)
}

func (m InitialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.loading = false
		m.err = msg.error
		return m, nil
	case gotDnsServersMsg:
		list := NewListModel(msg.dnsServers)
		return list, list.Init()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m InitialModel) View() string {
	if m.loading {
		return boldPrimaryForeground(m.spinner.View()) + " Gathering repositories..." + SingleOptionHelp("q", "quit")
	}
	if m.err != nil {
		return errorView("Error gathering the repository list", m.err)
	}
	return ""
}

type gotDnsServersMsg struct {
	dnsServers chgdns.DnsServers
}

func getDnsServers() tea.Cmd {
	return func() tea.Msg {
		dnsServers, err := chgdns.ReadDnsServers()
		if err != nil {
			return errMsg{err}
		}
		return gotDnsServersMsg{dnsServers}
	}
}
