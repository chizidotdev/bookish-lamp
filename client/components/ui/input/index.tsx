import { Input as ChakraInput, FormControl, FormErrorMessage, FormLabel } from '@chakra-ui/react';
import { ComponentProps, forwardRef } from 'react';

type InputProps = Omit<ComponentProps<'input'>, 'size'> & {
    label?: string;
    isInvalid?: boolean;
};

export const Input = forwardRef<HTMLInputElement, InputProps>(
    ({ label, isInvalid, ...props }, ref) => {
        return (
            <FormControl isInvalid={isInvalid}>
                <FormLabel>{label}</FormLabel>
                <ChakraInput {...props} ref={ref} />
                {/* <FormErrorMessage>{props.name} is invalid.</FormErrorMessage> */}
            </FormControl>
        );
    }
);

Input.displayName = 'Input';

export default Input;
