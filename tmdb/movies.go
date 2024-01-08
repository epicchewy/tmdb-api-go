package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type MoviesService interface {
	GetDetails(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieDetailsResponse, error)
	GetAccountStates(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieAccountStatesResponse, error)
	GetAlternativeTitles(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieAlternativeTitlesResponse, error)
	GetChanges(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieChangesResponse, error)
	GetCredits(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieCreditsResponse, error)
	GetExternalIDs(ctx context.Context, movieID int) (*MovieExternalIDsResponse, error)
	GetImages(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieImagesResponse, error)
	GetKeywords(ctx context.Context, movieID int) (*MovieKeywordsResponse, error)
	GetLatest(ctx context.Context) (*MovieLatestResponse, error)
	GetLists(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieListsResponse, error)
	GetRecommendations(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieRecommendationsResponse, error)
	GetReleaseDates(ctx context.Context, movieID int) (*MovieReleaseDatesResponse, error)
	GetReviews(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieReviewsResponse, error)
	GetSimilar(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieSimilarMoviesResponse, error)
	GetTranslations(ctx context.Context, movieID int) (*MovieTranslationsResponse, error)
	GetVideos(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieVideosResponse, error)
	GetWatchProviders(ctx context.Context, movieID int) (*MovieWatchProvidersResponse, error)
}

type MoviesClient struct {
	baseClient *Client
}

type MovieDetailsResponse struct {
	Adult               bool   `json:"adult"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection string `json:"belongs_to_collection"`
	Budget              int    `json:"budget"`
	Genres              []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string  `json:"homepage"`
	ID                  int     `json:"id"`
	ImdbID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso_3166_1 string `json:"iso_3166_1"`
		Name       string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int    `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso_639_1   string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

type MovieAccountStatesResponse struct {
	ID       int  `json:"id"`
	Favorite bool `json:"favorite"`
	Rated    struct {
		Value float64 `json:"value"`
	} `json:"rated"`
	Watchlist bool `json:"watchlist"`
}

type MovieAlternativeTitlesResponse struct {
	ID     int `json:"id"`
	Titles []struct {
		Iso_3166_1 string `json:"iso_3166_1"`
		Title      string `json:"title"`
		Type       string `json:"type"`
	} `json:"titles"`
}

// TODO: figure out this response
type MovieChangesResponse struct {
	Changes []struct{} `json:"changes"`
}

type MovieCreditsResponse struct {
	ID   int `json:"id"`
	Cast []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		CastID             int32   `json:"cast_id"`
		Character          string  `json:"character"`
		CreditID           string  `json:"credit_id"`
		Order              int32   `json:"order"`
	} `json:"cast"`
	Crew []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		CreditID           string  `json:"credit_id"`
		Department         string  `json:"department"`
		Job                string  `json:"job"`
	} `json:"crew"`
}

type MovieExternalIDsResponse struct {
	ID          int    `json:"id"`
	ImdbID      string `json:"imdb_id"`
	WikidataID  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
}

type MovieImage struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	Iso_639_1   string  `json:"iso_639_1"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}

type MovieImagesResponse struct {
	ID        int          `json:"id"`
	Backdrops []MovieImage `json:"backdrops"`
	Logos     []MovieImage `json:"logos"`
	Posters   []MovieImage `json:"posters"`
}

type MovieKeywordsResponse struct {
	ID       int `json:"id"`
	Keywords []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"keywords"`
}

type MovieLatestResponse struct {
	Adult               bool   `json:"adult"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection string `json:"belongs_to_collection"`
	Budget              int    `json:"budget"`
	// Genres              []Genre `json:"genres"` TODO figure this out
	Homepage            string  `json:"homepage"`
	ID                  int     `json:"id"`
	ImdbID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso_3166_1 string `json:"iso_3166_1"`
		Name       string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int    `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso_639_1   string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

type MovieListsResponse struct {
	ID      int `json:"id"`
	Page    int `json:"page"`
	Results []struct {
		Description   string `json:"description"`
		FavoriteCount int    `json:"favorite_count"`
		ID            int    `json:"id"`
		ItemCount     int    `json:"item_count"`
		Iso_639_1     string `json:"iso_639_1"`
		ListType      string `json:"list_type"`
		Name          string `json:"name"`
		PosterPath    string `json:"poster_path"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type MovieRecommendationsResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		MediaType        string  `json:"media_type"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type MovieReleaseDatesResponse struct {
	ID      int `json:"id"`
	Results []struct {
		Iso_3166_1   string `json:"iso_3166_1"`
		ReleaseDates []struct {
			Certification string `json:"certification"`
			// Descriptors []struct {} `json:"descriptors"` // TODO figure this out
			Iso_639_1   string `json:"iso_639_1"`
			Note        string `json:"note"`
			ReleaseDate string `json:"release_date"`
			Type        int    `json:"type"`
		} `json:"release_dates"`
	} `json:"results"`
}

type MovieReviewsResponse struct {
	ID      int `json:"id"`
	Page    int `json:"page"`
	Results []struct {
		Author        string `json:"author"`
		AuthorDetails struct {
			Name       string  `json:"name"`
			Username   string  `json:"username"`
			AvatarPath string  `json:"avatar_path"`
			Rating     float64 `json:"rating"`
		} `json:"author_details"`
		Content   string `json:"content"`
		CreatedAt string `json:"created_at"`
		ID        string `json:"id"`
		UpdatedAt string `json:"updated_at"`
		URL       string `json:"url"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type MovieSimilarMoviesResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDS         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type MovieTranslationsResponse struct {
	ID           int `json:"id"`
	Translations []struct {
		Iso_3166_1  string `json:"iso_3166_1"`
		Iso_639_1   string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Title    string `json:"title"`
			Overview string `json:"overview"`
			Homepage string `json:"homepage"`
			Tagline  string `json:"tagline"`
			Runtime  int    `json:"runtime"`
		} `json:"data"`
	} `json:"translations"`
}

type MovieVideosResponse struct {
	ID      int `json:"id"`
	Results []struct {
		ID          string `json:"id"`
		Iso_639_1   string `json:"iso_639_1"`
		Iso_3166_1  string `json:"iso_3166_1"`
		Key         string `json:"key"`
		Name        string `json:"name"`
		Site        string `json:"site"`
		Size        int    `json:"size"`
		Type        string `json:"type"`
		Official    bool   `json:"official"`
		PublishedAt bool   `json:"published_at"`
	} `json:"results"`
}

type WatchProvider struct {
	DisplayPriority int    `json:"display_priority"`
	LogoPath        string `json:"logo_path"`
	ProviderID      int    `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
}

type MovieWatchProvidersResponse struct {
	ID      int `json:"id"`
	Results struct {
		AE struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"AE"`
		AL struct {
			Link string          `json:"link"`
			Buy  []WatchProvider `json:"buy"`
		} `json:"AL"`
		AR struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"AR"`
		AT struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"AT"`
		AU struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"AU"`
		BA struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"BA"`
		BB struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"BB"`
		BE struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"BE"`
		BG struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"BG"`
		BH struct {
			Link string          `json:"link"`
			Buy  []WatchProvider `json:"buy"`
		} `json:"BH"`
		BO struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"BO"`
		BR struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"BR"`
		BS struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"BS"`
		CA struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"CA"`
		CH struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"CH"`
		CL struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"CL"`
		CO struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"CO"`
		CR struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"CR"`
		CY struct {
			Link string          `json:"link"`
			Buy  []WatchProvider `json:"buy"`
			Rent []WatchProvider `json:"rent"`
		} `json:"CY"`
		CZ struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"CZ"`
		DE struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"DE"`
		DK struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"DK"`
		DO struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"DO"`
		EC struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"EC"`
		EE struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"EE"`
		EG struct {
			Link string          `json:"link"`
			Buy  []WatchProvider `json:"buy"`
			Rent []WatchProvider `json:"rent"`
		} `json:"EG"`
		ES struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
			Ads      []WatchProvider `json:"ads"`
		} `json:"ES"`
		FI struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		}
		FJ struct {
			Link string          `json:"link"`
			Buy  []WatchProvider `json:"buy"`
		} `json:"FJ"`
		FR struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"FR"`
		GB struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"GB"`
		GF struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"GF"`
		GI struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"GI"`
		GR struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"GR"`
		GT struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"GT"`
		HK struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"HK"`
		HN struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"HN"`
		HR struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Ads      []WatchProvider `json:"ads"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"HR"`
		HU struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"HU"`
		ID struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"ID"`
		IE struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"IE"`
		IL struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"IL"`
		IN struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"IN"`
		IQ struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"IQ"`
		IS struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"IS"`
		IT struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"IT"`
		JM struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"JM"`
		JO struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"JO"`
		JP struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"JP"`
		KR struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"KR"`
		KW struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"KW"`
		LB struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"LB"`
		LI struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"LI"`
		LT struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"LT"`
		LV struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"LV"`
		MD struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"MD"`
		MK struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"MK"`
		MT struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"MT"`
		MU struct {
			Link string          `json:"link"`
			Buy  []WatchProvider `json:"buy"`
			Rent []WatchProvider `json:"rent"`
		} `json:"MU"`
		MX struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"MX"`
		MY struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"MY"`
		MZ struct {
			Link string          `json:"link"`
			Buy  []WatchProvider `json:"buy"`
			Rent []WatchProvider `json:"rent"`
		} `json:"MZ"`
		NL struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"NL"`
		NO struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"NO"`
		NZ struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"NZ"`
		OM struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"OM"`
		PA struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"PA"`
		PE struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"PE"`
		PH struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"PH"`
		PK struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"PK"`
		PL struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"PL"`
		PS struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"PS"`
		PT struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"PT"`
		PY struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"PY"`
		QA struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"QA"`
		RO struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"RO"`
		RS struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"RS"`
		RU struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"RU"`
		SA struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"SA"`
		SE struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"SE"`
		SG struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"SG"`
		SI struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"SI"`
		SK struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"SK"`
		SM struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"SM"`
		SV struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"SV"`
		TH struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"TH"`
		TR struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"TR"`
		TT struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"TT"`
		TW struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"TW"`
		UG struct {
			Link string          `json:"link"`
			Buy  []WatchProvider `json:"buy"`
			Rent []WatchProvider `json:"rent"`
		} `json:"UG"`
		US struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"US"`
		UY struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"UY"`
		VE struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"VE"`
		YE struct {
			Link     string          `json:"link"`
			Flatrate []WatchProvider `json:"flatrate"`
		} `json:"YE"`
		ZA struct {
			Link     string          `json:"link"`
			Buy      []WatchProvider `json:"buy"`
			Flatrate []WatchProvider `json:"flatrate"`
			Rent     []WatchProvider `json:"rent"`
		} `json:"ZA"`
	} `json:"results"`
}

func (c *MoviesClient) GetDetails(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieDetailsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d", movieID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var details MovieDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, err
	}
	return &details, nil
}

func (c *MoviesClient) GetAccountStates(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieAccountStatesResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/account_states", movieID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var states MovieAccountStatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&states); err != nil {
		return nil, err
	}
	return &states, nil
}

func (c *MoviesClient) GetAlternativeTitles(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieAlternativeTitlesResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/alternative_titles", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var titles MovieAlternativeTitlesResponse
	if err := json.NewDecoder(resp.Body).Decode(&titles); err != nil {
		return nil, err
	}
	return &titles, nil
}

func (c *MoviesClient) GetChanges(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieChangesResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/changes", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var changes MovieChangesResponse
	if err := json.NewDecoder(resp.Body).Decode(&changes); err != nil {
		return nil, err
	}
	return &changes, nil
}

func (c *MoviesClient) GetCredits(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieCreditsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/credits", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var credits MovieCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&credits); err != nil {
		return nil, err
	}
	return &credits, nil
}

func (c *MoviesClient) GetExternalIDs(ctx context.Context, movieID int) (*MovieExternalIDsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/external_ids", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ids MovieExternalIDsResponse
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}
	return &ids, nil
}

func (c *MoviesClient) GetImages(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieImagesResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/images", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var images MovieImagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
		return nil, err
	}
	return &images, nil
}

func (c *MoviesClient) GetKeywords(ctx context.Context, movieID int) (*MovieKeywordsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/keywords", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var keywords MovieKeywordsResponse
	if err := json.NewDecoder(resp.Body).Decode(&keywords); err != nil {
		return nil, err
	}
	return &keywords, nil
}

func (c *MoviesClient) GetLatest(ctx context.Context) (*MovieLatestResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, "/movie/latest")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var latest MovieLatestResponse
	if err := json.NewDecoder(resp.Body).Decode(&latest); err != nil {
		return nil, err
	}
	return &latest, nil
}

func (c *MoviesClient) GetLists(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieListsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/lists", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lists MovieListsResponse
	if err := json.NewDecoder(resp.Body).Decode(&lists); err != nil {
		return nil, err
	}
	return &lists, nil
}

func (c *MoviesClient) GetRecommendations(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieRecommendationsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/recommendations", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recommendations MovieRecommendationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&recommendations); err != nil {
		return nil, err
	}
	return &recommendations, nil
}

func (c *MoviesClient) GetReleaseDates(ctx context.Context, movieID int) (*MovieReleaseDatesResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/release_dates", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dates MovieReleaseDatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&dates); err != nil {
		return nil, err
	}
	return &dates, nil
}

func (c *MoviesClient) GetReviews(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieReviewsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/reviews", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reviews MovieReviewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&reviews); err != nil {
		return nil, err
	}
	return &reviews, nil
}

func (c *MoviesClient) GetSimilar(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieSimilarMoviesResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/similar", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var similar MovieSimilarMoviesResponse
	if err := json.NewDecoder(resp.Body).Decode(&similar); err != nil {
		return nil, err
	}
	return &similar, nil
}

func (c *MoviesClient) GetTranslations(ctx context.Context, movieID int) (*MovieTranslationsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/translations", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var translations MovieTranslationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&translations); err != nil {
		return nil, err
	}
	return &translations, nil
}

func (c *MoviesClient) GetVideos(ctx context.Context, movieID int, queryParams ...queryParam) (*MovieVideosResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/videos", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var videos MovieVideosResponse
	if err := json.NewDecoder(resp.Body).Decode(&videos); err != nil {
		return nil, err
	}
	return &videos, nil
}

func (c *MoviesClient) GetWatchProviders(ctx context.Context, movieID int) (*MovieWatchProvidersResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/movie/%d/watch/providers", movieID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var providers MovieWatchProvidersResponse
	if err := json.NewDecoder(resp.Body).Decode(&providers); err != nil {
		return nil, err
	}
	return &providers, nil
}
