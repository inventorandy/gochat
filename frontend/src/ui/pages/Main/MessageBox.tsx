import * as React from 'react';
import { TimeSinceMessage } from '../../../app/helpers/date';
import { Message } from '../../../app/types/conversation';

type Props = {
  message: Message;
}

const MessageBox: React.FC<Props> = (props: Props) => {
  // Render
  return(
    <div className="message box-sizing">
      <div className="author bold">{props.message.author_name}</div>
      <div className="date right-align">{TimeSinceMessage(props.message.created_at)}</div>
      <div className="content spaced-line">{props.message.message}</div>
    </div>
  );
}

export default MessageBox;