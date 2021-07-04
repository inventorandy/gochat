import { combineReducers } from "redux";
import userReducer from "./reducers/user";

export const rootReducer = combineReducers({
  userState: userReducer
});

// Export the App State of type rootReducer
export type AppState = ReturnType<typeof rootReducer>;