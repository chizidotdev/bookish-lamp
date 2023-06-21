import axios from 'axios';
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

export const signup = async ({ email, password }: LoginRequest): Promise<string> => {
    const response = await fetch(`${BASE_URL}/signup`, {
        method: 'POST',
        body: JSON.stringify({ email, password }),
    });
    const data = await response.json();
    return data;
};

export const login = async ({ email, password }: LoginRequest): Promise<{ token: string }> => {
    // const response = await fetch(`${BASE_URL}/login`, {
    //     method: 'POST',
    //     body: JSON.stringify({ email, password }),
    // });
    const response = await axios.post(
        `${BASE_URL}/login`,
        { email, password },
        {
            headers: {
                'Content-Type': 'application/json',
            },
            withCredentials: true,
        }
    );
    if (response.status === 200) {
        localStorage.setItem('token', response.data.token);
        window.dispatchEvent(new Event('storage'));
    }
    return response.data;
};

export const logout = async (): Promise<any> => {
    const response = await fetch(`${BASE_URL}/logout`);
    const data = await response.json();
    if (response.status === 200) {
        localStorage.removeItem('token');
        window.dispatchEvent(new Event('storage'));
    }
    return data;
};
