package response

type HelloResponse struct {
	ClientIp string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}
