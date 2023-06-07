import React from 'react';

type Props = React.ComponentPropsWithoutRef<'textarea'> & {
    name: string;
    className?: never;
};

const Textarea = (props: Props) => {
    return (
        <textarea
            {...props}
            className='w-full border-transparent bg-transparent focus:outline-none'
            cols={10}
            rows={4}
        />
    );
};

export default Textarea;
