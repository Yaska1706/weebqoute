package api

import "fmt"

type QouteAnime struct {
	Anime     string
	Character string
	Qoute     string
}

const ANIME_API_URL = "https://animechan.vercel.app/api/"

// GetAllAnime returns all available anime
func GetAllAnime() []string {

	var animes = GetAllAnimeRequest(ANIME_API_URL + "available/anime")
	return animes
}

// GetSpecificAnime returns a specific anime
func GetSpecificAnime(animeTitle string) string {
	var result string
	animes := GetAllAnime()
	for _, anime := range animes {
		if animeTitle == anime {
			result = fmt.Sprintf("%v found", animeTitle)

		} else {
			var ErrAnimeNotFound = "Anime not Found"
			result = ErrAnimeNotFound
		}

	}
	return result
}

//GetRandom returns a random quote
func GetRandom() QouteAnime {
	var quoteanime = GetSingleRequest(ANIME_API_URL + "random")
	return quoteanime
}

// GetTenRandom returns 10 random quotes
func GetTenRandom() []QouteAnime {
	var quotesanime = GetAllRequest(ANIME_API_URL + "quotes")
	return quotesanime

}

// GetByAnime returns quotes from a specific anime
func GetByAnime(anime string) []QouteAnime {
	var quotesanime = GetAllRequest(ANIME_API_URL + "quotes/anime?title=" + anime)
	return quotesanime

}

// GetByCharacter returns quotes by anime character
func GetByCharacter(character string) []QouteAnime {
	var quotesanime = GetAllRequest(ANIME_API_URL + "quotes/character?name=" + character)
	return quotesanime
}
