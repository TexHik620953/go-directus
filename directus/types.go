package directus

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type IDirectusObject interface {
	DeepCopy() IDirectusObject
	Diff(old IDirectusObject) map[string]interface{}
	Track() []IDirectusObject
	GetId() string
	CollectionName() string
}

type CheckEvents struct {
	IDirectusObject
	Check       *PlayerChecks `json:"check"`
	DateCreated *time.Time    `json:"date_created"`
	Event       any           `json:"event"`
	Id          uuid.UUID     `json:"id"`
}

func (cf *CheckEvents) UnmarshalJSON(data []byte) error {
	type checkevents_internal struct {
		Check       *PlayerChecks `json:"check"`
		DateCreated *time.Time    `json:"date_created"`
		Event       any           `json:"event"`
		Id          uuid.UUID     `json:"id"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj checkevents_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Check = _obj.Check
		cf.DateCreated = _obj.DateCreated
		cf.Event = _obj.Event
		cf.Id = _obj.Id
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf CheckEvents) DeepCopy() IDirectusObject {
	new_obj := &CheckEvents{}
	if cf.Check != nil {
		new_obj.Check = (*cf.Check).DeepCopy().(*PlayerChecks)
	}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Event = cf.Event
	new_obj.Id = cf.Id
	return new_obj
}
func (cf CheckEvents) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*CheckEvents).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*CheckEvents).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*CheckEvents).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Event != old.(*CheckEvents).Event {
		diff["event"] = cf.Event
	}

	if cf.Id != old.(*CheckEvents).Id {
		diff["id"] = cf.Id
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf CheckEvents) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Check != nil {
		trakingList = append(trakingList, cf.Check)
		trakingList = append(trakingList, cf.Check.Track()...)
	}

	return trakingList
}

func (cf CheckEvents) GetId() string {
	return cf.Id.String()
}

func (cf CheckEvents) CollectionName() string {
	return "CheckEvents"
}

type CheckMessages struct {
	IDirectusObject
	Check       *PlayerChecks  `json:"check"`
	DateCreated *time.Time     `json:"date_created"`
	Fromplayer  *ServerPlayers `json:"fromplayer"`
	Fromuser    *DirectusUsers `json:"fromuser"`
	Id          uuid.UUID      `json:"id"`
	Message     string         `json:"message"`
}

func (cf *CheckMessages) UnmarshalJSON(data []byte) error {
	type checkmessages_internal struct {
		Check       *PlayerChecks  `json:"check"`
		DateCreated *time.Time     `json:"date_created"`
		Fromplayer  *ServerPlayers `json:"fromplayer"`
		Fromuser    *DirectusUsers `json:"fromuser"`
		Id          uuid.UUID      `json:"id"`
		Message     string         `json:"message"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj checkmessages_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Check = _obj.Check
		cf.DateCreated = _obj.DateCreated
		cf.Fromplayer = _obj.Fromplayer
		cf.Fromuser = _obj.Fromuser
		cf.Id = _obj.Id
		cf.Message = _obj.Message
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf CheckMessages) DeepCopy() IDirectusObject {
	new_obj := &CheckMessages{}
	if cf.Check != nil {
		new_obj.Check = (*cf.Check).DeepCopy().(*PlayerChecks)
	}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.Fromplayer != nil {
		new_obj.Fromplayer = (*cf.Fromplayer).DeepCopy().(*ServerPlayers)
	}
	if cf.Fromuser != nil {
		new_obj.Fromuser = (*cf.Fromuser).DeepCopy().(*DirectusUsers)
	}
	new_obj.Id = cf.Id
	new_obj.Message = cf.Message
	return new_obj
}
func (cf CheckMessages) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*CheckMessages).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*CheckMessages).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*CheckMessages).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*CheckMessages).Id {
		diff["id"] = cf.Id
	}

	if cf.Message != old.(*CheckMessages).Message {
		diff["message"] = cf.Message
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf CheckMessages) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Check != nil {
		trakingList = append(trakingList, cf.Check)
		trakingList = append(trakingList, cf.Check.Track()...)
	}

	if cf.Fromplayer != nil {
		trakingList = append(trakingList, cf.Fromplayer)
		trakingList = append(trakingList, cf.Fromplayer.Track()...)
	}
	if cf.Fromuser != nil {
		trakingList = append(trakingList, cf.Fromuser)
		trakingList = append(trakingList, cf.Fromuser.Track()...)
	}

	return trakingList
}

func (cf CheckMessages) GetId() string {
	return cf.Id.String()
}

func (cf CheckMessages) CollectionName() string {
	return "CheckMessages"
}

type DirectusActivity struct {
	IDirectusObject
	Action     string              `json:"action"`
	Collection string              `json:"collection"`
	Comment    *string             `json:"comment"`
	Id         int                 `json:"id"`
	Ip         *string             `json:"ip"`
	Item       string              `json:"item"`
	Origin     *string             `json:"origin"`
	Revisions  []DirectusRevisions `json:"revisions"`
	Timestamp  time.Time           `json:"timestamp"`
	User       *DirectusUsers      `json:"user"`
	UserAgent  *string             `json:"user_agent"`
}

func (cf *DirectusActivity) UnmarshalJSON(data []byte) error {
	type directusactivity_internal struct {
		Action     string              `json:"action"`
		Collection string              `json:"collection"`
		Comment    *string             `json:"comment"`
		Id         int                 `json:"id"`
		Ip         *string             `json:"ip"`
		Item       string              `json:"item"`
		Origin     *string             `json:"origin"`
		Revisions  []DirectusRevisions `json:"revisions"`
		Timestamp  time.Time           `json:"timestamp"`
		User       *DirectusUsers      `json:"user"`
		UserAgent  *string             `json:"user_agent"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusactivity_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Action = _obj.Action
		cf.Collection = _obj.Collection
		cf.Comment = _obj.Comment
		cf.Id = _obj.Id
		cf.Ip = _obj.Ip
		cf.Item = _obj.Item
		cf.Origin = _obj.Origin
		cf.Revisions = _obj.Revisions
		cf.Timestamp = _obj.Timestamp
		cf.User = _obj.User
		cf.UserAgent = _obj.UserAgent
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusActivity) DeepCopy() IDirectusObject {
	new_obj := &DirectusActivity{}
	new_obj.Action = cf.Action
	new_obj.Collection = cf.Collection
	if cf.Comment != nil {
		temp := ""
		new_obj.Comment = &temp
		*new_obj.Comment = *cf.Comment
	}
	new_obj.Id = cf.Id
	if cf.Ip != nil {
		temp := ""
		new_obj.Ip = &temp
		*new_obj.Ip = *cf.Ip
	}
	new_obj.Item = cf.Item
	if cf.Origin != nil {
		temp := ""
		new_obj.Origin = &temp
		*new_obj.Origin = *cf.Origin
	}
	if cf.Revisions != nil {
		new_obj.Revisions = make([]DirectusRevisions, len(cf.Revisions))
		copy(new_obj.Revisions, cf.Revisions)
	}
	new_obj.Timestamp = cf.Timestamp
	if cf.User != nil {
		new_obj.User = (*cf.User).DeepCopy().(*DirectusUsers)
	}
	if cf.UserAgent != nil {
		temp := ""
		new_obj.UserAgent = &temp
		*new_obj.UserAgent = *cf.UserAgent
	}
	return new_obj
}
func (cf DirectusActivity) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Action != old.(*DirectusActivity).Action {
		diff["action"] = cf.Action
	}

	if cf.Collection != old.(*DirectusActivity).Collection {
		diff["collection"] = cf.Collection
	}
	if cf.Comment == nil {
		if old.(*DirectusActivity).Comment != nil {
			diff["comment"] = nil
		}
	} else {
		if old.(*DirectusActivity).Comment == nil {
			diff["comment"] = cf.Comment
		} else {
			if *cf.Comment != *old.(*DirectusActivity).Comment {
				diff["comment"] = cf.Comment
			}
		}
	}

	if cf.Id != old.(*DirectusActivity).Id {
		diff["id"] = cf.Id
	}
	if cf.Ip == nil {
		if old.(*DirectusActivity).Ip != nil {
			diff["ip"] = nil
		}
	} else {
		if old.(*DirectusActivity).Ip == nil {
			diff["ip"] = cf.Ip
		} else {
			if *cf.Ip != *old.(*DirectusActivity).Ip {
				diff["ip"] = cf.Ip
			}
		}
	}

	if cf.Item != old.(*DirectusActivity).Item {
		diff["item"] = cf.Item
	}
	if cf.Origin == nil {
		if old.(*DirectusActivity).Origin != nil {
			diff["origin"] = nil
		}
	} else {
		if old.(*DirectusActivity).Origin == nil {
			diff["origin"] = cf.Origin
		} else {
			if *cf.Origin != *old.(*DirectusActivity).Origin {
				diff["origin"] = cf.Origin
			}
		}
	}

	if cf.Timestamp != old.(*DirectusActivity).Timestamp {
		diff["timestamp"] = cf.Timestamp
	}

	if cf.UserAgent == nil {
		if old.(*DirectusActivity).UserAgent != nil {
			diff["user_agent"] = nil
		}
	} else {
		if old.(*DirectusActivity).UserAgent == nil {
			diff["user_agent"] = cf.UserAgent
		} else {
			if *cf.UserAgent != *old.(*DirectusActivity).UserAgent {
				diff["user_agent"] = cf.UserAgent
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusActivity) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Revisions != nil {
		for _, iter := range cf.Revisions {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.User != nil {
		trakingList = append(trakingList, cf.User)
		trakingList = append(trakingList, cf.User.Track()...)
	}

	return trakingList
}

func (cf DirectusActivity) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusActivity) CollectionName() string {
	return "directus_activity"
}

type DirectusDashboards struct {
	IDirectusObject
	Color       *string          `json:"color"`
	DateCreated *time.Time       `json:"date_created"`
	Icon        string           `json:"icon"`
	Id          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Note        *string          `json:"note"`
	Panels      []DirectusPanels `json:"panels"`
	UserCreated *DirectusUsers   `json:"user_created"`
}

func (cf *DirectusDashboards) UnmarshalJSON(data []byte) error {
	type directusdashboards_internal struct {
		Color       *string          `json:"color"`
		DateCreated *time.Time       `json:"date_created"`
		Icon        string           `json:"icon"`
		Id          uuid.UUID        `json:"id"`
		Name        string           `json:"name"`
		Note        *string          `json:"note"`
		Panels      []DirectusPanels `json:"panels"`
		UserCreated *DirectusUsers   `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusdashboards_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Color = _obj.Color
		cf.DateCreated = _obj.DateCreated
		cf.Icon = _obj.Icon
		cf.Id = _obj.Id
		cf.Name = _obj.Name
		cf.Note = _obj.Note
		cf.Panels = _obj.Panels
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusDashboards) DeepCopy() IDirectusObject {
	new_obj := &DirectusDashboards{}
	if cf.Color != nil {
		temp := ""
		new_obj.Color = &temp
		*new_obj.Color = *cf.Color
	}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Icon = cf.Icon
	new_obj.Id = cf.Id
	new_obj.Name = cf.Name
	if cf.Note != nil {
		temp := ""
		new_obj.Note = &temp
		*new_obj.Note = *cf.Note
	}
	if cf.Panels != nil {
		new_obj.Panels = make([]DirectusPanels, len(cf.Panels))
		copy(new_obj.Panels, cf.Panels)
	}
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf DirectusDashboards) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Color == nil {
		if old.(*DirectusDashboards).Color != nil {
			diff["color"] = nil
		}
	} else {
		if old.(*DirectusDashboards).Color == nil {
			diff["color"] = cf.Color
		} else {
			if *cf.Color != *old.(*DirectusDashboards).Color {
				diff["color"] = cf.Color
			}
		}
	}
	if cf.DateCreated == nil {
		if old.(*DirectusDashboards).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*DirectusDashboards).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*DirectusDashboards).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Icon != old.(*DirectusDashboards).Icon {
		diff["icon"] = cf.Icon
	}

	if cf.Id != old.(*DirectusDashboards).Id {
		diff["id"] = cf.Id
	}

	if cf.Name != old.(*DirectusDashboards).Name {
		diff["name"] = cf.Name
	}
	if cf.Note == nil {
		if old.(*DirectusDashboards).Note != nil {
			diff["note"] = nil
		}
	} else {
		if old.(*DirectusDashboards).Note == nil {
			diff["note"] = cf.Note
		} else {
			if *cf.Note != *old.(*DirectusDashboards).Note {
				diff["note"] = cf.Note
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusDashboards) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Panels != nil {
		for _, iter := range cf.Panels {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf DirectusDashboards) GetId() string {
	return cf.Id.String()
}

func (cf DirectusDashboards) CollectionName() string {
	return "directus_dashboards"
}

type DirectusFields struct {
	IDirectusObject
	Conditions        any             `json:"conditions"`
	Display           *string         `json:"display"`
	DisplayOptions    any             `json:"display_options"`
	Field             string          `json:"field"`
	Group             *DirectusFields `json:"group"`
	Hidden            bool            `json:"hidden"`
	Id                int             `json:"id"`
	Interface         *string         `json:"interface"`
	Note              *string         `json:"note"`
	Options           any             `json:"options"`
	Readonly          bool            `json:"readonly"`
	Required          *bool           `json:"required"`
	Sort              *int            `json:"sort"`
	Special           any             `json:"special"`
	Translations      any             `json:"translations"`
	Validation        any             `json:"validation"`
	ValidationMessage *string         `json:"validation_message"`
	Width             *string         `json:"width"`
}

func (cf *DirectusFields) UnmarshalJSON(data []byte) error {
	type directusfields_internal struct {
		Conditions        any             `json:"conditions"`
		Display           *string         `json:"display"`
		DisplayOptions    any             `json:"display_options"`
		Field             string          `json:"field"`
		Group             *DirectusFields `json:"group"`
		Hidden            bool            `json:"hidden"`
		Id                int             `json:"id"`
		Interface         *string         `json:"interface"`
		Note              *string         `json:"note"`
		Options           any             `json:"options"`
		Readonly          bool            `json:"readonly"`
		Required          *bool           `json:"required"`
		Sort              *int            `json:"sort"`
		Special           any             `json:"special"`
		Translations      any             `json:"translations"`
		Validation        any             `json:"validation"`
		ValidationMessage *string         `json:"validation_message"`
		Width             *string         `json:"width"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusfields_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Conditions = _obj.Conditions
		cf.Display = _obj.Display
		cf.DisplayOptions = _obj.DisplayOptions
		cf.Field = _obj.Field
		cf.Group = _obj.Group
		cf.Hidden = _obj.Hidden
		cf.Id = _obj.Id
		cf.Interface = _obj.Interface
		cf.Note = _obj.Note
		cf.Options = _obj.Options
		cf.Readonly = _obj.Readonly
		cf.Required = _obj.Required
		cf.Sort = _obj.Sort
		cf.Special = _obj.Special
		cf.Translations = _obj.Translations
		cf.Validation = _obj.Validation
		cf.ValidationMessage = _obj.ValidationMessage
		cf.Width = _obj.Width
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusFields) DeepCopy() IDirectusObject {
	new_obj := &DirectusFields{}
	new_obj.Conditions = cf.Conditions
	if cf.Display != nil {
		temp := ""
		new_obj.Display = &temp
		*new_obj.Display = *cf.Display
	}
	new_obj.DisplayOptions = cf.DisplayOptions
	new_obj.Field = cf.Field
	if cf.Group != nil {
		new_obj.Group = (*cf.Group).DeepCopy().(*DirectusFields)
	}
	new_obj.Hidden = cf.Hidden
	new_obj.Id = cf.Id
	if cf.Interface != nil {
		temp := ""
		new_obj.Interface = &temp
		*new_obj.Interface = *cf.Interface
	}
	if cf.Note != nil {
		temp := ""
		new_obj.Note = &temp
		*new_obj.Note = *cf.Note
	}
	new_obj.Options = cf.Options
	new_obj.Readonly = cf.Readonly
	if cf.Required != nil {
		temp := false
		new_obj.Required = &temp
		*new_obj.Required = *cf.Required
	}
	if cf.Sort != nil {
		temp := 0
		new_obj.Sort = &temp
		*new_obj.Sort = *cf.Sort
	}
	new_obj.Special = cf.Special
	new_obj.Translations = cf.Translations
	new_obj.Validation = cf.Validation
	if cf.ValidationMessage != nil {
		temp := ""
		new_obj.ValidationMessage = &temp
		*new_obj.ValidationMessage = *cf.ValidationMessage
	}
	if cf.Width != nil {
		temp := ""
		new_obj.Width = &temp
		*new_obj.Width = *cf.Width
	}
	return new_obj
}
func (cf DirectusFields) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Conditions != old.(*DirectusFields).Conditions {
		diff["conditions"] = cf.Conditions
	}
	if cf.Display == nil {
		if old.(*DirectusFields).Display != nil {
			diff["display"] = nil
		}
	} else {
		if old.(*DirectusFields).Display == nil {
			diff["display"] = cf.Display
		} else {
			if *cf.Display != *old.(*DirectusFields).Display {
				diff["display"] = cf.Display
			}
		}
	}

	if cf.DisplayOptions != old.(*DirectusFields).DisplayOptions {
		diff["display_options"] = cf.DisplayOptions
	}

	if cf.Field != old.(*DirectusFields).Field {
		diff["field"] = cf.Field
	}

	if cf.Hidden != old.(*DirectusFields).Hidden {
		diff["hidden"] = cf.Hidden
	}

	if cf.Id != old.(*DirectusFields).Id {
		diff["id"] = cf.Id
	}
	if cf.Interface == nil {
		if old.(*DirectusFields).Interface != nil {
			diff["interface"] = nil
		}
	} else {
		if old.(*DirectusFields).Interface == nil {
			diff["interface"] = cf.Interface
		} else {
			if *cf.Interface != *old.(*DirectusFields).Interface {
				diff["interface"] = cf.Interface
			}
		}
	}
	if cf.Note == nil {
		if old.(*DirectusFields).Note != nil {
			diff["note"] = nil
		}
	} else {
		if old.(*DirectusFields).Note == nil {
			diff["note"] = cf.Note
		} else {
			if *cf.Note != *old.(*DirectusFields).Note {
				diff["note"] = cf.Note
			}
		}
	}

	if cf.Options != old.(*DirectusFields).Options {
		diff["options"] = cf.Options
	}

	if cf.Readonly != old.(*DirectusFields).Readonly {
		diff["readonly"] = cf.Readonly
	}
	if cf.Required == nil {
		if old.(*DirectusFields).Required != nil {
			diff["required"] = nil
		}
	} else {
		if old.(*DirectusFields).Required == nil {
			diff["required"] = cf.Required
		} else {
			if *cf.Required != *old.(*DirectusFields).Required {
				diff["required"] = cf.Required
			}
		}
	}
	if cf.Sort == nil {
		if old.(*DirectusFields).Sort != nil {
			diff["sort"] = nil
		}
	} else {
		if old.(*DirectusFields).Sort == nil {
			diff["sort"] = cf.Sort
		} else {
			if *cf.Sort != *old.(*DirectusFields).Sort {
				diff["sort"] = cf.Sort
			}
		}
	}

	if cf.Special != old.(*DirectusFields).Special {
		diff["special"] = cf.Special
	}

	if cf.Translations != old.(*DirectusFields).Translations {
		diff["translations"] = cf.Translations
	}

	if cf.Validation != old.(*DirectusFields).Validation {
		diff["validation"] = cf.Validation
	}
	if cf.ValidationMessage == nil {
		if old.(*DirectusFields).ValidationMessage != nil {
			diff["validation_message"] = nil
		}
	} else {
		if old.(*DirectusFields).ValidationMessage == nil {
			diff["validation_message"] = cf.ValidationMessage
		} else {
			if *cf.ValidationMessage != *old.(*DirectusFields).ValidationMessage {
				diff["validation_message"] = cf.ValidationMessage
			}
		}
	}
	if cf.Width == nil {
		if old.(*DirectusFields).Width != nil {
			diff["width"] = nil
		}
	} else {
		if old.(*DirectusFields).Width == nil {
			diff["width"] = cf.Width
		} else {
			if *cf.Width != *old.(*DirectusFields).Width {
				diff["width"] = cf.Width
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusFields) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Group != nil {
		trakingList = append(trakingList, cf.Group)
		trakingList = append(trakingList, cf.Group.Track()...)
	}

	return trakingList
}

func (cf DirectusFields) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusFields) CollectionName() string {
	return "directus_fields"
}

type DirectusFiles struct {
	IDirectusObject
	Charset          *string          `json:"charset"`
	Description      *string          `json:"description"`
	Duration         *int             `json:"duration"`
	Embed            *string          `json:"embed"`
	FilenameDisk     *string          `json:"filename_disk"`
	FilenameDownload string           `json:"filename_download"`
	Filesize         *string          `json:"filesize"`
	Folder           *DirectusFolders `json:"folder"`
	Height           *int             `json:"height"`
	Id               uuid.UUID        `json:"id"`
	Location         *string          `json:"location"`
	Metadata         any              `json:"metadata"`
	ModifiedBy       *DirectusUsers   `json:"modified_by"`
	ModifiedOn       time.Time        `json:"modified_on"`
	Storage          string           `json:"storage"`
	StorageDivider   any              `json:"storage_divider"`
	Tags             any              `json:"tags"`
	Title            *string          `json:"title"`
	Type             *string          `json:"type"`
	UploadedBy       *DirectusUsers   `json:"uploaded_by"`
	UploadedOn       time.Time        `json:"uploaded_on"`
	Width            *int             `json:"width"`
}

func (cf *DirectusFiles) UnmarshalJSON(data []byte) error {
	type directusfiles_internal struct {
		Charset          *string          `json:"charset"`
		Description      *string          `json:"description"`
		Duration         *int             `json:"duration"`
		Embed            *string          `json:"embed"`
		FilenameDisk     *string          `json:"filename_disk"`
		FilenameDownload string           `json:"filename_download"`
		Filesize         *string          `json:"filesize"`
		Folder           *DirectusFolders `json:"folder"`
		Height           *int             `json:"height"`
		Id               uuid.UUID        `json:"id"`
		Location         *string          `json:"location"`
		Metadata         any              `json:"metadata"`
		ModifiedBy       *DirectusUsers   `json:"modified_by"`
		ModifiedOn       time.Time        `json:"modified_on"`
		Storage          string           `json:"storage"`
		StorageDivider   any              `json:"storage_divider"`
		Tags             any              `json:"tags"`
		Title            *string          `json:"title"`
		Type             *string          `json:"type"`
		UploadedBy       *DirectusUsers   `json:"uploaded_by"`
		UploadedOn       time.Time        `json:"uploaded_on"`
		Width            *int             `json:"width"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusfiles_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Charset = _obj.Charset
		cf.Description = _obj.Description
		cf.Duration = _obj.Duration
		cf.Embed = _obj.Embed
		cf.FilenameDisk = _obj.FilenameDisk
		cf.FilenameDownload = _obj.FilenameDownload
		cf.Filesize = _obj.Filesize
		cf.Folder = _obj.Folder
		cf.Height = _obj.Height
		cf.Id = _obj.Id
		cf.Location = _obj.Location
		cf.Metadata = _obj.Metadata
		cf.ModifiedBy = _obj.ModifiedBy
		cf.ModifiedOn = _obj.ModifiedOn
		cf.Storage = _obj.Storage
		cf.StorageDivider = _obj.StorageDivider
		cf.Tags = _obj.Tags
		cf.Title = _obj.Title
		cf.Type = _obj.Type
		cf.UploadedBy = _obj.UploadedBy
		cf.UploadedOn = _obj.UploadedOn
		cf.Width = _obj.Width
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusFiles) DeepCopy() IDirectusObject {
	new_obj := &DirectusFiles{}
	if cf.Charset != nil {
		temp := ""
		new_obj.Charset = &temp
		*new_obj.Charset = *cf.Charset
	}
	if cf.Description != nil {
		temp := ""
		new_obj.Description = &temp
		*new_obj.Description = *cf.Description
	}
	if cf.Duration != nil {
		temp := 0
		new_obj.Duration = &temp
		*new_obj.Duration = *cf.Duration
	}
	if cf.Embed != nil {
		temp := ""
		new_obj.Embed = &temp
		*new_obj.Embed = *cf.Embed
	}
	if cf.FilenameDisk != nil {
		temp := ""
		new_obj.FilenameDisk = &temp
		*new_obj.FilenameDisk = *cf.FilenameDisk
	}
	new_obj.FilenameDownload = cf.FilenameDownload
	if cf.Filesize != nil {
		temp := ""
		new_obj.Filesize = &temp
		*new_obj.Filesize = *cf.Filesize
	}
	if cf.Folder != nil {
		new_obj.Folder = (*cf.Folder).DeepCopy().(*DirectusFolders)
	}
	if cf.Height != nil {
		temp := 0
		new_obj.Height = &temp
		*new_obj.Height = *cf.Height
	}
	new_obj.Id = cf.Id
	if cf.Location != nil {
		temp := ""
		new_obj.Location = &temp
		*new_obj.Location = *cf.Location
	}
	new_obj.Metadata = cf.Metadata
	if cf.ModifiedBy != nil {
		new_obj.ModifiedBy = (*cf.ModifiedBy).DeepCopy().(*DirectusUsers)
	}
	new_obj.ModifiedOn = cf.ModifiedOn
	new_obj.Storage = cf.Storage
	new_obj.StorageDivider = cf.StorageDivider
	new_obj.Tags = cf.Tags
	if cf.Title != nil {
		temp := ""
		new_obj.Title = &temp
		*new_obj.Title = *cf.Title
	}
	if cf.Type != nil {
		temp := ""
		new_obj.Type = &temp
		*new_obj.Type = *cf.Type
	}
	if cf.UploadedBy != nil {
		new_obj.UploadedBy = (*cf.UploadedBy).DeepCopy().(*DirectusUsers)
	}
	new_obj.UploadedOn = cf.UploadedOn
	if cf.Width != nil {
		temp := 0
		new_obj.Width = &temp
		*new_obj.Width = *cf.Width
	}
	return new_obj
}
func (cf DirectusFiles) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Charset == nil {
		if old.(*DirectusFiles).Charset != nil {
			diff["charset"] = nil
		}
	} else {
		if old.(*DirectusFiles).Charset == nil {
			diff["charset"] = cf.Charset
		} else {
			if *cf.Charset != *old.(*DirectusFiles).Charset {
				diff["charset"] = cf.Charset
			}
		}
	}
	if cf.Description == nil {
		if old.(*DirectusFiles).Description != nil {
			diff["description"] = nil
		}
	} else {
		if old.(*DirectusFiles).Description == nil {
			diff["description"] = cf.Description
		} else {
			if *cf.Description != *old.(*DirectusFiles).Description {
				diff["description"] = cf.Description
			}
		}
	}
	if cf.Duration == nil {
		if old.(*DirectusFiles).Duration != nil {
			diff["duration"] = nil
		}
	} else {
		if old.(*DirectusFiles).Duration == nil {
			diff["duration"] = cf.Duration
		} else {
			if *cf.Duration != *old.(*DirectusFiles).Duration {
				diff["duration"] = cf.Duration
			}
		}
	}
	if cf.Embed == nil {
		if old.(*DirectusFiles).Embed != nil {
			diff["embed"] = nil
		}
	} else {
		if old.(*DirectusFiles).Embed == nil {
			diff["embed"] = cf.Embed
		} else {
			if *cf.Embed != *old.(*DirectusFiles).Embed {
				diff["embed"] = cf.Embed
			}
		}
	}
	if cf.FilenameDisk == nil {
		if old.(*DirectusFiles).FilenameDisk != nil {
			diff["filename_disk"] = nil
		}
	} else {
		if old.(*DirectusFiles).FilenameDisk == nil {
			diff["filename_disk"] = cf.FilenameDisk
		} else {
			if *cf.FilenameDisk != *old.(*DirectusFiles).FilenameDisk {
				diff["filename_disk"] = cf.FilenameDisk
			}
		}
	}

	if cf.FilenameDownload != old.(*DirectusFiles).FilenameDownload {
		diff["filename_download"] = cf.FilenameDownload
	}
	if cf.Filesize == nil {
		if old.(*DirectusFiles).Filesize != nil {
			diff["filesize"] = nil
		}
	} else {
		if old.(*DirectusFiles).Filesize == nil {
			diff["filesize"] = cf.Filesize
		} else {
			if *cf.Filesize != *old.(*DirectusFiles).Filesize {
				diff["filesize"] = cf.Filesize
			}
		}
	}

	if cf.Height == nil {
		if old.(*DirectusFiles).Height != nil {
			diff["height"] = nil
		}
	} else {
		if old.(*DirectusFiles).Height == nil {
			diff["height"] = cf.Height
		} else {
			if *cf.Height != *old.(*DirectusFiles).Height {
				diff["height"] = cf.Height
			}
		}
	}

	if cf.Id != old.(*DirectusFiles).Id {
		diff["id"] = cf.Id
	}
	if cf.Location == nil {
		if old.(*DirectusFiles).Location != nil {
			diff["location"] = nil
		}
	} else {
		if old.(*DirectusFiles).Location == nil {
			diff["location"] = cf.Location
		} else {
			if *cf.Location != *old.(*DirectusFiles).Location {
				diff["location"] = cf.Location
			}
		}
	}

	if cf.Metadata != old.(*DirectusFiles).Metadata {
		diff["metadata"] = cf.Metadata
	}

	if cf.ModifiedOn != old.(*DirectusFiles).ModifiedOn {
		diff["modified_on"] = cf.ModifiedOn
	}

	if cf.Storage != old.(*DirectusFiles).Storage {
		diff["storage"] = cf.Storage
	}

	if cf.StorageDivider != old.(*DirectusFiles).StorageDivider {
		diff["storage_divider"] = cf.StorageDivider
	}

	if cf.Tags != old.(*DirectusFiles).Tags {
		diff["tags"] = cf.Tags
	}
	if cf.Title == nil {
		if old.(*DirectusFiles).Title != nil {
			diff["title"] = nil
		}
	} else {
		if old.(*DirectusFiles).Title == nil {
			diff["title"] = cf.Title
		} else {
			if *cf.Title != *old.(*DirectusFiles).Title {
				diff["title"] = cf.Title
			}
		}
	}
	if cf.Type == nil {
		if old.(*DirectusFiles).Type != nil {
			diff["type"] = nil
		}
	} else {
		if old.(*DirectusFiles).Type == nil {
			diff["type"] = cf.Type
		} else {
			if *cf.Type != *old.(*DirectusFiles).Type {
				diff["type"] = cf.Type
			}
		}
	}

	if cf.UploadedOn != old.(*DirectusFiles).UploadedOn {
		diff["uploaded_on"] = cf.UploadedOn
	}
	if cf.Width == nil {
		if old.(*DirectusFiles).Width != nil {
			diff["width"] = nil
		}
	} else {
		if old.(*DirectusFiles).Width == nil {
			diff["width"] = cf.Width
		} else {
			if *cf.Width != *old.(*DirectusFiles).Width {
				diff["width"] = cf.Width
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusFiles) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Folder != nil {
		trakingList = append(trakingList, cf.Folder)
		trakingList = append(trakingList, cf.Folder.Track()...)
	}

	if cf.ModifiedBy != nil {
		trakingList = append(trakingList, cf.ModifiedBy)
		trakingList = append(trakingList, cf.ModifiedBy.Track()...)
	}

	if cf.UploadedBy != nil {
		trakingList = append(trakingList, cf.UploadedBy)
		trakingList = append(trakingList, cf.UploadedBy.Track()...)
	}

	return trakingList
}

func (cf DirectusFiles) GetId() string {
	return cf.Id.String()
}

func (cf DirectusFiles) CollectionName() string {
	return "directus_files"
}

type DirectusFlows struct {
	IDirectusObject
	Accountability *string              `json:"accountability"`
	Color          *string              `json:"color"`
	DateCreated    *time.Time           `json:"date_created"`
	Description    *string              `json:"description"`
	Icon           *string              `json:"icon"`
	Id             uuid.UUID            `json:"id"`
	Name           string               `json:"name"`
	Operation      *DirectusOperations  `json:"operation"`
	Operations     []DirectusOperations `json:"operations"`
	Options        any                  `json:"options"`
	Status         string               `json:"status"`
	Trigger        *string              `json:"trigger"`
	UserCreated    *DirectusUsers       `json:"user_created"`
}

func (cf *DirectusFlows) UnmarshalJSON(data []byte) error {
	type directusflows_internal struct {
		Accountability *string              `json:"accountability"`
		Color          *string              `json:"color"`
		DateCreated    *time.Time           `json:"date_created"`
		Description    *string              `json:"description"`
		Icon           *string              `json:"icon"`
		Id             uuid.UUID            `json:"id"`
		Name           string               `json:"name"`
		Operation      *DirectusOperations  `json:"operation"`
		Operations     []DirectusOperations `json:"operations"`
		Options        any                  `json:"options"`
		Status         string               `json:"status"`
		Trigger        *string              `json:"trigger"`
		UserCreated    *DirectusUsers       `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusflows_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Accountability = _obj.Accountability
		cf.Color = _obj.Color
		cf.DateCreated = _obj.DateCreated
		cf.Description = _obj.Description
		cf.Icon = _obj.Icon
		cf.Id = _obj.Id
		cf.Name = _obj.Name
		cf.Operation = _obj.Operation
		cf.Operations = _obj.Operations
		cf.Options = _obj.Options
		cf.Status = _obj.Status
		cf.Trigger = _obj.Trigger
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusFlows) DeepCopy() IDirectusObject {
	new_obj := &DirectusFlows{}
	if cf.Accountability != nil {
		temp := ""
		new_obj.Accountability = &temp
		*new_obj.Accountability = *cf.Accountability
	}
	if cf.Color != nil {
		temp := ""
		new_obj.Color = &temp
		*new_obj.Color = *cf.Color
	}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.Description != nil {
		temp := ""
		new_obj.Description = &temp
		*new_obj.Description = *cf.Description
	}
	if cf.Icon != nil {
		temp := ""
		new_obj.Icon = &temp
		*new_obj.Icon = *cf.Icon
	}
	new_obj.Id = cf.Id
	new_obj.Name = cf.Name
	if cf.Operation != nil {
		new_obj.Operation = (*cf.Operation).DeepCopy().(*DirectusOperations)
	}
	if cf.Operations != nil {
		new_obj.Operations = make([]DirectusOperations, len(cf.Operations))
		copy(new_obj.Operations, cf.Operations)
	}
	new_obj.Options = cf.Options
	new_obj.Status = cf.Status
	if cf.Trigger != nil {
		temp := ""
		new_obj.Trigger = &temp
		*new_obj.Trigger = *cf.Trigger
	}
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf DirectusFlows) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Accountability == nil {
		if old.(*DirectusFlows).Accountability != nil {
			diff["accountability"] = nil
		}
	} else {
		if old.(*DirectusFlows).Accountability == nil {
			diff["accountability"] = cf.Accountability
		} else {
			if *cf.Accountability != *old.(*DirectusFlows).Accountability {
				diff["accountability"] = cf.Accountability
			}
		}
	}
	if cf.Color == nil {
		if old.(*DirectusFlows).Color != nil {
			diff["color"] = nil
		}
	} else {
		if old.(*DirectusFlows).Color == nil {
			diff["color"] = cf.Color
		} else {
			if *cf.Color != *old.(*DirectusFlows).Color {
				diff["color"] = cf.Color
			}
		}
	}
	if cf.DateCreated == nil {
		if old.(*DirectusFlows).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*DirectusFlows).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*DirectusFlows).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}
	if cf.Description == nil {
		if old.(*DirectusFlows).Description != nil {
			diff["description"] = nil
		}
	} else {
		if old.(*DirectusFlows).Description == nil {
			diff["description"] = cf.Description
		} else {
			if *cf.Description != *old.(*DirectusFlows).Description {
				diff["description"] = cf.Description
			}
		}
	}
	if cf.Icon == nil {
		if old.(*DirectusFlows).Icon != nil {
			diff["icon"] = nil
		}
	} else {
		if old.(*DirectusFlows).Icon == nil {
			diff["icon"] = cf.Icon
		} else {
			if *cf.Icon != *old.(*DirectusFlows).Icon {
				diff["icon"] = cf.Icon
			}
		}
	}

	if cf.Id != old.(*DirectusFlows).Id {
		diff["id"] = cf.Id
	}

	if cf.Name != old.(*DirectusFlows).Name {
		diff["name"] = cf.Name
	}

	if cf.Options != old.(*DirectusFlows).Options {
		diff["options"] = cf.Options
	}

	if cf.Status != old.(*DirectusFlows).Status {
		diff["status"] = cf.Status
	}
	if cf.Trigger == nil {
		if old.(*DirectusFlows).Trigger != nil {
			diff["trigger"] = nil
		}
	} else {
		if old.(*DirectusFlows).Trigger == nil {
			diff["trigger"] = cf.Trigger
		} else {
			if *cf.Trigger != *old.(*DirectusFlows).Trigger {
				diff["trigger"] = cf.Trigger
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusFlows) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Operation != nil {
		trakingList = append(trakingList, cf.Operation)
		trakingList = append(trakingList, cf.Operation.Track()...)
	}
	if cf.Operations != nil {
		for _, iter := range cf.Operations {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf DirectusFlows) GetId() string {
	return cf.Id.String()
}

func (cf DirectusFlows) CollectionName() string {
	return "directus_flows"
}

type DirectusFolders struct {
	IDirectusObject
	Id     uuid.UUID        `json:"id"`
	Name   string           `json:"name"`
	Parent *DirectusFolders `json:"parent"`
}

func (cf *DirectusFolders) UnmarshalJSON(data []byte) error {
	type directusfolders_internal struct {
		Id     uuid.UUID        `json:"id"`
		Name   string           `json:"name"`
		Parent *DirectusFolders `json:"parent"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusfolders_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Id = _obj.Id
		cf.Name = _obj.Name
		cf.Parent = _obj.Parent
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusFolders) DeepCopy() IDirectusObject {
	new_obj := &DirectusFolders{}
	new_obj.Id = cf.Id
	new_obj.Name = cf.Name
	if cf.Parent != nil {
		new_obj.Parent = (*cf.Parent).DeepCopy().(*DirectusFolders)
	}
	return new_obj
}
func (cf DirectusFolders) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Id != old.(*DirectusFolders).Id {
		diff["id"] = cf.Id
	}

	if cf.Name != old.(*DirectusFolders).Name {
		diff["name"] = cf.Name
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusFolders) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Parent != nil {
		trakingList = append(trakingList, cf.Parent)
		trakingList = append(trakingList, cf.Parent.Track()...)
	}
	return trakingList
}

func (cf DirectusFolders) GetId() string {
	return cf.Id.String()
}

func (cf DirectusFolders) CollectionName() string {
	return "directus_folders"
}

type DirectusNotifications struct {
	IDirectusObject
	Collection *string        `json:"collection"`
	Id         int            `json:"id"`
	Item       *string        `json:"item"`
	Message    *string        `json:"message"`
	Recipient  *DirectusUsers `json:"recipient"`
	Sender     *DirectusUsers `json:"sender"`
	Status     *string        `json:"status"`
	Subject    string         `json:"subject"`
	Timestamp  *time.Time     `json:"timestamp"`
}

func (cf *DirectusNotifications) UnmarshalJSON(data []byte) error {
	type directusnotifications_internal struct {
		Collection *string        `json:"collection"`
		Id         int            `json:"id"`
		Item       *string        `json:"item"`
		Message    *string        `json:"message"`
		Recipient  *DirectusUsers `json:"recipient"`
		Sender     *DirectusUsers `json:"sender"`
		Status     *string        `json:"status"`
		Subject    string         `json:"subject"`
		Timestamp  *time.Time     `json:"timestamp"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusnotifications_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Collection = _obj.Collection
		cf.Id = _obj.Id
		cf.Item = _obj.Item
		cf.Message = _obj.Message
		cf.Recipient = _obj.Recipient
		cf.Sender = _obj.Sender
		cf.Status = _obj.Status
		cf.Subject = _obj.Subject
		cf.Timestamp = _obj.Timestamp
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusNotifications) DeepCopy() IDirectusObject {
	new_obj := &DirectusNotifications{}
	if cf.Collection != nil {
		temp := ""
		new_obj.Collection = &temp
		*new_obj.Collection = *cf.Collection
	}
	new_obj.Id = cf.Id
	if cf.Item != nil {
		temp := ""
		new_obj.Item = &temp
		*new_obj.Item = *cf.Item
	}
	if cf.Message != nil {
		temp := ""
		new_obj.Message = &temp
		*new_obj.Message = *cf.Message
	}
	if cf.Recipient != nil {
		new_obj.Recipient = (*cf.Recipient).DeepCopy().(*DirectusUsers)
	}
	if cf.Sender != nil {
		new_obj.Sender = (*cf.Sender).DeepCopy().(*DirectusUsers)
	}
	if cf.Status != nil {
		temp := ""
		new_obj.Status = &temp
		*new_obj.Status = *cf.Status
	}
	new_obj.Subject = cf.Subject
	if cf.Timestamp != nil {
		temp := time.Time{}
		new_obj.Timestamp = &temp
		*new_obj.Timestamp = *cf.Timestamp
	}
	return new_obj
}
func (cf DirectusNotifications) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Collection == nil {
		if old.(*DirectusNotifications).Collection != nil {
			diff["collection"] = nil
		}
	} else {
		if old.(*DirectusNotifications).Collection == nil {
			diff["collection"] = cf.Collection
		} else {
			if *cf.Collection != *old.(*DirectusNotifications).Collection {
				diff["collection"] = cf.Collection
			}
		}
	}

	if cf.Id != old.(*DirectusNotifications).Id {
		diff["id"] = cf.Id
	}
	if cf.Item == nil {
		if old.(*DirectusNotifications).Item != nil {
			diff["item"] = nil
		}
	} else {
		if old.(*DirectusNotifications).Item == nil {
			diff["item"] = cf.Item
		} else {
			if *cf.Item != *old.(*DirectusNotifications).Item {
				diff["item"] = cf.Item
			}
		}
	}
	if cf.Message == nil {
		if old.(*DirectusNotifications).Message != nil {
			diff["message"] = nil
		}
	} else {
		if old.(*DirectusNotifications).Message == nil {
			diff["message"] = cf.Message
		} else {
			if *cf.Message != *old.(*DirectusNotifications).Message {
				diff["message"] = cf.Message
			}
		}
	}

	if cf.Status == nil {
		if old.(*DirectusNotifications).Status != nil {
			diff["status"] = nil
		}
	} else {
		if old.(*DirectusNotifications).Status == nil {
			diff["status"] = cf.Status
		} else {
			if *cf.Status != *old.(*DirectusNotifications).Status {
				diff["status"] = cf.Status
			}
		}
	}

	if cf.Subject != old.(*DirectusNotifications).Subject {
		diff["subject"] = cf.Subject
	}
	if cf.Timestamp == nil {
		if old.(*DirectusNotifications).Timestamp != nil {
			diff["timestamp"] = nil
		}
	} else {
		if old.(*DirectusNotifications).Timestamp == nil {
			diff["timestamp"] = cf.Timestamp
		} else {
			if *cf.Timestamp != *old.(*DirectusNotifications).Timestamp {
				diff["timestamp"] = cf.Timestamp
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusNotifications) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Recipient != nil {
		trakingList = append(trakingList, cf.Recipient)
		trakingList = append(trakingList, cf.Recipient.Track()...)
	}
	if cf.Sender != nil {
		trakingList = append(trakingList, cf.Sender)
		trakingList = append(trakingList, cf.Sender.Track()...)
	}

	return trakingList
}

func (cf DirectusNotifications) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusNotifications) CollectionName() string {
	return "directus_notifications"
}

type DirectusOperations struct {
	IDirectusObject
	DateCreated *time.Time          `json:"date_created"`
	Flow        *DirectusFlows      `json:"flow"`
	Id          uuid.UUID           `json:"id"`
	Key         string              `json:"key"`
	Name        *string             `json:"name"`
	Options     any                 `json:"options"`
	PositionX   int                 `json:"position_x"`
	PositionY   int                 `json:"position_y"`
	Reject      *DirectusOperations `json:"reject"`
	Resolve     *DirectusOperations `json:"resolve"`
	Type        string              `json:"type"`
	UserCreated *DirectusUsers      `json:"user_created"`
}

func (cf *DirectusOperations) UnmarshalJSON(data []byte) error {
	type directusoperations_internal struct {
		DateCreated *time.Time          `json:"date_created"`
		Flow        *DirectusFlows      `json:"flow"`
		Id          uuid.UUID           `json:"id"`
		Key         string              `json:"key"`
		Name        *string             `json:"name"`
		Options     any                 `json:"options"`
		PositionX   int                 `json:"position_x"`
		PositionY   int                 `json:"position_y"`
		Reject      *DirectusOperations `json:"reject"`
		Resolve     *DirectusOperations `json:"resolve"`
		Type        string              `json:"type"`
		UserCreated *DirectusUsers      `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusoperations_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DateCreated = _obj.DateCreated
		cf.Flow = _obj.Flow
		cf.Id = _obj.Id
		cf.Key = _obj.Key
		cf.Name = _obj.Name
		cf.Options = _obj.Options
		cf.PositionX = _obj.PositionX
		cf.PositionY = _obj.PositionY
		cf.Reject = _obj.Reject
		cf.Resolve = _obj.Resolve
		cf.Type = _obj.Type
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusOperations) DeepCopy() IDirectusObject {
	new_obj := &DirectusOperations{}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.Flow != nil {
		new_obj.Flow = (*cf.Flow).DeepCopy().(*DirectusFlows)
	}
	new_obj.Id = cf.Id
	new_obj.Key = cf.Key
	if cf.Name != nil {
		temp := ""
		new_obj.Name = &temp
		*new_obj.Name = *cf.Name
	}
	new_obj.Options = cf.Options
	new_obj.PositionX = cf.PositionX
	new_obj.PositionY = cf.PositionY
	if cf.Reject != nil {
		new_obj.Reject = (*cf.Reject).DeepCopy().(*DirectusOperations)
	}
	if cf.Resolve != nil {
		new_obj.Resolve = (*cf.Resolve).DeepCopy().(*DirectusOperations)
	}
	new_obj.Type = cf.Type
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf DirectusOperations) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*DirectusOperations).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*DirectusOperations).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*DirectusOperations).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*DirectusOperations).Id {
		diff["id"] = cf.Id
	}

	if cf.Key != old.(*DirectusOperations).Key {
		diff["key"] = cf.Key
	}
	if cf.Name == nil {
		if old.(*DirectusOperations).Name != nil {
			diff["name"] = nil
		}
	} else {
		if old.(*DirectusOperations).Name == nil {
			diff["name"] = cf.Name
		} else {
			if *cf.Name != *old.(*DirectusOperations).Name {
				diff["name"] = cf.Name
			}
		}
	}

	if cf.Options != old.(*DirectusOperations).Options {
		diff["options"] = cf.Options
	}

	if cf.PositionX != old.(*DirectusOperations).PositionX {
		diff["position_x"] = cf.PositionX
	}

	if cf.PositionY != old.(*DirectusOperations).PositionY {
		diff["position_y"] = cf.PositionY
	}

	if cf.Type != old.(*DirectusOperations).Type {
		diff["type"] = cf.Type
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusOperations) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Flow != nil {
		trakingList = append(trakingList, cf.Flow)
		trakingList = append(trakingList, cf.Flow.Track()...)
	}

	if cf.Reject != nil {
		trakingList = append(trakingList, cf.Reject)
		trakingList = append(trakingList, cf.Reject.Track()...)
	}
	if cf.Resolve != nil {
		trakingList = append(trakingList, cf.Resolve)
		trakingList = append(trakingList, cf.Resolve.Track()...)
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf DirectusOperations) GetId() string {
	return cf.Id.String()
}

func (cf DirectusOperations) CollectionName() string {
	return "directus_operations"
}

type DirectusPanels struct {
	IDirectusObject
	Color       *string             `json:"color"`
	Dashboard   *DirectusDashboards `json:"dashboard"`
	DateCreated *time.Time          `json:"date_created"`
	Height      int                 `json:"height"`
	Icon        *string             `json:"icon"`
	Id          uuid.UUID           `json:"id"`
	Name        *string             `json:"name"`
	Note        *string             `json:"note"`
	Options     any                 `json:"options"`
	PositionX   int                 `json:"position_x"`
	PositionY   int                 `json:"position_y"`
	ShowHeader  bool                `json:"show_header"`
	Type        string              `json:"type"`
	UserCreated *DirectusUsers      `json:"user_created"`
	Width       int                 `json:"width"`
}

func (cf *DirectusPanels) UnmarshalJSON(data []byte) error {
	type directuspanels_internal struct {
		Color       *string             `json:"color"`
		Dashboard   *DirectusDashboards `json:"dashboard"`
		DateCreated *time.Time          `json:"date_created"`
		Height      int                 `json:"height"`
		Icon        *string             `json:"icon"`
		Id          uuid.UUID           `json:"id"`
		Name        *string             `json:"name"`
		Note        *string             `json:"note"`
		Options     any                 `json:"options"`
		PositionX   int                 `json:"position_x"`
		PositionY   int                 `json:"position_y"`
		ShowHeader  bool                `json:"show_header"`
		Type        string              `json:"type"`
		UserCreated *DirectusUsers      `json:"user_created"`
		Width       int                 `json:"width"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directuspanels_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Color = _obj.Color
		cf.Dashboard = _obj.Dashboard
		cf.DateCreated = _obj.DateCreated
		cf.Height = _obj.Height
		cf.Icon = _obj.Icon
		cf.Id = _obj.Id
		cf.Name = _obj.Name
		cf.Note = _obj.Note
		cf.Options = _obj.Options
		cf.PositionX = _obj.PositionX
		cf.PositionY = _obj.PositionY
		cf.ShowHeader = _obj.ShowHeader
		cf.Type = _obj.Type
		cf.UserCreated = _obj.UserCreated
		cf.Width = _obj.Width
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusPanels) DeepCopy() IDirectusObject {
	new_obj := &DirectusPanels{}
	if cf.Color != nil {
		temp := ""
		new_obj.Color = &temp
		*new_obj.Color = *cf.Color
	}
	if cf.Dashboard != nil {
		new_obj.Dashboard = (*cf.Dashboard).DeepCopy().(*DirectusDashboards)
	}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Height = cf.Height
	if cf.Icon != nil {
		temp := ""
		new_obj.Icon = &temp
		*new_obj.Icon = *cf.Icon
	}
	new_obj.Id = cf.Id
	if cf.Name != nil {
		temp := ""
		new_obj.Name = &temp
		*new_obj.Name = *cf.Name
	}
	if cf.Note != nil {
		temp := ""
		new_obj.Note = &temp
		*new_obj.Note = *cf.Note
	}
	new_obj.Options = cf.Options
	new_obj.PositionX = cf.PositionX
	new_obj.PositionY = cf.PositionY
	new_obj.ShowHeader = cf.ShowHeader
	new_obj.Type = cf.Type
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	new_obj.Width = cf.Width
	return new_obj
}
func (cf DirectusPanels) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Color == nil {
		if old.(*DirectusPanels).Color != nil {
			diff["color"] = nil
		}
	} else {
		if old.(*DirectusPanels).Color == nil {
			diff["color"] = cf.Color
		} else {
			if *cf.Color != *old.(*DirectusPanels).Color {
				diff["color"] = cf.Color
			}
		}
	}

	if cf.DateCreated == nil {
		if old.(*DirectusPanels).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*DirectusPanels).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*DirectusPanels).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Height != old.(*DirectusPanels).Height {
		diff["height"] = cf.Height
	}
	if cf.Icon == nil {
		if old.(*DirectusPanels).Icon != nil {
			diff["icon"] = nil
		}
	} else {
		if old.(*DirectusPanels).Icon == nil {
			diff["icon"] = cf.Icon
		} else {
			if *cf.Icon != *old.(*DirectusPanels).Icon {
				diff["icon"] = cf.Icon
			}
		}
	}

	if cf.Id != old.(*DirectusPanels).Id {
		diff["id"] = cf.Id
	}
	if cf.Name == nil {
		if old.(*DirectusPanels).Name != nil {
			diff["name"] = nil
		}
	} else {
		if old.(*DirectusPanels).Name == nil {
			diff["name"] = cf.Name
		} else {
			if *cf.Name != *old.(*DirectusPanels).Name {
				diff["name"] = cf.Name
			}
		}
	}
	if cf.Note == nil {
		if old.(*DirectusPanels).Note != nil {
			diff["note"] = nil
		}
	} else {
		if old.(*DirectusPanels).Note == nil {
			diff["note"] = cf.Note
		} else {
			if *cf.Note != *old.(*DirectusPanels).Note {
				diff["note"] = cf.Note
			}
		}
	}

	if cf.Options != old.(*DirectusPanels).Options {
		diff["options"] = cf.Options
	}

	if cf.PositionX != old.(*DirectusPanels).PositionX {
		diff["position_x"] = cf.PositionX
	}

	if cf.PositionY != old.(*DirectusPanels).PositionY {
		diff["position_y"] = cf.PositionY
	}

	if cf.ShowHeader != old.(*DirectusPanels).ShowHeader {
		diff["show_header"] = cf.ShowHeader
	}

	if cf.Type != old.(*DirectusPanels).Type {
		diff["type"] = cf.Type
	}

	if cf.Width != old.(*DirectusPanels).Width {
		diff["width"] = cf.Width
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusPanels) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Dashboard != nil {
		trakingList = append(trakingList, cf.Dashboard)
		trakingList = append(trakingList, cf.Dashboard.Track()...)
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}

	return trakingList
}

func (cf DirectusPanels) GetId() string {
	return cf.Id.String()
}

func (cf DirectusPanels) CollectionName() string {
	return "directus_panels"
}

type DirectusPermissions struct {
	IDirectusObject
	Action      string         `json:"action"`
	Collection  string         `json:"collection"`
	Fields      any            `json:"fields"`
	Id          int            `json:"id"`
	Permissions any            `json:"permissions"`
	Presets     any            `json:"presets"`
	Role        *DirectusRoles `json:"role"`
	Validation  any            `json:"validation"`
}

func (cf *DirectusPermissions) UnmarshalJSON(data []byte) error {
	type directuspermissions_internal struct {
		Action      string         `json:"action"`
		Collection  string         `json:"collection"`
		Fields      any            `json:"fields"`
		Id          int            `json:"id"`
		Permissions any            `json:"permissions"`
		Presets     any            `json:"presets"`
		Role        *DirectusRoles `json:"role"`
		Validation  any            `json:"validation"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directuspermissions_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Action = _obj.Action
		cf.Collection = _obj.Collection
		cf.Fields = _obj.Fields
		cf.Id = _obj.Id
		cf.Permissions = _obj.Permissions
		cf.Presets = _obj.Presets
		cf.Role = _obj.Role
		cf.Validation = _obj.Validation
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusPermissions) DeepCopy() IDirectusObject {
	new_obj := &DirectusPermissions{}
	new_obj.Action = cf.Action
	new_obj.Collection = cf.Collection
	new_obj.Fields = cf.Fields
	new_obj.Id = cf.Id
	new_obj.Permissions = cf.Permissions
	new_obj.Presets = cf.Presets
	if cf.Role != nil {
		new_obj.Role = (*cf.Role).DeepCopy().(*DirectusRoles)
	}
	new_obj.Validation = cf.Validation
	return new_obj
}
func (cf DirectusPermissions) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Action != old.(*DirectusPermissions).Action {
		diff["action"] = cf.Action
	}

	if cf.Collection != old.(*DirectusPermissions).Collection {
		diff["collection"] = cf.Collection
	}

	if cf.Fields != old.(*DirectusPermissions).Fields {
		diff["fields"] = cf.Fields
	}

	if cf.Id != old.(*DirectusPermissions).Id {
		diff["id"] = cf.Id
	}

	if cf.Permissions != old.(*DirectusPermissions).Permissions {
		diff["permissions"] = cf.Permissions
	}

	if cf.Presets != old.(*DirectusPermissions).Presets {
		diff["presets"] = cf.Presets
	}

	if cf.Validation != old.(*DirectusPermissions).Validation {
		diff["validation"] = cf.Validation
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusPermissions) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Role != nil {
		trakingList = append(trakingList, cf.Role)
		trakingList = append(trakingList, cf.Role.Track()...)
	}

	return trakingList
}

func (cf DirectusPermissions) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusPermissions) CollectionName() string {
	return "directus_permissions"
}

type DirectusPresets struct {
	IDirectusObject
	Bookmark        *string        `json:"bookmark"`
	Collection      *string        `json:"collection"`
	Color           *string        `json:"color"`
	Filter          any            `json:"filter"`
	Icon            *string        `json:"icon"`
	Id              int            `json:"id"`
	Layout          *string        `json:"layout"`
	LayoutOptions   any            `json:"layout_options"`
	LayoutQuery     any            `json:"layout_query"`
	RefreshInterval *int           `json:"refresh_interval"`
	Role            *DirectusRoles `json:"role"`
	Search          *string        `json:"search"`
	User            *DirectusUsers `json:"user"`
}

func (cf *DirectusPresets) UnmarshalJSON(data []byte) error {
	type directuspresets_internal struct {
		Bookmark        *string        `json:"bookmark"`
		Collection      *string        `json:"collection"`
		Color           *string        `json:"color"`
		Filter          any            `json:"filter"`
		Icon            *string        `json:"icon"`
		Id              int            `json:"id"`
		Layout          *string        `json:"layout"`
		LayoutOptions   any            `json:"layout_options"`
		LayoutQuery     any            `json:"layout_query"`
		RefreshInterval *int           `json:"refresh_interval"`
		Role            *DirectusRoles `json:"role"`
		Search          *string        `json:"search"`
		User            *DirectusUsers `json:"user"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directuspresets_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Bookmark = _obj.Bookmark
		cf.Collection = _obj.Collection
		cf.Color = _obj.Color
		cf.Filter = _obj.Filter
		cf.Icon = _obj.Icon
		cf.Id = _obj.Id
		cf.Layout = _obj.Layout
		cf.LayoutOptions = _obj.LayoutOptions
		cf.LayoutQuery = _obj.LayoutQuery
		cf.RefreshInterval = _obj.RefreshInterval
		cf.Role = _obj.Role
		cf.Search = _obj.Search
		cf.User = _obj.User
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusPresets) DeepCopy() IDirectusObject {
	new_obj := &DirectusPresets{}
	if cf.Bookmark != nil {
		temp := ""
		new_obj.Bookmark = &temp
		*new_obj.Bookmark = *cf.Bookmark
	}
	if cf.Collection != nil {
		temp := ""
		new_obj.Collection = &temp
		*new_obj.Collection = *cf.Collection
	}
	if cf.Color != nil {
		temp := ""
		new_obj.Color = &temp
		*new_obj.Color = *cf.Color
	}
	new_obj.Filter = cf.Filter
	if cf.Icon != nil {
		temp := ""
		new_obj.Icon = &temp
		*new_obj.Icon = *cf.Icon
	}
	new_obj.Id = cf.Id
	if cf.Layout != nil {
		temp := ""
		new_obj.Layout = &temp
		*new_obj.Layout = *cf.Layout
	}
	new_obj.LayoutOptions = cf.LayoutOptions
	new_obj.LayoutQuery = cf.LayoutQuery
	if cf.RefreshInterval != nil {
		temp := 0
		new_obj.RefreshInterval = &temp
		*new_obj.RefreshInterval = *cf.RefreshInterval
	}
	if cf.Role != nil {
		new_obj.Role = (*cf.Role).DeepCopy().(*DirectusRoles)
	}
	if cf.Search != nil {
		temp := ""
		new_obj.Search = &temp
		*new_obj.Search = *cf.Search
	}
	if cf.User != nil {
		new_obj.User = (*cf.User).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf DirectusPresets) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Bookmark == nil {
		if old.(*DirectusPresets).Bookmark != nil {
			diff["bookmark"] = nil
		}
	} else {
		if old.(*DirectusPresets).Bookmark == nil {
			diff["bookmark"] = cf.Bookmark
		} else {
			if *cf.Bookmark != *old.(*DirectusPresets).Bookmark {
				diff["bookmark"] = cf.Bookmark
			}
		}
	}
	if cf.Collection == nil {
		if old.(*DirectusPresets).Collection != nil {
			diff["collection"] = nil
		}
	} else {
		if old.(*DirectusPresets).Collection == nil {
			diff["collection"] = cf.Collection
		} else {
			if *cf.Collection != *old.(*DirectusPresets).Collection {
				diff["collection"] = cf.Collection
			}
		}
	}
	if cf.Color == nil {
		if old.(*DirectusPresets).Color != nil {
			diff["color"] = nil
		}
	} else {
		if old.(*DirectusPresets).Color == nil {
			diff["color"] = cf.Color
		} else {
			if *cf.Color != *old.(*DirectusPresets).Color {
				diff["color"] = cf.Color
			}
		}
	}

	if cf.Filter != old.(*DirectusPresets).Filter {
		diff["filter"] = cf.Filter
	}
	if cf.Icon == nil {
		if old.(*DirectusPresets).Icon != nil {
			diff["icon"] = nil
		}
	} else {
		if old.(*DirectusPresets).Icon == nil {
			diff["icon"] = cf.Icon
		} else {
			if *cf.Icon != *old.(*DirectusPresets).Icon {
				diff["icon"] = cf.Icon
			}
		}
	}

	if cf.Id != old.(*DirectusPresets).Id {
		diff["id"] = cf.Id
	}
	if cf.Layout == nil {
		if old.(*DirectusPresets).Layout != nil {
			diff["layout"] = nil
		}
	} else {
		if old.(*DirectusPresets).Layout == nil {
			diff["layout"] = cf.Layout
		} else {
			if *cf.Layout != *old.(*DirectusPresets).Layout {
				diff["layout"] = cf.Layout
			}
		}
	}

	if cf.LayoutOptions != old.(*DirectusPresets).LayoutOptions {
		diff["layout_options"] = cf.LayoutOptions
	}

	if cf.LayoutQuery != old.(*DirectusPresets).LayoutQuery {
		diff["layout_query"] = cf.LayoutQuery
	}
	if cf.RefreshInterval == nil {
		if old.(*DirectusPresets).RefreshInterval != nil {
			diff["refresh_interval"] = nil
		}
	} else {
		if old.(*DirectusPresets).RefreshInterval == nil {
			diff["refresh_interval"] = cf.RefreshInterval
		} else {
			if *cf.RefreshInterval != *old.(*DirectusPresets).RefreshInterval {
				diff["refresh_interval"] = cf.RefreshInterval
			}
		}
	}

	if cf.Search == nil {
		if old.(*DirectusPresets).Search != nil {
			diff["search"] = nil
		}
	} else {
		if old.(*DirectusPresets).Search == nil {
			diff["search"] = cf.Search
		} else {
			if *cf.Search != *old.(*DirectusPresets).Search {
				diff["search"] = cf.Search
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusPresets) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Role != nil {
		trakingList = append(trakingList, cf.Role)
		trakingList = append(trakingList, cf.Role.Track()...)
	}

	if cf.User != nil {
		trakingList = append(trakingList, cf.User)
		trakingList = append(trakingList, cf.User.Track()...)
	}
	return trakingList
}

func (cf DirectusPresets) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusPresets) CollectionName() string {
	return "directus_presets"
}

type DirectusRelations struct {
	IDirectusObject
	Id                    int     `json:"id"`
	JunctionField         *string `json:"junction_field"`
	ManyCollection        string  `json:"many_collection"`
	ManyField             string  `json:"many_field"`
	OneAllowedCollections any     `json:"one_allowed_collections"`
	OneCollection         *string `json:"one_collection"`
	OneCollectionField    *string `json:"one_collection_field"`
	OneDeselectAction     string  `json:"one_deselect_action"`
	OneField              *string `json:"one_field"`
	SortField             *string `json:"sort_field"`
}

func (cf *DirectusRelations) UnmarshalJSON(data []byte) error {
	type directusrelations_internal struct {
		Id                    int     `json:"id"`
		JunctionField         *string `json:"junction_field"`
		ManyCollection        string  `json:"many_collection"`
		ManyField             string  `json:"many_field"`
		OneAllowedCollections any     `json:"one_allowed_collections"`
		OneCollection         *string `json:"one_collection"`
		OneCollectionField    *string `json:"one_collection_field"`
		OneDeselectAction     string  `json:"one_deselect_action"`
		OneField              *string `json:"one_field"`
		SortField             *string `json:"sort_field"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusrelations_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Id = _obj.Id
		cf.JunctionField = _obj.JunctionField
		cf.ManyCollection = _obj.ManyCollection
		cf.ManyField = _obj.ManyField
		cf.OneAllowedCollections = _obj.OneAllowedCollections
		cf.OneCollection = _obj.OneCollection
		cf.OneCollectionField = _obj.OneCollectionField
		cf.OneDeselectAction = _obj.OneDeselectAction
		cf.OneField = _obj.OneField
		cf.SortField = _obj.SortField
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusRelations) DeepCopy() IDirectusObject {
	new_obj := &DirectusRelations{}
	new_obj.Id = cf.Id
	if cf.JunctionField != nil {
		temp := ""
		new_obj.JunctionField = &temp
		*new_obj.JunctionField = *cf.JunctionField
	}
	new_obj.ManyCollection = cf.ManyCollection
	new_obj.ManyField = cf.ManyField
	new_obj.OneAllowedCollections = cf.OneAllowedCollections
	if cf.OneCollection != nil {
		temp := ""
		new_obj.OneCollection = &temp
		*new_obj.OneCollection = *cf.OneCollection
	}
	if cf.OneCollectionField != nil {
		temp := ""
		new_obj.OneCollectionField = &temp
		*new_obj.OneCollectionField = *cf.OneCollectionField
	}
	new_obj.OneDeselectAction = cf.OneDeselectAction
	if cf.OneField != nil {
		temp := ""
		new_obj.OneField = &temp
		*new_obj.OneField = *cf.OneField
	}
	if cf.SortField != nil {
		temp := ""
		new_obj.SortField = &temp
		*new_obj.SortField = *cf.SortField
	}
	return new_obj
}
func (cf DirectusRelations) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Id != old.(*DirectusRelations).Id {
		diff["id"] = cf.Id
	}
	if cf.JunctionField == nil {
		if old.(*DirectusRelations).JunctionField != nil {
			diff["junction_field"] = nil
		}
	} else {
		if old.(*DirectusRelations).JunctionField == nil {
			diff["junction_field"] = cf.JunctionField
		} else {
			if *cf.JunctionField != *old.(*DirectusRelations).JunctionField {
				diff["junction_field"] = cf.JunctionField
			}
		}
	}

	if cf.ManyCollection != old.(*DirectusRelations).ManyCollection {
		diff["many_collection"] = cf.ManyCollection
	}

	if cf.ManyField != old.(*DirectusRelations).ManyField {
		diff["many_field"] = cf.ManyField
	}

	if cf.OneAllowedCollections != old.(*DirectusRelations).OneAllowedCollections {
		diff["one_allowed_collections"] = cf.OneAllowedCollections
	}
	if cf.OneCollection == nil {
		if old.(*DirectusRelations).OneCollection != nil {
			diff["one_collection"] = nil
		}
	} else {
		if old.(*DirectusRelations).OneCollection == nil {
			diff["one_collection"] = cf.OneCollection
		} else {
			if *cf.OneCollection != *old.(*DirectusRelations).OneCollection {
				diff["one_collection"] = cf.OneCollection
			}
		}
	}
	if cf.OneCollectionField == nil {
		if old.(*DirectusRelations).OneCollectionField != nil {
			diff["one_collection_field"] = nil
		}
	} else {
		if old.(*DirectusRelations).OneCollectionField == nil {
			diff["one_collection_field"] = cf.OneCollectionField
		} else {
			if *cf.OneCollectionField != *old.(*DirectusRelations).OneCollectionField {
				diff["one_collection_field"] = cf.OneCollectionField
			}
		}
	}

	if cf.OneDeselectAction != old.(*DirectusRelations).OneDeselectAction {
		diff["one_deselect_action"] = cf.OneDeselectAction
	}
	if cf.OneField == nil {
		if old.(*DirectusRelations).OneField != nil {
			diff["one_field"] = nil
		}
	} else {
		if old.(*DirectusRelations).OneField == nil {
			diff["one_field"] = cf.OneField
		} else {
			if *cf.OneField != *old.(*DirectusRelations).OneField {
				diff["one_field"] = cf.OneField
			}
		}
	}
	if cf.SortField == nil {
		if old.(*DirectusRelations).SortField != nil {
			diff["sort_field"] = nil
		}
	} else {
		if old.(*DirectusRelations).SortField == nil {
			diff["sort_field"] = cf.SortField
		} else {
			if *cf.SortField != *old.(*DirectusRelations).SortField {
				diff["sort_field"] = cf.SortField
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusRelations) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	return trakingList
}

func (cf DirectusRelations) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusRelations) CollectionName() string {
	return "directus_relations"
}

type DirectusRevisions struct {
	IDirectusObject
	Activity   *DirectusActivity  `json:"activity"`
	Collection string             `json:"collection"`
	Data       any                `json:"data"`
	Delta      any                `json:"delta"`
	Id         int                `json:"id"`
	Item       string             `json:"item"`
	Parent     *DirectusRevisions `json:"parent"`
	Version    *DirectusVersions  `json:"version"`
}

func (cf *DirectusRevisions) UnmarshalJSON(data []byte) error {
	type directusrevisions_internal struct {
		Activity   *DirectusActivity  `json:"activity"`
		Collection string             `json:"collection"`
		Data       any                `json:"data"`
		Delta      any                `json:"delta"`
		Id         int                `json:"id"`
		Item       string             `json:"item"`
		Parent     *DirectusRevisions `json:"parent"`
		Version    *DirectusVersions  `json:"version"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusrevisions_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Activity = _obj.Activity
		cf.Collection = _obj.Collection
		cf.Data = _obj.Data
		cf.Delta = _obj.Delta
		cf.Id = _obj.Id
		cf.Item = _obj.Item
		cf.Parent = _obj.Parent
		cf.Version = _obj.Version
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusRevisions) DeepCopy() IDirectusObject {
	new_obj := &DirectusRevisions{}
	if cf.Activity != nil {
		new_obj.Activity = (*cf.Activity).DeepCopy().(*DirectusActivity)
	}
	new_obj.Collection = cf.Collection
	new_obj.Data = cf.Data
	new_obj.Delta = cf.Delta
	new_obj.Id = cf.Id
	new_obj.Item = cf.Item
	if cf.Parent != nil {
		new_obj.Parent = (*cf.Parent).DeepCopy().(*DirectusRevisions)
	}
	if cf.Version != nil {
		new_obj.Version = (*cf.Version).DeepCopy().(*DirectusVersions)
	}
	return new_obj
}
func (cf DirectusRevisions) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Collection != old.(*DirectusRevisions).Collection {
		diff["collection"] = cf.Collection
	}

	if cf.Data != old.(*DirectusRevisions).Data {
		diff["data"] = cf.Data
	}

	if cf.Delta != old.(*DirectusRevisions).Delta {
		diff["delta"] = cf.Delta
	}

	if cf.Id != old.(*DirectusRevisions).Id {
		diff["id"] = cf.Id
	}

	if cf.Item != old.(*DirectusRevisions).Item {
		diff["item"] = cf.Item
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusRevisions) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Activity != nil {
		trakingList = append(trakingList, cf.Activity)
		trakingList = append(trakingList, cf.Activity.Track()...)
	}

	if cf.Parent != nil {
		trakingList = append(trakingList, cf.Parent)
		trakingList = append(trakingList, cf.Parent.Track()...)
	}
	if cf.Version != nil {
		trakingList = append(trakingList, cf.Version)
		trakingList = append(trakingList, cf.Version.Track()...)
	}
	return trakingList
}

func (cf DirectusRevisions) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusRevisions) CollectionName() string {
	return "directus_revisions"
}

type DirectusRoles struct {
	IDirectusObject
	AdminAccess bool            `json:"admin_access"`
	AppAccess   bool            `json:"app_access"`
	Description *string         `json:"description"`
	EnforceTfa  bool            `json:"enforce_tfa"`
	Icon        string          `json:"icon"`
	Id          uuid.UUID       `json:"id"`
	IpAccess    any             `json:"ip_access"`
	Name        string          `json:"name"`
	Users       []DirectusUsers `json:"users"`
}

func (cf *DirectusRoles) UnmarshalJSON(data []byte) error {
	type directusroles_internal struct {
		AdminAccess bool            `json:"admin_access"`
		AppAccess   bool            `json:"app_access"`
		Description *string         `json:"description"`
		EnforceTfa  bool            `json:"enforce_tfa"`
		Icon        string          `json:"icon"`
		Id          uuid.UUID       `json:"id"`
		IpAccess    any             `json:"ip_access"`
		Name        string          `json:"name"`
		Users       []DirectusUsers `json:"users"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusroles_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.AdminAccess = _obj.AdminAccess
		cf.AppAccess = _obj.AppAccess
		cf.Description = _obj.Description
		cf.EnforceTfa = _obj.EnforceTfa
		cf.Icon = _obj.Icon
		cf.Id = _obj.Id
		cf.IpAccess = _obj.IpAccess
		cf.Name = _obj.Name
		cf.Users = _obj.Users
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusRoles) DeepCopy() IDirectusObject {
	new_obj := &DirectusRoles{}
	new_obj.AdminAccess = cf.AdminAccess
	new_obj.AppAccess = cf.AppAccess
	if cf.Description != nil {
		temp := ""
		new_obj.Description = &temp
		*new_obj.Description = *cf.Description
	}
	new_obj.EnforceTfa = cf.EnforceTfa
	new_obj.Icon = cf.Icon
	new_obj.Id = cf.Id
	new_obj.IpAccess = cf.IpAccess
	new_obj.Name = cf.Name
	if cf.Users != nil {
		new_obj.Users = make([]DirectusUsers, len(cf.Users))
		copy(new_obj.Users, cf.Users)
	}
	return new_obj
}
func (cf DirectusRoles) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.AdminAccess != old.(*DirectusRoles).AdminAccess {
		diff["admin_access"] = cf.AdminAccess
	}

	if cf.AppAccess != old.(*DirectusRoles).AppAccess {
		diff["app_access"] = cf.AppAccess
	}
	if cf.Description == nil {
		if old.(*DirectusRoles).Description != nil {
			diff["description"] = nil
		}
	} else {
		if old.(*DirectusRoles).Description == nil {
			diff["description"] = cf.Description
		} else {
			if *cf.Description != *old.(*DirectusRoles).Description {
				diff["description"] = cf.Description
			}
		}
	}

	if cf.EnforceTfa != old.(*DirectusRoles).EnforceTfa {
		diff["enforce_tfa"] = cf.EnforceTfa
	}

	if cf.Icon != old.(*DirectusRoles).Icon {
		diff["icon"] = cf.Icon
	}

	if cf.Id != old.(*DirectusRoles).Id {
		diff["id"] = cf.Id
	}

	if cf.IpAccess != old.(*DirectusRoles).IpAccess {
		diff["ip_access"] = cf.IpAccess
	}

	if cf.Name != old.(*DirectusRoles).Name {
		diff["name"] = cf.Name
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusRoles) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Users != nil {
		for _, iter := range cf.Users {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	return trakingList
}

func (cf DirectusRoles) GetId() string {
	return cf.Id.String()
}

func (cf DirectusRoles) CollectionName() string {
	return "directus_roles"
}

type DirectusSettings struct {
	IDirectusObject
	AuthLoginAttempts     *int             `json:"auth_login_attempts"`
	AuthPasswordPolicy    *string          `json:"auth_password_policy"`
	Basemaps              any              `json:"basemaps"`
	BrandingDivider       any              `json:"branding_divider"`
	CustomAspectRatios    any              `json:"custom_aspect_ratios"`
	CustomCss             *string          `json:"custom_css"`
	DefaultAppearance     string           `json:"default_appearance"`
	DefaultLanguage       string           `json:"default_language"`
	DefaultThemeDark      *string          `json:"default_theme_dark"`
	DefaultThemeLight     *string          `json:"default_theme_light"`
	FilesDivider          any              `json:"files_divider"`
	Id                    int              `json:"id"`
	ImageEditor           any              `json:"image_editor"`
	MapDivider            any              `json:"map_divider"`
	MapboxKey             *string          `json:"mapbox_key"`
	ModuleBar             any              `json:"module_bar"`
	ModulesDivider        any              `json:"modules_divider"`
	ProjectColor          string           `json:"project_color"`
	ProjectDescriptor     *string          `json:"project_descriptor"`
	ProjectLogo           *DirectusFiles   `json:"project_logo"`
	ProjectName           string           `json:"project_name"`
	ProjectUrl            *string          `json:"project_url"`
	PublicBackground      *DirectusFiles   `json:"public_background"`
	PublicFavicon         *uuid.UUID       `json:"public_favicon"`
	PublicForeground      *DirectusFiles   `json:"public_foreground"`
	PublicNote            *string          `json:"public_note"`
	SecurityDivider       any              `json:"security_divider"`
	StorageAssetPresets   any              `json:"storage_asset_presets"`
	StorageAssetTransform *string          `json:"storage_asset_transform"`
	StorageDefaultFolder  *DirectusFolders `json:"storage_default_folder"`
	ThemeDarkOverrides    any              `json:"theme_dark_overrides"`
	ThemeLightOverrides   any              `json:"theme_light_overrides"`
	ThemingDivider        any              `json:"theming_divider"`
	ThemingGroup          any              `json:"theming_group"`
}

func (cf *DirectusSettings) UnmarshalJSON(data []byte) error {
	type directussettings_internal struct {
		AuthLoginAttempts     *int             `json:"auth_login_attempts"`
		AuthPasswordPolicy    *string          `json:"auth_password_policy"`
		Basemaps              any              `json:"basemaps"`
		BrandingDivider       any              `json:"branding_divider"`
		CustomAspectRatios    any              `json:"custom_aspect_ratios"`
		CustomCss             *string          `json:"custom_css"`
		DefaultAppearance     string           `json:"default_appearance"`
		DefaultLanguage       string           `json:"default_language"`
		DefaultThemeDark      *string          `json:"default_theme_dark"`
		DefaultThemeLight     *string          `json:"default_theme_light"`
		FilesDivider          any              `json:"files_divider"`
		Id                    int              `json:"id"`
		ImageEditor           any              `json:"image_editor"`
		MapDivider            any              `json:"map_divider"`
		MapboxKey             *string          `json:"mapbox_key"`
		ModuleBar             any              `json:"module_bar"`
		ModulesDivider        any              `json:"modules_divider"`
		ProjectColor          string           `json:"project_color"`
		ProjectDescriptor     *string          `json:"project_descriptor"`
		ProjectLogo           *DirectusFiles   `json:"project_logo"`
		ProjectName           string           `json:"project_name"`
		ProjectUrl            *string          `json:"project_url"`
		PublicBackground      *DirectusFiles   `json:"public_background"`
		PublicFavicon         *uuid.UUID       `json:"public_favicon"`
		PublicForeground      *DirectusFiles   `json:"public_foreground"`
		PublicNote            *string          `json:"public_note"`
		SecurityDivider       any              `json:"security_divider"`
		StorageAssetPresets   any              `json:"storage_asset_presets"`
		StorageAssetTransform *string          `json:"storage_asset_transform"`
		StorageDefaultFolder  *DirectusFolders `json:"storage_default_folder"`
		ThemeDarkOverrides    any              `json:"theme_dark_overrides"`
		ThemeLightOverrides   any              `json:"theme_light_overrides"`
		ThemingDivider        any              `json:"theming_divider"`
		ThemingGroup          any              `json:"theming_group"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directussettings_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.AuthLoginAttempts = _obj.AuthLoginAttempts
		cf.AuthPasswordPolicy = _obj.AuthPasswordPolicy
		cf.Basemaps = _obj.Basemaps
		cf.BrandingDivider = _obj.BrandingDivider
		cf.CustomAspectRatios = _obj.CustomAspectRatios
		cf.CustomCss = _obj.CustomCss
		cf.DefaultAppearance = _obj.DefaultAppearance
		cf.DefaultLanguage = _obj.DefaultLanguage
		cf.DefaultThemeDark = _obj.DefaultThemeDark
		cf.DefaultThemeLight = _obj.DefaultThemeLight
		cf.FilesDivider = _obj.FilesDivider
		cf.Id = _obj.Id
		cf.ImageEditor = _obj.ImageEditor
		cf.MapDivider = _obj.MapDivider
		cf.MapboxKey = _obj.MapboxKey
		cf.ModuleBar = _obj.ModuleBar
		cf.ModulesDivider = _obj.ModulesDivider
		cf.ProjectColor = _obj.ProjectColor
		cf.ProjectDescriptor = _obj.ProjectDescriptor
		cf.ProjectLogo = _obj.ProjectLogo
		cf.ProjectName = _obj.ProjectName
		cf.ProjectUrl = _obj.ProjectUrl
		cf.PublicBackground = _obj.PublicBackground
		cf.PublicFavicon = _obj.PublicFavicon
		cf.PublicForeground = _obj.PublicForeground
		cf.PublicNote = _obj.PublicNote
		cf.SecurityDivider = _obj.SecurityDivider
		cf.StorageAssetPresets = _obj.StorageAssetPresets
		cf.StorageAssetTransform = _obj.StorageAssetTransform
		cf.StorageDefaultFolder = _obj.StorageDefaultFolder
		cf.ThemeDarkOverrides = _obj.ThemeDarkOverrides
		cf.ThemeLightOverrides = _obj.ThemeLightOverrides
		cf.ThemingDivider = _obj.ThemingDivider
		cf.ThemingGroup = _obj.ThemingGroup
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusSettings) DeepCopy() IDirectusObject {
	new_obj := &DirectusSettings{}
	if cf.AuthLoginAttempts != nil {
		temp := 0
		new_obj.AuthLoginAttempts = &temp
		*new_obj.AuthLoginAttempts = *cf.AuthLoginAttempts
	}
	if cf.AuthPasswordPolicy != nil {
		temp := ""
		new_obj.AuthPasswordPolicy = &temp
		*new_obj.AuthPasswordPolicy = *cf.AuthPasswordPolicy
	}
	new_obj.Basemaps = cf.Basemaps
	new_obj.BrandingDivider = cf.BrandingDivider
	new_obj.CustomAspectRatios = cf.CustomAspectRatios
	if cf.CustomCss != nil {
		temp := ""
		new_obj.CustomCss = &temp
		*new_obj.CustomCss = *cf.CustomCss
	}
	new_obj.DefaultAppearance = cf.DefaultAppearance
	new_obj.DefaultLanguage = cf.DefaultLanguage
	if cf.DefaultThemeDark != nil {
		temp := ""
		new_obj.DefaultThemeDark = &temp
		*new_obj.DefaultThemeDark = *cf.DefaultThemeDark
	}
	if cf.DefaultThemeLight != nil {
		temp := ""
		new_obj.DefaultThemeLight = &temp
		*new_obj.DefaultThemeLight = *cf.DefaultThemeLight
	}
	new_obj.FilesDivider = cf.FilesDivider
	new_obj.Id = cf.Id
	new_obj.ImageEditor = cf.ImageEditor
	new_obj.MapDivider = cf.MapDivider
	if cf.MapboxKey != nil {
		temp := ""
		new_obj.MapboxKey = &temp
		*new_obj.MapboxKey = *cf.MapboxKey
	}
	new_obj.ModuleBar = cf.ModuleBar
	new_obj.ModulesDivider = cf.ModulesDivider
	new_obj.ProjectColor = cf.ProjectColor
	if cf.ProjectDescriptor != nil {
		temp := ""
		new_obj.ProjectDescriptor = &temp
		*new_obj.ProjectDescriptor = *cf.ProjectDescriptor
	}
	if cf.ProjectLogo != nil {
		new_obj.ProjectLogo = (*cf.ProjectLogo).DeepCopy().(*DirectusFiles)
	}
	new_obj.ProjectName = cf.ProjectName
	if cf.ProjectUrl != nil {
		temp := ""
		new_obj.ProjectUrl = &temp
		*new_obj.ProjectUrl = *cf.ProjectUrl
	}
	if cf.PublicBackground != nil {
		new_obj.PublicBackground = (*cf.PublicBackground).DeepCopy().(*DirectusFiles)
	}
	if cf.PublicFavicon != nil {
		temp := uuid.Nil
		new_obj.PublicFavicon = &temp
		*new_obj.PublicFavicon = *cf.PublicFavicon
	}
	if cf.PublicForeground != nil {
		new_obj.PublicForeground = (*cf.PublicForeground).DeepCopy().(*DirectusFiles)
	}
	if cf.PublicNote != nil {
		temp := ""
		new_obj.PublicNote = &temp
		*new_obj.PublicNote = *cf.PublicNote
	}
	new_obj.SecurityDivider = cf.SecurityDivider
	new_obj.StorageAssetPresets = cf.StorageAssetPresets
	if cf.StorageAssetTransform != nil {
		temp := ""
		new_obj.StorageAssetTransform = &temp
		*new_obj.StorageAssetTransform = *cf.StorageAssetTransform
	}
	if cf.StorageDefaultFolder != nil {
		new_obj.StorageDefaultFolder = (*cf.StorageDefaultFolder).DeepCopy().(*DirectusFolders)
	}
	new_obj.ThemeDarkOverrides = cf.ThemeDarkOverrides
	new_obj.ThemeLightOverrides = cf.ThemeLightOverrides
	new_obj.ThemingDivider = cf.ThemingDivider
	new_obj.ThemingGroup = cf.ThemingGroup
	return new_obj
}
func (cf DirectusSettings) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.AuthLoginAttempts == nil {
		if old.(*DirectusSettings).AuthLoginAttempts != nil {
			diff["auth_login_attempts"] = nil
		}
	} else {
		if old.(*DirectusSettings).AuthLoginAttempts == nil {
			diff["auth_login_attempts"] = cf.AuthLoginAttempts
		} else {
			if *cf.AuthLoginAttempts != *old.(*DirectusSettings).AuthLoginAttempts {
				diff["auth_login_attempts"] = cf.AuthLoginAttempts
			}
		}
	}
	if cf.AuthPasswordPolicy == nil {
		if old.(*DirectusSettings).AuthPasswordPolicy != nil {
			diff["auth_password_policy"] = nil
		}
	} else {
		if old.(*DirectusSettings).AuthPasswordPolicy == nil {
			diff["auth_password_policy"] = cf.AuthPasswordPolicy
		} else {
			if *cf.AuthPasswordPolicy != *old.(*DirectusSettings).AuthPasswordPolicy {
				diff["auth_password_policy"] = cf.AuthPasswordPolicy
			}
		}
	}

	if cf.Basemaps != old.(*DirectusSettings).Basemaps {
		diff["basemaps"] = cf.Basemaps
	}

	if cf.BrandingDivider != old.(*DirectusSettings).BrandingDivider {
		diff["branding_divider"] = cf.BrandingDivider
	}

	if cf.CustomAspectRatios != old.(*DirectusSettings).CustomAspectRatios {
		diff["custom_aspect_ratios"] = cf.CustomAspectRatios
	}
	if cf.CustomCss == nil {
		if old.(*DirectusSettings).CustomCss != nil {
			diff["custom_css"] = nil
		}
	} else {
		if old.(*DirectusSettings).CustomCss == nil {
			diff["custom_css"] = cf.CustomCss
		} else {
			if *cf.CustomCss != *old.(*DirectusSettings).CustomCss {
				diff["custom_css"] = cf.CustomCss
			}
		}
	}

	if cf.DefaultAppearance != old.(*DirectusSettings).DefaultAppearance {
		diff["default_appearance"] = cf.DefaultAppearance
	}

	if cf.DefaultLanguage != old.(*DirectusSettings).DefaultLanguage {
		diff["default_language"] = cf.DefaultLanguage
	}
	if cf.DefaultThemeDark == nil {
		if old.(*DirectusSettings).DefaultThemeDark != nil {
			diff["default_theme_dark"] = nil
		}
	} else {
		if old.(*DirectusSettings).DefaultThemeDark == nil {
			diff["default_theme_dark"] = cf.DefaultThemeDark
		} else {
			if *cf.DefaultThemeDark != *old.(*DirectusSettings).DefaultThemeDark {
				diff["default_theme_dark"] = cf.DefaultThemeDark
			}
		}
	}
	if cf.DefaultThemeLight == nil {
		if old.(*DirectusSettings).DefaultThemeLight != nil {
			diff["default_theme_light"] = nil
		}
	} else {
		if old.(*DirectusSettings).DefaultThemeLight == nil {
			diff["default_theme_light"] = cf.DefaultThemeLight
		} else {
			if *cf.DefaultThemeLight != *old.(*DirectusSettings).DefaultThemeLight {
				diff["default_theme_light"] = cf.DefaultThemeLight
			}
		}
	}

	if cf.FilesDivider != old.(*DirectusSettings).FilesDivider {
		diff["files_divider"] = cf.FilesDivider
	}

	if cf.Id != old.(*DirectusSettings).Id {
		diff["id"] = cf.Id
	}

	if cf.ImageEditor != old.(*DirectusSettings).ImageEditor {
		diff["image_editor"] = cf.ImageEditor
	}

	if cf.MapDivider != old.(*DirectusSettings).MapDivider {
		diff["map_divider"] = cf.MapDivider
	}
	if cf.MapboxKey == nil {
		if old.(*DirectusSettings).MapboxKey != nil {
			diff["mapbox_key"] = nil
		}
	} else {
		if old.(*DirectusSettings).MapboxKey == nil {
			diff["mapbox_key"] = cf.MapboxKey
		} else {
			if *cf.MapboxKey != *old.(*DirectusSettings).MapboxKey {
				diff["mapbox_key"] = cf.MapboxKey
			}
		}
	}

	if cf.ModuleBar != old.(*DirectusSettings).ModuleBar {
		diff["module_bar"] = cf.ModuleBar
	}

	if cf.ModulesDivider != old.(*DirectusSettings).ModulesDivider {
		diff["modules_divider"] = cf.ModulesDivider
	}

	if cf.ProjectColor != old.(*DirectusSettings).ProjectColor {
		diff["project_color"] = cf.ProjectColor
	}
	if cf.ProjectDescriptor == nil {
		if old.(*DirectusSettings).ProjectDescriptor != nil {
			diff["project_descriptor"] = nil
		}
	} else {
		if old.(*DirectusSettings).ProjectDescriptor == nil {
			diff["project_descriptor"] = cf.ProjectDescriptor
		} else {
			if *cf.ProjectDescriptor != *old.(*DirectusSettings).ProjectDescriptor {
				diff["project_descriptor"] = cf.ProjectDescriptor
			}
		}
	}

	if cf.ProjectName != old.(*DirectusSettings).ProjectName {
		diff["project_name"] = cf.ProjectName
	}
	if cf.ProjectUrl == nil {
		if old.(*DirectusSettings).ProjectUrl != nil {
			diff["project_url"] = nil
		}
	} else {
		if old.(*DirectusSettings).ProjectUrl == nil {
			diff["project_url"] = cf.ProjectUrl
		} else {
			if *cf.ProjectUrl != *old.(*DirectusSettings).ProjectUrl {
				diff["project_url"] = cf.ProjectUrl
			}
		}
	}

	if cf.PublicFavicon == nil {
		if old.(*DirectusSettings).PublicFavicon != nil {
			diff["public_favicon"] = nil
		}
	} else {
		if old.(*DirectusSettings).PublicFavicon == nil {
			diff["public_favicon"] = cf.PublicFavicon
		} else {
			if *cf.PublicFavicon != *old.(*DirectusSettings).PublicFavicon {
				diff["public_favicon"] = cf.PublicFavicon
			}
		}
	}

	if cf.PublicNote == nil {
		if old.(*DirectusSettings).PublicNote != nil {
			diff["public_note"] = nil
		}
	} else {
		if old.(*DirectusSettings).PublicNote == nil {
			diff["public_note"] = cf.PublicNote
		} else {
			if *cf.PublicNote != *old.(*DirectusSettings).PublicNote {
				diff["public_note"] = cf.PublicNote
			}
		}
	}

	if cf.SecurityDivider != old.(*DirectusSettings).SecurityDivider {
		diff["security_divider"] = cf.SecurityDivider
	}

	if cf.StorageAssetPresets != old.(*DirectusSettings).StorageAssetPresets {
		diff["storage_asset_presets"] = cf.StorageAssetPresets
	}
	if cf.StorageAssetTransform == nil {
		if old.(*DirectusSettings).StorageAssetTransform != nil {
			diff["storage_asset_transform"] = nil
		}
	} else {
		if old.(*DirectusSettings).StorageAssetTransform == nil {
			diff["storage_asset_transform"] = cf.StorageAssetTransform
		} else {
			if *cf.StorageAssetTransform != *old.(*DirectusSettings).StorageAssetTransform {
				diff["storage_asset_transform"] = cf.StorageAssetTransform
			}
		}
	}

	if cf.ThemeDarkOverrides != old.(*DirectusSettings).ThemeDarkOverrides {
		diff["theme_dark_overrides"] = cf.ThemeDarkOverrides
	}

	if cf.ThemeLightOverrides != old.(*DirectusSettings).ThemeLightOverrides {
		diff["theme_light_overrides"] = cf.ThemeLightOverrides
	}

	if cf.ThemingDivider != old.(*DirectusSettings).ThemingDivider {
		diff["theming_divider"] = cf.ThemingDivider
	}

	if cf.ThemingGroup != old.(*DirectusSettings).ThemingGroup {
		diff["theming_group"] = cf.ThemingGroup
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusSettings) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.ProjectLogo != nil {
		trakingList = append(trakingList, cf.ProjectLogo)
		trakingList = append(trakingList, cf.ProjectLogo.Track()...)
	}

	if cf.PublicBackground != nil {
		trakingList = append(trakingList, cf.PublicBackground)
		trakingList = append(trakingList, cf.PublicBackground.Track()...)
	}

	if cf.PublicForeground != nil {
		trakingList = append(trakingList, cf.PublicForeground)
		trakingList = append(trakingList, cf.PublicForeground.Track()...)
	}

	if cf.StorageDefaultFolder != nil {
		trakingList = append(trakingList, cf.StorageDefaultFolder)
		trakingList = append(trakingList, cf.StorageDefaultFolder.Track()...)
	}

	return trakingList
}

func (cf DirectusSettings) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusSettings) CollectionName() string {
	return "directus_settings"
}

type DirectusShares struct {
	IDirectusObject
	DateCreated *time.Time     `json:"date_created"`
	DateEnd     *time.Time     `json:"date_end"`
	DateStart   *time.Time     `json:"date_start"`
	Id          uuid.UUID      `json:"id"`
	Item        string         `json:"item"`
	MaxUses     *int           `json:"max_uses"`
	Name        *string        `json:"name"`
	Password    *string        `json:"password"`
	Role        *DirectusRoles `json:"role"`
	TimesUsed   *int           `json:"times_used"`
	UserCreated *DirectusUsers `json:"user_created"`
}

func (cf *DirectusShares) UnmarshalJSON(data []byte) error {
	type directusshares_internal struct {
		DateCreated *time.Time     `json:"date_created"`
		DateEnd     *time.Time     `json:"date_end"`
		DateStart   *time.Time     `json:"date_start"`
		Id          uuid.UUID      `json:"id"`
		Item        string         `json:"item"`
		MaxUses     *int           `json:"max_uses"`
		Name        *string        `json:"name"`
		Password    *string        `json:"password"`
		Role        *DirectusRoles `json:"role"`
		TimesUsed   *int           `json:"times_used"`
		UserCreated *DirectusUsers `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusshares_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DateCreated = _obj.DateCreated
		cf.DateEnd = _obj.DateEnd
		cf.DateStart = _obj.DateStart
		cf.Id = _obj.Id
		cf.Item = _obj.Item
		cf.MaxUses = _obj.MaxUses
		cf.Name = _obj.Name
		cf.Password = _obj.Password
		cf.Role = _obj.Role
		cf.TimesUsed = _obj.TimesUsed
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusShares) DeepCopy() IDirectusObject {
	new_obj := &DirectusShares{}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.DateEnd != nil {
		temp := time.Time{}
		new_obj.DateEnd = &temp
		*new_obj.DateEnd = *cf.DateEnd
	}
	if cf.DateStart != nil {
		temp := time.Time{}
		new_obj.DateStart = &temp
		*new_obj.DateStart = *cf.DateStart
	}
	new_obj.Id = cf.Id
	new_obj.Item = cf.Item
	if cf.MaxUses != nil {
		temp := 0
		new_obj.MaxUses = &temp
		*new_obj.MaxUses = *cf.MaxUses
	}
	if cf.Name != nil {
		temp := ""
		new_obj.Name = &temp
		*new_obj.Name = *cf.Name
	}
	if cf.Password != nil {
		temp := ""
		new_obj.Password = &temp
		*new_obj.Password = *cf.Password
	}
	if cf.Role != nil {
		new_obj.Role = (*cf.Role).DeepCopy().(*DirectusRoles)
	}
	if cf.TimesUsed != nil {
		temp := 0
		new_obj.TimesUsed = &temp
		*new_obj.TimesUsed = *cf.TimesUsed
	}
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf DirectusShares) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*DirectusShares).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*DirectusShares).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*DirectusShares).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}
	if cf.DateEnd == nil {
		if old.(*DirectusShares).DateEnd != nil {
			diff["date_end"] = nil
		}
	} else {
		if old.(*DirectusShares).DateEnd == nil {
			diff["date_end"] = cf.DateEnd
		} else {
			if *cf.DateEnd != *old.(*DirectusShares).DateEnd {
				diff["date_end"] = cf.DateEnd
			}
		}
	}
	if cf.DateStart == nil {
		if old.(*DirectusShares).DateStart != nil {
			diff["date_start"] = nil
		}
	} else {
		if old.(*DirectusShares).DateStart == nil {
			diff["date_start"] = cf.DateStart
		} else {
			if *cf.DateStart != *old.(*DirectusShares).DateStart {
				diff["date_start"] = cf.DateStart
			}
		}
	}

	if cf.Id != old.(*DirectusShares).Id {
		diff["id"] = cf.Id
	}

	if cf.Item != old.(*DirectusShares).Item {
		diff["item"] = cf.Item
	}
	if cf.MaxUses == nil {
		if old.(*DirectusShares).MaxUses != nil {
			diff["max_uses"] = nil
		}
	} else {
		if old.(*DirectusShares).MaxUses == nil {
			diff["max_uses"] = cf.MaxUses
		} else {
			if *cf.MaxUses != *old.(*DirectusShares).MaxUses {
				diff["max_uses"] = cf.MaxUses
			}
		}
	}
	if cf.Name == nil {
		if old.(*DirectusShares).Name != nil {
			diff["name"] = nil
		}
	} else {
		if old.(*DirectusShares).Name == nil {
			diff["name"] = cf.Name
		} else {
			if *cf.Name != *old.(*DirectusShares).Name {
				diff["name"] = cf.Name
			}
		}
	}
	if cf.Password == nil {
		if old.(*DirectusShares).Password != nil {
			diff["password"] = nil
		}
	} else {
		if old.(*DirectusShares).Password == nil {
			diff["password"] = cf.Password
		} else {
			if *cf.Password != *old.(*DirectusShares).Password {
				diff["password"] = cf.Password
			}
		}
	}

	if cf.TimesUsed == nil {
		if old.(*DirectusShares).TimesUsed != nil {
			diff["times_used"] = nil
		}
	} else {
		if old.(*DirectusShares).TimesUsed == nil {
			diff["times_used"] = cf.TimesUsed
		} else {
			if *cf.TimesUsed != *old.(*DirectusShares).TimesUsed {
				diff["times_used"] = cf.TimesUsed
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusShares) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Role != nil {
		trakingList = append(trakingList, cf.Role)
		trakingList = append(trakingList, cf.Role.Track()...)
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf DirectusShares) GetId() string {
	return cf.Id.String()
}

func (cf DirectusShares) CollectionName() string {
	return "directus_shares"
}

type DirectusTranslations struct {
	IDirectusObject
	Id       uuid.UUID `json:"id"`
	Key      string    `json:"key"`
	Language string    `json:"language"`
	Value    string    `json:"value"`
}

func (cf *DirectusTranslations) UnmarshalJSON(data []byte) error {
	type directustranslations_internal struct {
		Id       uuid.UUID `json:"id"`
		Key      string    `json:"key"`
		Language string    `json:"language"`
		Value    string    `json:"value"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directustranslations_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Id = _obj.Id
		cf.Key = _obj.Key
		cf.Language = _obj.Language
		cf.Value = _obj.Value
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusTranslations) DeepCopy() IDirectusObject {
	new_obj := &DirectusTranslations{}
	new_obj.Id = cf.Id
	new_obj.Key = cf.Key
	new_obj.Language = cf.Language
	new_obj.Value = cf.Value
	return new_obj
}
func (cf DirectusTranslations) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Id != old.(*DirectusTranslations).Id {
		diff["id"] = cf.Id
	}

	if cf.Key != old.(*DirectusTranslations).Key {
		diff["key"] = cf.Key
	}

	if cf.Language != old.(*DirectusTranslations).Language {
		diff["language"] = cf.Language
	}

	if cf.Value != old.(*DirectusTranslations).Value {
		diff["value"] = cf.Value
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusTranslations) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	return trakingList
}

func (cf DirectusTranslations) GetId() string {
	return cf.Id.String()
}

func (cf DirectusTranslations) CollectionName() string {
	return "directus_translations"
}

type DirectusUsers struct {
	IDirectusObject
	AdminDivider        any            `json:"admin_divider"`
	Appearance          *string        `json:"appearance"`
	AuthData            any            `json:"auth_data"`
	Avatar              *DirectusFiles `json:"avatar"`
	Description         *string        `json:"description"`
	Email               *string        `json:"email"`
	EmailNotifications  *bool          `json:"email_notifications"`
	ExternalIdentifier  *string        `json:"external_identifier"`
	FirstName           *string        `json:"first_name"`
	Id                  uuid.UUID      `json:"id"`
	Language            *string        `json:"language"`
	LastAccess          *time.Time     `json:"last_access"`
	LastName            *string        `json:"last_name"`
	LastPage            *string        `json:"last_page"`
	Location            *string        `json:"location"`
	Password            *string        `json:"password"`
	PreferencesDivider  any            `json:"preferences_divider"`
	Provider            string         `json:"provider"`
	Role                *DirectusRoles `json:"role"`
	Status              string         `json:"status"`
	Tags                any            `json:"tags"`
	TfaSecret           *string        `json:"tfa_secret"`
	ThemeDark           *string        `json:"theme_dark"`
	ThemeDarkOverrides  any            `json:"theme_dark_overrides"`
	ThemeLight          *string        `json:"theme_light"`
	ThemeLightOverrides any            `json:"theme_light_overrides"`
	ThemingDivider      any            `json:"theming_divider"`
	Title               *string        `json:"title"`
	Token               *string        `json:"token"`
}

func (cf *DirectusUsers) UnmarshalJSON(data []byte) error {
	type directususers_internal struct {
		AdminDivider        any            `json:"admin_divider"`
		Appearance          *string        `json:"appearance"`
		AuthData            any            `json:"auth_data"`
		Avatar              *DirectusFiles `json:"avatar"`
		Description         *string        `json:"description"`
		Email               *string        `json:"email"`
		EmailNotifications  *bool          `json:"email_notifications"`
		ExternalIdentifier  *string        `json:"external_identifier"`
		FirstName           *string        `json:"first_name"`
		Id                  uuid.UUID      `json:"id"`
		Language            *string        `json:"language"`
		LastAccess          *time.Time     `json:"last_access"`
		LastName            *string        `json:"last_name"`
		LastPage            *string        `json:"last_page"`
		Location            *string        `json:"location"`
		Password            *string        `json:"password"`
		PreferencesDivider  any            `json:"preferences_divider"`
		Provider            string         `json:"provider"`
		Role                *DirectusRoles `json:"role"`
		Status              string         `json:"status"`
		Tags                any            `json:"tags"`
		TfaSecret           *string        `json:"tfa_secret"`
		ThemeDark           *string        `json:"theme_dark"`
		ThemeDarkOverrides  any            `json:"theme_dark_overrides"`
		ThemeLight          *string        `json:"theme_light"`
		ThemeLightOverrides any            `json:"theme_light_overrides"`
		ThemingDivider      any            `json:"theming_divider"`
		Title               *string        `json:"title"`
		Token               *string        `json:"token"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directususers_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.AdminDivider = _obj.AdminDivider
		cf.Appearance = _obj.Appearance
		cf.AuthData = _obj.AuthData
		cf.Avatar = _obj.Avatar
		cf.Description = _obj.Description
		cf.Email = _obj.Email
		cf.EmailNotifications = _obj.EmailNotifications
		cf.ExternalIdentifier = _obj.ExternalIdentifier
		cf.FirstName = _obj.FirstName
		cf.Id = _obj.Id
		cf.Language = _obj.Language
		cf.LastAccess = _obj.LastAccess
		cf.LastName = _obj.LastName
		cf.LastPage = _obj.LastPage
		cf.Location = _obj.Location
		cf.Password = _obj.Password
		cf.PreferencesDivider = _obj.PreferencesDivider
		cf.Provider = _obj.Provider
		cf.Role = _obj.Role
		cf.Status = _obj.Status
		cf.Tags = _obj.Tags
		cf.TfaSecret = _obj.TfaSecret
		cf.ThemeDark = _obj.ThemeDark
		cf.ThemeDarkOverrides = _obj.ThemeDarkOverrides
		cf.ThemeLight = _obj.ThemeLight
		cf.ThemeLightOverrides = _obj.ThemeLightOverrides
		cf.ThemingDivider = _obj.ThemingDivider
		cf.Title = _obj.Title
		cf.Token = _obj.Token
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusUsers) DeepCopy() IDirectusObject {
	new_obj := &DirectusUsers{}
	new_obj.AdminDivider = cf.AdminDivider
	if cf.Appearance != nil {
		temp := ""
		new_obj.Appearance = &temp
		*new_obj.Appearance = *cf.Appearance
	}
	new_obj.AuthData = cf.AuthData
	if cf.Avatar != nil {
		new_obj.Avatar = (*cf.Avatar).DeepCopy().(*DirectusFiles)
	}
	if cf.Description != nil {
		temp := ""
		new_obj.Description = &temp
		*new_obj.Description = *cf.Description
	}
	if cf.Email != nil {
		temp := ""
		new_obj.Email = &temp
		*new_obj.Email = *cf.Email
	}
	if cf.EmailNotifications != nil {
		temp := false
		new_obj.EmailNotifications = &temp
		*new_obj.EmailNotifications = *cf.EmailNotifications
	}
	if cf.ExternalIdentifier != nil {
		temp := ""
		new_obj.ExternalIdentifier = &temp
		*new_obj.ExternalIdentifier = *cf.ExternalIdentifier
	}
	if cf.FirstName != nil {
		temp := ""
		new_obj.FirstName = &temp
		*new_obj.FirstName = *cf.FirstName
	}
	new_obj.Id = cf.Id
	if cf.Language != nil {
		temp := ""
		new_obj.Language = &temp
		*new_obj.Language = *cf.Language
	}
	if cf.LastAccess != nil {
		temp := time.Time{}
		new_obj.LastAccess = &temp
		*new_obj.LastAccess = *cf.LastAccess
	}
	if cf.LastName != nil {
		temp := ""
		new_obj.LastName = &temp
		*new_obj.LastName = *cf.LastName
	}
	if cf.LastPage != nil {
		temp := ""
		new_obj.LastPage = &temp
		*new_obj.LastPage = *cf.LastPage
	}
	if cf.Location != nil {
		temp := ""
		new_obj.Location = &temp
		*new_obj.Location = *cf.Location
	}
	if cf.Password != nil {
		temp := ""
		new_obj.Password = &temp
		*new_obj.Password = *cf.Password
	}
	new_obj.PreferencesDivider = cf.PreferencesDivider
	new_obj.Provider = cf.Provider
	if cf.Role != nil {
		new_obj.Role = (*cf.Role).DeepCopy().(*DirectusRoles)
	}
	new_obj.Status = cf.Status
	new_obj.Tags = cf.Tags
	if cf.TfaSecret != nil {
		temp := ""
		new_obj.TfaSecret = &temp
		*new_obj.TfaSecret = *cf.TfaSecret
	}
	if cf.ThemeDark != nil {
		temp := ""
		new_obj.ThemeDark = &temp
		*new_obj.ThemeDark = *cf.ThemeDark
	}
	new_obj.ThemeDarkOverrides = cf.ThemeDarkOverrides
	if cf.ThemeLight != nil {
		temp := ""
		new_obj.ThemeLight = &temp
		*new_obj.ThemeLight = *cf.ThemeLight
	}
	new_obj.ThemeLightOverrides = cf.ThemeLightOverrides
	new_obj.ThemingDivider = cf.ThemingDivider
	if cf.Title != nil {
		temp := ""
		new_obj.Title = &temp
		*new_obj.Title = *cf.Title
	}
	if cf.Token != nil {
		temp := ""
		new_obj.Token = &temp
		*new_obj.Token = *cf.Token
	}
	return new_obj
}
func (cf DirectusUsers) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.AdminDivider != old.(*DirectusUsers).AdminDivider {
		diff["admin_divider"] = cf.AdminDivider
	}
	if cf.Appearance == nil {
		if old.(*DirectusUsers).Appearance != nil {
			diff["appearance"] = nil
		}
	} else {
		if old.(*DirectusUsers).Appearance == nil {
			diff["appearance"] = cf.Appearance
		} else {
			if *cf.Appearance != *old.(*DirectusUsers).Appearance {
				diff["appearance"] = cf.Appearance
			}
		}
	}

	if cf.AuthData != old.(*DirectusUsers).AuthData {
		diff["auth_data"] = cf.AuthData
	}

	if cf.Description == nil {
		if old.(*DirectusUsers).Description != nil {
			diff["description"] = nil
		}
	} else {
		if old.(*DirectusUsers).Description == nil {
			diff["description"] = cf.Description
		} else {
			if *cf.Description != *old.(*DirectusUsers).Description {
				diff["description"] = cf.Description
			}
		}
	}
	if cf.Email == nil {
		if old.(*DirectusUsers).Email != nil {
			diff["email"] = nil
		}
	} else {
		if old.(*DirectusUsers).Email == nil {
			diff["email"] = cf.Email
		} else {
			if *cf.Email != *old.(*DirectusUsers).Email {
				diff["email"] = cf.Email
			}
		}
	}
	if cf.EmailNotifications == nil {
		if old.(*DirectusUsers).EmailNotifications != nil {
			diff["email_notifications"] = nil
		}
	} else {
		if old.(*DirectusUsers).EmailNotifications == nil {
			diff["email_notifications"] = cf.EmailNotifications
		} else {
			if *cf.EmailNotifications != *old.(*DirectusUsers).EmailNotifications {
				diff["email_notifications"] = cf.EmailNotifications
			}
		}
	}
	if cf.ExternalIdentifier == nil {
		if old.(*DirectusUsers).ExternalIdentifier != nil {
			diff["external_identifier"] = nil
		}
	} else {
		if old.(*DirectusUsers).ExternalIdentifier == nil {
			diff["external_identifier"] = cf.ExternalIdentifier
		} else {
			if *cf.ExternalIdentifier != *old.(*DirectusUsers).ExternalIdentifier {
				diff["external_identifier"] = cf.ExternalIdentifier
			}
		}
	}
	if cf.FirstName == nil {
		if old.(*DirectusUsers).FirstName != nil {
			diff["first_name"] = nil
		}
	} else {
		if old.(*DirectusUsers).FirstName == nil {
			diff["first_name"] = cf.FirstName
		} else {
			if *cf.FirstName != *old.(*DirectusUsers).FirstName {
				diff["first_name"] = cf.FirstName
			}
		}
	}

	if cf.Id != old.(*DirectusUsers).Id {
		diff["id"] = cf.Id
	}
	if cf.Language == nil {
		if old.(*DirectusUsers).Language != nil {
			diff["language"] = nil
		}
	} else {
		if old.(*DirectusUsers).Language == nil {
			diff["language"] = cf.Language
		} else {
			if *cf.Language != *old.(*DirectusUsers).Language {
				diff["language"] = cf.Language
			}
		}
	}
	if cf.LastAccess == nil {
		if old.(*DirectusUsers).LastAccess != nil {
			diff["last_access"] = nil
		}
	} else {
		if old.(*DirectusUsers).LastAccess == nil {
			diff["last_access"] = cf.LastAccess
		} else {
			if *cf.LastAccess != *old.(*DirectusUsers).LastAccess {
				diff["last_access"] = cf.LastAccess
			}
		}
	}
	if cf.LastName == nil {
		if old.(*DirectusUsers).LastName != nil {
			diff["last_name"] = nil
		}
	} else {
		if old.(*DirectusUsers).LastName == nil {
			diff["last_name"] = cf.LastName
		} else {
			if *cf.LastName != *old.(*DirectusUsers).LastName {
				diff["last_name"] = cf.LastName
			}
		}
	}
	if cf.LastPage == nil {
		if old.(*DirectusUsers).LastPage != nil {
			diff["last_page"] = nil
		}
	} else {
		if old.(*DirectusUsers).LastPage == nil {
			diff["last_page"] = cf.LastPage
		} else {
			if *cf.LastPage != *old.(*DirectusUsers).LastPage {
				diff["last_page"] = cf.LastPage
			}
		}
	}
	if cf.Location == nil {
		if old.(*DirectusUsers).Location != nil {
			diff["location"] = nil
		}
	} else {
		if old.(*DirectusUsers).Location == nil {
			diff["location"] = cf.Location
		} else {
			if *cf.Location != *old.(*DirectusUsers).Location {
				diff["location"] = cf.Location
			}
		}
	}
	if cf.Password == nil {
		if old.(*DirectusUsers).Password != nil {
			diff["password"] = nil
		}
	} else {
		if old.(*DirectusUsers).Password == nil {
			diff["password"] = cf.Password
		} else {
			if *cf.Password != *old.(*DirectusUsers).Password {
				diff["password"] = cf.Password
			}
		}
	}

	if cf.PreferencesDivider != old.(*DirectusUsers).PreferencesDivider {
		diff["preferences_divider"] = cf.PreferencesDivider
	}

	if cf.Provider != old.(*DirectusUsers).Provider {
		diff["provider"] = cf.Provider
	}

	if cf.Status != old.(*DirectusUsers).Status {
		diff["status"] = cf.Status
	}

	if cf.Tags != old.(*DirectusUsers).Tags {
		diff["tags"] = cf.Tags
	}
	if cf.TfaSecret == nil {
		if old.(*DirectusUsers).TfaSecret != nil {
			diff["tfa_secret"] = nil
		}
	} else {
		if old.(*DirectusUsers).TfaSecret == nil {
			diff["tfa_secret"] = cf.TfaSecret
		} else {
			if *cf.TfaSecret != *old.(*DirectusUsers).TfaSecret {
				diff["tfa_secret"] = cf.TfaSecret
			}
		}
	}
	if cf.ThemeDark == nil {
		if old.(*DirectusUsers).ThemeDark != nil {
			diff["theme_dark"] = nil
		}
	} else {
		if old.(*DirectusUsers).ThemeDark == nil {
			diff["theme_dark"] = cf.ThemeDark
		} else {
			if *cf.ThemeDark != *old.(*DirectusUsers).ThemeDark {
				diff["theme_dark"] = cf.ThemeDark
			}
		}
	}

	if cf.ThemeDarkOverrides != old.(*DirectusUsers).ThemeDarkOverrides {
		diff["theme_dark_overrides"] = cf.ThemeDarkOverrides
	}
	if cf.ThemeLight == nil {
		if old.(*DirectusUsers).ThemeLight != nil {
			diff["theme_light"] = nil
		}
	} else {
		if old.(*DirectusUsers).ThemeLight == nil {
			diff["theme_light"] = cf.ThemeLight
		} else {
			if *cf.ThemeLight != *old.(*DirectusUsers).ThemeLight {
				diff["theme_light"] = cf.ThemeLight
			}
		}
	}

	if cf.ThemeLightOverrides != old.(*DirectusUsers).ThemeLightOverrides {
		diff["theme_light_overrides"] = cf.ThemeLightOverrides
	}

	if cf.ThemingDivider != old.(*DirectusUsers).ThemingDivider {
		diff["theming_divider"] = cf.ThemingDivider
	}
	if cf.Title == nil {
		if old.(*DirectusUsers).Title != nil {
			diff["title"] = nil
		}
	} else {
		if old.(*DirectusUsers).Title == nil {
			diff["title"] = cf.Title
		} else {
			if *cf.Title != *old.(*DirectusUsers).Title {
				diff["title"] = cf.Title
			}
		}
	}
	if cf.Token == nil {
		if old.(*DirectusUsers).Token != nil {
			diff["token"] = nil
		}
	} else {
		if old.(*DirectusUsers).Token == nil {
			diff["token"] = cf.Token
		} else {
			if *cf.Token != *old.(*DirectusUsers).Token {
				diff["token"] = cf.Token
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusUsers) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Avatar != nil {
		trakingList = append(trakingList, cf.Avatar)
		trakingList = append(trakingList, cf.Avatar.Track()...)
	}

	if cf.Role != nil {
		trakingList = append(trakingList, cf.Role)
		trakingList = append(trakingList, cf.Role.Track()...)
	}

	return trakingList
}

func (cf DirectusUsers) GetId() string {
	return cf.Id.String()
}

func (cf DirectusUsers) CollectionName() string {
	return "directus_users"
}

type DirectusVersions struct {
	IDirectusObject
	DateCreated *time.Time     `json:"date_created"`
	DateUpdated *time.Time     `json:"date_updated"`
	Hash        *string        `json:"hash"`
	Id          uuid.UUID      `json:"id"`
	Item        string         `json:"item"`
	Key         string         `json:"key"`
	Name        *string        `json:"name"`
	UserCreated *DirectusUsers `json:"user_created"`
	UserUpdated *DirectusUsers `json:"user_updated"`
}

func (cf *DirectusVersions) UnmarshalJSON(data []byte) error {
	type directusversions_internal struct {
		DateCreated *time.Time     `json:"date_created"`
		DateUpdated *time.Time     `json:"date_updated"`
		Hash        *string        `json:"hash"`
		Id          uuid.UUID      `json:"id"`
		Item        string         `json:"item"`
		Key         string         `json:"key"`
		Name        *string        `json:"name"`
		UserCreated *DirectusUsers `json:"user_created"`
		UserUpdated *DirectusUsers `json:"user_updated"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directusversions_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DateCreated = _obj.DateCreated
		cf.DateUpdated = _obj.DateUpdated
		cf.Hash = _obj.Hash
		cf.Id = _obj.Id
		cf.Item = _obj.Item
		cf.Key = _obj.Key
		cf.Name = _obj.Name
		cf.UserCreated = _obj.UserCreated
		cf.UserUpdated = _obj.UserUpdated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusVersions) DeepCopy() IDirectusObject {
	new_obj := &DirectusVersions{}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.DateUpdated != nil {
		temp := time.Time{}
		new_obj.DateUpdated = &temp
		*new_obj.DateUpdated = *cf.DateUpdated
	}
	if cf.Hash != nil {
		temp := ""
		new_obj.Hash = &temp
		*new_obj.Hash = *cf.Hash
	}
	new_obj.Id = cf.Id
	new_obj.Item = cf.Item
	new_obj.Key = cf.Key
	if cf.Name != nil {
		temp := ""
		new_obj.Name = &temp
		*new_obj.Name = *cf.Name
	}
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	if cf.UserUpdated != nil {
		new_obj.UserUpdated = (*cf.UserUpdated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf DirectusVersions) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*DirectusVersions).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*DirectusVersions).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*DirectusVersions).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}
	if cf.DateUpdated == nil {
		if old.(*DirectusVersions).DateUpdated != nil {
			diff["date_updated"] = nil
		}
	} else {
		if old.(*DirectusVersions).DateUpdated == nil {
			diff["date_updated"] = cf.DateUpdated
		} else {
			if *cf.DateUpdated != *old.(*DirectusVersions).DateUpdated {
				diff["date_updated"] = cf.DateUpdated
			}
		}
	}
	if cf.Hash == nil {
		if old.(*DirectusVersions).Hash != nil {
			diff["hash"] = nil
		}
	} else {
		if old.(*DirectusVersions).Hash == nil {
			diff["hash"] = cf.Hash
		} else {
			if *cf.Hash != *old.(*DirectusVersions).Hash {
				diff["hash"] = cf.Hash
			}
		}
	}

	if cf.Id != old.(*DirectusVersions).Id {
		diff["id"] = cf.Id
	}

	if cf.Item != old.(*DirectusVersions).Item {
		diff["item"] = cf.Item
	}

	if cf.Key != old.(*DirectusVersions).Key {
		diff["key"] = cf.Key
	}
	if cf.Name == nil {
		if old.(*DirectusVersions).Name != nil {
			diff["name"] = nil
		}
	} else {
		if old.(*DirectusVersions).Name == nil {
			diff["name"] = cf.Name
		} else {
			if *cf.Name != *old.(*DirectusVersions).Name {
				diff["name"] = cf.Name
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusVersions) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	if cf.UserUpdated != nil {
		trakingList = append(trakingList, cf.UserUpdated)
		trakingList = append(trakingList, cf.UserUpdated.Track()...)
	}
	return trakingList
}

func (cf DirectusVersions) GetId() string {
	return cf.Id.String()
}

func (cf DirectusVersions) CollectionName() string {
	return "directus_versions"
}

type DirectusWebhooks struct {
	IDirectusObject
	Actions         any    `json:"actions"`
	Collections     any    `json:"collections"`
	Data            bool   `json:"data"`
	Headers         any    `json:"headers"`
	Id              int    `json:"id"`
	Method          string `json:"method"`
	Name            string `json:"name"`
	Status          string `json:"status"`
	TriggersDivider any    `json:"triggers_divider"`
	Url             string `json:"url"`
}

func (cf *DirectusWebhooks) UnmarshalJSON(data []byte) error {
	type directuswebhooks_internal struct {
		Actions         any    `json:"actions"`
		Collections     any    `json:"collections"`
		Data            bool   `json:"data"`
		Headers         any    `json:"headers"`
		Id              int    `json:"id"`
		Method          string `json:"method"`
		Name            string `json:"name"`
		Status          string `json:"status"`
		TriggersDivider any    `json:"triggers_divider"`
		Url             string `json:"url"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj directuswebhooks_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Actions = _obj.Actions
		cf.Collections = _obj.Collections
		cf.Data = _obj.Data
		cf.Headers = _obj.Headers
		cf.Id = _obj.Id
		cf.Method = _obj.Method
		cf.Name = _obj.Name
		cf.Status = _obj.Status
		cf.TriggersDivider = _obj.TriggersDivider
		cf.Url = _obj.Url
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf DirectusWebhooks) DeepCopy() IDirectusObject {
	new_obj := &DirectusWebhooks{}
	new_obj.Actions = cf.Actions
	new_obj.Collections = cf.Collections
	new_obj.Data = cf.Data
	new_obj.Headers = cf.Headers
	new_obj.Id = cf.Id
	new_obj.Method = cf.Method
	new_obj.Name = cf.Name
	new_obj.Status = cf.Status
	new_obj.TriggersDivider = cf.TriggersDivider
	new_obj.Url = cf.Url
	return new_obj
}
func (cf DirectusWebhooks) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Actions != old.(*DirectusWebhooks).Actions {
		diff["actions"] = cf.Actions
	}

	if cf.Collections != old.(*DirectusWebhooks).Collections {
		diff["collections"] = cf.Collections
	}

	if cf.Data != old.(*DirectusWebhooks).Data {
		diff["data"] = cf.Data
	}

	if cf.Headers != old.(*DirectusWebhooks).Headers {
		diff["headers"] = cf.Headers
	}

	if cf.Id != old.(*DirectusWebhooks).Id {
		diff["id"] = cf.Id
	}

	if cf.Method != old.(*DirectusWebhooks).Method {
		diff["method"] = cf.Method
	}

	if cf.Name != old.(*DirectusWebhooks).Name {
		diff["name"] = cf.Name
	}

	if cf.Status != old.(*DirectusWebhooks).Status {
		diff["status"] = cf.Status
	}

	if cf.TriggersDivider != old.(*DirectusWebhooks).TriggersDivider {
		diff["triggers_divider"] = cf.TriggersDivider
	}

	if cf.Url != old.(*DirectusWebhooks).Url {
		diff["url"] = cf.Url
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf DirectusWebhooks) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	return trakingList
}

func (cf DirectusWebhooks) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf DirectusWebhooks) CollectionName() string {
	return "directus_webhooks"
}

type ExternalFiles struct {
	IDirectusObject
	Bucket      string         `json:"bucket"`
	DateCreated *time.Time     `json:"date_created"`
	Fileext     *string        `json:"fileext"`
	Filehash    string         `json:"filehash"`
	Id          uuid.UUID      `json:"id"`
	UserCreated *DirectusUsers `json:"user_created"`
}

func (cf *ExternalFiles) UnmarshalJSON(data []byte) error {
	type externalfiles_internal struct {
		Bucket      string         `json:"bucket"`
		DateCreated *time.Time     `json:"date_created"`
		Fileext     *string        `json:"fileext"`
		Filehash    string         `json:"filehash"`
		Id          uuid.UUID      `json:"id"`
		UserCreated *DirectusUsers `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj externalfiles_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Bucket = _obj.Bucket
		cf.DateCreated = _obj.DateCreated
		cf.Fileext = _obj.Fileext
		cf.Filehash = _obj.Filehash
		cf.Id = _obj.Id
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf ExternalFiles) DeepCopy() IDirectusObject {
	new_obj := &ExternalFiles{}
	new_obj.Bucket = cf.Bucket
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.Fileext != nil {
		temp := ""
		new_obj.Fileext = &temp
		*new_obj.Fileext = *cf.Fileext
	}
	new_obj.Filehash = cf.Filehash
	new_obj.Id = cf.Id
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf ExternalFiles) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Bucket != old.(*ExternalFiles).Bucket {
		diff["bucket"] = cf.Bucket
	}
	if cf.DateCreated == nil {
		if old.(*ExternalFiles).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*ExternalFiles).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*ExternalFiles).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}
	if cf.Fileext == nil {
		if old.(*ExternalFiles).Fileext != nil {
			diff["fileext"] = nil
		}
	} else {
		if old.(*ExternalFiles).Fileext == nil {
			diff["fileext"] = cf.Fileext
		} else {
			if *cf.Fileext != *old.(*ExternalFiles).Fileext {
				diff["fileext"] = cf.Fileext
			}
		}
	}

	if cf.Filehash != old.(*ExternalFiles).Filehash {
		diff["filehash"] = cf.Filehash
	}

	if cf.Id != old.(*ExternalFiles).Id {
		diff["id"] = cf.Id
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf ExternalFiles) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf ExternalFiles) GetId() string {
	return cf.Id.String()
}

func (cf ExternalFiles) CollectionName() string {
	return "ExternalFiles"
}

type GameProjects struct {
	IDirectusObject
	Annotation  *string        `json:"annotation"`
	Balance     float32        `json:"balance"`
	DateCreated *time.Time     `json:"date_created"`
	GameServers []GameServers  `json:"GameServers"`
	Id          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	UserCreated *DirectusUsers `json:"user_created"`
}

func (cf *GameProjects) UnmarshalJSON(data []byte) error {
	type gameprojects_internal struct {
		Annotation  *string        `json:"annotation"`
		Balance     float32        `json:"balance"`
		DateCreated *time.Time     `json:"date_created"`
		GameServers []GameServers  `json:"GameServers"`
		Id          uuid.UUID      `json:"id"`
		Name        string         `json:"name"`
		UserCreated *DirectusUsers `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj gameprojects_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Annotation = _obj.Annotation
		cf.Balance = _obj.Balance
		cf.DateCreated = _obj.DateCreated
		cf.GameServers = _obj.GameServers
		cf.Id = _obj.Id
		cf.Name = _obj.Name
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf GameProjects) DeepCopy() IDirectusObject {
	new_obj := &GameProjects{}
	if cf.Annotation != nil {
		temp := ""
		new_obj.Annotation = &temp
		*new_obj.Annotation = *cf.Annotation
	}
	new_obj.Balance = cf.Balance
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.GameServers != nil {
		new_obj.GameServers = make([]GameServers, len(cf.GameServers))
		copy(new_obj.GameServers, cf.GameServers)
	}
	new_obj.Id = cf.Id
	new_obj.Name = cf.Name
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf GameProjects) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Annotation == nil {
		if old.(*GameProjects).Annotation != nil {
			diff["annotation"] = nil
		}
	} else {
		if old.(*GameProjects).Annotation == nil {
			diff["annotation"] = cf.Annotation
		} else {
			if *cf.Annotation != *old.(*GameProjects).Annotation {
				diff["annotation"] = cf.Annotation
			}
		}
	}

	if cf.Balance != old.(*GameProjects).Balance {
		diff["balance"] = cf.Balance
	}
	if cf.DateCreated == nil {
		if old.(*GameProjects).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*GameProjects).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*GameProjects).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*GameProjects).Id {
		diff["id"] = cf.Id
	}

	if cf.Name != old.(*GameProjects).Name {
		diff["name"] = cf.Name
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf GameProjects) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.GameServers != nil {
		for _, iter := range cf.GameServers {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf GameProjects) GetId() string {
	return cf.Id.String()
}

func (cf GameProjects) CollectionName() string {
	return "GameProjects"
}

type GameServers struct {
	IDirectusObject
	Address            *string                 `json:"address"`
	DateCreated        *time.Time              `json:"date_created"`
	Id                 uuid.UUID               `json:"id"`
	Moderators         []GameServersModerators `json:"moderators"`
	Name               string                  `json:"name"`
	Project            *GameProjects           `json:"project"`
	ServerToken        *string                 `json:"server_token"`
	ServerChatMessages []ServerChatMessages    `json:"ServerChatMessages"`
	ServerPlayers      []ServerPlayers         `json:"ServerPlayers"`
	UserCreated        *DirectusUsers          `json:"user_created"`
}

func (cf *GameServers) UnmarshalJSON(data []byte) error {
	type gameservers_internal struct {
		Address            *string                 `json:"address"`
		DateCreated        *time.Time              `json:"date_created"`
		Id                 uuid.UUID               `json:"id"`
		Moderators         []GameServersModerators `json:"moderators"`
		Name               string                  `json:"name"`
		Project            *GameProjects           `json:"project"`
		ServerToken        *string                 `json:"server_token"`
		ServerChatMessages []ServerChatMessages    `json:"ServerChatMessages"`
		ServerPlayers      []ServerPlayers         `json:"ServerPlayers"`
		UserCreated        *DirectusUsers          `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj gameservers_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Address = _obj.Address
		cf.DateCreated = _obj.DateCreated
		cf.Id = _obj.Id
		cf.Moderators = _obj.Moderators
		cf.Name = _obj.Name
		cf.Project = _obj.Project
		cf.ServerToken = _obj.ServerToken
		cf.ServerChatMessages = _obj.ServerChatMessages
		cf.ServerPlayers = _obj.ServerPlayers
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf GameServers) DeepCopy() IDirectusObject {
	new_obj := &GameServers{}
	if cf.Address != nil {
		temp := ""
		new_obj.Address = &temp
		*new_obj.Address = *cf.Address
	}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Id = cf.Id
	if cf.Moderators != nil {
		new_obj.Moderators = make([]GameServersModerators, len(cf.Moderators))
		copy(new_obj.Moderators, cf.Moderators)
	}
	new_obj.Name = cf.Name
	if cf.Project != nil {
		new_obj.Project = (*cf.Project).DeepCopy().(*GameProjects)
	}
	if cf.ServerToken != nil {
		temp := ""
		new_obj.ServerToken = &temp
		*new_obj.ServerToken = *cf.ServerToken
	}
	if cf.ServerChatMessages != nil {
		new_obj.ServerChatMessages = make([]ServerChatMessages, len(cf.ServerChatMessages))
		copy(new_obj.ServerChatMessages, cf.ServerChatMessages)
	}
	if cf.ServerPlayers != nil {
		new_obj.ServerPlayers = make([]ServerPlayers, len(cf.ServerPlayers))
		copy(new_obj.ServerPlayers, cf.ServerPlayers)
	}
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf GameServers) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Address == nil {
		if old.(*GameServers).Address != nil {
			diff["address"] = nil
		}
	} else {
		if old.(*GameServers).Address == nil {
			diff["address"] = cf.Address
		} else {
			if *cf.Address != *old.(*GameServers).Address {
				diff["address"] = cf.Address
			}
		}
	}
	if cf.DateCreated == nil {
		if old.(*GameServers).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*GameServers).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*GameServers).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*GameServers).Id {
		diff["id"] = cf.Id
	}

	if cf.Name != old.(*GameServers).Name {
		diff["name"] = cf.Name
	}

	if cf.ServerToken == nil {
		if old.(*GameServers).ServerToken != nil {
			diff["server_token"] = nil
		}
	} else {
		if old.(*GameServers).ServerToken == nil {
			diff["server_token"] = cf.ServerToken
		} else {
			if *cf.ServerToken != *old.(*GameServers).ServerToken {
				diff["server_token"] = cf.ServerToken
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf GameServers) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Moderators != nil {
		for _, iter := range cf.Moderators {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.Project != nil {
		trakingList = append(trakingList, cf.Project)
		trakingList = append(trakingList, cf.Project.Track()...)
	}

	if cf.ServerChatMessages != nil {
		for _, iter := range cf.ServerChatMessages {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.ServerPlayers != nil {
		for _, iter := range cf.ServerPlayers {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf GameServers) GetId() string {
	return cf.Id.String()
}

func (cf GameServers) CollectionName() string {
	return "GameServers"
}

type GameServersModerators struct {
	IDirectusObject
	DirectusUsersId *DirectusUsers `json:"directus_users_id"`
	GameServersId   *GameServers   `json:"GameServers_id"`
	Id              int            `json:"id"`
}

func (cf *GameServersModerators) UnmarshalJSON(data []byte) error {
	type gameserversmoderators_internal struct {
		DirectusUsersId *DirectusUsers `json:"directus_users_id"`
		GameServersId   *GameServers   `json:"GameServers_id"`
		Id              int            `json:"id"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj gameserversmoderators_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DirectusUsersId = _obj.DirectusUsersId
		cf.GameServersId = _obj.GameServersId
		cf.Id = _obj.Id
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf GameServersModerators) DeepCopy() IDirectusObject {
	new_obj := &GameServersModerators{}
	if cf.DirectusUsersId != nil {
		new_obj.DirectusUsersId = (*cf.DirectusUsersId).DeepCopy().(*DirectusUsers)
	}
	if cf.GameServersId != nil {
		new_obj.GameServersId = (*cf.GameServersId).DeepCopy().(*GameServers)
	}
	new_obj.Id = cf.Id
	return new_obj
}
func (cf GameServersModerators) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Id != old.(*GameServersModerators).Id {
		diff["id"] = cf.Id
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf GameServersModerators) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.DirectusUsersId != nil {
		trakingList = append(trakingList, cf.DirectusUsersId)
		trakingList = append(trakingList, cf.DirectusUsersId.Track()...)
	}
	if cf.GameServersId != nil {
		trakingList = append(trakingList, cf.GameServersId)
		trakingList = append(trakingList, cf.GameServersId.Track()...)
	}

	return trakingList
}

func (cf GameServersModerators) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf GameServersModerators) CollectionName() string {
	return "GameServers_Moderators"
}

type KnownIps struct {
	IDirectusObject
	DateCreated *time.Time   `json:"date_created"`
	Id          uuid.UUID    `json:"id"`
	Ip          string       `json:"ip"`
	Location    string       `json:"location"`
	SteamUsers  []SteamUsers `json:"SteamUsers"`
}

func (cf *KnownIps) UnmarshalJSON(data []byte) error {
	type knownips_internal struct {
		DateCreated *time.Time   `json:"date_created"`
		Id          uuid.UUID    `json:"id"`
		Ip          string       `json:"ip"`
		Location    string       `json:"location"`
		SteamUsers  []SteamUsers `json:"SteamUsers"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj knownips_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DateCreated = _obj.DateCreated
		cf.Id = _obj.Id
		cf.Ip = _obj.Ip
		cf.Location = _obj.Location
		cf.SteamUsers = _obj.SteamUsers
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf KnownIps) DeepCopy() IDirectusObject {
	new_obj := &KnownIps{}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Id = cf.Id
	new_obj.Ip = cf.Ip
	new_obj.Location = cf.Location
	if cf.SteamUsers != nil {
		new_obj.SteamUsers = make([]SteamUsers, len(cf.SteamUsers))
		copy(new_obj.SteamUsers, cf.SteamUsers)
	}
	return new_obj
}
func (cf KnownIps) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*KnownIps).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*KnownIps).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*KnownIps).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*KnownIps).Id {
		diff["id"] = cf.Id
	}

	if cf.Ip != old.(*KnownIps).Ip {
		diff["ip"] = cf.Ip
	}

	if cf.Location != old.(*KnownIps).Location {
		diff["location"] = cf.Location
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf KnownIps) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.SteamUsers != nil {
		for _, iter := range cf.SteamUsers {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	return trakingList
}

func (cf KnownIps) GetId() string {
	return cf.Id.String()
}

func (cf KnownIps) CollectionName() string {
	return "KnownIps"
}

type PlayerBans struct {
	IDirectusObject
	BannedTill   time.Time                 `json:"banned_till"`
	DateCreated  *time.Time                `json:"date_created"`
	Files        []PlayerBansExternalFiles `json:"Files"`
	Id           uuid.UUID                 `json:"id"`
	Project      *GameProjects             `json:"project"`
	ProjectBan   bool                      `json:"project_ban"`
	Reason       *string                   `json:"reason"`
	TargetIp     *KnownIps                 `json:"target_ip"`
	TargetPlayer *ServerPlayers            `json:"target_player"`
	TargetSteam  *SteamUsers               `json:"target_steam"`
	Unbanned     bool                      `json:"unbanned"`
	UserCreated  *DirectusUsers            `json:"user_created"`
}

func (cf *PlayerBans) UnmarshalJSON(data []byte) error {
	type playerbans_internal struct {
		BannedTill   time.Time                 `json:"banned_till"`
		DateCreated  *time.Time                `json:"date_created"`
		Files        []PlayerBansExternalFiles `json:"Files"`
		Id           uuid.UUID                 `json:"id"`
		Project      *GameProjects             `json:"project"`
		ProjectBan   bool                      `json:"project_ban"`
		Reason       *string                   `json:"reason"`
		TargetIp     *KnownIps                 `json:"target_ip"`
		TargetPlayer *ServerPlayers            `json:"target_player"`
		TargetSteam  *SteamUsers               `json:"target_steam"`
		Unbanned     bool                      `json:"unbanned"`
		UserCreated  *DirectusUsers            `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj playerbans_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.BannedTill = _obj.BannedTill
		cf.DateCreated = _obj.DateCreated
		cf.Files = _obj.Files
		cf.Id = _obj.Id
		cf.Project = _obj.Project
		cf.ProjectBan = _obj.ProjectBan
		cf.Reason = _obj.Reason
		cf.TargetIp = _obj.TargetIp
		cf.TargetPlayer = _obj.TargetPlayer
		cf.TargetSteam = _obj.TargetSteam
		cf.Unbanned = _obj.Unbanned
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf PlayerBans) DeepCopy() IDirectusObject {
	new_obj := &PlayerBans{}
	new_obj.BannedTill = cf.BannedTill
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.Files != nil {
		new_obj.Files = make([]PlayerBansExternalFiles, len(cf.Files))
		copy(new_obj.Files, cf.Files)
	}
	new_obj.Id = cf.Id
	if cf.Project != nil {
		new_obj.Project = (*cf.Project).DeepCopy().(*GameProjects)
	}
	new_obj.ProjectBan = cf.ProjectBan
	if cf.Reason != nil {
		temp := ""
		new_obj.Reason = &temp
		*new_obj.Reason = *cf.Reason
	}
	if cf.TargetIp != nil {
		new_obj.TargetIp = (*cf.TargetIp).DeepCopy().(*KnownIps)
	}
	if cf.TargetPlayer != nil {
		new_obj.TargetPlayer = (*cf.TargetPlayer).DeepCopy().(*ServerPlayers)
	}
	if cf.TargetSteam != nil {
		new_obj.TargetSteam = (*cf.TargetSteam).DeepCopy().(*SteamUsers)
	}
	new_obj.Unbanned = cf.Unbanned
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf PlayerBans) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.BannedTill != old.(*PlayerBans).BannedTill {
		diff["banned_till"] = cf.BannedTill
	}
	if cf.DateCreated == nil {
		if old.(*PlayerBans).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*PlayerBans).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*PlayerBans).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*PlayerBans).Id {
		diff["id"] = cf.Id
	}

	if cf.ProjectBan != old.(*PlayerBans).ProjectBan {
		diff["project_ban"] = cf.ProjectBan
	}
	if cf.Reason == nil {
		if old.(*PlayerBans).Reason != nil {
			diff["reason"] = nil
		}
	} else {
		if old.(*PlayerBans).Reason == nil {
			diff["reason"] = cf.Reason
		} else {
			if *cf.Reason != *old.(*PlayerBans).Reason {
				diff["reason"] = cf.Reason
			}
		}
	}

	if cf.Unbanned != old.(*PlayerBans).Unbanned {
		diff["unbanned"] = cf.Unbanned
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf PlayerBans) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Files != nil {
		for _, iter := range cf.Files {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.Project != nil {
		trakingList = append(trakingList, cf.Project)
		trakingList = append(trakingList, cf.Project.Track()...)
	}

	if cf.TargetIp != nil {
		trakingList = append(trakingList, cf.TargetIp)
		trakingList = append(trakingList, cf.TargetIp.Track()...)
	}
	if cf.TargetPlayer != nil {
		trakingList = append(trakingList, cf.TargetPlayer)
		trakingList = append(trakingList, cf.TargetPlayer.Track()...)
	}
	if cf.TargetSteam != nil {
		trakingList = append(trakingList, cf.TargetSteam)
		trakingList = append(trakingList, cf.TargetSteam.Track()...)
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf PlayerBans) GetId() string {
	return cf.Id.String()
}

func (cf PlayerBans) CollectionName() string {
	return "PlayerBans"
}

type PlayerBansExternalFiles struct {
	IDirectusObject
	ExternalFilesId *ExternalFiles `json:"ExternalFiles_id"`
	Id              int            `json:"id"`
	PlayerBansId    *PlayerBans    `json:"PlayerBans_id"`
}

func (cf *PlayerBansExternalFiles) UnmarshalJSON(data []byte) error {
	type playerbansexternalfiles_internal struct {
		ExternalFilesId *ExternalFiles `json:"ExternalFiles_id"`
		Id              int            `json:"id"`
		PlayerBansId    *PlayerBans    `json:"PlayerBans_id"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj playerbansexternalfiles_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.ExternalFilesId = _obj.ExternalFilesId
		cf.Id = _obj.Id
		cf.PlayerBansId = _obj.PlayerBansId
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf PlayerBansExternalFiles) DeepCopy() IDirectusObject {
	new_obj := &PlayerBansExternalFiles{}
	if cf.ExternalFilesId != nil {
		new_obj.ExternalFilesId = (*cf.ExternalFilesId).DeepCopy().(*ExternalFiles)
	}
	new_obj.Id = cf.Id
	if cf.PlayerBansId != nil {
		new_obj.PlayerBansId = (*cf.PlayerBansId).DeepCopy().(*PlayerBans)
	}
	return new_obj
}
func (cf PlayerBansExternalFiles) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Id != old.(*PlayerBansExternalFiles).Id {
		diff["id"] = cf.Id
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf PlayerBansExternalFiles) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.ExternalFilesId != nil {
		trakingList = append(trakingList, cf.ExternalFilesId)
		trakingList = append(trakingList, cf.ExternalFilesId.Track()...)
	}

	if cf.PlayerBansId != nil {
		trakingList = append(trakingList, cf.PlayerBansId)
		trakingList = append(trakingList, cf.PlayerBansId.Track()...)
	}
	return trakingList
}

func (cf PlayerBansExternalFiles) GetId() string {
	return fmt.Sprintf("%d", cf.Id)
}

func (cf PlayerBansExternalFiles) CollectionName() string {
	return "PlayerBans_ExternalFiles"
}

type PlayerChecks struct {
	IDirectusObject
	DateCreated *time.Time     `json:"date_created"`
	Id          uuid.UUID      `json:"id"`
	Inspector   *DirectusUsers `json:"inspector"`
	Target      *ServerPlayers `json:"target"`
	UserCreated *DirectusUsers `json:"user_created"`
}

func (cf *PlayerChecks) UnmarshalJSON(data []byte) error {
	type playerchecks_internal struct {
		DateCreated *time.Time     `json:"date_created"`
		Id          uuid.UUID      `json:"id"`
		Inspector   *DirectusUsers `json:"inspector"`
		Target      *ServerPlayers `json:"target"`
		UserCreated *DirectusUsers `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj playerchecks_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DateCreated = _obj.DateCreated
		cf.Id = _obj.Id
		cf.Inspector = _obj.Inspector
		cf.Target = _obj.Target
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf PlayerChecks) DeepCopy() IDirectusObject {
	new_obj := &PlayerChecks{}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Id = cf.Id
	if cf.Inspector != nil {
		new_obj.Inspector = (*cf.Inspector).DeepCopy().(*DirectusUsers)
	}
	if cf.Target != nil {
		new_obj.Target = (*cf.Target).DeepCopy().(*ServerPlayers)
	}
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf PlayerChecks) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*PlayerChecks).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*PlayerChecks).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*PlayerChecks).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*PlayerChecks).Id {
		diff["id"] = cf.Id
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf PlayerChecks) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Inspector != nil {
		trakingList = append(trakingList, cf.Inspector)
		trakingList = append(trakingList, cf.Inspector.Track()...)
	}
	if cf.Target != nil {
		trakingList = append(trakingList, cf.Target)
		trakingList = append(trakingList, cf.Target.Track()...)
	}
	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf PlayerChecks) GetId() string {
	return cf.Id.String()
}

func (cf PlayerChecks) CollectionName() string {
	return "PlayerChecks"
}

type PlayerEvents struct {
	IDirectusObject
	DateCreated *time.Time     `json:"date_created"`
	Event       any            `json:"event"`
	Id          uuid.UUID      `json:"id"`
	Player      *ServerPlayers `json:"player"`
}

func (cf *PlayerEvents) UnmarshalJSON(data []byte) error {
	type playerevents_internal struct {
		DateCreated *time.Time     `json:"date_created"`
		Event       any            `json:"event"`
		Id          uuid.UUID      `json:"id"`
		Player      *ServerPlayers `json:"player"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj playerevents_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DateCreated = _obj.DateCreated
		cf.Event = _obj.Event
		cf.Id = _obj.Id
		cf.Player = _obj.Player
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf PlayerEvents) DeepCopy() IDirectusObject {
	new_obj := &PlayerEvents{}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Event = cf.Event
	new_obj.Id = cf.Id
	if cf.Player != nil {
		new_obj.Player = (*cf.Player).DeepCopy().(*ServerPlayers)
	}
	return new_obj
}
func (cf PlayerEvents) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*PlayerEvents).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*PlayerEvents).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*PlayerEvents).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Event != old.(*PlayerEvents).Event {
		diff["event"] = cf.Event
	}

	if cf.Id != old.(*PlayerEvents).Id {
		diff["id"] = cf.Id
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf PlayerEvents) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Player != nil {
		trakingList = append(trakingList, cf.Player)
		trakingList = append(trakingList, cf.Player.Track()...)
	}
	return trakingList
}

func (cf PlayerEvents) GetId() string {
	return cf.Id.String()
}

func (cf PlayerEvents) CollectionName() string {
	return "PlayerEvents"
}

type PlayerReports struct {
	IDirectusObject
	DateCreated *time.Time     `json:"date_created"`
	Id          uuid.UUID      `json:"id"`
	Reason      *string        `json:"reason"`
	Reporter    *ServerPlayers `json:"reporter"`
	Target      *ServerPlayers `json:"target"`
}

func (cf *PlayerReports) UnmarshalJSON(data []byte) error {
	type playerreports_internal struct {
		DateCreated *time.Time     `json:"date_created"`
		Id          uuid.UUID      `json:"id"`
		Reason      *string        `json:"reason"`
		Reporter    *ServerPlayers `json:"reporter"`
		Target      *ServerPlayers `json:"target"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj playerreports_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DateCreated = _obj.DateCreated
		cf.Id = _obj.Id
		cf.Reason = _obj.Reason
		cf.Reporter = _obj.Reporter
		cf.Target = _obj.Target
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf PlayerReports) DeepCopy() IDirectusObject {
	new_obj := &PlayerReports{}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Id = cf.Id
	if cf.Reason != nil {
		temp := ""
		new_obj.Reason = &temp
		*new_obj.Reason = *cf.Reason
	}
	if cf.Reporter != nil {
		new_obj.Reporter = (*cf.Reporter).DeepCopy().(*ServerPlayers)
	}
	if cf.Target != nil {
		new_obj.Target = (*cf.Target).DeepCopy().(*ServerPlayers)
	}
	return new_obj
}
func (cf PlayerReports) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*PlayerReports).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*PlayerReports).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*PlayerReports).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*PlayerReports).Id {
		diff["id"] = cf.Id
	}
	if cf.Reason == nil {
		if old.(*PlayerReports).Reason != nil {
			diff["reason"] = nil
		}
	} else {
		if old.(*PlayerReports).Reason == nil {
			diff["reason"] = cf.Reason
		} else {
			if *cf.Reason != *old.(*PlayerReports).Reason {
				diff["reason"] = cf.Reason
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf PlayerReports) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Reporter != nil {
		trakingList = append(trakingList, cf.Reporter)
		trakingList = append(trakingList, cf.Reporter.Track()...)
	}
	if cf.Target != nil {
		trakingList = append(trakingList, cf.Target)
		trakingList = append(trakingList, cf.Target.Track()...)
	}
	return trakingList
}

func (cf PlayerReports) GetId() string {
	return cf.Id.String()
}

func (cf PlayerReports) CollectionName() string {
	return "PlayerReports"
}

type ServerChatMessages struct {
	IDirectusObject
	Content     string         `json:"content"`
	DateCreated *time.Time     `json:"date_created"`
	Id          uuid.UUID      `json:"id"`
	Prefix      *string        `json:"prefix"`
	Sender      *ServerPlayers `json:"sender"`
	Server      *GameServers   `json:"server"`
}

func (cf *ServerChatMessages) UnmarshalJSON(data []byte) error {
	type serverchatmessages_internal struct {
		Content     string         `json:"content"`
		DateCreated *time.Time     `json:"date_created"`
		Id          uuid.UUID      `json:"id"`
		Prefix      *string        `json:"prefix"`
		Sender      *ServerPlayers `json:"sender"`
		Server      *GameServers   `json:"server"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj serverchatmessages_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Content = _obj.Content
		cf.DateCreated = _obj.DateCreated
		cf.Id = _obj.Id
		cf.Prefix = _obj.Prefix
		cf.Sender = _obj.Sender
		cf.Server = _obj.Server
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf ServerChatMessages) DeepCopy() IDirectusObject {
	new_obj := &ServerChatMessages{}
	new_obj.Content = cf.Content
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Id = cf.Id
	if cf.Prefix != nil {
		temp := ""
		new_obj.Prefix = &temp
		*new_obj.Prefix = *cf.Prefix
	}
	if cf.Sender != nil {
		new_obj.Sender = (*cf.Sender).DeepCopy().(*ServerPlayers)
	}
	if cf.Server != nil {
		new_obj.Server = (*cf.Server).DeepCopy().(*GameServers)
	}
	return new_obj
}
func (cf ServerChatMessages) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Content != old.(*ServerChatMessages).Content {
		diff["content"] = cf.Content
	}
	if cf.DateCreated == nil {
		if old.(*ServerChatMessages).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*ServerChatMessages).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*ServerChatMessages).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*ServerChatMessages).Id {
		diff["id"] = cf.Id
	}
	if cf.Prefix == nil {
		if old.(*ServerChatMessages).Prefix != nil {
			diff["prefix"] = nil
		}
	} else {
		if old.(*ServerChatMessages).Prefix == nil {
			diff["prefix"] = cf.Prefix
		} else {
			if *cf.Prefix != *old.(*ServerChatMessages).Prefix {
				diff["prefix"] = cf.Prefix
			}
		}
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf ServerChatMessages) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Sender != nil {
		trakingList = append(trakingList, cf.Sender)
		trakingList = append(trakingList, cf.Sender.Track()...)
	}
	if cf.Server != nil {
		trakingList = append(trakingList, cf.Server)
		trakingList = append(trakingList, cf.Server.Track()...)
	}
	return trakingList
}

func (cf ServerChatMessages) GetId() string {
	return cf.Id.String()
}

func (cf ServerChatMessages) CollectionName() string {
	return "ServerChatMessages"
}

type ServerPlayers struct {
	IDirectusObject
	DateCreated        *time.Time           `json:"date_created"`
	Id                 uuid.UUID            `json:"id"`
	Online             bool                 `json:"online"`
	PlayerChecks       []PlayerChecks       `json:"PlayerChecks"`
	PlayerEvents       []PlayerEvents       `json:"PlayerEvents"`
	Reports            []PlayerReports      `json:"reports"`
	SentReports        []PlayerReports      `json:"sent_reports"`
	Server             *GameServers         `json:"server"`
	ServerChatMessages []ServerChatMessages `json:"ServerChatMessages"`
	Steamuser          *SteamUsers          `json:"steamuser"`
}

func (cf *ServerPlayers) UnmarshalJSON(data []byte) error {
	type serverplayers_internal struct {
		DateCreated        *time.Time           `json:"date_created"`
		Id                 uuid.UUID            `json:"id"`
		Online             bool                 `json:"online"`
		PlayerChecks       []PlayerChecks       `json:"PlayerChecks"`
		PlayerEvents       []PlayerEvents       `json:"PlayerEvents"`
		Reports            []PlayerReports      `json:"reports"`
		SentReports        []PlayerReports      `json:"sent_reports"`
		Server             *GameServers         `json:"server"`
		ServerChatMessages []ServerChatMessages `json:"ServerChatMessages"`
		Steamuser          *SteamUsers          `json:"steamuser"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj serverplayers_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.DateCreated = _obj.DateCreated
		cf.Id = _obj.Id
		cf.Online = _obj.Online
		cf.PlayerChecks = _obj.PlayerChecks
		cf.PlayerEvents = _obj.PlayerEvents
		cf.Reports = _obj.Reports
		cf.SentReports = _obj.SentReports
		cf.Server = _obj.Server
		cf.ServerChatMessages = _obj.ServerChatMessages
		cf.Steamuser = _obj.Steamuser
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf ServerPlayers) DeepCopy() IDirectusObject {
	new_obj := &ServerPlayers{}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	new_obj.Id = cf.Id
	new_obj.Online = cf.Online
	if cf.PlayerChecks != nil {
		new_obj.PlayerChecks = make([]PlayerChecks, len(cf.PlayerChecks))
		copy(new_obj.PlayerChecks, cf.PlayerChecks)
	}
	if cf.PlayerEvents != nil {
		new_obj.PlayerEvents = make([]PlayerEvents, len(cf.PlayerEvents))
		copy(new_obj.PlayerEvents, cf.PlayerEvents)
	}
	if cf.Reports != nil {
		new_obj.Reports = make([]PlayerReports, len(cf.Reports))
		copy(new_obj.Reports, cf.Reports)
	}
	if cf.SentReports != nil {
		new_obj.SentReports = make([]PlayerReports, len(cf.SentReports))
		copy(new_obj.SentReports, cf.SentReports)
	}
	if cf.Server != nil {
		new_obj.Server = (*cf.Server).DeepCopy().(*GameServers)
	}
	if cf.ServerChatMessages != nil {
		new_obj.ServerChatMessages = make([]ServerChatMessages, len(cf.ServerChatMessages))
		copy(new_obj.ServerChatMessages, cf.ServerChatMessages)
	}
	if cf.Steamuser != nil {
		new_obj.Steamuser = (*cf.Steamuser).DeepCopy().(*SteamUsers)
	}
	return new_obj
}
func (cf ServerPlayers) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*ServerPlayers).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*ServerPlayers).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*ServerPlayers).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}

	if cf.Id != old.(*ServerPlayers).Id {
		diff["id"] = cf.Id
	}

	if cf.Online != old.(*ServerPlayers).Online {
		diff["online"] = cf.Online
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf ServerPlayers) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.PlayerChecks != nil {
		for _, iter := range cf.PlayerChecks {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.PlayerEvents != nil {
		for _, iter := range cf.PlayerEvents {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.Reports != nil {
		for _, iter := range cf.Reports {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.SentReports != nil {
		for _, iter := range cf.SentReports {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.Server != nil {
		trakingList = append(trakingList, cf.Server)
		trakingList = append(trakingList, cf.Server.Track()...)
	}
	if cf.ServerChatMessages != nil {
		for _, iter := range cf.ServerChatMessages {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.Steamuser != nil {
		trakingList = append(trakingList, cf.Steamuser)
		trakingList = append(trakingList, cf.Steamuser.Track()...)
	}
	return trakingList
}

func (cf ServerPlayers) GetId() string {
	return cf.Id.String()
}

func (cf ServerPlayers) CollectionName() string {
	return "ServerPlayers"
}

type SteamUsers struct {
	IDirectusObject
	Avatar        *ExternalFiles  `json:"avatar"`
	DateCreated   *time.Time      `json:"date_created"`
	DateUpdated   *time.Time      `json:"date_updated"`
	Id            uuid.UUID       `json:"id"`
	Ip            *KnownIps       `json:"ip"`
	Nickname      string          `json:"nickname"`
	ServerPlayers []ServerPlayers `json:"ServerPlayers"`
	Steamid       string          `json:"steamid"`
	UserCreated   *DirectusUsers  `json:"user_created"`
}

func (cf *SteamUsers) UnmarshalJSON(data []byte) error {
	type steamusers_internal struct {
		Avatar        *ExternalFiles  `json:"avatar"`
		DateCreated   *time.Time      `json:"date_created"`
		DateUpdated   *time.Time      `json:"date_updated"`
		Id            uuid.UUID       `json:"id"`
		Ip            *KnownIps       `json:"ip"`
		Nickname      string          `json:"nickname"`
		ServerPlayers []ServerPlayers `json:"ServerPlayers"`
		Steamid       string          `json:"steamid"`
		UserCreated   *DirectusUsers  `json:"user_created"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj steamusers_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Avatar = _obj.Avatar
		cf.DateCreated = _obj.DateCreated
		cf.DateUpdated = _obj.DateUpdated
		cf.Id = _obj.Id
		cf.Ip = _obj.Ip
		cf.Nickname = _obj.Nickname
		cf.ServerPlayers = _obj.ServerPlayers
		cf.Steamid = _obj.Steamid
		cf.UserCreated = _obj.UserCreated
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf SteamUsers) DeepCopy() IDirectusObject {
	new_obj := &SteamUsers{}
	if cf.Avatar != nil {
		new_obj.Avatar = (*cf.Avatar).DeepCopy().(*ExternalFiles)
	}
	if cf.DateCreated != nil {
		temp := time.Time{}
		new_obj.DateCreated = &temp
		*new_obj.DateCreated = *cf.DateCreated
	}
	if cf.DateUpdated != nil {
		temp := time.Time{}
		new_obj.DateUpdated = &temp
		*new_obj.DateUpdated = *cf.DateUpdated
	}
	new_obj.Id = cf.Id
	if cf.Ip != nil {
		new_obj.Ip = (*cf.Ip).DeepCopy().(*KnownIps)
	}
	new_obj.Nickname = cf.Nickname
	if cf.ServerPlayers != nil {
		new_obj.ServerPlayers = make([]ServerPlayers, len(cf.ServerPlayers))
		copy(new_obj.ServerPlayers, cf.ServerPlayers)
	}
	new_obj.Steamid = cf.Steamid
	if cf.UserCreated != nil {
		new_obj.UserCreated = (*cf.UserCreated).DeepCopy().(*DirectusUsers)
	}
	return new_obj
}
func (cf SteamUsers) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.DateCreated == nil {
		if old.(*SteamUsers).DateCreated != nil {
			diff["date_created"] = nil
		}
	} else {
		if old.(*SteamUsers).DateCreated == nil {
			diff["date_created"] = cf.DateCreated
		} else {
			if *cf.DateCreated != *old.(*SteamUsers).DateCreated {
				diff["date_created"] = cf.DateCreated
			}
		}
	}
	if cf.DateUpdated == nil {
		if old.(*SteamUsers).DateUpdated != nil {
			diff["date_updated"] = nil
		}
	} else {
		if old.(*SteamUsers).DateUpdated == nil {
			diff["date_updated"] = cf.DateUpdated
		} else {
			if *cf.DateUpdated != *old.(*SteamUsers).DateUpdated {
				diff["date_updated"] = cf.DateUpdated
			}
		}
	}

	if cf.Id != old.(*SteamUsers).Id {
		diff["id"] = cf.Id
	}

	if cf.Nickname != old.(*SteamUsers).Nickname {
		diff["nickname"] = cf.Nickname
	}

	if cf.Steamid != old.(*SteamUsers).Steamid {
		diff["steamid"] = cf.Steamid
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf SteamUsers) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	if cf.Avatar != nil {
		trakingList = append(trakingList, cf.Avatar)
		trakingList = append(trakingList, cf.Avatar.Track()...)
	}

	if cf.Ip != nil {
		trakingList = append(trakingList, cf.Ip)
		trakingList = append(trakingList, cf.Ip.Track()...)
	}

	if cf.ServerPlayers != nil {
		for _, iter := range cf.ServerPlayers {
			trakingList = append(trakingList, iter.Track()...)
		}
	}

	if cf.UserCreated != nil {
		trakingList = append(trakingList, cf.UserCreated)
		trakingList = append(trakingList, cf.UserCreated.Track()...)
	}
	return trakingList
}

func (cf SteamUsers) GetId() string {
	return cf.Id.String()
}

func (cf SteamUsers) CollectionName() string {
	return "SteamUsers"
}

type Test struct {
	IDirectusObject
	Id uuid.UUID `json:"id"`
}

func (cf *Test) UnmarshalJSON(data []byte) error {
	type test_internal struct {
		Id uuid.UUID `json:"id"`
	}
	if data[0] == '"' { //Data is a string
		return json.Unmarshal(data, &cf.Id)
	} else if data[0] == '{' { //Data is an object
		var _obj test_internal
		err := json.Unmarshal(data, &_obj)
		if err != nil {
			return err
		}
		cf.Id = _obj.Id
	} else {
		//Number or unkown, probably id
		return json.Unmarshal(data, &cf.Id)
	}
	return nil
}

func (cf Test) DeepCopy() IDirectusObject {
	new_obj := &Test{}
	new_obj.Id = cf.Id
	return new_obj
}
func (cf Test) Diff(old IDirectusObject) map[string]interface{} {
	diff := make(map[string]interface{})

	if cf.Id != old.(*Test).Id {
		diff["id"] = cf.Id
	}

	if len(diff) == 0 {
		return nil
	}
	return diff
}
func (cf Test) Track() []IDirectusObject {
	trakingList := make([]IDirectusObject, 0)

	return trakingList
}

func (cf Test) GetId() string {
	return cf.Id.String()
}

func (cf Test) CollectionName() string {
	return "test"
}
