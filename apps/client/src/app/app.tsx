import { QueryClient, QueryClientProvider } from 'react-query';
import { RoutesConfig } from './routes';
import { BrowserRouter } from 'react-router-dom';

const queryClient = new QueryClient();

function App() {
    return (
        <QueryClientProvider client={queryClient}>
            <BrowserRouter>
                <RoutesConfig />
            </BrowserRouter>
        </QueryClientProvider>
    );
}

export default App;
