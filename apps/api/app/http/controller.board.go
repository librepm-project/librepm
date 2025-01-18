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
	boards, err := c.BoardService.All()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, BoardSerializer{}.ModelToResponseMany(*boards))
}

func (c BoardController) Show(w http.ResponseWriter, r *http.Request) {
	var board_id, _ = http_utils.GetParamUUID(r, "board_id")
	board, err := c.BoardService.Show(board_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, BoardSerializer{}.ModelToResponse(*board))
}

func (c BoardController) Create(w http.ResponseWriter, r *http.Request) {
	var board_request BoardRequest
	json.NewDecoder(r.Body).Decode(&board_request)
	board := BoardSerializer{}.RequestToModel(board_request)
	err := c.BoardService.Create(&board)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusCreated, BoardSerializer{}.ModelToResponse(board))
}

func (c BoardController) Update(w http.ResponseWriter, r *http.Request) {
	board_id, _ := http_utils.GetParamUUID(r, "board_id")
	var board_request BoardRequest
	json.NewDecoder(r.Body).Decode(&board_request)
	board := BoardSerializer{}.RequestToModel(board_request)
	err := c.BoardService.Update(board_id, &board)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c BoardController) Destroy(w http.ResponseWriter, r *http.Request) {
	board_id, _ := http_utils.GetParamUUID(r, "board_id")
	err := c.BoardService.Destroy(board_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
