import { combineReducers } from 'redux';
import conversationReducer from './reducers/conversation';
import userReducer from './reducers/user';

// Combine all of the reducers into a single root reducer
export const rootReducer = combineReducers({
  userState: userReducer,
  conversationState: conversationReducer,
});

// Export the App State of the type rootReducer
export type AppState = ReturnType<typeof rootReducer>;
