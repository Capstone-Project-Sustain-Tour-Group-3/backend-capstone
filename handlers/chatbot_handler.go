package handlers

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type chatbotHandler struct {
	usecase  usecases.ChatbotUsecase
	upgrader websocket.Upgrader
}

func NewChatbotHandler(usecase usecases.ChatbotUsecase, upgrader websocket.Upgrader) *chatbotHandler {
	return &chatbotHandler{
		usecase:  usecase,
		upgrader: upgrader,
	}
}

func (h *chatbotHandler) Chatbot(ctx echo.Context) error {
	uid, ok := ctx.Get("userId").(*uuid.UUID)
	if !ok {
		return errorHandlers.HandleError(
			ctx,
			&errorHandlers.UnAuthorizedError{Message: "user tidak dikenali"},
		)
	}

	if err := h.usecase.FlushChatHistory(*uid); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	ws, err := h.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, msg, readErr := ws.ReadMessage()
		if readErr != nil {
			ctx.Logger().Error(readErr)
			break
		}

		if len(msg) != 0 {
			request := &dto.ChatbotRequest{
				Message: string(msg),
			}

			res := h.usecase.GenerateAnswer(request, *uid)

			err = ws.WriteMessage(websocket.TextMessage, []byte(res.Answer))
			if err != nil {
				ctx.Logger().Error(err)
			}
		}
	}

	return nil
}
