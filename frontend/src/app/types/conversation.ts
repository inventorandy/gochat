export interface Message {
  id: string;
  author_id: string;
  author_name: string;
  conversation_id: string;
  created_at: string;
  message: string;
}
export interface Conversation {
  id: string;
  last_message_on: string;
  label: string;
  is_public: boolean;
  messages: Message[];
  has_new: boolean;
}

// Conversation Websocket Message
export interface ConverationWebsocketMessage {
  type?: string;
  event_type?: string;
  data?: Conversation | Message;
}

export enum ConversationActions {
  GET_PUBLIC_CONVERSATIONS = 'GET_PUBLIC_CONVERSATIONS',
  SET_CURRENT_CONVERSATION = 'SET_CURRENT_CONVERSATION',
}

interface GetPublicConversationsAction {
  type: typeof ConversationActions.GET_PUBLIC_CONVERSATIONS;
  conversations: Conversation[];
}

interface SetCurrentConversationAction {
  type: typeof ConversationActions.SET_CURRENT_CONVERSATION;
  conversation: Conversation;
}

export type ConversationActionTypes =
  | GetPublicConversationsAction
  | SetCurrentConversationAction;
