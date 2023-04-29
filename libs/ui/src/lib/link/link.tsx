import { Link as ReactRouterLink, LinkProps } from 'react-router-dom';

export function Link(props: Omit<LinkProps, 'className'>) {
    return <ReactRouterLink className="" {...props} />;
}

export default Link;
