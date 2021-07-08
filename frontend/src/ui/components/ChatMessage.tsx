import * as React from 'react';
import { rfc3339ToSystemDate } from '../../app/helpers/rfcDate';
import { Message } from '../../app/types/conversation';

type ChatMessageProps = {
  message: Message;
}

export const ChatMessage: React.FC<ChatMessageProps> = (props: ChatMessageProps) => {
  return(
    <div className="chat-message">
      <h4>
        <span className="message-author">
          {props.message.author_name}
        </span>
        <span className="message-info">
          {rfc3339ToSystemDate(props.message.created_at)}
        </span>
      </h4>
      <div className="content">
        {props.message.message}
      </div>
    </div>
  );
}