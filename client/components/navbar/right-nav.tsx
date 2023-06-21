import {
    Avatar,
    Button,
    Flex,
    Menu,
    MenuButton,
    MenuDivider,
    MenuItem,
    MenuList,
    Stack,
} from '@chakra-ui/react';
import { logout } from '~api/user';
import { User } from '~lib/types';
import { useUser } from '~store/user-store';

export default function RightNav() {
    const { user } = useUser();
    const handleLogout = async () => {
        await logout();
    };

    return (
        <Stack flex={{ base: 1, md: 0 }} justify={'flex-end'} direction={'row'} spacing={6}>
            <Flex alignItems={'center'}>
                <Menu>
                    <MenuButton
                        as={Button}
                        rounded={'full'}
                        variant={'link'}
                        cursor={'pointer'}
                        minW={0}
                    >
                        <Avatar
                            size="sm"
                            src="https://avatars.dicebear.com/api/adventurer-neutral/mail%40ashallendesign.co.uk.svg"
                        />
                    </MenuButton>
                    <MenuList>
                        <MenuItem onClick={handleLogout}>Logout</MenuItem>
                        <MenuDivider />
                        <MenuItem>{(user as User)?.email}</MenuItem>
                    </MenuList>
                </Menu>
            </Flex>
        </Stack>
    );
}
