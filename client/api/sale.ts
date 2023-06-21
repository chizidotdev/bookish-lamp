import axios from 'axios';
import { BASE_URL } from '~lib/constants';
import type { Sale, SaleBase } from '~lib/types';

const getSalesURL = (itemID?: string) => {
    if (!itemID) return `${BASE_URL}/sales`;

    return `${BASE_URL}/items/${itemID}/sales`;
};

export const getSales = async (itemID?: string): Promise<Sale[] | undefined> => {
    const response = await axios.get(getSalesURL(itemID));
    return response.data;
};

export const getSaleById = async (saleID: string): Promise<Sale | undefined> => {
    if (!saleID) return;
    const response = await axios.get(`${getSalesURL()}/${saleID}`);
    return response.data;
};

export const newSale = async (sale: SaleBase & { itemID: string }): Promise<Sale> => {
    const { itemID, ...rest } = sale;
    const response = await axios.post(getSalesURL(itemID), rest);
    return response.data;
};

export const updateSale = async (sale: SaleBase & { saleID: string }): Promise<Sale> => {
    const { saleID, ...rest } = sale;
    const response = await axios.put(`${BASE_URL}/sales/${saleID}`, rest);
    return response.data;
};

export const deleteSale = async (id: string): Promise<string> => {
    const response = await fetch(`${BASE_URL}/sales/${id}`, {
        method: 'DELETE',
    });
    console.log('response', response);
    const data = await response.json();
    return data;
};
