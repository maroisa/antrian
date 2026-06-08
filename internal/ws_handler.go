package internal

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

	log.Println("user baru terkoneksi")

	ctx := r.Context()
	s.Broadcast(ctx)

	for {
		msgType, msg, err := c.Read(ctx)
		if err != nil {
			log.Printf("Error membaca pesan: %v", err)
			break
		}

		if msgType != websocket.MessageText {
			log.Println("Permintaan tidak valid!")
			http.Error(w, "Permintaan tidak valid!", http.StatusInternalServerError)
			continue
		}

		s.data.Tambah(string(msg))
		s.Broadcast(ctx)
	}

	s.mu.Lock()
	delete(s.clients, c)
	s.mu.Unlock()
}

func (s *Server) Broadcast(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for client := range s.clients {
		err := wsjson.Write(ctx, client, s.data)
		if err != nil {
			log.Println(err)
			client.Close(websocket.StatusInternalError, "Gagal menerima pesan")
			delete(s.clients, client)
		}
	}
}
