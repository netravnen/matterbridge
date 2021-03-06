// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Event undocumented
type Event struct {
	// OutlookItem is the base model of Event
	OutlookItem
	// OriginalStartTimeZone undocumented
	OriginalStartTimeZone *string `json:"originalStartTimeZone,omitempty"`
	// OriginalEndTimeZone undocumented
	OriginalEndTimeZone *string `json:"originalEndTimeZone,omitempty"`
	// ResponseStatus undocumented
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
	// UID undocumented
	UID *string `json:"uid,omitempty"`
	// ReminderMinutesBeforeStart undocumented
	ReminderMinutesBeforeStart *int `json:"reminderMinutesBeforeStart,omitempty"`
	// IsReminderOn undocumented
	IsReminderOn *bool `json:"isReminderOn,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// BodyPreview undocumented
	BodyPreview *string `json:"bodyPreview,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// Sensitivity undocumented
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`
	// Start undocumented
	Start *DateTimeTimeZone `json:"start,omitempty"`
	// OriginalStart undocumented
	OriginalStart *time.Time `json:"originalStart,omitempty"`
	// End undocumented
	End *DateTimeTimeZone `json:"end,omitempty"`
	// Location undocumented
	Location *Location `json:"location,omitempty"`
	// Locations undocumented
	Locations []Location `json:"locations,omitempty"`
	// IsAllDay undocumented
	IsAllDay *bool `json:"isAllDay,omitempty"`
	// IsCancelled undocumented
	IsCancelled *bool `json:"isCancelled,omitempty"`
	// IsOrganizer undocumented
	IsOrganizer *bool `json:"isOrganizer,omitempty"`
	// Recurrence undocumented
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// ResponseRequested undocumented
	ResponseRequested *bool `json:"responseRequested,omitempty"`
	// SeriesMasterID undocumented
	SeriesMasterID *string `json:"seriesMasterId,omitempty"`
	// ShowAs undocumented
	ShowAs *FreeBusyStatus `json:"showAs,omitempty"`
	// Type undocumented
	Type *EventType `json:"type,omitempty"`
	// Attendees undocumented
	Attendees []Attendee `json:"attendees,omitempty"`
	// Organizer undocumented
	Organizer *Recipient `json:"organizer,omitempty"`
	// WebLink undocumented
	WebLink *string `json:"webLink,omitempty"`
	// OnlineMeetingURL undocumented
	OnlineMeetingURL *string `json:"onlineMeetingUrl,omitempty"`
	// IsOnlineMeeting undocumented
	IsOnlineMeeting *bool `json:"isOnlineMeeting,omitempty"`
	// OnlineMeetingProvider undocumented
	OnlineMeetingProvider *OnlineMeetingProviderType `json:"onlineMeetingProvider,omitempty"`
	// OnlineMeeting undocumented
	OnlineMeeting *OnlineMeetingInfo `json:"onlineMeeting,omitempty"`
	// AllowNewTimeProposals undocumented
	AllowNewTimeProposals *bool `json:"allowNewTimeProposals,omitempty"`
	// Attachments undocumented
	Attachments []Attachment `json:"attachments,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// Calendar undocumented
	Calendar *Calendar `json:"calendar,omitempty"`
	// Instances undocumented
	Instances []Event `json:"instances,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
}
