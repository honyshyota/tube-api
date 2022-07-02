package apiserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/honyshyota/tube-api-go/internal/app/model"
	"github.com/honyshyota/tube-api-go/internal/app/store"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	ctxKeyVideo ctxKey = iota
	ctxKeyRequestID
)

const MaxResult = 10

type ctxKey int8

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	store   store.Store
	youtube *youtube.Service
}

func newServer(store store.Store) *server {
	ytClient, err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YOUTUBE_KEY")))
	if err != nil {
		logrus.Fatalln(err)
	}

	srv := &server{
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		store:   store,
		youtube: ytClient,
	}

	srv.configureRouter()

	return srv
}

func (srv *server) configureRouter() {
	srv.router.Use(srv.setRequestID)
	srv.router.Use(srv.logRequest)
	srv.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	srv.router.HandleFunc("/search", srv.handleChannelSearch()).Methods("GET")
	srv.router.HandleFunc("/video", srv.handleVideoSearch()).Methods("GET")
	srv.router.HandleFunc("/playlist", srv.handlePlaylistSearch()).Methods("GET")
}

func (srv *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (srv *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := srv.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}

func (srv *server) handleChannelSearch() http.HandlerFunc {
	type request struct {
		Keyword string `json:"key"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		var resultSch []*model.Channel

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			srv.error(w, r, http.StatusBadRequest, err)
			return
		}

		list := make([]string, 2)
		list[0] = "id"
		list[1] = "snippet"
		call := srv.youtube.Search.List(list).
			Q(req.Keyword).
			Type("channel").
			MaxResults(int64(MaxResult))
		response, err := call.Do()
		if err != nil {
			log.Println(err)
		}

		for _, item := range response.Items {
			sch := &model.Channel{}
			sch.ChannelID = item.Snippet.ChannelId
			sch.ChannelName = item.Snippet.ChannelTitle
			sch.ChannelInfo = item.Snippet.Description
			resultSch = append(resultSch, sch)
			srv.store.Repo().CreateChannels(sch)
		}

		srv.respond(w, r, http.StatusFound, resultSch)

	}
}

func (srv *server) handleVideoSearch() http.HandlerFunc {
	type request struct {
		ID int `json:"id,string"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			srv.error(w, r, http.StatusBadRequest, err)
			return
		}

		channel, err := srv.store.Repo().FindChannel(req.ID)
		if err != nil {
			srv.error(w, r, http.StatusNotFound, err)
			return
		}

		list := make([]string, 2)
		list[0] = "id"
		list[1] = "snippet"

		call := srv.youtube.Search.List(list).
			ChannelId(channel.ChannelID).
			Type("video").
			MaxResults(int64(MaxResult))
		response, err := call.Do()
		if err != nil {
			srv.error(w, r, http.StatusBadRequest, err)
			return
		}

		var resultSch []*model.Video

		for _, item := range response.Items {
			sch := &model.Video{}
			sch.VideoID = item.Id.VideoId
			sch.VideoTitle = item.Snippet.Title
			sch.PublishDate = item.Snippet.PublishedAt
			sch.Description = item.Snippet.Description
			resultSch = append(resultSch, sch)
			srv.store.Repo().CreateVideos(sch)
		}

		srv.respond(w, r, http.StatusFound, resultSch)
	}
}

func (srv *server) handlePlaylistSearch() http.HandlerFunc {
	type request struct {
		ID int `json:"id,string"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			srv.error(w, r, http.StatusBadRequest, err)
			return
		}

		channel, err := srv.store.Repo().FindChannel(req.ID)
		if err != nil {
			srv.error(w, r, http.StatusNotFound, err)
			return
		}

		list := make([]string, 4)
		list[0] = "id"
		list[1] = "snippet"
		list[2] = "player"
		list[3] = "contentDetails"

		call := srv.youtube.Playlists.List(list).
			ChannelId(channel.ChannelID).
			MaxResults(int64(MaxResult))

		response, err := call.Do()
		if err != nil {
			srv.error(w, r, http.StatusBadRequest, err)
			return
		}

		var resultSch []*model.Playlist

		for _, item := range response.Items {
			sch := &model.Playlist{}
			sch.PlaylistID = item.Id
			sch.PlaylistTitle = item.Snippet.Title
			sch.EmbededHTML = item.Player.EmbedHtml
			sch.VideoCount = int(item.ContentDetails.ItemCount)
			resultSch = append(resultSch, sch)
			srv.store.Repo().CreatePlaylist(sch)
		}

		srv.respond(w, r, http.StatusFound, resultSch)
	}
}

func (srv *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	srv.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (srv *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
