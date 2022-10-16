import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '../store';
import { User } from '../types/user';

// UserState interface
interface UserState {
  loggedInUser?: User;
}

// Define the Initial State
const initialState: UserState = {
  loggedInUser: undefined,
};

// Define the Methods for Managing UserState
export const UserSlice = createSlice({
  name: 'User',
  initialState,
  reducers: {
    setCurrentUser: (state, action: PayloadAction<User>) => {
      state.loggedInUser = action.payload;
    },
  },
});

// Export the Actions
export const { setCurrentUser } = UserSlice.actions;

// Export the Selectors
export const currentUser = (state: RootState) => state.user.loggedInUser;

export default UserSlice.reducer;
