import { useRouter } from 'next/router';
import React, { useEffect } from 'react';
import { useUser } from '~store/user-store';

export function ProtectedLayout({ children }: { children: React.ReactNode }) {
    const { user, isLoading } = useUser();
    const { push } = useRouter();

    useEffect(() => {
        if (!user && !isLoading) {
            push('/auth/login');
        }
    }, [user, isLoading, push]);

    return <>{children}</>;
}

export default ProtectedLayout;