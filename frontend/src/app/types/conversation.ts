import { User } from "./user";

// Messages API Object
export interface Message {
  id?: string;
  conversation_id?: string;
  author_id?: string;
  message?: string;
}

// Conversation API Object
export interface Conversation {
  id?: string;
  last_message_on?: string;
  label?: string;
  is_public?: boolean;
  participants?: User[];
  messages?: Message[];
}

export enum ConversationActions {
  GET_PUBLIC_CONVERSATIONS = 1,
  GET_PRIVATE_CONVERSATIONS = 2,
  SET_CURRENT_CONVERSATION = 3
}

// Get Public Conversations Action Type
interface GetPublicConversations {
  type: typeof ConversationActions.GET_PUBLIC_CONVERSATIONS,
  conversations: Conversation[],
}

// Get Private Conversations Action Type
interface GetPrivateConversations {
  type: typeof ConversationActions.GET_PRIVATE_CONVERSATIONS,
  conversations: Conversation[],
}

// Set Current Conversation Action Type
interface SetCurrentConversation {
  type: typeof ConversationActions.SET_CURRENT_CONVERSATION,
  conversation: Conversation,
}

// Export the Action Types
export type ConversationActionTypes = GetPublicConversations | GetPrivateConversations | SetCurrentConversation;