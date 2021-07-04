import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import './scss/gochat.scss';
import PrivateRoute from './ui/components/PrivateRoute';
import ChatInterface from './ui/pages/ChatInterface';
import LoginPage from './ui/pages/Login';
import RegisterPage from './ui/pages/Register';

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/auth/login">
          <LoginPage />
        </Route>
        <Route path="/auth/create-account">
          <RegisterPage />
        </Route>
        <PrivateRoute path="/" component={ChatInterface}>
          <LoginPage />
        </PrivateRoute>
      </Switch>
    </Router>
  );
}

export default App;
