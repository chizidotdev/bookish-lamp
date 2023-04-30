import { Route, Routes } from 'react-router-dom';
import { Navbar } from '@copia/ui';
import { Items } from './items';
import { CreateItem } from './create-item';
import { DeleteItem } from './delete-item';
import { EditItem } from './edit-item';
import { ItemSales } from './item-sales';
import { CreateSale } from './create-sale';

export function RoutesConfig() {
    return (
        <>
            <Navbar navItems={navItems} />
            <div className="my-5">
            <Routes>
                <Route path="/" element={<Items />}>
                    <Route path="items/delete/:id" element={<DeleteItem />} />
                </Route>
                <Route path="items/create" element={<CreateItem />} />
                <Route path="items/edit/:id" element={<EditItem />} />
                <Route path="items/:id/sales" element={<ItemSales />} />
                <Route path="sales/create" element={<CreateSale />} />
            </Routes>
            </div>
        </>
    );
}

export const navItems = [
    {
        name: 'Items',
        path: '/',
    },
    {
        name: 'Sales',
        path: '/sales',
    },
];
