// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/service/acs/v1"
)

func (dispatcher *EventReqDispatcher) OnAccessRecordCreatedV1(handler func(ctx context.Context, event *acs.AccessRecordCreatedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["acs.access_record.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "acs.access_record.created_v1")
	}
	dispatcher.eventType2EventHandler["acs.access_record.created_v1"] = acs.NewAccessRecordCreatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnUserUpdatedV1(handler func(ctx context.Context, event *acs.UserUpdatedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["acs.user.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "acs.user.updated_v1")
	}
	dispatcher.eventType2EventHandler["acs.user.updated_v1"] = acs.NewUserUpdatedEventHandler(handler)
	return dispatcher
}
