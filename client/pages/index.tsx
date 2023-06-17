import { Box, Button, ButtonGroup } from '@chakra-ui/react';
import Link from 'next/link';
import { Text } from '~components';
import { useUser } from '~store/user-store';

export default function Home() {
    const { user } = useUser();

    let body = (
        <>
            <Text>Welcome to Copia</Text>
            <Text>Log in with your Copia account to continue</Text>

            <ButtonGroup mt='5'>
                <Link href='/auth/login'>
                    <Button variant='outline'>Login</Button>
                </Link>
                <Link href='/auth/signup'>
                    <Button>Get started</Button>
                </Link>
            </ButtonGroup>
        </>
    );

    if (user) {
        body = (
            <>
                <Text>Signed in as {user.email}</Text>
                <Text>Continue where you left off</Text>

                <Box>
                    <div className='mt-5 flex justify-center gap-4'>
                        <Link href='/dashboard'>
                            <Button>Dashboard</Button>
                        </Link>
                    </div>
                </Box>
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
