package api

type QouteAnime struct {
	Anime     string
	Character string
	Qoute     string
}

const ANIME_API_URL = "https://animechan.vercel.app/api/"

// GetAllAnime returns all available anime
func GetAllAnime() QouteAnime {
	var quoteanime QouteAnime

	return quoteanime
}

// GetSpecificAnime returns a specific anime
func GetSpecificAnime(anime string) string {

	return ""
}

//GetRandom returns a random quote
func GetRandom() QouteAnime {
	var quoteanime = GetSingleRequest(ANIME_API_URL + "random")
	return quoteanime
}

// GetTenRandom returns 10 random quotes
func GetTenRandom() []QouteAnime {
	var quotesanime = GetRequest(ANIME_API_URL + "quotes")
	return quotesanime

}

// GetByAnime returns quotes from a specific anime
func GetByAnime(anime string) []QouteAnime {
	var quotesanime = GetRequest(ANIME_API_URL + "quotes/anime?title=" + anime)
	return quotesanime

}

// GetByCharacter returns quotes by anime character
func GetByCharacter(character string) []QouteAnime {
	var quotesanime = GetRequest(ANIME_API_URL + "quotes/character?name=" + character)
	return quotesanime
}
