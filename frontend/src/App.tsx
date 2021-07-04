import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom';
import './scss/gochat.scss';
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
        <Route exact path="/">
          <LoginPage />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
