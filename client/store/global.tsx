import type { GetServerSideProps } from 'next';
import { getUser } from '~api/user';
import { User } from '~lib/types';

export const getServerSideProps: GetServerSideProps<{
    user: User;
}> = async () => {
    const user = await getUser();
    console.log(user)

    if (!user)
        return {
            redirect: {
                destination: '/auth/login',
                permanent: false,
            },
        };

    return {
        props: { user },
    };
};
