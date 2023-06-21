import { BASE_URL } from '~lib/constants';
import { Inventory } from '~lib/types';

export const getInventory = async (): Promise<Inventory> => {
    const response = await fetch(`${BASE_URL}/inventory`, {
        credentials: 'include',
    });
    const data = await response.json();
    return data;
};
