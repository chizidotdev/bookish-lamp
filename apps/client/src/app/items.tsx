import { useQuery, useQueryClient } from 'react-query';

export const Items = () => {
    const { isLoading, error, data } = useQuery('get-items', getItems);
    const queryClient = useQueryClient();

    function handleInvalidate() {
        // Invalidate and refetch any query with the key 'get-items'
        queryClient.invalidateQueries('get-items')
    }

    return (
        <div>
            <h1>Items</h1>
            <button onClick={handleInvalidate} className='btn'>Invalidate</button>

            {isLoading && <div>Loading...</div>}
            {error && <div>Error: {error}</div>}

            {data && (
                <div>
                    {data.map((item) => (
                        <div key={item.id}>
                            <div>{item.name}</div>
                            <div>{item.buying_price}</div>
                            <div>{item.selling_price}</div>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

type Item = {
    id: string;
    name: string;
    buying_price: number;
    selling_price: number;
};

const getItems = async (): Promise<Item[]> => {
    const response = await fetch('http://localhost:3333/api/v1/items');
    const data = await response.json();
    return data;
};
