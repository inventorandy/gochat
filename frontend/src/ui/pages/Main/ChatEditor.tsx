import * as React from 'react';

const ChatEditor: React.FC = () => {
  // Render
  return (
    <form className='chat-editor box-sizing'>
      <textarea
        id='message'
        className='bdr-none'
        placeholder='Type your message in here...'
      />
      <button className='bdr-rounded bdr-solid'>Send</button>
    </form>
  );
};

export default ChatEditor;
