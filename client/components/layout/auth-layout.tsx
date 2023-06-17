import { Box, useColorModeValue } from '@chakra-ui/react';
import React from 'react';
import AppLogo from '~components/app-logo';

export function AuthLayout({ children }: { children: React.ReactNode }) {
    const bgColor = useColorModeValue('white', 'gray.700');

    return (
        <div>
            <div className="text-center pb-10">
                <AppLogo />
            </div>

            <Box rounded="lg" bg={bgColor} maxW="md" mx="auto" boxShadow="lg" p={8}>
                {children}
            </Box>
        </div>
    );
}

export default AuthLayout;
