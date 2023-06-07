import { BASE_URL } from '~lib/constants';
import { User } from '~lib/types';

type LoginRequest = {
    email: string;
    password: string;
};

export const getUser = async (): Promise<User | false> => {
    const response = await fetch(`${BASE_URL}/validateToken`);
    if (response.status === 401) {
        return false;
    }

    const data = await response.json();
    return data;
};

export const login = async ({ email, password }: LoginRequest): Promise<{ token: string }> => {
    const response = await fetch(`${BASE_URL}/login`, {
        method: 'POST',
        body: JSON.stringify({ email, password }),
    });
    const data = await response.json();
    return data;
};

export const logout = async (): Promise<any> => {
    const response = await fetch(`${BASE_URL}/logout`);
    const data = await response.json();
    return data;
};
