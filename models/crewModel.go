package models

type Crew struct {
	Name      string   `json:"name"`
	Agency    string   `json:"agency"`
	Image     string   `json:"image"`
	Wikipedia string   `json:"wikipedia"`
	Launches  []string `json:"launches"`
	Status    string   `json:"status"`
	ID        string   `json:"id"`
}
