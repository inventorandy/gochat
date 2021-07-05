import { Dispatch } from "redux";
import { appAPI } from "../apiConn";
import { Conversation, ConversationActions, Message } from "../types/conversation";

/**
 * Get a list of public conversation channels
 */
export const GetPublicConversations = () => async(dispatch: Dispatch) => {
  await appAPI.get("/conversation?public=true", { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
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
export const GetPrivateConversations = () => async(dispatch: Dispatch) => {
  await appAPI.get("/conversations", { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
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
export const SetCurrentConversation = (conversation: Conversation) => async(dispatch: Dispatch) => {
  // Do the Dispatch
  dispatch({
    type: ConversationActions.SET_CURRENT_CONVERSATION,
    conversation: conversation,
  })
}

/**
 * Send a Message
 */
export const SendMessage = (message: Message) => async(dispatch: Dispatch) => {
  await appAPI.post("/message", message, { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
  }).catch(error => {});
}