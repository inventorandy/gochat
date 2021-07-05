import * as React from 'react';
import { Message } from '../../app/types/conversation';

type ChatMessageProps = {
  message: Message;
}

export const ChatMessage: React.FC<ChatMessageProps> = (props: ChatMessageProps) => {
  return(
    <div className="chat-message">
      <h4>Message Sender</h4>
      <div className="content">
        {props.message.message}
      </div>
    </div>
  );
}