import { ComponentProps } from 'react';
import { cva, type VariantProps } from 'class-variance-authority';

export type ButtonProps = VariantProps<typeof buttonStyles> &
    Omit<ComponentProps<'button'>, 'className'> & {
        children: React.ReactNode;
    };

const buttonStyles = cva(['btn'], {
    variants: {
        loading: { true: 'loading' },
        size: {
            small: ['btn-sm'],
        },
    },
    defaultVariants: {
        loading: false,
    },
});

export function Button({ children, loading, size }: ButtonProps) {
    return (
        <button className={buttonStyles({ loading, size })}>{children}</button>
    );
}

export default Button;
