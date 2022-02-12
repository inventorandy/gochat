import * as React from 'react';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { ConnectConversationWebsocket } from '../../../app/actions/conversation';
import ChatInterface from './ChatInterface';
import SideNavigation from './SideNavigation';

const MainPage: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // useEffect Hook handles state changes
  useEffect(() => {
    dispatch(ConnectConversationWebsocket());
  }, [dispatch]);

  // Render
  return(
    <div className='main-application'>
      <SideNavigation />
      <ChatInterface />
    </div>
  );
}

export default MainPage;