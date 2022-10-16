import { Route, Routes } from 'react-router-dom';
import RequireAuth from './ui/components/auth/RequireAuth';
import Dashboard from './ui/components/dashboard/Dashboard';
import ConversationPage from './ui/pages/Conversation';
import LoginPage from './ui/pages/LoginPage';

function App() {
  return (
    <Routes>
      <Route path='/auth/login' element={<LoginPage />} />
      <Route path='/' element={<Dashboard />}>
        <Route
          index
          element={
            <RequireAuth>
              <h1>Hello, World!</h1>
            </RequireAuth>
          }
        />
        <Route
          path='/channels/:id'
          element={
            <RequireAuth>
              <ConversationPage />
            </RequireAuth>
          }
        />
      </Route>
    </Routes>
  );
}

export default App;
