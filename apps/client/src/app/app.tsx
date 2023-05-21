import { QueryClient, QueryClientProvider } from 'react-query';
import { RoutesConfig } from './routes';

const queryClient = new QueryClient();

function App() {
    return (
        <div className="max-w-4xl mx-auto px-2 sm:px-5">
            <QueryClientProvider client={queryClient}>
                <RoutesConfig />
            </QueryClientProvider>
        </div>
    );
}

export default App;
