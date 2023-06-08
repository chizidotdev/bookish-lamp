import { type AppType } from 'next/app';
import { QueryClient, QueryClientProvider } from 'react-query';
// import ReactGA from "react-ga";

import '../styles/globals.css';
import { RootLayout } from '~components';
import { UserProvider } from '~store/user-store';

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            refetchOnWindowFocus: false,
        },
    },
});

const MyApp: AppType = ({ Component, pageProps }) => {
    // ReactGA.initialize("G-FGNLC0J6Q3");

    return (
        <QueryClientProvider client={queryClient}>
            <UserProvider>
                <RootLayout>
                    <Component {...pageProps} />
                </RootLayout>
            </UserProvider>
        </QueryClientProvider>
    );
};

export default MyApp;
