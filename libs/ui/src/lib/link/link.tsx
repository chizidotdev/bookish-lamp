import { Link as ReactRouterLink, LinkProps } from 'react-router-dom';

export function Link(props: Omit<LinkProps, 'className'>) {
    return (
        <ReactRouterLink
            className="text-primary font-bold underline underline-offset-4"
            {...props}
        />
    );
}

export default Link;
