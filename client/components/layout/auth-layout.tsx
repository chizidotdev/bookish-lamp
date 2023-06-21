import { Box, useColorModeValue } from '@chakra-ui/react';
import React from 'react';
import AppLogo from '~components/app-logo';

export function AuthLayout({ children }: { children: React.ReactNode }) {
    const bgColor = useColorModeValue('white', 'gray.700');

    return (
        <Box>
            <Box paddingBlock="12" textAlign="center">
                <AppLogo />
            </Box>

            <Box rounded="lg" bg={bgColor} maxW="md" mx="auto" boxShadow="lg" p={8}>
                {children}
            </Box>
        </Box>
    );
}

export default AuthLayout;
