import React, { type MouseEventHandler } from 'react';
import { cva } from 'class-variance-authority';
import type { VariantProps } from 'class-variance-authority';
import { ThreeDots } from 'react-loader-spinner';

type Props = VariantProps<typeof buttonStyles> & {
    children: React.ReactNode;
    onClick?: React.MouseEventHandler<HTMLButtonElement>;
    type?: 'button' | 'submit' | 'reset';
    loading?: boolean;
};

const buttonStyles = cva(['btn'], {
    variants: {
        variant: {
            primary: 'btn-primary',
            secondary: 'btn-secondary',
            accent: 'btn-accent',
        },
        size: {
            small: ['btn-sm'],
            medium: ['btn-md'],
        },
        autoWidth: {
            true: ['w-full'],
        },
    },
});

const Button = ({ children, loading, onClick, type, ...props }: Props) => {
    const handleClick: MouseEventHandler<HTMLButtonElement> = (event) => {
        if (loading || !onClick) return;

        onClick(event);
    };

    const content = loading ? (
        <ThreeDots
            height='10'
            width='40'
            radius='10'
            color='#fff'
            ariaLabel='three-dots-loading'
            visible={true}
        />
    ) : (
        children
    );

    return (
        <button className={buttonStyles(props)} onClick={handleClick} type={type}>
            {content}
        </button>
    );
};

export default Button;
