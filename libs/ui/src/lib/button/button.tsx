import { ComponentProps } from 'react';
import { cva, type VariantProps } from 'class-variance-authority';

export type ButtonProps = VariantProps<typeof buttonStyles> &
    Omit<ComponentProps<'button'>, 'className'> & {
        children: React.ReactNode;
    };

const buttonStyles = cva(['btn'], {
    variants: {
        variant: {
            primary: ['btn-primary'],
            secondary: ['btn-secondary'],
            danger: ['btn-error'],
        },
        loading: { true: 'loading' },
        size: {
            small: ['btn-sm'],
        },
    },
    defaultVariants: {
        loading: false,
    },
});

export function Button({
    children,
    loading,
    size,
    variant,
    ...props
}: ButtonProps) {
    return (
        <button className={buttonStyles({ loading, size, variant })} {...props}>
            {children}
        </button>
    );
}

export default Button;
