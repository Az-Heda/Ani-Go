package scraper

type ScraperAnime struct {
	Id               string
	Title            string
	AlternativeTitle string
	Aired            int64
	Duration         int
	Broadcast        int
	Url              string
	Season_ID        string
	Type_ID          string
	CurrentStatus    int
}

type ScraperDescription struct {
	Id           string
	Description  string
	Anime_ID     string
	Episode_ID   string
	Character_ID string
}

type ScraperCharacter struct {
	Id    string
	Name  string
	Url   string
	Image ScraperImage
}

type ScraperEpisode struct {
	Id       string
	Number   int
	Title    string
	Aired    int64
	Duration int64
	Anime_ID string
}

type ScraperGenre struct {
	Id   string
	Name string
}

type ScraperTheme struct {
	Id   string
	Name string
}

type ScraperImage struct {
	Id  string
	Url string
}

type ScraperSeason struct {
	Id     string
	Season string
}

type ScraperStudio struct {
	Id   string
	Name string
}

type ScraperTypes struct {
	Id   string
	Name string
}

type ScraperAnimeCharacter struct {
	Anime_ID     string
	Character_ID string
}

type ScraperAnimeType struct {
	Anime_ID string
	Type_ID  string
}

// type ScraperAnimeDescription struct {
// 	Anime_ID       string
// 	Description_ID string
// }

type ScraperAnimeGenre struct {
	Anime_ID string
	Genre_ID string
}

type ScraperAnimeTheme struct {
	Anime_ID string
	Theme_ID string
}

type ScraperAnimeImage struct {
	Anime_ID string
	Image_ID string
}

type ScraperAnimeStudio struct {
	Anime_ID  string
	Studio_ID string
}

// type ScraperCharacterDescription struct {
// 	Character_ID   string
// 	Description_ID string
// }

type ScraperCharacterImage struct {
	Character_ID string
	Image_ID     string
}

// type ScraperEpisodeDescription struct {
// 	Episode_ID     string
// 	Description_ID string
// }

type ScrapedAnime struct {
	Anime            ScraperAnime
	AnimeSeason      ScraperSeason
	AnimeType        ScraperTypes
	AnimeDescription []ScraperDescription
	AnimeStudios     []ScraperStudio
	AnimeGenre       []ScraperGenre
	AnimeTheme       []ScraperTheme

	Join_Anime_Studio []ScraperAnimeStudio
	Join_Anime_Genre  []ScraperAnimeGenre
	Join_Anime_Theme  []ScraperAnimeTheme
}

type ScrapedAnimePics struct {
	AnimeImages       []ScraperImage
	Join_Anime_Images []ScraperAnimeImage
}

type ScrapedAnimeCharacters struct {
	AnimeCharacters       []ScraperCharacter
	AnimeDescriptions     []ScraperDescription
	CharacterImages       []ScraperImage
	Join_Character_Image  []ScraperCharacterImage
	Join_Anime_Characters []ScraperAnimeCharacter
}

type ScrapedAnimeEpisodes struct {
	AnimeEpisodes     []ScraperEpisode
	AnimeDescriptions []ScraperDescription
}

type ScrapedAll struct {
	ScrapedAnime           ScrapedAnime
	ScrapedAnimePics       ScrapedAnimePics
	ScrapedAnimeCharacters ScrapedAnimeCharacters
	ScrapedAnimeEpisodes   ScrapedAnimeEpisodes
}
