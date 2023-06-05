import Link, { LinkProps } from 'next/link';
import React from 'react';

type CustomLinkProps = LinkProps & {
    children: React.ReactNode;
};

export function CustomLink({ children, ...props }: CustomLinkProps) {
    return (
        <Link {...props}>
            <span className='btn-link'>{children}</span>
        </Link>
    );
}

export default CustomLink;
