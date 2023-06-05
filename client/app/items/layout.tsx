import { getUser } from "~api/user";

export const metadata = {
    title: 'Items - Copia',
    description: 'Inventory items for your store',
};

export default async function RootLayout({ children }: { children: React.ReactNode }) {
    const user = await getUser();
    console.log('user ==== ', user)

    return (
        <div className='grid xl:grid-cols-[1fr_3fr]'>
            <aside className='' />
            <main className='container mx-auto p-5 min-h-screen'>{children}</main>
        </div>
    );
}
