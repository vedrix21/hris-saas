package models

type Account struct {
    ID           uint
    Code         string
    CompanyName  string
	Package		 string
	UserLimit	 int
    IsActive	 bool
    
    LogoURL      string
    ThemeColor	 string
    Environment	 string

    Features     string // JSON string
	
}