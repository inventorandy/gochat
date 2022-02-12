import * as React from 'react';
import FormInput from '../Forms/TextInput';
import { Modal, ModalContent, ModalHeader } from '../Modal/Modal';

const CreatePublicChannelDialog: React.FC = () => {
  // Render
  return(
    <Modal id='create-public-channel' hideOnLoad={false}>
      <ModalHeader dialog='create-public-channel' title='Create Public Channel' />
      <ModalContent>
        <form>
          <FormInput type='text' id='channel-name' placeholder='mychannel' />
        </form>
      </ModalContent>
    </Modal>
  );
}

export default CreatePublicChannelDialog;