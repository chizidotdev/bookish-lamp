import { useGetItems } from '~hooks/items';
import { useEffect, useState } from 'react';

export type UseSelectSalesItemReturnType = ReturnType<
    typeof useSelectSalesItem
>;

export function useSelectSalesItem() {
    const { data, isLoading } = useGetItems();
    const items = data ?? [];
    const [selected, setSelected] = useState(items[0]);
    const [query, setQuery] = useState('');

    useEffect(() => {
        if (!isLoading && data) {
            setSelected(data[0]);
        }
    }, [isLoading, data]);

    const filteredItems =
        query === ''
            ? items
            : items?.filter((item) =>
                  item.title
                      .toLowerCase()
                      .replace(/\s+/g, '')
                      .includes(
                          query.toLowerCase().replace(/\s+/g, '')
                      )
              );

    return {
        items,
        filteredItems,
        selected,
        setSelected,
        query,
        setQuery,
    };
}
