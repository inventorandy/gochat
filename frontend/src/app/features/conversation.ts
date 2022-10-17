import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '../store';
import { Conversation, Message } from '../types/conversation';

// ConversationState interface
interface ConversationState {
  currentConversation?: Conversation;
  publicConversations?: Conversation[];
  privateConversations?: Conversation[];
}

// Define the Initial State
const initialState: ConversationState = {
  currentConversation: undefined,
  publicConversations: undefined,
  privateConversations: undefined,
};

// Define the Methods for Managing ConversationState
export const ConversationSlice = createSlice({
  name: 'Conversation',
  initialState,
  reducers: {
    setCurrentConversation: (state, action: PayloadAction<Conversation>) => {
      state.currentConversation = action.payload;
    },
    setPublicConversations: (state, action: PayloadAction<Conversation[]>) => {
      state.publicConversations = action.payload;
    },
    setPrivateConversations: (state, action: PayloadAction<Conversation[]>) => {
      state.privateConversations = action.payload;
    },
    addConversation: (state, action: PayloadAction<Conversation>) => {
      if (action.payload.is_public) {
        state.publicConversations?.push(action.payload);
      } else {
        state.privateConversations?.push(action.payload);
      }
    },
    addMessageToConversation: (state, action: PayloadAction<Message>) => {
      console.log(action.payload);
      let msg = action.payload;
      // Check if the message belongs to the current conversation
      if (msg.conversation_id === state.currentConversation?.id) {
        let conv = state.currentConversation;
        if (conv) {
          if (conv.messages === undefined || conv.messages === null) {
            conv.messages = [];
          }
          conv.messages.push(msg);
          state.currentConversation = conv;
        }
      }
    },
  },
});

// Export the Actions
export const {
  setCurrentConversation,
  setPublicConversations,
  setPrivateConversations,
  addConversation,
  addMessageToConversation,
} = ConversationSlice.actions;

// Export the Selectors
export const currentConversation = (state: RootState) =>
  state.conversation.currentConversation;
export const publicConversations = (state: RootState) =>
  state.conversation.publicConversations;
export const privateConversations = (state: RootState) =>
  state.conversation.privateConversations;

export default ConversationSlice.reducer;
