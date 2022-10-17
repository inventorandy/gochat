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
  const [errorMsg, setErrorMsg] = useState<string>('');

  useEffect(() => {
    if (id) {
      if (conversation === undefined || conversation.id !== id) {
        dispatch(
          SetCurrentConversation(id, (err) => {
            setErrorMsg(err.message);
          })
        );
      }
    }
  }, [id, conversation, dispatch]);

  // Render
  return (
    <>
      {errorMsg === '' && (
        <div className='conversation'>
          <div className='header'>
            <h1>{conversation?.label}</h1>
          </div>
          <MessageList />
          <ChatBox />
        </div>
      )}
      {errorMsg !== '' && (
        <div className='welcome'>
          <h1>Channel Not Found</h1>
          <p>{errorMsg}</p>
        </div>
      )}
    </>
  );
};

export default ConversationPage;
