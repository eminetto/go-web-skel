package model

import (
	"time"
)

//Applicant dados do candidato
type Applicant struct {
	ID          int64     `valid:"-" json: "id"`
	Nickname    string    `valid:"required" json: "nickname"`
	Picture     string    `valid:"url,required" json: "picture"`
	Email       string    `valid:"email,required" json: "email"`
	Name        string    `valid:"required" json: "name"`
	Phone       string    `valid:"required" json: "phone"`
	Birthday    time.Time `valid:"-" json: "birthday"`
	CurrentJob  string    `valid:"-" json: "current_job"`
	Course      string    `valid:"-" json: "course"`
	Semester    string    `valid:"-" json: "semester"`
	DevYears    string    `valid:"required" json: "dev_years"`
	ResumeURL   string    `valid:"required" json: "resume_url"`
	CoverLetter string    `valid:"required" json: "cover_letter"`
	Orientation string    `valid:"required" json: "orientation"`
	Aptitude    string    `valid:"required" json: "aptitude"`
	Linkedin    string    `valid:"url" json: "linkedin"`
	LikeToWork  string    `valid:"required" "json:"like_to_work"`
	Evaluation  Evaluation
}

//Company dados da empresa
type Company struct {
	ID          int64  `valid:"-"`
	Name        string `valid:"required"`
	Logo        string `valid:"url,required"`
	Email       string `valid:"email,required"`
	City        string `valid:"required"`
	URL         string `valid:"url,required"`
	Description string `valid:"-"`
}

//Batch dados do batch
type Batch struct {
	ID          int64     `valid:"-"`
	Name        string    `valid:"required"`
	DateStart   time.Time `valid:"-"`
	DateEnd     time.Time `valid:"-"`
	City        string    `valid:"required"`
	URL         string    `valid:"required"`
	Description string    `valid:"-"`
	Lang        string    `valid:"required"`
}

//CompanyBatch relação entre as entidades
type CompanyBatch struct {
	CompanyID int64 `valid:"required"`
	BatchID   int64 `valid:"required"`
	Positions int   `valid:"required"`
}

//ApplicantBatch relação entre as entidades
type ApplicantBatch struct {
	ApplicantID  int64   `valid:"-"`
	BatchID      int64   `valid:"required"`
	Salary       float64 `valid:"required"`
	CompanyID    int64   `valid:"-"` //se possui id é um funcionário da empresa
	SelectedBy   int64   `valid:"-"` //id da empresa que selecionou o candidato
	SuggestedTo  int64   `valid:"-"` //id da empresa que a curadoria indicou
	Observations string  `valid:"string"`
}

//User dados dos usuários que podem logar no sistema
type User struct {
	ID       int64  `valid:"-"`
	Name     string `valid:"required"`
	Picture  string `valid:"url,optional"`
	Email    string `valid:"email,required"`
	Password string `valid: required`
	IsAdmin  bool   `valid: required`
}

//UserCompany empresas que o usuário pode administrar
type UserCompany struct {
	UserID    int64 `valid:"required"`
	CompanyID int64 `valid:"required"`
}

//Evaluation avaliação do candidato
type Evaluation struct {
	QualityAssurance float64 `json: quality_assurance`
	FastLearning     float64 `json: fast_learning`
	ProblemSolving   float64 `json: problem_solving`
	PotentialIndex   float64 `json: potential_index`
}
