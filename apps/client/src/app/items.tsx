import { useQuery } from 'react-query';
import { ItemCard, Link, NotFound } from '@copia/ui';
import { getItems } from './api';
import { Outlet } from 'react-router-dom';

export const Items = () => {
    const { isLoading, data, isSuccess } = useQuery({
        queryKey: ['get-items'],
        queryFn: getItems,
        refetchOnWindowFocus: false,
    });

    return (
        <>
            <h1 className="text-2xl font-bold mb-5">Items</h1>
            {isLoading && <div>Loading...</div>}
            {isSuccess && data?.length === 0 && (
                <NotFound>
                    No items found...&nbsp;
                    <Link to="/items/create">Add Item</Link>
                </NotFound>
            )}

            <div className="flex flex-col gap-3">
                {data &&
                    data.map((item) => <ItemCard key={item.ID} item={item} />)}
            </div>
            <Outlet />
        </>
    );
};
