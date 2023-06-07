import { useRouter } from 'next/router';
import React from 'react';
import { useUser } from '~store/user-store';

export function ItemsLayout({ children }: { children: React.ReactNode }) {
    const { user } = useUser();
    const { push } = useRouter();

    if (!user) {
        push('/auth/login');
    }

    return <div>{children}</div>;
}

export default ItemsLayout;
