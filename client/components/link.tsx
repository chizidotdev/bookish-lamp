import { Button } from '@chakra-ui/react';
import Link, { LinkProps } from 'next/link';
import React from 'react';

type CustomLinkProps = LinkProps & {
    children: React.ReactNode;
};

export function CustomLink({ children, ...props }: CustomLinkProps) {
    return (
        <Link {...props}>
            <Button variant='link'>{children}</Button>
        </Link>
    );
}

export default CustomLink;
