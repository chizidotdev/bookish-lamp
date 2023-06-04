import './globals.css';
import { Balsamiq_Sans } from 'next/font/google';
import Navbar from './navbar';

const balsamiq = Balsamiq_Sans({
    weight: ['400', '700'],
    subsets: ['latin'],
    variable: '--font-balsamiq',
});

export const metadata = {
    title: 'Copia',
    description: 'Generated by create next app',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang='en'>
            <body className={`${balsamiq.variable} font-balsamiq`}>
                <Navbar />
                <main className='container mx-auto p-5 min-h-screen'>{children}</main>
            </body>
        </html>
    );
}