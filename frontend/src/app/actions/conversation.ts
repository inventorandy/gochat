import { Dispatch } from 'redux';
import { appAPI, headers } from '../apiConn';
import { getCookie } from '../helpers/cookie';
import { ConversationState } from '../reducers/conversation';
import { store } from '../store';
import {
  ConverationWebsocketMessage,
  Conversation,
  ConversationActions,
  Message,
} from '../types/conversation';
import { APIError } from '../types/generic';

// Global Variable for the Websocket Connection
var conversationSocket: WebSocket;

/**
 * Get a list of public conversations.
 * @param onSuccess
 * @param onError
 * @returns
 */
export const GetPublicConversations =
  (
    onSuccess: (conversations: Conversation[]) => void,
    onError: (err: APIError) => void
  ) =>
  async (dispatch: Dispatch) => {
    // Call the API
    await appAPI
      .get('/conversation?public=true', headers())
      .then((res) => {
        // Get the Conversation Info
        let conversations: Conversation[] = res.data;

        // Call the Dispatch
        dispatch({
          type: ConversationActions.GET_PUBLIC_CONVERSATIONS,
          conversations: conversations,
        });

        // Call the onSuccess Method
        onSuccess(conversations);
      })
      .catch((err) => {
        // Get the Error Object
        let error: APIError = err.response.data;

        // Handle Error
        onError(error);
      })
      .finally(() => {
        // Create an Error Object
        let error: APIError = {
          code: 500,
          message: 'Severe API Error.',
        };

        // Handle Error
        onError(error);
      });
  };

/**
 * Get a conversation by its ID.
 * @param id
 * @param onError
 * @returns
 */
export const GetConversation =
  (id: string, onError: (err: APIError) => void) =>
  async (dispatch: Dispatch) => {
    // Call the API
    await appAPI
      .get('/conversation/' + id, headers())
      .then((res) => {
        // Get the Conversation Info
        let conversation: Conversation = res.data;

        // Call the Dispatch
        dispatch({
          type: ConversationActions.SET_CURRENT_CONVERSATION,
          conversation: conversation,
        });
      })
      .catch((err) => {
        // Get the Error Object
        let error: APIError = err.response.data;

        // Handle Error
        onError(error);
      })
      .finally(() => {
        // Create an Error Object
        let error: APIError = {
          code: 500,
          message: 'Severe API Error.',
        };

        // Handle Error
        onError(error);
      });
  };

/**
 * Connects to the Conversation WebSocket endpoint.
 * @returns 
 */
export const ConnectConversationWebsocket =
  () => async (dispatch: Dispatch) => {
    // Create the WebSocket
    const token = getCookie('token');
    let conversationSocket = new WebSocket(
      `${process.env.REACT_APP_WS_URL}/conversations`
    );

    // Create the Event Listeners
    conversationSocket.onopen = (event) => {};
    conversationSocket.onmessage = (event) => {
      // Get the Event
      let wsEvent: ConverationWebsocketMessage = JSON.parse(event.data);
      switch (wsEvent.event_type) {
        case 'MESSAGE_CREATED':
          ProcessNewMessage(wsEvent.data as Message, dispatch);
          break;
        case 'CONVERSATION_CREATED':
          break;
        default:
      }
    };
    conversationSocket.onclose = (event) => {};
    conversationSocket.onerror = (error) => {};
  };

const ProcessNewMessage = (message: Message, dispatch: Dispatch) => {
  // Define the API State
  const convState: ConversationState = store.getState().conversationState;

  // Check if the message belongs to the current conversation
  if (message.conversation_id === convState.currentConversation?.id) {
    // Get the current conversation
    let conversation = convState.currentConversation;

    // Check if we have messages
    if (conversation.messages === null || conversation.messages.length === 0) {
      conversation.messages = [];
    }

    // Add the Message
    conversation.messages.push(message);

    // Do the Dispatch
    dispatch({
      type: ConversationActions.SET_CURRENT_CONVERSATION,
      conversation: conversation,
    });
  }
};

export const CloseConversationWebsocket = () => {
  conversationSocket.close();
};
