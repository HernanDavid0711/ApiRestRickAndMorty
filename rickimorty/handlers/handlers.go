package handlers

import (
	"fmt"
	"net/http"
	"prubarickmorti/db"
	"prubarickmorti/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CharactersLista(ctx *gin.Context) {
	data, err := db.CharactersLista()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	var resp []models.CharacterName

	for i := 0; i < len(data); i++ {
		curr := models.CharacterName{
			ID:    data[i].Id,
			Name:  data[i].Name,
			Photo: data[i].ImageUrl,
		}
		resp = append(resp, curr)
	}

	ctx.JSON(http.StatusOK, resp)
}

func Characters(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	data, err := db.Characters(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, data)
}

func SyncCharacters(ctx *gin.Context) {
	err := db.SyncCharacters()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, models.StandardResponse{
		Message: "Ok",
	})
}

func SyncEpisodes(ctx *gin.Context) {
	err := db.SyncEpisodes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, models.StandardResponse{
		Message: "Ok",
	})
}

func DelCharacters(ctx *gin.Context) {
	err := db.DelCharacters()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, models.StandardResponse{
		Message: "Ok",
	})
}
