import { useRouteError } from 'react-router-dom';

export default function ErrorPage() {
    const error = useRouteError() as Error | undefined;
    console.error(error);

    return (
        <div id="error-page">
            <h1 className='text-xl font-bold'>Oops!</h1>
            <p>Sorry, an unexpected error has occurred.</p>
            <p>
                <i>{error?.message}</i>
            </p>
        </div>
    );
}
