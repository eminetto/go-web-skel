package model

import (
	valid "github.com/asaskevich/govalidator"
	"gitlab.com/thecodenation/thecodenation/model"
	"testing"
)

func TestApplicantValidation(t *testing.T) {
	NewApplicant := model.Applicant{
		Nickname:    "eminetto",
		Picture:     "https://avatars0.githubusercontent.com/u/197939?:3",
		Email:       "eminetto@gmail.com",
		Name:        "Elton Minetto",
		Phone:       "479918962345",
		CurrentJob:  "Coderockr",
		Course:      "Ciencia da Computação",
		Semester:    "Primeiro",
		DevYears:    "1 ano",
		ResumeURL:   "file:///Users/eminetto/Documents/Projects/thecodenation/data/resume/eminetto/eminetto.pdf",
		CoverLetter: "teste de teste",
		Orientation: "Company Oriented - Family Oriented",
		Aptitude:    "Data Science",
		Linkedin:    "http://linkedin.com/eminetto",
		LikeToWork:  "Abroad",
	}
	_, err := valid.ValidateStruct(NewApplicant)
	if err != nil {
		t.Errorf("expected %s result %s", nil, err)
	}
}

func TestApplicantBatchValidation(t *testing.T) {
	ApplicantBatch := model.ApplicantBatch{
		Salary:  10000.50,
		BatchID: 1,
	}
	_, err := valid.ValidateStruct(ApplicantBatch)
	if err != nil {
		t.Errorf("expected %s result %s", nil, err)
	}
}
