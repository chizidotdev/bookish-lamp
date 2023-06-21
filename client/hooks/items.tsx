import { useQuery } from 'react-query';
import { getItemById, getItems } from '~api/item';

export function useGetItemById(itemId: string) {
    return useQuery(['items', itemId], {
        queryFn: () => getItemById(itemId),
    });
}

export function useGetItems() {
    return useQuery(['items'], {
        queryFn: getItems,
    });
}
