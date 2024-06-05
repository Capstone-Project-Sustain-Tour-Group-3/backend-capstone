package handlers

import (
	"fmt"
	"net/http"

	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/usecases"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type profileHandler struct {
	usecase usecases.ProfileUsecase
}

func NewProfileHandler(uc usecases.ProfileUsecase) *profileHandler {
	return &profileHandler{usecase: uc}
}

func (h *profileHandler) GetDetailUser(ctx echo.Context) error {
	uid, ok := ctx.Get("userId").(*uuid.UUID)
	if !ok {
		return errorHandlers.HandleError(ctx, &errorHandlers.UnAuthorizedError{Message: "user tidak dikenali"})
	}

	user, err := h.usecase.GetDetailUser(*uid)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	findResponse := dto.ToFindByIdResponse(user)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data profil berhasil didapatkan",
		Data:       findResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *profileHandler) InsertUserDetail(ctx echo.Context) error {
	uid, ok := ctx.Get("userId").(*uuid.UUID)
	if !ok {
		return errorHandlers.HandleError(ctx, &errorHandlers.UnAuthorizedError{Message: "user tidak dikenali"})
	}

	var req dto.UserDetailRequest
	img, err := ctx.FormFile("foto_profil")
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "data yang dimasukkan tidak valid"})
	}
	req.FotoProfil = img
	if err = ctx.Bind(&req); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{Message: "data yang dimasukkan tidak valid"})
	}

	if err := helpers.ValidateRequest(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:  "failed",
			Message: "permintaan tidak valid. silakan periksa kembali data yang anda masukkan.",
			Errors:  err,
		})
	}

	fmt.Println(req)

	if err = h.usecase.InsertUserDetail(&req, *uid); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "data profil berhasil diperbarui",
	})
	return ctx.JSON(http.StatusOK, response)
}
