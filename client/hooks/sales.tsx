import { useRouter } from 'next/router';
import { useMutation, useQuery } from 'react-query';
import { getSaleById, getSales, newSale, updateSale } from '~api/sale';

export function useGetSales(itemId?: string) {
    return useQuery({
        queryKey: ['sales', itemId],
        queryFn: () => getSales(itemId),
    });
}

export function useGetSaleById(saleId: string) {
    return useQuery({
        queryKey: ['sale', saleId],
        queryFn: () => getSaleById(saleId),
    });
}

export function useCreateSale() {
    const { back } = useRouter();
    return useMutation(newSale, {
        onSuccess: () => {
            back();
        },
    });
}

export function useUpdateSale() {
    const { back } = useRouter();
    return useMutation(updateSale, {
        onSuccess: () => {
            back();
        },
    });
}