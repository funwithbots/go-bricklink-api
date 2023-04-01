package go_bricklink_api

import "github.com/funwithbots/go-bricklink-api/member"

// MemberAPI provides an interface for interacting with the Member API.
type MemberAPI interface {
	GetNote() (*member.Note, error)
	UpsertNote(note string) (*member.Note, error)
	DeleteNote() error
}
