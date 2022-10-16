import { DateTime } from 'luxon';
import * as React from 'react';
import { currentConversation } from '../../../app/features/conversation';
import { useAppSelector } from '../../../app/hooks';

const MessageList: React.FC = () => {
  // Get the Conversation
  const conversation = useAppSelector(currentConversation);
  // Render
  return (
    <div className='message-list'>
      {conversation?.messages &&
        conversation.messages.map((message) => {
          let dt = DateTime.fromISO(message.created_at || '').toFormat(
            'dd-LLL hh:mm'
          );
          // Render
          return (
            <div key={message.id}>
              <h5>{message.author_name}</h5>
              <h6>Posted on {dt}</h6>
              <p>{message.message}</p>
            </div>
          );
        })}
    </div>
  );
};

export default MessageList;
