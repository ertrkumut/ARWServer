package main

type ARWEvents struct {
	Connection        ARWEvent
	Login             ARWEvent
	Room_Create       ARWEvent
	Join_Room         ARWEvent
	Join_Any_Room     ARWEvent
	User_Enter_Room   ARWEvent
	User_Exit_Room    ARWEvent
	Extension_Request ARWEvent
	allEvents         []ARWEvent
}

const (
	Exception_Error    = "EXCEPTION_ERROR"
	Connection_Success = "CONNECTION_SUCCESS"
	Login              = "LOGIN"
	Login_Error        = "LOGIN_ERROR"
	Room_Create        = "ROOM_CREATE"
	Join_Room          = "ROOM_JOIN"
	Join_Any_Room      = "ANY_ROOM_JOIN"
	User_Enter_Room    = "USER_ENTER_ROOM"
	User_Exit_Room     = "USER_EXIT_ROOM"
	Extension_Response = "EXTENTION_REQUEST"
)

func (events *ARWEvents) Initialize() {

	events.Connection.eventName = Connection_Success
	events.Connection.Private_Handler = P_ConnectionSuccess

	events.Login.eventName = Login
	events.Login.Private_Handler = P_LoginEvent

	events.Room_Create.eventName = Room_Create

	events.Join_Room.eventName = Join_Room

	events.User_Enter_Room.eventName = User_Enter_Room

	events.User_Exit_Room.eventName = User_Exit_Room

	events.Join_Any_Room.eventName = Join_Any_Room
	events.Join_Any_Room.Private_Handler = P_JoinAnyRoom

	events.Extension_Request.Private_Handler = P_ExtensionResponse
	events.Extension_Request.eventName = Extension_Response

	events.allEvents = make([]ARWEvent, 0, 10)
	events.allEvents = append(events.allEvents, events.Connection)
	events.allEvents = append(events.allEvents, events.Login)
	events.allEvents = append(events.allEvents, events.Room_Create)
	events.allEvents = append(events.allEvents, events.Join_Room)
	events.allEvents = append(events.allEvents, events.Join_Any_Room)
	events.allEvents = append(events.allEvents, events.User_Enter_Room)
	events.allEvents = append(events.allEvents, events.User_Exit_Room)
	events.allEvents = append(events.allEvents, events.Extension_Request)
}
