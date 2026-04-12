package models

type Account struct {
    ID           uint
    Code         string
    CompanyName  string
	Package		 string
	UserLimit	 int
    ThemeColor	 string
    LogoURL      string
	IsActive	 bool
}