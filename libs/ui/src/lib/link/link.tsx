import { Link as ReactRouterLink, LinkProps } from 'react-router-dom';

export function Link(props: Omit<LinkProps, 'className'>) {
    return <ReactRouterLink className="btn btn-link" {...props} />;
}

export default Link;
