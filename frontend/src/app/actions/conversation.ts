import { Dispatch } from 'redux';
import { appAPI, getHeaderConfig } from '../apiConn';
import {
  addMessageToConversation,
  setCurrentConversation,
  setPrivateConversations,
  setPublicConversations,
} from '../features/conversation';
import {
  ConverationWebsocketMessage,
  Conversation,
  Message,
} from '../types/conversation';
import { APIError } from '../types/error';

var conversationSocket: WebSocket;

export const SetCurrentConversation =
  (id: string, onError: (err: APIError) => void) =>
  async (dispatch: Dispatch) => {
    appAPI
      .get('/conversation/' + id, getHeaderConfig())
      .then((res) => {
        // Get the Conversation from the Response
        const conv: Conversation = res.data;

        // Set the Current Conversation
        dispatch(setCurrentConversation(conv));
      })
      .catch((err) => {
        // Get the Error Message
        let error: APIError = err.response.data;

        // Call the Error Method
        onError(error);
      });
  };

export const GetPublicConversations =
  (
    onSuccess: (conversations: Conversation[]) => void,
    onError: (err: APIError) => void
  ) =>
  async (dispatch: Dispatch) => {
    appAPI
      .get('/conversation?public=true', getHeaderConfig())
      .then((res) => {
        // Get the Conversations from the Response
        const convos: Conversation[] = res.data;

        // Set the Public Conversations
        dispatch(setPublicConversations(convos));

        // Call the Success Method
        onSuccess(convos);
      })
      .catch((err) => {
        // Get the Error Message
        let error: APIError = err.response.data;

        // Call the Error Method
        onError(error);
      });
  };

export const GetPrivateConversations =
  (
    onSuccess: (conversations: Conversation[]) => void,
    onError: (err: APIError) => void
  ) =>
  async (dispatch: Dispatch) => {
    appAPI
      .get('/conversation', getHeaderConfig())
      .then((res) => {
        // Get the Conversations from the Response
        const convos: Conversation[] = res.data;

        // Set the Private Conversations
        dispatch(setPrivateConversations(convos));

        // Call the Success Method
        onSuccess(convos);
      })
      .catch((err) => {
        // Get the Error Message
        let error: APIError = err.response.data;

        // Call the Error Method
        onError(error);
      });
  };

export const SendMessage = (message: Message) => (dispatch: Dispatch) => {
  appAPI
    .post('/message', message, getHeaderConfig())
    .then((res) => {})
    .catch((error) => {});
};

export const ConnectConversationWebsocket = () => (dispatch: Dispatch) => {
  if (!conversationSocket) {
    // Create the WebSocket
    const token = localStorage.getItem('authToken');
    document.cookie = `authToken=${token}; path=/`;
    conversationSocket = new WebSocket(
      `${process.env.REACT_APP_WS_URL}/conversations`
    );

    // Create the Event Listeners
    conversationSocket.onopen = (event) => {
      console.log('connected websocket');
    };
    conversationSocket.onmessage = (event) => {
      // Get the Event
      let wsEvent: ConverationWebsocketMessage = JSON.parse(event.data);
      switch (wsEvent.event_type) {
        case 'MESSAGE_CREATED':
          ProcessNewMessage(wsEvent.data as Message, dispatch);
          break;
        case 'CONVERSATION_CREATED':
          //ProcessNewConversation(wsEvent.data as Conversation, dispatch);
          break;
        default:
      }
    };
    conversationSocket.onclose = (event) => {};
    conversationSocket.onerror = (error) => {};
  }
};

const ProcessNewMessage = (msg: Message, dispatch: Dispatch) => {
  console.log('received a new message');
  dispatch(addMessageToConversation(msg));
};
