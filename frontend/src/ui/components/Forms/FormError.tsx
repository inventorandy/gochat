import * as React from 'react';
import { useState } from 'react';

type Props = {
  error: string;
}

const FormError: React.FC<Props> = (props: Props) => {
  // Render
  return(
    <div className={props.error === '' ? 'form-error hidden' : 'form-error'}>
      {props.error}
    </div>
  );
}

export default FormError;