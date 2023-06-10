import { Bars } from 'react-loader-spinner';

export function Loading() {
    return (
        <div className='flex flex-col items-center justify-center h-full w-full'>
            <Bars
                height='50'
                width='50'
                color='hsl(var(--s) / var(--tw-text-opacity))'
                ariaLabel='bars-loading'
                wrapperStyle={{}}
                wrapperClass='text-base-300'
                visible={true}
            />
        </div>
    );
}

export default Loading;
