import { useRouter } from 'next/router';
import React from 'react';
import AppLogo from '~components/app-logo';
import { useUser } from '~store/user-store';

export function AuthLayout({ children }: { children: React.ReactNode }) {
    const { user } = useUser();
    const { push } = useRouter();

    if (user) {
        push('/');
    }

    return (
        <div className='min-h-screen py-32'>
            <div className='text-center pb-10'>
                <AppLogo />
            </div>

            <div className='max-w-md mx-auto bg-base-100 px-3 sm:px-10 py-12 rounded-2xl'>
                {children}
            </div>
        </div>
    );
}

export default AuthLayout;
