import { useRouter } from 'next/router';
import { useMutation, useQuery } from 'react-query';
import { getSaleById, getSales, newSale, updateSale } from '~api/sale';
import {useToast} from "@chakra-ui/react";

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
    const toast = useToast()
    return useMutation(newSale, {
        onSuccess: () => {
            toast({
                description: 'Sale created successfully',
                status: 'success',
                duration: 3000,
            });
            back();
        },
    });
}

export function useUpdateSale() {
    const { back } = useRouter();
    const toast = useToast()
    return useMutation(updateSale, {
        onSuccess: () => {
            toast({
                description: 'Sale updated successfully',
                status: 'success',
                duration: 3000,
            });
            back();
        },
    });
}