import { useQuery } from 'react-query';
import { ItemCard } from '@copia/ui';
import { getItems } from './api';
import { Outlet } from 'react-router-dom';

export const Items = () => {
    const { isLoading, data } = useQuery({
        queryKey: ['get-items'],
        queryFn: getItems,
        refetchOnWindowFocus: false,
    });

    return (
        <div className="my-5">
            {isLoading && <div>Loading...</div>}

            <div className="flex flex-col gap-3">
                {data &&
                    data.map((item) => <ItemCard key={item.ID} item={item} />)}
            </div>
            <Outlet />
        </div>
    );
};
