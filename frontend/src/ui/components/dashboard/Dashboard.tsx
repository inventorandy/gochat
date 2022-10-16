import * as React from 'react';
import { useEffect } from 'react';
import { Outlet, useNavigate } from 'react-router';
import { ConnectConversationWebsocket } from '../../../app/actions/conversation';
import { clearSession } from '../../../app/apiConn';
import { useAppDispatch } from '../../../app/hooks';
import '../../scss/dashboard.scss';
import gopher from '../../scss/images/gopher.png';
import PrivateConversationList from './PrivateConversationList';
import PublicConversationList from './PublicConversationList';

const Dashboard: React.FC = () => {
  const dispatch = useAppDispatch();

  // Set the Browser History
  const navigate = useNavigate();

  // useEffect Hook
  useEffect(() => {
    dispatch(ConnectConversationWebsocket());
  }, [dispatch]);

  // Log Out Handler
  const onLogout = (e: React.MouseEvent<HTMLButtonElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Clear the Session
    clearSession();

    // Navigate to the Dashboard
    navigate('/auth/login', { replace: true });
  };

  // Render
  return (
    <div className='dashboard'>
      <div className='sidebar'>
        <img src={gopher} alt='GoChat Gopher' />
        <PublicConversationList />
        <PrivateConversationList />
        <button className='logout-button' onClick={onLogout}>
          Log Out
        </button>
      </div>
      <Outlet />
    </div>
  );
};

export default Dashboard;
