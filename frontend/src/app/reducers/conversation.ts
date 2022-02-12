import {
  Conversation,
  ConversationActions,
  ConversationActionTypes,
} from '../types/conversation';

// ConversationState
export interface ConversationState {
  publicConversations?: Conversation[];
  privateConversations?: Conversation[];
  currentConversation?: Conversation;
}

// Set the initial (empty) state
const initialState: ConversationState = {
  publicConversations: undefined,
  privateConversations: undefined,
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
      };
    case ConversationActions.CREATE_CONVERSATION:
      if (action.conversation.is_public) {
        let pubConvs = state.publicConversations || [];
        pubConvs.push(action.conversation);
        return {
          ...state,
          publicConversations: pubConvs,
        };
      } else {
        let priConvs = state.privateConversations || [];
        priConvs.push(action.conversation);
        return {
          ...state,
          privateConversations: priConvs,
        };
      }
    default:
      return state;
  }
};

export default conversationReducer;
