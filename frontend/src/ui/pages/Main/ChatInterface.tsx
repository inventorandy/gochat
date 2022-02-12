import * as React from 'react';
import { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { GetConversation } from '../../../app/actions/conversation';
import { AppState } from '../../../app/rootReducer';
import { Conversation } from '../../../app/types/conversation';
import ChatEditor from './ChatEditor';
import MessageBox from './MessageBox';

const ChatInterface: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Get the API States
  const convState = useSelector((state: AppState) => state.conversationState);

  // Set the Conversation
  const [channel, setChannel] = useState<Conversation | undefined>(
    convState.currentConversation
  );
  const [error, setError] = useState<string | undefined>();

  // useEffect Hook (called when the conversation state is changed)
  useEffect(() => {
    if (convState.currentConversation) {
      setChannel(convState.currentConversation);
      setError('');
    } else {
      convState.publicConversations?.forEach(conversation => {
        if (conversation.label === "general") {
          dispatch(GetConversation(conversation.id, err => {
            // Handle the Error
            setError(err.message);
          }));
        }
      });
    }
  }, [dispatch, convState]);

  // Render
  return (
    <div className='main-content'>
      <div className='conversation-header'>
        <h1>{channel?.label}</h1>
      </div>
      <div className='conversation-stream'>
        {error &&
          <div className='error align-center'>{error}</div>
        }
        {channel?.messages?.map((message) => {
          return <MessageBox key={message.id} message={message} />;
        })}
      </div>
      <ChatEditor />
    </div>
  );
};

export default ChatInterface;
