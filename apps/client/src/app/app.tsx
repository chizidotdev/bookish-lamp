import { QueryClient, QueryClientProvider } from 'react-query';
import { RoutesConfig } from './routes';
import { BrowserRouter } from 'react-router-dom';

const queryClient = new QueryClient();

function App() {
    return (
        <div className="max-w-4xl mx-auto">
            <QueryClientProvider client={queryClient}>
                <BrowserRouter>
                    <RoutesConfig />
                </BrowserRouter>
            </QueryClientProvider>
        </div>
    );
}

export default App;
