package vpe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
)

type PolicyEngineModule string

//const UpStreamSimpleModule PolicyEngineModule = "upstream-simple"
//const AuthApiKeyModule PolicyEngineModule = "auth-apikey"

type Endpoint struct {
	Name    string    `json:"name"`
	Host    string    `json:"host"`
	Path    string    `json:"string"`
	Modules []*Module `json:"modules"`
}

type Module struct {
	Name      PolicyEngineModule    `json:"name"`
	Order     int                   `json:"order"`
	Consumers map[string]*Arguments `json:"consumers,omitempty"`
	General   *General              `json:"general,omitempty"`
	Default   *Arguments            `json:"default,omitempty"`
}

// General options?
type General struct {
	Optional bool   `json:"optional,omitempty"`
	Realm    string `json:"realm,omitempty"`
	// HMAC
	Skew string `json:"skew,omitempty"`
	// oAuth2
	CallBack string `json:"callback,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	// Client secret from provider
	Secret   string `json:"secret,omitempty"`
	Provider string `json:"google,omitempty"`
	// Custom JWT signing secret
	JWTSecret string `json:"jwt_secret,omitempty"`
	TTL       string `json:"ttl,omitempty"`
	Cookie    string `json:"cookie,omitempty"`
}

type Arguments struct {
	// Endpoint backend related
	Url          string `json:"url,omitempty"`
	PreserveHost bool   `json:"preserve_host,omitempty"`
	StripPath    bool   `json:"strip_path,omitempty"`
	// Authentication related
	Secret   string `json:"secret,omitempty"`
	Password string `json:"password,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	// IP or HTTP verb allow deny
	Order string `json:"order,omitempty"`
	Allow string `json:"allow,omitempty"`
	Deny  string `json:"allow,omitempty"`
	// Rate limit arguments
	Hour   int `json:"hour,omitempty"`
	Minute int `json:"minute,omitempty"`
	Second int `json:"second,omitempty"`
}

type Consumer struct {
	Name  string `json:"name"`
	EMail string `json:"email,omitempty"`
}

func (obj *Endpoint) MarshalJSON() string {
	var out bytes.Buffer
	ba, err := json.Marshal(obj)
	if err != nil {
		glog.Error(err)
	}
	err = json.Indent(&out, ba, "", "  ")
	if err != nil {
		glog.Error(err)
	}
	return out.String()
}

func ParseConsumer(f string) (error, *Consumer) {

	return nil, nil
}

func ParseEndpoint(f string) (error, *Endpoint) {
	fmt.Println("in parse", f)
	return nil, nil
}
