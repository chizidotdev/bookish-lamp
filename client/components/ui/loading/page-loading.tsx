import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { Loading } from '~components';

export function PageLoading() {
    const router = useRouter();
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        router.events.on('routeChangeStart', () => setLoading(true));
        router.events.on('routeChangeComplete', () => setTimeout(() => setLoading(false), 500));
        router.events.on('routeChangeError', () => setTimeout(() => setLoading(false), 500));
    }, [loading, router.asPath, router.events]);
    return null

    return (
        <>
            {loading && (
                <div className='fixed inset-0 z-[9999] flex h-screen w-screen items-center justify-center bg-base-100 from-blue to-dark'>
                    <Loading />
                </div>
            )}
        </>
    );
}

export default PageLoading;
