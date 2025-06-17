package constants

type WatchlistRoutes struct {
	GET_WATCHLISTS   string
	CREATE_WATCHLIST string
}

var WATCHLIST_ROUTES = WatchlistRoutes{
	GET_WATCHLISTS:   "/watchlists",
	CREATE_WATCHLIST: "/watchlists",
}
