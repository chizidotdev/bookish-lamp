import React from 'react';
import AppLogo from '~components/app-logo';

export function AuthLayout({ children }: { children: React.ReactNode }) {
    return (
        <div>
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
