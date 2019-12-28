package model

type Error struct {
	Code        int32  `json:"code,omitempty"`
	Description string `json:"description,omitmpty"`
}
