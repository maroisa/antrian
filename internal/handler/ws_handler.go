package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

func (s *Server) WSHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "Gagal mendapatkan koneksi!", http.StatusInternalServerError)
		return
	}
	defer c.Close(websocket.StatusInternalError, "Koneksi terputus.")

	s.mu.Lock()
	s.clients[c] = true
	s.mu.Unlock()

	log.Println("user connected")

	ctx := r.Context()

	allLoket, err := s.db.GetAllLoket(ctx)
	s.Broadcast(ctx, map[string]interface{}{
		"data": allLoket,
	})

	<-ctx.Done()

	log.Println("user disconnected")

	s.mu.Lock()
	delete(s.clients, c)
	s.mu.Unlock()
}

func (s *Server) Broadcast(ctx context.Context, msg any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for client := range s.clients {
		err := wsjson.Write(ctx, client, msg)
		if err != nil {
			log.Println(err)
			client.Close(websocket.StatusInternalError, "Gagal menerima pesan")
			delete(s.clients, client)
		}
	}
}
