import { Route, Routes } from 'react-router-dom';
import { Navbar } from '@copia/ui';
import { Items } from './items';
import { CreateItem } from './create-item';
import { DeleteItem } from './delete-item';
import { EditItem } from './edit-item';

export function RoutesConfig() {
    return (
        <>
            <Navbar navItems={navItems} />
            <Routes>
                <Route path="/" element={<Items />}>
                    <Route path="items/delete/:id" element={<DeleteItem />} />
                </Route>
                <Route path="items/create" element={<CreateItem />} />
                <Route path="items/edit/:id" element={<EditItem />} />
            </Routes>
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
