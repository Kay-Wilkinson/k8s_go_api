package models

import "gorm.io/gorm"

type Fact struct {
    gorm.Model
    Question string `json:"question" gorm:"text;not null;default:null"`
    Answer string `json:"answer" gorm:"text;not null;default:null"`
}

// Struct tags are smll pieces of metadata attached to fields of a struct. The Fact struct
// has JSON keys corresponding with each field of the Fact struct.
// Each column has the declared type of text, neither column can be null.
// Initially set to null so validation can fail if values are not supplied by user during write.