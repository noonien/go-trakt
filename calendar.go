package trakt

import (
	"fmt"
	"sort"
	"time"
)

type CalendarEpisode struct {
	AirsAt  time.Time `json:"airs_at"`
	Episode Episode   `json:"episode"`
	Show    Show      `json:"show"`
}

type calSort []CalendarEpisode

func (a calSort) Len() int           { return len(a) }
func (a calSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a calSort) Less(i, j int) bool { return a[i].AirsAt.Before(a[j].AirsAt) }

type CalendarService struct {
	client *Client
}

func (c *CalendarService) Shows(start string, days int) ([]CalendarEpisode, error) {
	urlStr := fmt.Sprintf("/calendars/shows/%s/%d", start, days)

	req, err := c.client.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	calEps := map[string][]CalendarEpisode{}
	_, err = c.client.Do(req, &calEps)
	if err != nil {
		return nil, err
	}

	var eps []CalendarEpisode
	for k := range calEps {
		eps = append(eps, calEps[k]...)
	}

	sort.Sort(calSort(eps))

	return eps, nil
}
