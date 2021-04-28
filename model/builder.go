package model

import "encoding/json"

type Builder interface {
	Build(map[string]interface{}) *MachineReview
}

type ContactBuilder struct{}

func NewContactBuilder() *ContactBuilder {
	return &ContactBuilder{}
}

func (ContactBuilder) Build(m map[string]interface{}) *MachineReview {
	review := &MachineReview{}
	review.OrderID = m["OrderId"].(string)

	emerContacts := make([]*UserEmerContacts, 0)
	emerInfo := m["EmergencyInfo"].(string)
	_ = json.Unmarshal([]byte(emerInfo), &emerContacts)
	review.UserEmerContacts = emerContacts


	contacts := make([]*Contact, 0)
	cont := m["Contact"].(string)
	_ = json.Unmarshal([]byte(cont), &contacts)
	review.DeviceInfo = &DeviceInfo{}
	review.DeviceInfo.Contact = contacts
	return review
}


