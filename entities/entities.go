package entities

import "time"

type Log struct {
        Time time.Time
        MessageID string
        Service string
        PayloadType  string
        Payload interface{} //`json:"request_payload"`
}
type CTAResponse struct{
        Text interface{} `json:"text"`
        Id interface{}   `json:"id"`
        Link interface{}  `json:"link"`
        Typ interface{}   `json:"typ"`
        IsEnabled interface{} `json:"isEnabled"`
        Label interface{}  `json:"label"`
        CTA_Type interface{} `json:"cta_type"`
}

