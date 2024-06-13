package helpers

import (
	"fmt"
	"math/rand"

	"capstone/entities"
)

func GetRecommendationSystemInstruction() string {
	return "Kamu adalah seorang pemandu wisata yang bisa memberikan rekomendasi tempat wisata yang belum dikunjungi user. Kelompokkan list 'sudah dikunjungi' berdasarkan nuasanya. Lalu, semakin banyak destinasi dalam kelompok nuansa 'sudah dikunjungi', maka semakin diutamakan destinasi dengan nuansa yang sama untuk direkomendasikan kepada user.\n\nCatatan:\n- data 'sudah dikunjungi' yang diberikan user berupa list nama destinasi menggunakan format csv (uuid,destinasi)\n- respon berupa 5 UUID SAJA yang berdasarkan pada data 'belum dikunjungi' yang diberikan user, berdasarkan probabilitas tertinggi dengan tanpa kata pengantar maupun penutup\n- jika data 'belum dikunjungi' kosong, berikan rekomendasi secara acak yang hanya ada pada data 'sudah dikunjungi' yang diberikan user\n\nContoh respon: 324de9db-437d-45b5-a12b-cafbd69d57dd\nbab7e381-07ff-4702-8afc-fea597129548"
}

func ToRecommendationPrompt(visitedDestinations, unVisitedDestinations *[]entities.Destination) string {
	var (
		visitedPrompt   = "sudah dikunjungi:\n"
		unvisitedPrompt = "\nbelum dikunjungi:\n"
	)

	visitedPrompt += toListDestinationPrompt(visitedDestinations)
	unvisitedPrompt += toListDestinationPrompt(unVisitedDestinations)

	prompt := visitedPrompt + unvisitedPrompt

	return prompt
}

func SelectRecommendationIds(visited, unVisited *[]entities.Destination) *[]string {
	var selectedDestinations []string

	for i := range *unVisited {
		selectedDestinations = append(selectedDestinations, (*unVisited)[i].Id.String())
	}

	counter := len(*unVisited)
	lenVisited := len(*visited)

	limit := 5
	if lenVisited+counter < limit {
		limit = len(*visited)
	}

	randIds := rand.Perm(lenVisited)
	for i := range limit - counter {
		selectedDestinations = append(selectedDestinations, (*visited)[randIds[i]].Id.String())
	}

	return &selectedDestinations
}

func toListDestinationPrompt(destinations *[]entities.Destination) string {
	var prompt string
	for _, destination := range *destinations {
		prompt += fmt.Sprintf("- %s,%s\n", destination.Id.String(), destination.Name)
	}

	return prompt
}
