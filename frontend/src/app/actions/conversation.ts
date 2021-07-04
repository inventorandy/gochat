import { Dispatch } from "redux";
import { appAPI } from "../apiConn";
import { Conversation, ConversationActions } from "../types/conversation";

export const GetPublicConversations = () => async(dispatch: Dispatch) => {
  await appAPI.get("/conversation?public=true", { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
    // Get the Conversation List from the Response
    let conversations: Conversation[] = res.data;
    console.log(conversations);

    // Do the Dispatch
    dispatch({
      type: ConversationActions.GET_PUBLIC_CONVERSATIONS,
	    conversations: conversations,
    })
  }).catch(error => {
    // Handle the Error Response
  })
}