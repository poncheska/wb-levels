package main

import (
	"bytes"
	"encoding/json"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/domain"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/handler"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/service"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

type eventResult struct {
	Res []domain.Event `json:"result"`
}

type doneResult struct {
	Res string `json:"result"`
}

type errResult struct {
	Err string `json:"error"`
}

var events = []domain.Event{
	{
		ID:          "id1",
		Header:      "header1",
		Date:        time.Now().Add(12 * time.Hour),
		Description: "description1",
		CreatorID:   "creator",
	},
	{
		ID:          "id2",
		Header:      "header2",
		Date:        time.Now().Add(2 * 24 * time.Hour),
		Description: "description2",
		CreatorID:   "creator",
	},
	{
		ID:          "id3",
		Header:      "header3",
		Date:        time.Now().Add(8 * 24 * time.Hour),
		Description: "description3",
		CreatorID:   "creator",
	},
	{
		ID:          "id4",
		Header:      "header4",
		Date:        time.Now().Add(32 * 24 * time.Hour),
		Description: "description4",
		CreatorID:   "creator4",
	},
}

var updEvent = domain.Event{
	ID:          "id4",
	Header:      "header4",
	Date:        time.Now().Add(12 * time.Hour),
	Description: "description4",
	CreatorID:   "creator",
}

const addr = "http://127.0.0.1:8080"

func TestServer(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	go func() {
		svc := service.NewEventService()
		h := handler.LoggerMiddleware(handler.New(svc))
		http.ListenAndServe(":8080", h)
	}()
	time.Sleep(100 * time.Millisecond)

	t.Run("create events", func(t *testing.T) {
		for _, v := range events {
			bs, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			r := bytes.NewReader(bs)
			resp, err := http.Post(addr+"/create_event", "application/json", r)
			if err != nil {
				panic(err)
			}
			res := doneResult{}
			bb, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(bb, &res)
			if err != nil {
				panic(err)
			}
			assert.Equal(t, "done", res.Res)
		}
	})

	t.Run("get events for day", func(t *testing.T) {
		resp, err := http.Get(addr + "/events_for_day?user_id=creator")
		if err != nil {
			panic(err)
		}
		res := eventResult{}
		bb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bb, &res)
		if err != nil {
			panic(err)
		}
		assert.NotNil(t, res.Res)
		assert.Equal(t, 1, len(res.Res))
	})

	t.Run("get events for week", func(t *testing.T) {
		resp, err := http.Get(addr + "/events_for_week?user_id=creator")
		if err != nil {
			panic(err)
		}
		res := eventResult{}
		bb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bb, &res)
		if err != nil {
			panic(err)
		}
		assert.NotNil(t, res.Res)
		assert.Equal(t, 2, len(res.Res))
	})

	t.Run("get events for month", func(t *testing.T) {
		resp, err := http.Get(addr + "/events_for_month?user_id=creator")
		if err != nil {
			panic(err)
		}
		res := eventResult{}
		bb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bb, &res)
		if err != nil {
			panic(err)
		}
		assert.NotNil(t, res.Res)
		assert.Equal(t, 3, len(res.Res))
	})

	t.Run("delete event", func(t *testing.T) {
		resp, err := http.Post(addr+"/delete_event?event_id=id1",
			"application/json", bytes.NewReader([]byte{}))
		if err != nil {
			panic(err)
		}
		res := doneResult{}
		bb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bb, &res)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "done", res.Res)
	})

	t.Run("update event", func(t *testing.T) {
		bs, err := json.Marshal(updEvent)
		if err != nil {
			panic(err)
		}
		r := bytes.NewReader(bs)
		resp, err := http.Post(addr+"/update_event", "application/json", r)
		if err != nil {
			panic(err)
		}
		res := doneResult{}
		bb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bb, &res)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "done", res.Res)
	})

	t.Run("check update and delete", func(t *testing.T) {
		resp, err := http.Get(addr + "/events_for_day?user_id=creator")
		if err != nil {
			panic(err)
		}
		res := eventResult{}
		bb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bb, &res)
		if err != nil {
			panic(err)
		}
		assert.NotNil(t, res.Res)
		assert.Equal(t, 1, len(res.Res))
		assert.Equal(t, "id4", res.Res[0].ID)
	})
}
