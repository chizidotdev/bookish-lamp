export const metadata = {
    title: 'Items - Copia',
    description: 'Inventory items for your store',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <div className="grid md:grid-cols-[1fr_3fr]">
            <aside className="" />
            <main className="container mx-auto p-5 min-h-screen">{children}</main>
        </div>
    );
}
