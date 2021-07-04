import * as React from 'react';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { GetPublicConversations } from '../../app/actions/conversation';
import { AppState } from '../../app/rootReducer';
import { ConversationList } from '../components/ConversationList';

const ChatInterface: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Define the API States
  const conversationState = useSelector((state: AppState) => state.conversationState);

  // Load the Required API Data
  useEffect(() => {
    dispatch(GetPublicConversations())
  }, []);

  return(
    <div className="chat-interface-page">
      <div className="sidebar">
        <h2>GoChat!</h2>
        <h3>Channels</h3>
        <ConversationList conversations={conversationState.publicConversations} />
        <h3>Private Channels</h3>
      </div>
      <div className="main-chat-container">
        <p>This is the chat window</p>
      </div>
    </div>
  );
}

export default ChatInterface;