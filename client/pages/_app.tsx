import { type AppType } from 'next/app';
import { QueryClient, QueryClientProvider, QueryCache } from 'react-query';
// import ReactGA from "react-ga";

import '../styles/globals.css';
import { RootLayout } from '~components';
import { UserProvider } from '~store/user-store';
import { ChakraProvider } from '@chakra-ui/react';
import { customTheme } from '~styles/theme';

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            refetchOnWindowFocus: false,
        staleTime: 1000 * 60 * 5, // 5 minutes
        },
    },
});

const queryCache = new QueryCache()

const MyApp: AppType = ({ Component, pageProps }) => {
    // ReactGA.initialize("G-FGNLC0J6Q3");

    console.log(queryCache.findAll())

    return (
        <QueryClientProvider client={queryClient}>
            <ChakraProvider theme={customTheme}>
                <UserProvider>
                    <RootLayout>
                        <Component {...pageProps} />
                    </RootLayout>
                </UserProvider>
            </ChakraProvider>
        </QueryClientProvider>
    );
};

export default MyApp;
