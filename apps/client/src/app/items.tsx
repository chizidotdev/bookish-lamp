import { useQuery, useQueryClient } from 'react-query';

export const Items = () => {
    const { isLoading, error, data } = useQuery({
        queryKey: ['get-items'],
        queryFn: getItems,
    });
    const queryClient = useQueryClient();

    function handleInvalidate() {
        // Invalidate and refetch any query with the key 'get-items'
        queryClient.invalidateQueries({ queryKey: ['get-items'] });
    }

    return (
        <>
            <h1>Items</h1>
            <button onClick={handleInvalidate} className="btn mb-2">
                Refetch Items
            </button>

            {isLoading && <div>Loading...</div>}
            {error && <div>Error: {error}</div>}

            <div className='overflow-x-auto'>
            {data && (
                <table className='table w-full'>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Buying Price</th>
                            <th>Selling Price</th>
                        </tr>
                    </thead>
                    <tbody>
                        {data.map((item) => (
                            <tr key={item.id} className='hover'>
                                <td>{item.name}</td>
                                <td>{item.buying_price}</td>
                                <td>{item.selling_price}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            )}
            </div>
        </>
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
