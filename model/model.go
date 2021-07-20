package model

type Info struct {
	Message       string `json:"message,omitempty"`
	Branches      string `json:"branches,omitempty"`
	Employee      int    `json:"employee,omitempty"`
	ContainerName string `json:"container_name,omitempty"`
}
