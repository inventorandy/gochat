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
}

// Get Public Conversations Action Type
interface GetPublicConversations {
  type: typeof ConversationActions.GET_PUBLIC_CONVERSATIONS,
  conversations: Conversation[],
}

// Export the Action Types
export type ConversationActionTypes = GetPublicConversations;