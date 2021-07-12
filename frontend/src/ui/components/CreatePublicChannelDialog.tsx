import * as React from 'react';
import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { CreateConversation, GetPublicConversations } from '../../app/actions/conversation';
import { AppState } from '../../app/rootReducer';
import { Conversation } from '../../app/types/conversation';
import { ErrorMessage } from '../../app/types/error';
import ModalDialog, { closeDialog } from './ModalDialog';

const CreatePublicChannelDialog: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Define the API States
  const conversationState = useSelector((state: AppState) => state.conversationState);

  // Set the States
  const [channelName, setChannelName] = useState<string>();
  const [error, setError] = useState<string>();

  // Handle Channel Name Method
  const handleChannelName = (e: React.ChangeEvent<HTMLInputElement>) => {
    const channelName = (e.target as HTMLInputElement).value;
    setChannelName(channelName);
  }

  const createChannel = (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Create the Conversation Object
    let conversation: Conversation = {
      label: channelName,
      is_public: true
    }

    // Send the API Request
    dispatch(CreateConversation(conversation, channelCreated, channelCreateFailed));
  }

  // Callback for Successful Channel Creation
  const channelCreated = (conversation: Conversation) => {
    // Close the Dialog
    closeDialog("create-public-channel-dialog");
  }

  // Callback for Failed Channel Creation
  const channelCreateFailed = (error: ErrorMessage) => {
    // Get the Error Box Element
    let errorBox = document.getElementById("create-public-channel-error") as HTMLElement;
    setError(error.message);
    errorBox.classList.remove("hidden");
  }

  // Render Method
  return(
    <ModalDialog id="create-public-channel-dialog" title="Create Channel" hideOnLoad={true} className="create-public-channel-dialog">
      <form action="/create-public-channel" method="post" onSubmit={createChannel}>
      <p id="create-public-channel-error" className="notification error hidden">{error}</p>
        <p>
          <label htmlFor="public-channel-name">Channel Name</label>
          <input type="text" id="public-channel-name" name="public_channel_name" autoComplete="Off" onChange={handleChannelName} />
        </p>
        <p>
          <button className="primary">Create Channel</button>
        </p>
      </form>
    </ModalDialog>
  );
}

export default CreatePublicChannelDialog;