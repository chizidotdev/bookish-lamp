import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import ErrorPage from './404';
import { Navbar } from '@copia/ui';
import { Items } from './items';
import { CreateItem } from './create-item';
import { DeleteItem } from './delete-item';
import { EditItem } from './edit-item';
import { ItemSales } from './item-sales';
import { CreateSale } from './create-sale';

const router = createBrowserRouter([
    {
        path: '/',
        element: <Navbar />,
        errorElement: <ErrorPage />,
        children: [
            {
                index: true,
                element: <Items />,
            },
            {
                path: 'items/create',
                element: <CreateItem />,
            },
            {
                path: 'items/edit/:id',
                element: <EditItem />,
            },
            {
                path: 'items/delete/:id',
                element: <DeleteItem />,
            },
            {
                path: 'items/:id/sales',
                element: <ItemSales />,
            },
            {
                path: 'sales/create',
                element: <CreateSale />,
            },
        ],
    },
]);

export function RoutesConfig() {
    return <RouterProvider router={router} />;
}
