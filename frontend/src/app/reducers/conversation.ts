import {
  Conversation,
  ConversationActions,
  ConversationActionTypes,
} from '../types/conversation';

// ConversationState
export interface ConversationState {
  publicConversations?: Conversation[];
  currentConversation?: Conversation;
}

// Set the initial (empty) state
const initialState: ConversationState = {
  publicConversations: undefined,
  currentConversation: undefined,
};

// Conversation Reducer
const conversationReducer = (
  state: ConversationState = initialState,
  action: ConversationActionTypes
) => {
  switch (action.type) {
    case ConversationActions.GET_PUBLIC_CONVERSATIONS:
      return {
        ...state,
        publicConversations: action.conversations,
      };
    case ConversationActions.SET_CURRENT_CONVERSATION:
      return {
        ...state,
        currentConversation: action.conversation,
      }
    default:
      return state;
  }
};

export default conversationReducer;
