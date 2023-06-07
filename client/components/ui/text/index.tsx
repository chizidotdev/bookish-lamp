import React from 'react';
import { VariantProps, cva } from 'class-variance-authority';

type TextProps = VariantProps<typeof textStyles> & {
    children: React.ReactNode;
};

const textStyles = cva([''], {
    variants: {
        variant: {
            p: '',
            h1: 'text-3xl font-bold',
            h2: 'text-2xl font-bold',
            h3: 'text-xl font-bold',
        },
    },
    defaultVariants: {
        variant: 'p',
    },
});

export function Text({ children, ...props }: TextProps) {
    const Component = props.variant ?? 'p';

    return <Component className={textStyles(props)} >{children}</Component>;
}

export default Text;
