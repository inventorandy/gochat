import { Conversation, ConversationActions, ConversationActionTypes } from '../types/conversation';

// ConversationState
interface ConversationState {
  publicConversations?: Conversation[],
  error?: string | null
}

// Set the initial (empty) state
const initialState: ConversationState = {
  publicConversations: undefined,
  error: null
}

// Conversation Reducer
const conversationReducer = (state: ConversationState = initialState, action: ConversationActionTypes): ConversationState => {
  switch(action.type) {
    case ConversationActions.GET_PUBLIC_CONVERSATIONS:
      return {
        publicConversations: action.conversations,
      }
    default:
      return state;
  }
}

export default conversationReducer;