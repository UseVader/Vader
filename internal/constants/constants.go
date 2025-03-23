package constants

const (
	API_URL = "http://localhost:3170/v1"
	FE_URL  = "http://localhost:5173"
)

func URL(path string) string {
	return API_URL + "/" + path
}

func FEURL(path string) string {
	return FE_URL + "/" + path
}
