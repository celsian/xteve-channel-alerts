package channel

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannel_ParseM3U(t *testing.T) {
	tests := []struct {
		name   string
		m3u    []byte
		expect []Channel
	}{
		{
			"empty file",
			[]byte(""),
			nil,
		},
		{
			"single channel",
			[]byte(`#EXTINF:-1,Channel 1,tvg-name="Channel 1",tvg-id="1",group-title="Group 1"`),
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
			},
		},
		{
			"multiple channels",
			[]byte(`#EXTINF:-1,Channel 1,tvg-name="Channel 1",tvg-id="1",group-title="Group 1"
			#EXTINF:-1,Channel 2,tvg-name="Channel 2",tvg-id="2",group-title="Group 2"`),
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channels := ParseM3U(tt.m3u)
			assert.Equal(t, tt.expect, channels)
		})
	}
}

func TestChannel_CompareChannels(t *testing.T) {
	tests := []struct {
		name     string
		previous []Channel
		current  []Channel
		expect   []Channel
	}{
		{
			"channels match",
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 2"},
			},
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 2"},
			},
			nil,
		},
		{
			"different title",
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 1"},
			},
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 3", GroupTitle: "Group 1"},
			},
			[]Channel{{Number: "2", Title: "Channel 2", GroupTitle: "Group 1"}},
		},
		{
			"different channel number",
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 1"},
			},
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "3", Title: "Channel 2", GroupTitle: "Group 1"},
			},
			[]Channel{{Number: "2", Title: "Channel 2", GroupTitle: "Group 1"}},
		},
		{
			"different channel group",
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 1"},
			},
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 2"},
			},
			[]Channel{{Number: "2", Title: "Channel 2", GroupTitle: "Group 1"}},
		},
		{
			"all channels missing",
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 1"},
			},
			nil,
			[]Channel{
				{Number: "1", Title: "Channel 1", GroupTitle: "Group 1"},
				{Number: "2", Title: "Channel 2", GroupTitle: "Group 1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channels := CompareChannels(tt.previous, tt.current)
			assert.Equal(t, tt.expect, channels)
		})
	}
}

func TestChannel_Log(t *testing.T) {
	tests := []struct {
		name     string
		channel  Channel
		expected string
	}{
		{
			"log channel",
			Channel{
				Number: "1", Title: "Channel 1", GroupTitle: "Group 1",
			},
			"1 - Channel 1 - Group 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
			tt.channel.Log()
			assert.Contains(t, buf.String(), tt.expected)
		})
	}
}

func TestChannel_Warn(t *testing.T) {
	tests := []struct {
		name     string
		channel  Channel
		expected string
	}{
		{
			"log channel",
			Channel{
				Number: "1", Title: "Channel 1", GroupTitle: "Group 1",
			},
			"1 - Channel 1 - Group 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
			tt.channel.LogWarning()
			assert.Contains(t, buf.String(), tt.expected)
		})
	}
}

// Test comment
