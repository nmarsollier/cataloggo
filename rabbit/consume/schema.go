package consume

type ConsumeMessage struct {
	Version    int    `json:"version"`
	RoutingKey string `json:"routing_key" example:"Remote RoutingKey to Reply"`
	Exchange   string `json:"exchange"`
}
