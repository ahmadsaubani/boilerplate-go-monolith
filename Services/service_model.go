package Services

import (
	"boilerplate-go/Services/Emails"
	"boilerplate-go/Services/News"
)

type service struct {
	Email Emails.EmailServices
	News  News.NewsServices
}
