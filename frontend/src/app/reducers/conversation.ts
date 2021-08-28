import { SetCurrentConversation } from '../actions/conversation';
import { Conversation, ConversationActions, ConversationActionTypes } from '../types/conversation';

// ConversationState
interface ConversationState {
  publicConversations?: Conversation[],
  privateConversations?: Conversation[],
  currentConversation?: Conversation,
  error?: string | null
}

// Set the initial (empty) state
const initialState: ConversationState = {
  publicConversations: undefined,
  privateConversations: undefined,
  currentConversation: undefined,
  error: null
}

// Conversation Reducer
const conversationReducer = (state: ConversationState = initialState, action: ConversationActionTypes): ConversationState => {
  switch(action.type) {
    case ConversationActions.GET_PUBLIC_CONVERSATIONS:
      return {
        ...state,
        publicConversations: action.conversations,
      }
    case ConversationActions.GET_PRIVATE_CONVERSATIONS:
      return {
        ...state,
        privateConversations: action.conversations,
      }
    case ConversationActions.SET_CURRENT_CONVERSATION:
      return {
        ...state,
        currentConversation: action.conversation,
      }
    case ConversationActions.CREATE_CONVERSATION:
      if (action.conversation.is_public === true) {
        let conversations = state.publicConversations;
        conversations?.push(action.conversation);
        return {
          ...state,
          publicConversations: conversations
        }
      } else {
        let conversations = state.privateConversations;
        conversations?.push(action.conversation);
        return {
          ...state,
          privateConversations: conversations
        }
      }
    default:
      return state;
  }
}

export default conversationReducer;