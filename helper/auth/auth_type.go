package auth

type AuthData interface {
	GetUid() string
	GetName() string
	GetToken() string
	GetClient() ClientData
	GetUser() UserData
}

type authData struct {
	Uid        string
	Name       string
	Token      string
	clientData ClientData
	userData   UserData
}

func NewAuthData(tkn string, clientData ClientData, userData UserData) AuthData {
	r := &authData{
		clientData: clientData,
		userData:   userData,
		Token:      tkn,
	}
	if clientData != nil {
		r.Uid = clientData.GetUid()
		r.Name = clientData.GetName()
	}
	if userData != nil {
		r.Name = userData.GetName()
	}
	return r
}

func (s *authData) GetUid() string {
	return s.Uid
}

func (s *authData) GetName() string {
	return s.Name
}

func (s *authData) GetToken() string {
	return s.Token
}

func (s *authData) GetClient() ClientData {
	return s.clientData
}

func (s *authData) GetUser() UserData {
	return s.userData
}
