export const metadata = {
    title: 'Items - Copia',
    description: 'Inventory items for your store',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <div className="grid xl:grid-cols-[1fr_3fr]">
            <aside className="" />
            <main className="container mx-auto min-h-screen">{children}</main>
        </div>
    );
}
