// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/service/vc/v1"
)

func (dispatcher *EventReqDispatcher) OnMeetingJoinMeetingV1(handler func(ctx context.Context, event *vc.MeetingJoinMeetingEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.join_meeting_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.join_meeting_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.join_meeting_v1"] = vc.NewMeetingJoinMeetingEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMeetingLeaveMeetingV1(handler func(ctx context.Context, event *vc.MeetingLeaveMeetingEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.leave_meeting_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.leave_meeting_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.leave_meeting_v1"] = vc.NewMeetingLeaveMeetingEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMeetingMeetingEndedV1(handler func(ctx context.Context, event *vc.MeetingMeetingEndedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.meeting_ended_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.meeting_ended_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.meeting_ended_v1"] = vc.NewMeetingMeetingEndedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMeetingMeetingStartedV1(handler func(ctx context.Context, event *vc.MeetingMeetingStartedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.meeting_started_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.meeting_started_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.meeting_started_v1"] = vc.NewMeetingMeetingStartedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMeetingRecordingEndedV1(handler func(ctx context.Context, event *vc.MeetingRecordingEndedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.recording_ended_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.recording_ended_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.recording_ended_v1"] = vc.NewMeetingRecordingEndedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMeetingRecordingReadyV1(handler func(ctx context.Context, event *vc.MeetingRecordingReadyEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.recording_ready_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.recording_ready_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.recording_ready_v1"] = vc.NewMeetingRecordingReadyEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMeetingRecordingStartedV1(handler func(ctx context.Context, event *vc.MeetingRecordingStartedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.recording_started_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.recording_started_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.recording_started_v1"] = vc.NewMeetingRecordingStartedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMeetingShareEndedV1(handler func(ctx context.Context, event *vc.MeetingShareEndedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.share_ended_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.share_ended_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.share_ended_v1"] = vc.NewMeetingShareEndedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) OnMeetingShareStartedV1(handler func(ctx context.Context, event *vc.MeetingShareStartedEvent) error) *EventReqDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.share_started_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.share_started_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.share_started_v1"] = vc.NewMeetingShareStartedEventHandler(handler)
	return dispatcher
}
