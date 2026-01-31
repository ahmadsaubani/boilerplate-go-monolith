package ConfigStructs

import (
	"boilerplate-go/Modules"
	"boilerplate-go/Services/Emails"
)

type Utilities struct {
	Email   Emails.EmailServices
	Modules Modules.Modules
}
