// MEP-1013
package spacemanagement

type RoomTypeEnum string

const (
	ROOM_LECTURE_ROOM RoomTypeEnum = "Lecture Room"
	ROOM_LAB_ROOM     RoomTypeEnum = "Lab Room"
	ROOM_MEETING_ROOM RoomTypeEnum = "Meeting Room"
)

func (r RoomTypeEnum) TypeRoomString() string {
	switch r {
	case ROOM_LECTURE_ROOM:
		return "Lecture"
	case ROOM_LAB_ROOM:
		return "Laboratory"
	case ROOM_MEETING_ROOM:
		return "Office"
	default:
		return string(r)
	}
}
