import * as React from 'react';
import { useState } from 'react';
import { SendMessage } from '../../../app/actions/conversation';
import { currentConversation } from '../../../app/features/conversation';
import { currentUser } from '../../../app/features/user';
import { useAppDispatch, useAppSelector } from '../../../app/hooks';
import { Message } from '../../../app/types/conversation';

const ChatBox: React.FC = () => {
  const dispatch = useAppDispatch();

  // Get the User
  const user = useAppSelector(currentUser);

  // Get the Conversation
  const conversation = useAppSelector(currentConversation);

  // Set the Form States
  const [message, setMessage] = useState<string>('');

  // Input Handlers
  const handleMessage = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setMessage(e.target.value);
  };

  // Form Handler
  const onSendMessage = (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Check we are logged in
    if (user) {
      // Check we have a conversation selected
      if (conversation) {
        // Build the Message
        let msg: Message = {
          conversation_id: conversation.id,
          author_id: user.id,
          message: message,
        };

        // Send the Message
        dispatch(SendMessage(msg));
        setMessage('');
      }
    }
  };
  // Render
  return (
    <form className='chat-box' onSubmit={onSendMessage}>
      <textarea
        id='message'
        placeholder='Type message in here'
        value={message}
        onChange={handleMessage}
      />
      <button>Send</button>
    </form>
  );
};

export default ChatBox;
