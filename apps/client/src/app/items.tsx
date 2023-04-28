import { useQuery } from 'react-query';
import { ItemCard } from '@copia/ui';

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
                    data.map((item) => (
                        <ItemCard key={item.id} item={item} />
                    ))}
            </div>
        </div>
    );
};

type Item = {
    id: string;
    name: string;
    buying_price: number;
    selling_price: number;
    quantity: number;
    CreatedAt: string;
};

const getItems = async (): Promise<Item[]> => {
    const response = await fetch('http://localhost:3333/api/v1/items');
    const data = await response.json();
    return data;
};
