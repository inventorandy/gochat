import { Action, ThunkAction } from '@reduxjs/toolkit';
import { applyMiddleware, createStore } from 'redux';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import thunk from 'redux-thunk';
import { rootReducer } from './rootReducer';

// Persistent Store Configuration
export const persistConfig = {
  key: 'root',
  storage,
};

// Create the Persisted Reducer
const persistedReducer = persistReducer<any, any>(persistConfig, rootReducer);

// Export the Store
export const store = createStore(persistedReducer, applyMiddleware(thunk));
export const persistor = persistStore(store);

export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
