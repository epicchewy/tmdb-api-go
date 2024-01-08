package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type WatchProvidersService interface {
	GetAvailableRegions(ctx context.Context, queryParams ...queryParam) (*WatchProvidersAvailableRegionsResponse, error)
	GetMovieProviders(ctx context.Context, queryParams ...queryParam) (*WatchProvidersMovieProvidersResponse, error)
	GetTVProviders(ctx context.Context, queryParams ...queryParam) (*WatchProvidersTVProvidersResponse, error)
}

type WatchProvidersClient struct {
	baseClient *Client
}

type WatchProvidersAvailableRegionsResponse struct {
	Results []struct {
		EnglishName string `json:"english_name"`
		Iso_3166_1  string `json:"iso_3166_1"`
		Name        string `json:"name"`
	} `json:"results"`
}

type DisplayPriorities struct {
	AE int32 `json:"AE"`
	AR int32 `json:"AR"`
	AT int32 `json:"AT"`
	AU int32 `json:"AU"`
	BE int32 `json:"BE"`
	BG int32 `json:"BG"`
	BO int32 `json:"BO"`
	BR int32 `json:"BR"`
	CA int32 `json:"CA"`
	CH int32 `json:"CH"`
	CL int32 `json:"CL"`
	CO int32 `json:"CO"`
	CR int32 `json:"CR"`
	CV int32 `json:"CV"`
	CZ int32 `json:"CZ"`
	DE int32 `json:"DE"`
	DK int32 `json:"DK"`
	EC int32 `json:"EC"`
	EE int32 `json:"EE"`
	EG int32 `json:"EG"`
	ES int32 `json:"ES"`
	FI int32 `json:"FI"`
	FR int32 `json:"FR"`
	GB int32 `json:"GB"`
	GH int32 `json:"GH"`
	GR int32 `json:"GR"`
	GT int32 `json:"GT"`
	HK int32 `json:"HK"`
	HN int32 `json:"HN"`
	HU int32 `json:"HU"`
	ID int32 `json:"ID"`
	IE int32 `json:"IE"`
	IL int32 `json:"IL"`
	IN int32 `json:"IN"`
	IT int32 `json:"IT"`
	JP int32 `json:"JP"`
	LT int32 `json:"LT"`
	LV int32 `json:"LV"`
	MX int32 `json:"MX"`
	MY int32 `json:"MY"`
	MU int32 `json:"MU"`
	MZ int32 `json:"MZ"`
	NL int32 `json:"NL"`
	NO int32 `json:"NO"`
	NZ int32 `json:"NZ"`
	PE int32 `json:"PE"`
	PH int32 `json:"PH"`
	PL int32 `json:"PL"`
	PT int32 `json:"PT"`
	PY int32 `json:"PY"`
	RU int32 `json:"RU"`
	SA int32 `json:"SA"`
	SE int32 `json:"SE"`
	SG int32 `json:"SG"`
	SI int32 `json:"SI"`
	SK int32 `json:"SK"`
	TH int32 `json:"TH"`
	TR int32 `json:"TR"`
	TW int32 `json:"TW"`
	UG int32 `json:"UG"`
	US int32 `json:"US"`
	VE int32 `json:"VE"`
	ZA int32 `json:"ZA"`
}

type WatchProvidersMovieProvidersResponse struct {
	Results []struct {
		DisplayPriorities DisplayPriorities `json:"display_priorities"`
		DisplayPriority   int32             `json:"display_priority"`
		LogoPath          string            `json:"logo_path"`
		ProviderID        int32             `json:"provider_id"`
		ProviderName      string            `json:"provider_name"`
	} `json:"results"`
}

type WatchProvidersTVProvidersResponse struct {
	Results []struct {
		DisplayPriorities DisplayPriorities `json:"display_priorities"`
		DisplayPriority   int32             `json:"display_priority"`
		LogoPath          string            `json:"logo_path"`
		ProviderID        int32             `json:"provider_id"`
		ProviderName      string            `json:"provider_name"`
	} `json:"results"`
}

func (wpc *WatchProvidersClient) GetAvailableRegions(ctx context.Context, queryParams ...queryParam) (*WatchProvidersAvailableRegionsResponse, error) {
	resp, err := wpc.baseClient.request(ctx, http.MethodGet, "/watch/providers/regions", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WatchProvidersAvailableRegionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (wpc *WatchProvidersClient) GetMovieProviders(ctx context.Context, queryParams ...queryParam) (*WatchProvidersMovieProvidersResponse, error) {
	resp, err := wpc.baseClient.request(ctx, http.MethodGet, "/watch/providers/movie", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WatchProvidersMovieProvidersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (wpc *WatchProvidersClient) GetTVProviders(ctx context.Context, queryParams ...queryParam) (*WatchProvidersTVProvidersResponse, error) {
	resp, err := wpc.baseClient.request(ctx, http.MethodGet, "/watch/providers/tv", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WatchProvidersTVProvidersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
