package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
)

type BoardControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type BoardController struct {
	BoardService domain.BoardServiceInterface
}

func (c BoardController) Index(w http.ResponseWriter, r *http.Request) {
	boards := c.BoardService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, BoardSerializer{}.SerializeBoards(*boards))
}

func (c BoardController) Show(w http.ResponseWriter, r *http.Request) {
	var board_id, _ = http_utils.GetParamUUID(r, "board_id")
	board := c.BoardService.Show(board_id)
	http_utils.RespondWithJSON(w, http.StatusOK, BoardSerializer{}.SerializeBoard(*board))
}

func (c BoardController) Create(w http.ResponseWriter, r *http.Request) {
	var board domain.BoardModel
	json.NewDecoder(r.Body).Decode(&board)
	c.BoardService.Create(&board)
	http_utils.RespondWithJSON(w, http.StatusCreated, BoardSerializer{}.SerializeBoard(board))
}

func (c BoardController) Update(w http.ResponseWriter, r *http.Request) {
	board_id, _ := http_utils.GetParamUUID(r, "board_id")
	var board domain.BoardModel
	json.NewDecoder(r.Body).Decode(&board)
	c.BoardService.Update(board_id, &board)
	w.WriteHeader(http.StatusOK)
}

func (c BoardController) Destroy(w http.ResponseWriter, r *http.Request) {
	board_id, _ := http_utils.GetParamUUID(r, "board_id")
	c.BoardService.Destroy(board_id)
	w.WriteHeader(http.StatusNoContent)
}
