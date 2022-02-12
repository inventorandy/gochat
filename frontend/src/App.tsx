import React from 'react';
import { Suspense } from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route
} from 'react-router-dom';
import { PersistGate } from 'redux-persist/integration/react';
import { Provider } from 'react-redux';
import { store, persistor } from './app/store';
import './scss/gochat.scss';
import LoginPage from './ui/pages/Login/LoginPage';
import MainPage from './ui/pages/Main/MainPage';
import RequireAuth from './ui/components/Routing/RequireAuth';

function App() {
  // Render
  return(
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <Router>
          <Routes>
            <Route path="/welcome" element={<LoginPage />} />
            <Route path="/" element={
              <RequireAuth>
                <MainPage />
              </RequireAuth>
            } />
          </Routes>
        </Router>
      </PersistGate>
    </Provider>
  );
}

export default App;
