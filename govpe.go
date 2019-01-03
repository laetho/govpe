/*
Copyright 2019 The govpe Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package govpe

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func (obj *Endpoint) MarshalJSON() (error, string) {
	var out bytes.Buffer
	ba, err := json.Marshal(obj)
	if err != nil {
		return err, "Error"
	}
	err = json.Indent(&out, ba, "", "  ")
	if err != nil {
		return err, "Error"
	}
	return nil, out.String()
}

func ParseConsumer(file string) (error, *Consumer) {

	return nil, nil
}

func ParseEndpoint(file string) (error, *Endpoint) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err, &Endpoint{}
	}

	err = nil
	var ep Endpoint

	err = json.Unmarshal(b, &ep)
	if err != nil {
		return err, &Endpoint{}
	}
	return nil, &ep
}
