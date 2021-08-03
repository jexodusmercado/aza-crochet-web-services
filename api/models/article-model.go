package models

type Article struct {
    ID      		string      `json:"id,omitempty"`
    Title   		string      `json:"title,omitempty"`
    Description    	string      `json:"desc,omitempty"`
    Content 		string      `json:"content,omitempty"`
}
