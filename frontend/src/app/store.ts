import { ThunkAction, Action, createStore, applyMiddleware } from '@reduxjs/toolkit';
import { rootReducer } from './rootReducer';
import thunk from 'redux-thunk';

export const store = createStore(rootReducer, applyMiddleware(thunk));

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;