import Link from 'next/link';
import { Text } from '~components';
import { useUser } from '~store/user-store';

export default function Home() {
    const { user } = useUser();

    let body = (
        <>
            <Text>Welcome to Copia</Text>
            <Text>Log in with your Copia account to continue</Text>

            <div className='mt-5 flex justify-center gap-4'>
                <Link href='/auth/login' className='btn'>
                    Login
                </Link>
                <Link href='/auth/signup' className='btn btn-secondary'>
                    Get started
                </Link>
            </div>
        </>
    );

    if (user) {
        body = (
            <>
                <Text>Signed in as {user.email}</Text>
                <Text>Continue where you left off</Text>

                <div className='mt-5 flex justify-center gap-4'>
                    <Link href='/items' className='btn btn-secondary'>
                        Dashboard
                    </Link>
                </div>
            </>
        );
    }

    return (
        <>
            <div className='flex flex-col items-center pt-[20vh] px-10 py-24 sm:px-24 text-center'>
                <div className='mb-4'>
                    <Text variant='h1'>Copia</Text>
                </div>
                {body}
            </div>
        </>
    );
}
