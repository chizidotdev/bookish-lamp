import React, { ComponentProps, forwardRef } from 'react';

type CheckboxProps = ComponentProps<'input'> & {
    children: React.ReactNode;
};

export const Checkbox = forwardRef<HTMLInputElement, CheckboxProps>(
    ({ children, ...props }, ref) => {
        return (
            <label className='label cursor-pointer justify-start gap-3'>
                <input ref={ref} type='checkbox' className='checkbox' {...props} />
                {children && <span className='label-text'>{children}</span>}
            </label>
        );
    }
);

Checkbox.displayName = 'Checkbox';

export default Checkbox;

