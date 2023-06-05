import React, { ComponentProps, forwardRef } from 'react';

type InputProps = ComponentProps<'input'> & {
    label?: string;
};

export const Input = forwardRef<HTMLInputElement, InputProps>(({ label, ...props }, ref) => {
    return (
        <div className='w-full'>
            {label && (
                <label className='label'>
                    <span className='label-text'>{label}</span>
                </label>
            )}
            <input {...props} className='input input-bordered w-full' ref={ref} />
        </div>
    );
});

Input.displayName = 'Input';

export default Input;
