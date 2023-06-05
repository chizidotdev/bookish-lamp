import { BASE_URL } from '~lib/constants';

type LoginRequest = {
    email: string;
    password: string;
};

export const getUser = async (): Promise<any> => {
    const response = await fetch(`${BASE_URL}/validateToken`, {
        credentials: 'include',
    });
    const data = await response.json();
    return data;
};

export const login = async ({ email, password }: LoginRequest): Promise<any> => {
    const response = await fetch(`${BASE_URL}/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include',
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
