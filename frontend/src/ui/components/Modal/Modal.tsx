import * as React from 'react';
import { PropsWithChildren as PWC } from 'react';
import combineCSS from '../../../app/helpers/combineCSS';

/**
 * Open a specified dialog.
 * @param id HTML ID tag of the dialog.
 */
export const openDialog = (id: string) => {
  // Get the Elements
  const dialog = document.getElementById(id) as HTMLElement;
  const overlay = document.getElementById(id + '_overlay') as HTMLElement;

  // Remove the Hidden Class
  dialog.classList.remove('hidden');
  overlay.classList.remove('hidden');
}

/**
 * Close a specified dialog.
 * @param id HTML ID tag of the dialog.
 */
export const closeDialog = (id: string) => {
  // Get the Elements
  const dialog = document.getElementById(id);
  const overlay = document.getElementById(id + '_overlay');

  // Add the Hidden Class
  dialog?.classList.add('hidden');
  overlay?.classList.add('hidden');
}

/**
 * Modal Element Properties
 */
type ModalProps = {
  id: string;
  className?: string;
  hideOnLoad: boolean;
}

export const Modal: React.FC<PWC<ModalProps>> = (props: PWC<ModalProps>) => {
  // Render
  return(
    <>
    <div id={props.id + '_overlay'} className={combineCSS(props.hideOnLoad ? 'modal-overlay hidden' : 'modal-overlay', props.className)} />
    <div id={props.id} className={combineCSS(props.hideOnLoad ? 'modal hidden' : 'modal', props.className)}>
      {props.children}
    </div>
    </>
  );
}

/**
 * Modal Header Element Properties
 */
type ModalHeaderProps = {
  dialog: string;
  title: string;
}

export const ModalHeader: React.FC<PWC<ModalHeaderProps>> = (props: PWC<ModalHeaderProps>) => {
  // Close Dialog Handler
  const onClose = (e: React.MouseEvent<HTMLButtonElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Close the Dialog
    closeDialog(props.dialog);
  }

  // Render
  return(
    <div className="header btl-rounded btr-rounded bdr-solid">
      <h4>{props.title}</h4>
      <button onClick={onClose}>✕</button>
    </div>
  );
}

/**
 * Modal Content Element Properties
 */
type ModalContentProps = {}

export const ModalContent: React.FC<PWC<ModalContentProps>> = (props: PWC<ModalContentProps>) => {
  // Render
  return(
    <div className="content bbl-rounded bbr-rounded bdr-solid">
      {props.children}
    </div>
  );
}