import * as React from 'react';
import { useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';
import { CreateConversation } from '../../../app/actions/conversation';
import { ConversationInput } from '../../../app/types/conversation';
import FormError from '../Forms/FormError';
import FormInput from '../Forms/TextInput';
import { closeDialog, Modal, ModalContent, ModalHeader } from '../Modal/Modal';

const CreatePublicChannelDialog: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Set some states
  const [channelName, setChannelName] = useState<string>('');
  const [error, setError] = useState<string>('');

  // Input Handlers
  const handleChannelName = (e: React.ChangeEvent<HTMLInputElement>) => {
    // Set the Value
    setChannelName(e.target.value);
  };

  // Form Handler
  const onCreateChannel = (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Build the Channel Object
    const channel: ConversationInput = {
      label: channelName,
      is_public: true,
    };

    // Call the API Method
    dispatch(
      CreateConversation(
        channel,
        (conversation) => {
          // Close the Dialog
          closeDialog('create-public-channel');
        },
        (err) => {
          // Set the Error
          setError(err.message);
        }
      )
    );
  };

  // Render
  return (
    <Modal id='create-public-channel' hideOnLoad={false}>
      <ModalHeader
        dialog='create-public-channel'
        title='Create Public Channel'
      />
      <ModalContent>
        <form onSubmit={onCreateChannel}>
          <FormError error={error} />
          <FormInput
            type='text'
            id='channel-name'
            placeholder='mychannel'
            value={channelName}
            onChange={handleChannelName}
            required={true}
          />
        </form>
      </ModalContent>
    </Modal>
  );
};

export default CreatePublicChannelDialog;
