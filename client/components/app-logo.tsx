import Link from 'next/link';
import React from 'react';
import Text from '~components/text';

export function AppLogo() {
    return (
        <Link href='/'>
            <Text variant='h1'>Copia</Text>
        </Link>
    );
}

export default AppLogo;
