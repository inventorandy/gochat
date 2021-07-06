import { Dispatch } from "redux";
import { appAPI } from "../apiConn";
import { ConverationWebsocketMessage, Conversation, ConversationActions, Message } from "../types/conversation";

/**
 * Get a list of public conversation channels
 */
export const GetPublicConversations = () => (dispatch: Dispatch) => {
  appAPI.get("/conversation?public=true", { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
    // Get the Conversation List from the Response
    let conversations: Conversation[] = res.data;

    // Do the Dispatch
    dispatch({
      type: ConversationActions.GET_PUBLIC_CONVERSATIONS,
	    conversations: conversations,
    })
  }).catch(error => {
    // Handle the Error Response
  })
}

/**
 * Get a list of private conversation channels
 */
export const GetPrivateConversations = () => (dispatch: Dispatch) => {
  appAPI.get("/conversations", { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
    // Get the Conversation List from the Response
    let conversations: Conversation[] = res.data;

    // Do the Dispatch
    dispatch({
      type: ConversationActions.GET_PRIVATE_CONVERSATIONS,
	    conversations: conversations,
    })
  }).catch(error => {
    // Handle the Error Response
  })
}

/**
 * Set the Current Conversation
 */
export const SetCurrentConversation = (conversation: Conversation) => (dispatch: Dispatch) => {
  // Do the Dispatch
  dispatch({
    type: ConversationActions.SET_CURRENT_CONVERSATION,
    conversation: conversation,
  })
}

/**
 * Send a Message
 */
export const SendMessage = (message: Message) => (dispatch: Dispatch) => {
  appAPI.post("/message", message, { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
  }).catch(error => {});
}

export const ConnectConversationWebsocket = () => (dispatch: Dispatch) => {
  // Create the WebSocket
  const socket = new WebSocket(`${process.env.REACT_APP_WS_URL}/conversations`);

  // Create the Event Listeners
  socket.onopen = event => {
    console.log("Conversation websocket opened, getting fresh data...");
  }
  socket.onmessage = event => {
    // Get the Event
    let wsEvent: ConverationWebsocketMessage = JSON.parse(event.data);
    switch (wsEvent.type) {
      case "MESSAGE":
        console.log("message received...");
        break;
      default:
        console.log("unknown packet type...");
    }
  }
  socket.onclose = event => {
    console.log("Conversation websocket closed...");
  }
  socket.onerror = error => {
    console.log("Socket Error: ", error);
  }
}