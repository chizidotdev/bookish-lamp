import { useQuery } from 'react-query';
import { ItemCard } from '@copia/ui';
import { getItems } from './api';

export const Items = () => {
    const { isLoading, data } = useQuery({
        queryKey: ['get-items'],
        queryFn: getItems,
        refetchOnWindowFocus: false,
    });

    return (
        <div className="my-5">
            {isLoading && <div>Loading...</div>}

            <div className="flex flex-col gap-2 mx-2">
                {data &&
                    data.map((item) => <ItemCard key={item.id} item={item} />)}
            </div>
        </div>
    );
};
