import * as React from 'react';
import { useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { CloseConversationWebsocket, GetPrivateConversations, GetPublicConversations } from '../../app/actions/conversation';
import { ConnectConversationWebsocket } from '../../app/actions/conversation';
import { AppState } from '../../app/rootReducer';
import { ConversationList } from '../components/ConversationList';
import { MessageEditor } from '../components/MessageEditor';
import { MessageList } from '../components/MessageList';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faPlus } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import ModalDialog, { openDialog } from '../components/ModalDialog';
import CreatePublicChannelDialog from '../components/CreatePublicChannelDialog';

// Add the Required Icons to the Library
library.add( faPlus );

const ChatInterface: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Set the Browser History
  const history = useHistory();

  // Define the API States
  const conversationState = useSelector((state: AppState) => state.conversationState);

  // Load the Required API Data
  useEffect(() => {
    dispatch(GetPublicConversations());
    dispatch(GetPrivateConversations());
    dispatch(ConnectConversationWebsocket())
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // Logout Method (called when logout button pressed)
  const logout = () => {
    CloseConversationWebsocket();
    localStorage.removeItem("authToken");
    history.push("/auth/login");
  }

  return(
    <div className="chat-interface-page">
      <div className="sidebar">
        <h2>GoChat!</h2>
        <h3>
          <span className="channel-label">
            Channels
          </span>
          <button onClick={(e) => openDialog("create-public-channel-dialog")}>
            <FontAwesomeIcon icon="plus" />
          </button>
        </h3>
        <ConversationList conversations={conversationState.publicConversations} />
        <h3>Private Channels</h3>
        <ConversationList conversations={conversationState.privateConversations} />
        <hr />
        <button className="logout" onClick={logout}>Log Out</button>
      </div>
      <div className="main-chat-container">
        <div className="inner">
          <h1>{conversationState.currentConversation?.label}</h1>
          <MessageList conversation={conversationState.currentConversation} />
          <MessageEditor />
        </div>
      </div>
      <CreatePublicChannelDialog />
    </div>
  );
}

export default ChatInterface;