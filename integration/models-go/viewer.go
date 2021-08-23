package models

import "github.com/00security/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
