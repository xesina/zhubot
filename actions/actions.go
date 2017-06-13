package actions

import "github.com/eternnoir/go-wit-ai"

var Wit *gowitai.WitClient
var witConfidenceThreshold float64


func init() {
	token := "YOUR_WIT_TOKEN"
	Wit = gowitai.NewWitClient(token, nil, nil)
	Wit.Debug = true
	witConfidenceThreshold = 0.5
}
