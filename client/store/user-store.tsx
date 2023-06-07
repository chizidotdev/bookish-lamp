import React, { createContext, useContext, useEffect } from 'react';
import { useQuery } from 'react-query';
import { getUser } from '~api/user';
import { User } from '~lib/types';

type TUserProps = {
    children: React.ReactNode;
};

type UserContextProps = {
    user?: User | false;
    isLoading: boolean;
};

const UserContext = createContext<UserContextProps>({
    user: false,
    isLoading: true,
});

export const UserProvider = ({ children }: TUserProps) => {
    const {
        data: user,
        isLoading,
        refetch,
        isRefetching,
    } = useQuery('user', {
        queryFn: getUser,
        refetchOnWindowFocus: false,
        refetchOnMount: false,
    });

    console.log(isRefetching);
    useEffect(() => {
        const refetchUser = () => {
            console.log('storage change');
            refetch();
        };

        window.addEventListener('storage', refetchUser);

        return () => {
            window.removeEventListener('storage', refetchUser);
        };
    }, [refetch]);

    return <UserContext.Provider value={{ user, isLoading }}>{children}</UserContext.Provider>;
};

export const useUser = () => {
    if (!UserContext) {
        throw new Error('useUser should be used inside a UserProvider');
    }

    return useContext(UserContext);
};
