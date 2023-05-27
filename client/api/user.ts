import { BASE_URL } from "~lib/constants";

export const getUser = async (): Promise<any> => {
    const response = await fetch(`${BASE_URL}/user`);
    const data = await response.json();
    return data;
};

export const logout = async (): Promise<any> => {
    const response = await fetch(`${BASE_URL}/logout`);
    const data = await response.json();
    return data;
};