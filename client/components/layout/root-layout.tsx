import React, { useEffect } from 'react';
import { Navbar, PageLoading } from '~components';
import { interceptor } from '~api/interceptors';

export function RootLayout({ children }: { children: React.ReactNode }) {
    useEffect(() => {
        interceptor();
    }, []);

    return (
        <div className='font-balsamiq'>
            <PageLoading />
            <Navbar />
            <div className='max-w-7xl mx-auto px-5 sm:px-10 min-h-screen py-12 md:py-28'>{children}</div>
        </div>
    );
}

export default RootLayout;
