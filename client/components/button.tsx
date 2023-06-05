import React, { ComponentProps } from 'react';

type ButtonProps = ComponentProps<'button'> & {
    children: React.ReactNode;
};

export function Button({ children, ...props }: ButtonProps) {
    return (
        <button className='btn w-full' {...props}>
            {children}
        </button>
    );
}

export default Button;
