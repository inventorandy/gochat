import { combineReducers } from "redux";
import conversationReducer from "./reducers/conversation";
import userReducer from "./reducers/user";

export const rootReducer = combineReducers({
  userState: userReducer,
  conversationState: conversationReducer
});

// Export the App State of type rootReducer
export type AppState = ReturnType<typeof rootReducer>;