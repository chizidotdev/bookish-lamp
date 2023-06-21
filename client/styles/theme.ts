import { extendTheme, withDefaultColorScheme } from '@chakra-ui/react';

export const customTheme = extendTheme(
    {
        fonts: {
            heading: 'Balsamiq Sans',
            body: 'Balsamiq Sans',
        },
    },
    withDefaultColorScheme({
        colorScheme: 'blue',
    })
);
