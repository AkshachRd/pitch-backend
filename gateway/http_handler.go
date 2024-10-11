package main

import (
	"errors"
	"net/http"

	"github.com/AkshachRd/pitch-backend/common"
	pb "github.com/AkshachRd/pitch-backend/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	client pb.CardServiceClient
}

func NewHandler(client pb.CardServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/users/{userId}/cards", h.HandleCreateCard)
}

func (h *handler) HandleCreateCard(w http.ResponseWriter, r *http.Request) {
	// userId := r.PathValue("userId") // TODO: validate userId

	card := &pb.Card{}
	if err := common.ReadJSON(r, card); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateCard(card); err != nil {
	 	common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	c, err := h.client.CreateCard(r.Context(), &pb.CreateCardRequest{
		Card: card,
	})
	rStatus := status.Convert(err)
	if err != nil {
		if rStatus != nil {
			if rStatus.Code() != codes.InvalidArgument {
				common.WriteError(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusCreated, c)
}

func validateCard(card *pb.Card) error {
	if card.Id == "" {
		return errors.New("card id is required")
	}

	if card.FrontSide == "" {
		return errors.New("front side is required")
	}

	if card.BackSide == "" {
		return errors.New("back side is required")
	}

	if card.UserId == "" {
		return errors.New("user id is required")
	}

	return nil
}