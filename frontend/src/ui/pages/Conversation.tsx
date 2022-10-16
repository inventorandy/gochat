import * as React from 'react';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { SetCurrentConversation } from '../../app/actions/conversation';
import { currentConversation } from '../../app/features/conversation';
import { useAppDispatch, useAppSelector } from '../../app/hooks';
import ChatBox from '../components/conversation/ChatBox';
import MessageList from '../components/conversation/MessageList';
import '../scss/conversation.scss';

const ConversationPage: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useAppDispatch();

  // Get the Conversation ID
  const { id } = useParams();

  // Get the Conversation
  const conversation = useAppSelector(currentConversation);

  // Error State
  const [error, setError] = useState<string>('');

  useEffect(() => {
    // console.log('rendering');
    if (id) {
      if (conversation === undefined) {
        dispatch(
          SetCurrentConversation(id, (err) => {
            setError(err.message);
          })
        );
      }
    }
  }, [id, conversation, dispatch]);

  // Render
  return (
    <div className='conversation'>
      <div className='header'>
        <h1>{conversation?.label}</h1>
      </div>
      <MessageList />
      <ChatBox />
    </div>
  );
};

export default ConversationPage;
