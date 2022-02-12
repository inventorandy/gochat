import * as React from 'react';
import { PropsWithChildren as PWC } from 'react';

type Props = {
  id: string;
  label?: string;
  type: string;
  value?: string;
  placeholder?: string;
  required?: boolean;
  min?: number | string;
  max?: number | string;
  step?: number | string;
  default?: number | string;
  autoComplete?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const FormInput: React.FC<Props> = (props: Props) => {
  // Render
  return (
    <p className='form-input box-sizing'>
      {props.label && <label htmlFor={props.id}>{props.label}</label>}
      <input
        type={props.type}
        id={props.id}
        className='bdr-solid bdr-rounded text-input full-width box-sizing'
        min={props.min}
        max={props.max}
        defaultValue={props.default}
        step={props.step}
        value={props.value}
        placeholder={props.placeholder}
        required={props.required}
        autoComplete={props.autoComplete}
        onChange={props.onChange}
      />
    </p>
  );
};

export default FormInput;
